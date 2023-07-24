package currency

import (
	"math/big"

	"github.com/KyberNetwork/blockchain-toolkit/float"
)

// CalcAmountUSD returns amount in USD
// amountUSD = (amount / 10^decimals) * priceUSD
func CalcAmountUSD(amount *big.Int, decimals uint8, priceUSD float64) *big.Float {
	result := new(big.Float).SetInt(amount)

	result = result.Quo(result, float.TenPow(decimals))
	result = result.Mul(result, new(big.Float).SetFloat64(priceUSD))

	return result
}

// CalcGasCost calculate gas cost given gas limit (units) and gas price (wei)
func CalcGasCost(gasLimit uint64, gasPrice *big.Int) *big.Int {
	result := big.NewInt(int64(gasLimit))
	result = result.Mul(result, gasPrice)

	return result
}
