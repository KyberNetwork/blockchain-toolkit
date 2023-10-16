package dsmath

import (
	"math/big"

	"github.com/KyberNetwork/blockchain-toolkit/integer"
)

var (
	WAD *big.Int = integer.TenPow(18)
)

// WMul
// ((x * y) + (WAD / 2)) / WAD;
func WMul(x *big.Int, y *big.Int) *big.Int {
	return new(big.Int).Div(
		new(big.Int).Add(
			new(big.Int).Mul(x, y),
			new(big.Int).Div(WAD, integer.Two()),
		),
		WAD,
	)
}

// WDiv
// ((x * WAD) + (y / 2)) / y
func WDiv(x *big.Int, y *big.Int) *big.Int {
	return new(big.Int).Div(
		new(big.Int).Add(
			new(big.Int).Mul(x, WAD),
			new(big.Int).Div(y, integer.Two()),
		),
		y,
	)
}
