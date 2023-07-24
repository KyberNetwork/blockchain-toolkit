package integer

import "math/big"

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
