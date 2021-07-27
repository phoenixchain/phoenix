package systemcontracts

import (
	"bytes"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

const (
	// genesis contracts
	ValidatorHubContract = "0x0000000000000000000000000000000000001006"
)

func isHubTransition(to *common.Address, data []byte) bool {
	if to == nil || data == nil || len(data) < 4 {
		return false
	}
	systemAddr := common.HexToAddress(ValidatorHubContract)
	if bytes.Compare(to[:], systemAddr[:]) != 0 {
		return false
	}
	return true
}

func IsSlashTransition(to *common.Address, data []byte) bool {
	if isHubTransition(to, data) == false {
		return false
	}
	if len(data) == 36 && hexutil.Encode(data[:4]) == "0xc96be4cb" { //slash(address)
		return true
	}
	return false
}

func IsSyncHeaderTransition(to *common.Address, data []byte) bool {
	if isHubTransition(to, data) == false {
		return false
	}
	if len(data) == 36 && hexutil.Encode(data[:4]) == "0xffd8136e" { //syncTendermintHeader(uint256)
		return true
	}
	return false
}

func IsSystemTransition(to *common.Address, data []byte) bool {
	return IsSlashTransition(to, data) || IsSyncHeaderTransition(to, data)
}
