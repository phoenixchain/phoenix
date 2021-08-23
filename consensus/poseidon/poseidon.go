// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// Package clique implements the proof-of-authority consensus engine.
package poseidon

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/systemcontracts"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/vrf"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"io"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/misc"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/oqs/oqs_crypto"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/hashicorp/golang-lru"
	"golang.org/x/crypto/sha3"
)

const (
	inmemorySignatures = 4096 // Number of recent block signatures to keep in memory

	vrfExpectedSize = 3.0

	validatorBytesLength = common.AddressLength

	nonceSignSize = 255
	heartRate     = 100

	gasCap = 50000000
)

// Spos proof-of-authority protocol constants.
var (
	extraVanity = 32                         // Fixed number of extra-data prefix bytes reserved for signer vanity
	extraSeal   = crypto.SignatureLength // Fixed number of extra-data suffix bytes reserved for signer seal
	extraVrf    = 81

	uncleHash = types.CalcUncleHash(nil) // Always Keccak256(RLP([])) as uncles are meaningless outside of PoW.

	ether = new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
)

// Various error messages to mark blocks invalid. These should be private to
// prevent engine specific errors from being referenced in the remainder of the
// codebase, inherently breaking if the engine is swapped out. Please put common
// error types into the consensus package.
var (
	// errUnknownBlock is returned when the list of signers is requested for a block
	// that is not part of the local blockchain.
	errUnknownBlock = errors.New("unknown block")

	// errInvalidCheckpointBeneficiary is returned if a checkpoint/epoch transition
	// block has a beneficiary set to non-zeroes.
	errInvalidCheckpointBeneficiary = errors.New("beneficiary in checkpoint block non-zero")

	// errInvalidVote is returned if a nonce value is something else that the two
	// allowed constants of 0x00..0 or 0xff..f.
	errInvalidVote = errors.New("vote nonce not 0x00..0 or 0xff..f")

	// errInvalidCheckpointVote is returned if a checkpoint/epoch transition block
	// has a vote nonce set to non-zeroes.
	errInvalidCheckpointVote = errors.New("vote nonce in checkpoint block non-zero")

	// errMissingVanity is returned if a block's extra-data section is shorter than
	// 32 bytes, which is required to store the signer vanity.
	errMissingVanity = errors.New("extra-data 32 byte vanity prefix missing")

	// errMissingSignature is returned if a block's extra-data section doesn't seem
	// to contain a 65 byte secp256k1 signature.
	errMissingSignature = errors.New("extra-data 65 byte signature suffix missing")

	errMissingVrf = errors.New("extra-data 81 byte vrf suffix missing")

	// errExtraSigners is returned if non-checkpoint block contain signer data in
	// their extra-data fields.
	errExtraSigners = errors.New("non-checkpoint block contains extra signer list")

	// errInvalidCheckpointSigners is returned if a checkpoint block contains an
	// invalid list of signers (i.e. non divisible by 20 bytes).
	errInvalidCheckpointSigners = errors.New("invalid signer list on checkpoint block")

	// errMismatchingCheckpointSigners is returned if a checkpoint block contains a
	// list of signers different than the one the local node calculated.
	errMismatchingCheckpointSigners = errors.New("mismatching signer list on checkpoint block")

	// errInvalidMixDigest is returned if a block's mix digest is non-zero.
	errInvalidMixDigest = errors.New("non-zero mix digest")

	// errInvalidUncleHash is returned if a block contains an non-empty uncle list.
	errInvalidUncleHash = errors.New("non empty uncle hash")

	errInvalidVrfFn = errors.New("non empty vrf call")

	// errInvalidDifficulty is returned if the difficulty of a block neither 1 or 2.
	errInvalidDifficulty = errors.New("invalid difficulty")

	// errWrongDifficulty is returned if the difficulty of a block doesn't match the
	// turn of the signer.
	errWrongDifficulty = errors.New("wrong difficulty")

	// errInvalidTimestamp is returned if the timestamp of a block is lower than
	// the previous block's timestamp + the minimum block period.
	errInvalidTimestamp = errors.New("invalid timestamp")

	// errInvalidVotingChain is returned if an authorization list is attempted to
	// be modified via out-of-range or non-contiguous headers.
	errInvalidVotingChain = errors.New("invalid voting chain")

	// errUnauthorizedSigner is returned if a header is signed by a non-authorized entity.
	errUnauthorizedSigner = errors.New("unauthorized signer")

	errUnauthorizedProposer = errors.New("unauthorized proposer")

	// errRecentlySigned is returned if a header is signed by an authorized entity
	// that already signed a header recently, thus is temporarily not allowed to.
	errRecentlySigned = errors.New("recently signed")

	// errOutOfRangeChain is returned if an authorization list is attempted to
	// be modified via out-of-range or non-contiguous headers.
	errOutOfRangeChain = errors.New("out of range or non-contiguous chain")

	// errBlockHashInconsistent is returned if an authorization list is attempted to
	// insert an inconsistent block.
	errBlockHashInconsistent = errors.New("the block hash is inconsistent")

	// errUnauthorizedValidator is returned if a header is signed by a non-authorized entity.
	errUnauthorizedValidator = errors.New("unauthorized validator")

	// errCoinBaseMisMatch is returned if a header's coinbase do not match with signature
	//errCoinBaseMisMatch = errors.New("coinbase do not match with signature")

	// errRecentlySigned is returned if a header is signed by an authorized entity
	// that already signed a header recently, thus is temporarily not allowed to.
	//errRecentlySigned = errors.New("recently signed")
)

// SignerFn hashes and signs the data to be signed by a backing account.
type SignerFn func(signer accounts.Account, mimeType string, message []byte) ([]byte, error)
type SignerTxFn func(accounts.Account, *types.Transaction, *big.Int) (*types.Transaction, error)
type VrfProveFn func(alpha []byte) (beta, pi []byte, err error)

// ecrecover extracts the Ethereum account address from a signed header.
func ecrecover(header *types.Header, sigcache *lru.ARCCache, chainId *big.Int) ([]byte, common.Address, error) {
	// If the signature's already cached, return that
	hash := header.Hash()
	if data, known := sigcache.Get(hash); known {
		pubkey := data.([]byte)
		var addr common.Address
		copy(addr[:], crypto.Keccak256(pubkey[1:])[12:])
		return pubkey, addr, nil
	}
	// Retrieve the signature from the header extra-data
	if len(header.Extra) < extraSeal {
		return nil, common.Address{}, errMissingSignature
	}
	signature := header.Extra[len(header.Extra)-extraSeal:]

	// Recover the public key and the Ethereum address
	pubkey, err := oqs_crypto.Ecrecover(SealHash(header, chainId).Bytes(), signature)
	if err != nil {
		return nil, common.Address{}, err
	}
	var addr common.Address
	copy(addr[:], crypto.Keccak256(pubkey[1:])[12:])

	sigcache.Add(hash, pubkey)
	return pubkey, addr, nil
}

// Spos is the proof-of-authority consensus engine proposed to support the
// Ethereum testnet following the Ropsten attacks.
type Poseidon struct {
	chainConfig *params.ChainConfig    // Chain config
	config      *params.PoseidonConfig // Consensus engine configuration parameters
	genesisHash common.Hash
	db          ethdb.Database // Database to store and retrieve snapshot checkpoints

	beatcache  *lru.Cache
	signatures *lru.ARCCache // Signatures of recent blocks to speed up mining

	vrfFn    VrfProveFn
	signer   types.Signer
	val      common.Address // Ethereum address of the signing key
	signFn   SignerFn       // Signer function to authorize hashes with
	signTxFn SignerTxFn
	lock     sync.RWMutex // Protects the signer fields

	ethAPI    *ethapi.PublicBlockChainAPI
	txPoolAPI *ethapi.PublicTransactionPoolAPI

	validatorSetABI abi.ABI
}

// New creates a Spos proof-of-authority consensus engine with the initial
// signers set to the ones provided by the user.
func New(
	chainConfig *params.ChainConfig,
	db ethdb.Database,
	ethAPI *ethapi.PublicBlockChainAPI,
	genesisHash common.Hash,
) *Poseidon {
	// Set any missing consensus parameters to their defaults
	poseidonConfig := chainConfig.Poseidon

	// Allocate the snapshot caches and create the engine
	signatures, _ := lru.NewARC(inmemorySignatures)

	vABI, err := abi.JSON(strings.NewReader(validatorSetABI))
	if err != nil {
		panic(err)
	}
	beatCache, err := lru.New(1)
	if err != nil {
		panic(err)
	}

	return &Poseidon{
		chainConfig:     chainConfig,
		config:          poseidonConfig,
		genesisHash:     genesisHash,
		db:              db,
		ethAPI:          ethAPI,
		signatures:      signatures,
		validatorSetABI: vABI,
		signer:          types.NewEIP155Signer(chainConfig.ChainID),
		beatcache:       beatCache,
	}
}

func (p *Poseidon) SetTxPoolAPI(txPoolAPI *ethapi.PublicTransactionPoolAPI) {
	p.txPoolAPI = txPoolAPI
}

// Author implements consensus.Engine, returning the Ethereum address recovered
// from the signature in the header's extra-data section.
func (c *Poseidon) Author(header *types.Header) (common.Address, error) {
	_, signer, err := ecrecover(header, c.signatures, c.chainConfig.ChainID)
	return signer, err
}

// VerifyHeader checks whether a header conforms to the consensus rules.
func (c *Poseidon) VerifyHeader(chain consensus.ChainHeaderReader, header *types.Header, seal bool) error {
	return c.verifyHeader(chain, header, nil, seal)
}

// VerifyHeaders is similar to VerifyHeader, but verifies a batch of headers. The
// method returns a quit channel to abort the operations and a results channel to
// retrieve the async verifications (the order is that of the input slice).
func (c *Poseidon) VerifyHeaders(chain consensus.ChainHeaderReader, headers []*types.Header, seals []bool) (chan<- struct{}, <-chan error) {
	abort := make(chan struct{})
	results := make(chan error, len(headers))

	go func() {
		for i, header := range headers {
			err := c.verifyHeader(chain, header, headers[:i], seals[i])

			select {
			case <-abort:
				return
			case results <- err:
			}
		}
	}()
	return abort, results
}

// verifyHeader checks whether a header conforms to the consensus rules.The
// caller may optionally pass in a batch of parents (ascending order) to avoid
// looking those up from the database. This is useful for concurrently verifying
// a batch of new headers.
func (c *Poseidon) verifyHeader(chain consensus.ChainHeaderReader, header *types.Header, parents []*types.Header, seal bool) error {
	if header.Number == nil {
		return errUnknownBlock
	}
	number := header.Number.Uint64()

	// Don't waste time checking blocks from the future
	if header.Time > uint64(time.Now().Unix()) {
		return consensus.ErrFutureBlock
	}

	// Check that the extra-data contains both the vanity and signature
	if len(header.Extra) < extraVanity {
		return errMissingVanity
	}
	if len(header.Extra) < extraVanity+extraSeal {
		return errMissingSignature
	}
	if len(header.Extra) < extraVanity+extraSeal+extraVrf {
		return errMissingVrf
	}
	// Ensure that the extra-data contains a signer list on checkpoint, but none otherwise
	//signersBytes := len(header.Extra) - extraVanity - extraSeal

	// Ensure that the mix digest is zero as we don't have fork protection currently
	if header.MixDigest != (common.Hash{}) {
		return errInvalidMixDigest
	}
	// Ensure that the block doesn't contain any uncles which are meaningless in PoA
	if header.UncleHash != uncleHash {
		return errInvalidUncleHash
	}
	// Ensure that the block's difficulty is meaningful (may not be correct at this point)
	if number > 0 {
		if header.Difficulty == nil {
			return errInvalidDifficulty
		}
	}
	// Verify that the gas limit is <= 2^63-1
	cap := uint64(0x7fffffffffffffff)
	if header.GasLimit > cap {
		return fmt.Errorf("invalid gasLimit: have %v, max %v", header.GasLimit, cap)
	}
	// If all checks passed, validate any special fields for hard forks
	if err := misc.VerifyForkHashes(chain.Config(), header, false); err != nil {
		return err
	}
	// All basic checks passed, verify cascading fields
	return c.verifyCascadingFields(chain, header, parents, seal)
}

// verifyCascadingFields verifies all the header fields that are not standalone,
// rather depend on a batch of previous headers. The caller may optionally pass
// in a batch of parents (ascending order) to avoid looking those up from the
// database. This is useful for concurrently verifying a batch of new headers.
func (c *Poseidon) verifyCascadingFields(chain consensus.ChainHeaderReader, header *types.Header, parents []*types.Header, seal bool) error {
	// The genesis block is the always valid dead-end
	number := header.Number.Uint64()
	if number == 0 {
		return nil
	}
	// Ensure that the block's timestamp isn't too close to its parent
	var parent *types.Header
	if len(parents) > 0 {
		parent = parents[len(parents)-1]
	} else {
		parent = chain.GetHeader(header.ParentHash, number-1)
	}
	if parent == nil || parent.Number.Uint64() != number-1 || parent.Hash() != header.ParentHash {
		return consensus.ErrUnknownAncestor
	}
	if parent.Time+c.config.Period > header.Time {
		return errInvalidTimestamp
	}
	// Verify that the gasUsed is <= gasLimit
	if header.GasUsed > header.GasLimit {
		return fmt.Errorf("invalid gasUsed: have %d, gasLimit %d", header.GasUsed, header.GasLimit)
	}
	if !chain.Config().IsLondon(header.Number) {
		// Verify BaseFee not present before EIP-1559 fork.
		if header.BaseFee != nil {
			return fmt.Errorf("invalid baseFee before fork: have %d, want <nil>", header.BaseFee)
		}
		if err := misc.VerifyGaslimit(parent.GasLimit, header.GasLimit); err != nil {
			return err
		}
	} else if err := misc.VerifyEip1559Header(chain.Config(), parent, header); err != nil {
		// Verify the header's EIP-1559 attributes.
		return err
	}

	// All basic checks passed, verify the seal and return
	if seal == true {
		if err := c.verifySeal(chain, header, parents); err != nil {
			log.Warn("Poseidon verifySeal fail", "number", header.Number, "err", err)
			return err
		}
	}
	return nil
}

// VerifyUncles implements consensus.Engine, always returning an error for any
// uncles as this consensus mechanism doesn't permit uncles.
func (c *Poseidon) VerifyUncles(chain consensus.ChainReader, block *types.Block) error {
	if len(block.Uncles()) > 0 {
		return errors.New("uncles not allowed")
	}
	return nil
}

// verifySeal checks whether the signature contained in the header satisfies the
// consensus protocol requirements. The method accepts an optional list of parent
// headers that aren't yet part of the local blockchain to generate the snapshots
// from.
func (c *Poseidon) verifySeal(chain consensus.ChainHeaderReader, header *types.Header, parents []*types.Header) error {
	// Verifying the genesis block is not supported
	number := header.Number.Uint64()
	if number == 0 {
		return errUnknownBlock
	}

	// Resolve the authorization key and check against signers
	pubkey, signer, err := ecrecover(header, c.signatures, c.chainConfig.ChainID)
	if err != nil {
		return err
	}
	if isProposer, err := c.IsProposer(signer, header.Number); err != nil || isProposer == false {
		return errUnauthorizedProposer
	}
	info, err := c.GetValidatorInfo(signer, header.Number)
	if err != nil {
		return err
	}
	committeeSupply, err := c.GetCommitteeSupply(header.Number, signer)
	if err != nil {
		return err
	}
	pi := make([]byte, extraVrf)
	copy(pi, header.Extra[len(header.Extra)-extraSeal-extraVrf:len(header.Extra)-extraSeal])

	alpha := c.GetVrfAlpha(header.ParentHash, header.Nonce)
	publicKey, err := crypto.UnmarshalPubkey(pubkey)
	if err != nil {
		return err
	}
	_ = publicKey
	beta, err := vrf.Verify(nil, alpha, pi)
	if err != nil {
		return err
	}
	if err := c.checkDifficulty(chain, header, info, beta); err != nil {
		return err
	}
	if c.verifySort(info.TotalSupply, committeeSupply, header.Number, beta) == false {
		return errUnauthorizedSigner
	}
	return nil
}

// Prepare implements consensus.Engine, preparing all the consensus fields of the
// header for running the transactions on top.
func (c *Poseidon) Prepare(chain consensus.ChainHeaderReader, header *types.Header) error {
	// If the block isn't a checkpoint, cast a random vote (good enough for now)
	header.Nonce = types.BlockNonce{}

	number := header.Number.Uint64()

	// Ensure the extra data has all its components
	if len(header.Extra) < extraVanity {
		header.Extra = append(header.Extra, bytes.Repeat([]byte{0x00}, extraVanity-len(header.Extra))...)
	}

	header.Extra = header.Extra[:extraVanity]

	// Mix digest is reserved for now, set to empty
	header.MixDigest = common.Hash{}

	// Ensure the timestamp has the correct delay
	parent := chain.GetHeader(header.ParentHash, number-1)
	if parent == nil {
		return consensus.ErrUnknownAncestor
	}
	header.Time = parent.Time + c.config.Period
	if header.Time < uint64(time.Now().Unix()) {
		header.Time = uint64(time.Now().Unix())
	}

	info, err := c.GetValidatorInfo(c.val, header.Number)
	if err != nil {
		return err
	}
	header.Coinbase = info.RewardAddr

	header.Extra = append(header.Extra, make([]byte, extraVrf+extraSeal)...)
	header.Difficulty = common.Big0
	return nil
}

func (c *Poseidon) verifySort(money *big.Int, totalMoney *big.Int, blockNumber *big.Int, vrfOutput []byte) bool {
	expectedSize := vrfExpectedSize
	if money.Cmp(totalMoney) >= 0 {
		expectedSize = 1
	}
	return vrf.VerifySort(new(big.Int).Div(money, ether).Uint64(), new(big.Int).Div(totalMoney, ether).Uint64(), expectedSize, vrfOutput)
}

// Finalize implements consensus.Engine, ensuring no uncles are set, nor block
// rewards given.
func (c *Poseidon) Finalize(chain consensus.ChainHeaderReader, header *types.Header, state *state.StateDB, txs []*types.Transaction,
	uncles []*types.Header) {
	header.Root = state.IntermediateRoot(chain.Config().IsEIP158(header.Number))
	header.UncleHash = types.CalcUncleHash(nil)
}

// FinalizeAndAssemble implements consensus.Engine, ensuring no uncles are set,
// nor block rewards given, and returns the final block.
func (c *Poseidon) FinalizeAndAssemble(chain consensus.ChainHeaderReader, header *types.Header, state *state.StateDB,
	txs []*types.Transaction, uncles []*types.Header, receipts []*types.Receipt) (*types.Block, error) {
	if txs == nil {
		txs = make([]*types.Transaction, 0)
	}
	if receipts == nil {
		receipts = make([]*types.Receipt, 0)
	}

	// should not happen. Once happen, stop the node is better than broadcast the block
	if header.GasLimit < header.GasUsed {
		panic("Gas consumption of system txs exceed the gas limit")
	}
	header.Root = state.IntermediateRoot(chain.Config().IsEIP158(header.Number))
	header.UncleHash = types.CalcUncleHash(nil)

	// Assemble and return the final block for sealing
	return types.NewBlock(header, txs, nil, receipts, trie.NewStackTrie(nil)), nil
}

// Authorize injects a private key into the consensus engine to mint new blocks
// with.
func (c *Poseidon) Authorize(val common.Address, signFn SignerFn, signTxFn SignerTxFn, vrfFn VrfProveFn) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.val = val
	c.signFn = signFn
	c.vrfFn = vrfFn
	c.signTxFn = signTxFn
}

func (c *Poseidon) GetVrfAlpha(parentHash common.Hash, nonce types.BlockNonce) []byte {
	alpha := make([]byte, len(parentHash))
	copy(alpha, parentHash[:])
	alpha = append(alpha, nonce[:]...)
	return alpha
}

func (c *Poseidon) sortition(chain consensus.ChainHeaderReader, header *types.Header, info *ValidatorInfo, committeeSupply *big.Int, signer common.Address, signFn SignerFn) (bool, error) {
	alpha := c.GetVrfAlpha(header.ParentHash, header.Nonce)
	beta, pi, err := c.vrfFn(alpha)
	if err != nil {
		return false, err
	}
	copy(header.Extra[extraVanity:], pi)

	if c.verifySort(info.TotalSupply, committeeSupply, header.Number, beta) == false {
		return false, nil
	}
	// Set the correct difficulty
	header.Difficulty = calcDifficulty(header.Nonce, header.Number, info.TotalSupply, info.LastBlockHeight, beta)

	// Sign all the things!
	sighash, err := signFn(accounts.Account{Address: signer}, accounts.MimetypePoseidon, PoseidonRLP(header, c.chainConfig.ChainID))
	if err != nil {
		return false, err
	}
	copy(header.Extra[len(header.Extra)-extraSeal:], sighash)
	return true, nil
}

// Seal implements consensus.Engine, attempting to create a sealed block using
// the local signing credentials.
func (c *Poseidon) Seal(chain consensus.ChainHeaderReader, block *types.Block, results chan<- *types.Block, stop <-chan struct{}) error {
	if c.vrfFn == nil {
		return errInvalidVrfFn
	}
	header := block.Header()
	if isProposer, err := c.IsProposer(c.val, header.Number); err != nil || isProposer == false {
		return errUnauthorizedProposer
	}

	info, err := c.GetValidatorInfo(c.val, header.Number)
	if err != nil {
		return err
	}
	committeeSupply, err := c.GetCommitteeSupply(header.Number, c.val)
	if err != nil {
		return err
	}

	// Sealing the genesis block is not supported
	number := header.Number.Uint64()
	if number == 0 {
		return errUnknownBlock
	}
	// For 0-period chains, refuse to seal empty blocks (no reward but would spin sealing)
	if c.config.Period == 0 && len(block.Transactions()) == 0 {
		log.Info("Sealing paused, waiting for transactions")
		return nil
	}

	// Don't hold the signer fields for the entire sealing procedure
	c.lock.RLock()
	signer, signFn := c.val, c.signFn
	c.lock.RUnlock()

	// Sweet, the protocol permits us to sign the block, wait for our time
	delay := time.Unix(int64(header.Time), 0).Sub(time.Now()) // nolint: gosimple

	isSeal, err := c.sortition(chain, header, info, committeeSupply, signer, signFn)
	if err != nil {
		return err
	}
	// Wait until sealing is terminated or delay timeout.
	log.Trace("Waiting for slot to sign and propagate", "delay", common.PrettyDuration(delay))
	go func() {
		vrfTimer := time.NewTicker(time.Duration(c.config.Period/2) * time.Second)
		defer vrfTimer.Stop()
		afterTimer := time.NewTimer(delay)
		defer afterTimer.Stop()

		for {
			select {
			case <-stop:
				return
			case <-afterTimer.C:
				if isSeal {
					goto sealLabel
				}
				afterTimer.Reset(1 * time.Second)
			case <-vrfTimer.C:
				if isSeal {
					continue
				}
				header.Nonce = types.EncodeNonce(header.Nonce.Uint64() + 1)
				isSeal, err = c.sortition(chain, header, info, committeeSupply, signer, signFn)
				if err != nil {
					log.Warn("Block sealExtra failed", "err", err)
					return
				}
			}
		}
	sealLabel:
		select {
		case results <- block.WithSeal(header):
		default:
			log.Warn("Sealing result is not read by miner", "sealhash", SealHash(header, c.chainConfig.ChainID))
		}
	}()

	return nil
}

// CalcDifficulty is the difficulty adjustment algorithm. It returns the difficulty
// that a new block should have:
// * DIFF_NOTURN(2) if BLOCK_NUMBER % SIGNER_COUNT != SIGNER_INDEX
// * DIFF_INTURN(1) if BLOCK_NUMBER % SIGNER_COUNT == SIGNER_INDEX
func (c *Poseidon) CalcDifficulty(chain consensus.ChainHeaderReader, time uint64, parent *types.Header) *big.Int {
	header := chain.CurrentHeader()
	nonce := types.EncodeNonce(0)
	if header != nil {
		nonce = header.Nonce
	}
	info, err := c.GetValidatorInfo(c.val, header.Number)
	if err != nil {
		info = &ValidatorInfo{
			LastBlockHeight: big.NewInt(0),
			TotalSupply:     big.NewInt(0),
		}
	}
	return calcDifficulty(nonce, header.Number, info.TotalSupply, info.LastBlockHeight, make([]byte, 32))
}

func (c *Poseidon) checkDifficulty(chain consensus.ChainHeaderReader, header *types.Header, info *ValidatorInfo, beta []byte) error {
	diff := calcDifficulty(header.Nonce, header.Number, info.TotalSupply, info.LastBlockHeight, beta)
	if diff.Cmp(header.Difficulty) != 0 {
		return errInvalidDifficulty
	}
	return nil
}

func calcDifficulty(
	blockNonce types.BlockNonce,
	blockNumber *big.Int,
	totalSupply *big.Int,
	lastBlockHeight *big.Int,
	beta []byte,
) *big.Int {
	nonce := big.NewInt(0) //uint8
	if blockNonce.Uint64() < nonceSignSize {
		nonce = nonce.SetUint64(nonceSignSize - blockNonce.Uint64())
	}
	custom := big.NewInt(0)                                       //uint8
	diffNumber := big.NewInt(0).Sub(blockNumber, lastBlockHeight) //uint32
	if diffNumber.Cmp(big.NewInt(0)) < 0 {
		diffNumber = diffNumber.SetInt64(0)
	} else {
		diffNumber = diffNumber.Div(diffNumber, big.NewInt(256))
	}
	randBeta := big.NewInt(0).SetBytes(beta[:6]) //uint48
	diff := big.NewInt(0)
	diff = diff.Or(diff, nonce.Lsh(nonce, 88)).
		Or(diff, custom.Lsh(custom, 80)).
		Or(diff, diffNumber.Lsh(diffNumber, 48)).
		Or(diff, randBeta)

	return diff
}

// SealHash returns the hash of a block prior to it being sealed.
func (c *Poseidon) SealHash(header *types.Header) common.Hash {
	return SealHash(header, c.chainConfig.ChainID)
}

// Close implements consensus.Engine. It's a noop for clique as there are no background threads.
func (c *Poseidon) Close() error {
	return nil
}

// APIs implements consensus.Engine, returning the user facing RPC API to allow
// controlling the signer voting.
func (c *Poseidon) APIs(chain consensus.ChainHeaderReader) []rpc.API {
	return []rpc.API{{
		Namespace: "poseidon",
		Version:   "1.0",
		Service:   &API{chain: chain, poseidon: c},
		Public:    false,
	}}
}

// SealHash returns the hash of a block prior to it being sealed.
func SealHash(header *types.Header, chainId *big.Int) (hash common.Hash) {
	hasher := sha3.NewLegacyKeccak256()
	encodeSigHeader(hasher, header, chainId)
	hasher.Sum(hash[:0])
	return hash
}

// PoseidonRLP returns the rlp bytes which needs to be signed for the proof-of-authority
// sealing. The RLP to sign consists of the entire header apart from the 65 byte signature
// contained at the end of the extra data.
//
// Note, the method requires the extra data to be at least 65 bytes, otherwise it
// panics. This is done to avoid accidentally using both forms (signature present
// or not), which could be abused to produce different hashes for the same header.
func PoseidonRLP(header *types.Header, chainId *big.Int) []byte {
	b := new(bytes.Buffer)
	encodeSigHeader(b, header, chainId)
	return b.Bytes()
}

func encodeSigHeader(w io.Writer, header *types.Header, chainId *big.Int) {
	enc := []interface{}{
		chainId,
		header.ParentHash,
		header.UncleHash,
		header.Coinbase,
		header.Root,
		header.TxHash,
		header.ReceiptHash,
		header.Bloom,
		//header.Difficulty,
		header.Number,
		header.GasLimit,
		header.GasUsed,
		header.Time,
		header.Extra[:len(header.Extra)-crypto.SignatureLength-extraVrf], // Yes, this will panic if extra is too short
		header.MixDigest,
		//header.Nonce,
	}
	if header.BaseFee != nil {
		enc = append(enc, header.BaseFee)
	}
	if err := rlp.Encode(w, enc); err != nil {
		panic("can't encode: " + err.Error())
	}
}

func (c *Poseidon) Heartbeat(number *big.Int) error {
	currentHeight := number.Uint64()

	if value, ok := c.beatcache.Peek(c.val); ok {
		if cacheHeight, ok := value.(*big.Int); ok {
			subResult := big.NewInt(0)
			subResult.Sub(number, cacheHeight).Abs(subResult)

			if subResult.Cmp(common.Big3) <= 0 {
				return nil
			}
		}
	}

	info, err := c.GetValidatorInfo(c.val, number)
	if err != nil {
		return err
	}
	lastBlockHeight := info.LastBlockHeight.Uint64()

	if (currentHeight < lastBlockHeight) || (currentHeight-lastBlockHeight) < heartRate {
		return nil
	}

	// method
	method := "slash"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers

	//ctx := context.WithValue(context.Background(), "", "")

	data, err := c.validatorSetABI.Pack(method, c.val)
	if err != nil {
		log.Error("Unable to pack tx for slash", "error", err)
		return err
	}

	// call
	msgData := (hexutil.Bytes)(data)
	toAddress := common.HexToAddress(systemcontracts.ValidatorHubContract)
	gas := (hexutil.Uint64)(uint64(100000))

	//gasPrice := (*hexutil.Big)(big.NewInt(0))
	_, err = c.txPoolAPI.SendTransaction(ctx, ethapi.TransactionArgs{From: &c.val, To: &toAddress, Data: &msgData, Gas: &gas})
	if err != nil {
		return err
	}

	c.beatcache.Add(c.val, number)

	return nil
}

// totalFees computes total consumed miner fees in ETH. Block transactions and receipts have to have the same order.
func totalFees(header *types.Header, txs []*types.Transaction, receipts []*types.Receipt) *big.Int {
	feesWei := new(big.Int)
	for i, tx := range txs {
		minerFee, _ := tx.EffectiveGasTip(header.BaseFee)
		feesWei.Add(feesWei, new(big.Int).Mul(new(big.Int).SetUint64(receipts[i].GasUsed), minerFee))
	}
	return feesWei
}

// chain context
type chainContext struct {
	Chain    consensus.ChainHeaderReader
	poseidon consensus.Engine
}

func (c chainContext) Engine() consensus.Engine {
	return c.poseidon
}

func (c chainContext) GetHeader(hash common.Hash, number uint64) *types.Header {
	return c.Chain.GetHeader(hash, number)
}

func (p *Poseidon) GetSystemTransaction(signer types.Signer, state *state.StateDB, baseFee *big.Int, totalFee *big.Int) (*types.TransactionsByPriceAndNonce, error) {
	nonce := state.GetNonce(p.val)

	method := "syncTendermintHeader"
	// get packed data
	data, err := p.validatorSetABI.Pack(method,
		totalFee,
	)
	if err != nil {
		return nil, err
	}
	gasPrice := big.NewInt(0)
	if baseFee != nil {
		gasPrice = gasPrice.Set(baseFee)
	}
	tx := types.NewTransaction(nonce, common.HexToAddress(systemcontracts.ValidatorHubContract), common.Big0, 200000, gasPrice, data)
	//signtx
	expectedTx, err := p.signTxFn(accounts.Account{Address: p.val}, tx, p.chainConfig.ChainID)
	if err != nil {
		return nil, err
	}

	txs := make(map[common.Address]types.Transactions)
	txs[p.val] = types.Transactions{expectedTx}

	return types.NewTransactionsByPriceAndNonce(signer, txs, baseFee), nil
}
