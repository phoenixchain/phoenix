package poseidon

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
)

// API is a user facing RPC API to allow controlling the signer and voting
// mechanisms of the proof-of-authority scheme.
type API struct {
	chain    consensus.ChainHeaderReader
	poseidon *Poseidon
}

// GetSigners retrieves the list of authorized signers at the specified block.
func (api *API) GetSigners(number *rpc.BlockNumber) ([]common.Address, error) {
	//TODO:
	return nil, nil
}

func (api *API) IsValidator(validatorAddr common.Address, blockNumber *big.Int) (bool, error) {
	return api.poseidon.IsValidator(validatorAddr, blockNumber)
}

func (api *API) GetValidatorInfo(validatorAddr common.Address, blockNumber *big.Int) (*ValidatorInfo, error) {
	return api.poseidon.GetValidatorInfo(validatorAddr, blockNumber)
}

func (api *API) GetCommitteeSupply( blockNumber *big.Int) (*big.Int, error) {
	return api.poseidon.GetCommitteeSupply(blockNumber,common.Address{})
}

func (api *API) IsProposer(validatorAddr common.Address, blockNumber *big.Int) (bool, error) {
	return api.poseidon.IsProposer(validatorAddr, blockNumber)
}

func (api *API) GetValidators(blockNumber *big.Int) ([]common.Address, error) {
	return api.poseidon.GetValidators(blockNumber)
}
