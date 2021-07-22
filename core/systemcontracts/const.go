package systemcontracts

import (
	"bytes"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

const (
	// genesis contracts
	//ValidatorContract          = "0x0000000000000000000000000000000000001000"
	//SlashContract              = "0x0000000000000000000000000000000000001001"
	//SystemRewardContract       = "0x0000000000000000000000000000000000001002"
	//LightClientContract        = "0x0000000000000000000000000000000000001003"
	//TokenHubContract           = "0x0000000000000000000000000000000000001004"
	//RelayerIncentivizeContract = "0x0000000000000000000000000000000000001005"
	//RelayerHubContract         = "0x0000000000000000000000000000000000001006"
	//GovHubContract             = "0x0000000000000000000000000000000000001007"
	//TokenManagerContract       = "0x0000000000000000000000000000000000001008"
	//CrossChainContract         = "0x0000000000000000000000000000000000002000"

	ValidatorFactoryContract = "0x0000000000000000000000000000000000001009"
	ValidatorHubContract     = "0x0000000000000000000000000000000000001006"
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
