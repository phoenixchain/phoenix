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

func (api *API) GetValidatorInfo(validatorAddr common.Address, blockNumber *rpc.BlockNumber) (*ValidatorInfo, error) {
	info, err := api.poseidon.GetValidatorInfo(validatorAddr, big.NewInt(blockNumber.Int64()))
	if err != nil {
		return nil, err
	}

	return info, nil
}

func (api *API) GetCommitteeSupply(blockNumber *big.Int) (*big.Int, error) {
	return api.poseidon.GetCommitteeSupply(blockNumber)
}
