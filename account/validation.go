package account

import (
	"reflect"

	"github.com/ethereum/go-ethereum/common"
)

// IsZeroAddress validate if it's a 0 address
func IsZeroAddress(iAddress any) bool {
	var address common.Address
	switch v := iAddress.(type) {
	case string:
		address = common.HexToAddress(v)
	case common.Address:
		address = v
	default:
		return false
	}

	addressBytes := address.Bytes()
	return reflect.DeepEqual(addressBytes, ZeroAddress.Bytes())
}

// IsValidAddress validate hex address
func IsValidAddress(iAddress any) bool {
	switch v := iAddress.(type) {
	case string:
		return EthAddressRegex.MatchString(v)
	case common.Address:
		return EthAddressRegex.MatchString(v.Hex())
	default:
		return false
	}
}
