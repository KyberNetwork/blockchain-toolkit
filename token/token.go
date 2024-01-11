package token

import "github.com/ethereum/go-ethereum/common"

type Token struct {
	ChainID uint

	Address  common.Address
	Decimals uint8
	Name     string
	Symbol   string
}
