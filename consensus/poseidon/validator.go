package poseidon

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/systemcontracts"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
	"math"
	"math/big"
)

type ValidatorInfo struct {
	Name              string
	RewardAddr        common.Address
	TotalSupply       *big.Int
	PerProposerHeight *big.Int
}

// ==========================  interaction with contract/account =========

func (p *Poseidon) IsValidator(validator common.Address) (bool, error) {
	// method
	method := "isValidator"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers

	data, err := p.validatorSetABI.Pack(method, validator)
	if err != nil {
		log.Error("Unable to pack tx for isValidator", "error", err)
		return false, err
	}
	// call
	msgData := (hexutil.Bytes)(data)
	toAddress := common.HexToAddress(systemcontracts.ValidatorHubContract)
	gas := (hexutil.Uint64)(uint64(math.MaxUint64 / 2))
	result, err := p.ethAPI.Call(ctx, ethapi.TransactionArgs{
		Gas:  &gas,
		To:   &toAddress,
		Data: &msgData,
	}, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber), nil)
	if err != nil {
		return false, err
	}

	out := false

	if err := p.validatorSetABI.UnpackIntoInterface(&out, method, result); err != nil {
		return false, err
	}
	return out, nil
}

func (p *Poseidon) GetValidatorInfo(validator common.Address) (*ValidatorInfo, error) {
	// method
	method := "getValidatorInfo"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers

	data, err := p.validatorSetABI.Pack(method, validator)
	if err != nil {
		log.Error("Unable to pack tx for getValidatorInfo", "error", err)
		return nil, err
	}
	// call
	msgData := (hexutil.Bytes)(data)
	toAddress := common.HexToAddress(systemcontracts.ValidatorHubContract)
	gas := (hexutil.Uint64)(uint64(math.MaxUint64 / 2))
	result, err := p.ethAPI.Call(ctx, ethapi.TransactionArgs{
		Gas:  &gas,
		To:   &toAddress,
		Data: &msgData,
	}, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber), nil)
	if err != nil {
		return nil, err
	}

	out := new(ValidatorInfo)

	if err := p.validatorSetABI.UnpackIntoInterface(out, method, result); err != nil {
		return nil, err
	}
	return out, nil
}

func (p *Poseidon) IsProposer(validator common.Address) (bool, error) {
	// method
	method := "isProposer"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers

	data, err := p.validatorSetABI.Pack(method, validator)
	if err != nil {
		log.Error("Unable to pack tx for isValidator", "error", err)
		return false, err
	}
	// call
	msgData := (hexutil.Bytes)(data)
	toAddress := common.HexToAddress(systemcontracts.ValidatorHubContract)
	gas := (hexutil.Uint64)(uint64(math.MaxUint64 / 2))
	result, err := p.ethAPI.Call(ctx, ethapi.TransactionArgs{
		Gas:  &gas,
		To:   &toAddress,
		Data: &msgData,
	}, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber), nil)
	if err != nil {
		return false, err
	}

	out := false

	if err := p.validatorSetABI.UnpackIntoInterface(&out, method, result); err != nil {
		return false, err
	}
	return out, nil
}

func (p *Poseidon) GetCommitteeSupply() (*big.Int, error) {
	// method
	method := "getCommitteeSupply"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers

	data, err := p.validatorSetABI.Pack(method)
	if err != nil {
		log.Error("Unable to pack tx for getCommitteeSupply", "error", err)
		return nil, err
	}
	// call
	msgData := (hexutil.Bytes)(data)
	toAddress := common.HexToAddress(systemcontracts.ValidatorHubContract)
	gas := (hexutil.Uint64)(uint64(math.MaxUint64 / 2))
	result, err := p.ethAPI.Call(ctx, ethapi.TransactionArgs{
		Gas:  &gas,
		To:   &toAddress,
		Data: &msgData,
	}, rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber), nil)
	if err != nil {
		return nil, err
	}

	out := new(big.Int)

	if err := p.validatorSetABI.UnpackIntoInterface(out, method, result); err != nil {
		return nil, err
	}
	return out, nil
}
