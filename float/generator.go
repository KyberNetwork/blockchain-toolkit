package float

import "math/big"

// TenPow returns value of 10^decimal
func TenPow(decimals uint8) *big.Float {
	return new(big.Float).SetInt(
		new(big.Int).Exp(
			big.NewInt(10),
			big.NewInt(int64(decimals)),
			nil,
		),
	)
}

// Numbers

func Zero() *big.Float {
	return new(big.Float).Set(zero)
}

func One() *big.Float {
	return new(big.Float).Set(one)
}

func Two() *big.Float {
	return new(big.Float).Set(two)
}

func Three() *big.Float {
	return new(big.Float).Set(three)
}

func Four() *big.Float {
	return new(big.Float).Set(four)
}

func Five() *big.Float {
	return new(big.Float).Set(five)
}
