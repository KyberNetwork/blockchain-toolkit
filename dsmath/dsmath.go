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

// ToWad
// Convert x to WAD (18 decimals) from d decimals.
func ToWAD(x *big.Int, d uint8) *big.Int {
	if d < 18 {
		return new(big.Int).Mul(x, integer.TenPow(18-d))
	}

	if d > 18 {
		return new(big.Int).Div(x, integer.TenPow(d-18))
	}

	return new(big.Int).Set(x)
}

// FromWAD
// Convert x from WAD (18 decimals) to d decimals.
func FromWAD(x *big.Int, d uint8) *big.Int {
	if d < 18 {
		return new(big.Int).Div(x, integer.TenPow(18-d))
	}

	if d > 18 {
		return new(big.Int).Mul(x, integer.TenPow(d-18))
	}

	return new(big.Int).Set(x)
}
