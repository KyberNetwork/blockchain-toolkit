package unit

import (
	"math/big"

	"github.com/KyberNetwork/blockchain-toolkit/float"
	"github.com/KyberNetwork/blockchain-toolkit/integer"
)

// ToDecimal wei to decimals
func ToDecimal(value *big.Int, decimals uint8) *big.Float {
	return new(big.Float).Quo(
		new(big.Float).SetInt(value),
		float.TenPow(decimals),
	)
}

// ToWei decimals to wei
func ToWei(value *big.Int, decimals uint8) *big.Int {
	return new(big.Int).Mul(value, integer.TenPow(decimals))
}
