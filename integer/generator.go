package integer

import "math/big"

// NewBig10 returns a big int from a decimal string number
func NewBig10(s string) (res *big.Int) {
	res, _ = new(big.Int).SetString(s, 10)
	return res
}

// TenPow returns value of 10^decimal
func TenPow(decimals uint8) *big.Int {
	return new(big.Int).Exp(
		big.NewInt(10),
		big.NewInt(int64(decimals)),
		nil,
	)
}

// Numbers

func Zero() *big.Int {
	return new(big.Int).Set(zero)
}

func One() *big.Int {
	return new(big.Int).Set(one)
}

func Two() *big.Int {
	return new(big.Int).Set(two)
}

func Three() *big.Int {
	return new(big.Int).Set(three)
}

func Four() *big.Int {
	return new(big.Int).Set(four)
}

func Five() *big.Int {
	return new(big.Int).Set(five)
}
