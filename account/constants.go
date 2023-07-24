package account

import (
	"regexp"

	"github.com/ethereum/go-ethereum/common"
)

const (
	ethAddressRegexString = `^0x[0-9a-fA-F]{40}$`
)

var (
	ZeroAddress = common.HexToAddress("0x0000000000000000000000000000000000000000")

	EthAddressRegex = regexp.MustCompile(ethAddressRegexString)
)
