package openzeppelin

import "math/big"

/**
 * This formula follows implementation from openzeppelin
 * https://github.com/OpenZeppelin/openzeppelin-contracts/blob/94697be8a3f0dfcd95dfb13ffbd39b5973f5c65d/contracts/utils/math/Math.sol#L221
 **/
func Sqrt(value *big.Int) *big.Int {
	if value.Cmp(big.NewInt(0)) == 0 {
		return value
	}
	result := new(big.Int).Lsh(big.NewInt(1), log2(value)/2)
	tmp := new(big.Int)
	result.Rsh(tmp.Div(value, result).Add(tmp, result), 1)
	result.Rsh(tmp.Div(value, result).Add(result, tmp), 1)
	result.Rsh(tmp.Div(value, result).Add(result, tmp), 1)
	result.Rsh(tmp.Div(value, result).Add(result, tmp), 1)
	result.Rsh(tmp.Div(value, result).Add(result, tmp), 1)
	result.Rsh(tmp.Div(value, result).Add(result, tmp), 1)
	result.Rsh(tmp.Div(value, result).Add(result, tmp), 1)
	tmp = new(big.Int).Div(value, result)
	if result.Cmp(tmp) == -1 {
		return result
	}
	return tmp
}

func log2(value *big.Int) uint {
	var result uint = 0
	zero := big.NewInt(0)
	comparator := new(big.Int)
	tempValue := new(big.Int).Set(value)
	if comparator.Rsh(tempValue, 128).Cmp(zero) == 1 {
		tempValue.Set(comparator)
		result += 128
	}
	if comparator.Rsh(tempValue, 64).Cmp(zero) == 1 {
		tempValue.Set(comparator)
		result += 64
	}
	if comparator.Rsh(tempValue, 32).Cmp(zero) == 1 {
		tempValue.Set(comparator)
		result += 32
	}
	if comparator.Rsh(tempValue, 16).Cmp(zero) == 1 {
		tempValue.Set(comparator)
		result += 16
	}
	if comparator.Rsh(tempValue, 8).Cmp(zero) == 1 {
		tempValue.Set(comparator)
		result += 8
	}
	if comparator.Rsh(tempValue, 4).Cmp(zero) == 1 {
		tempValue.Set(comparator)
		result += 4
	}
	if comparator.Rsh(tempValue, 2).Cmp(zero) == 1 {
		tempValue.Set(comparator)
		result += 2
	}
	if comparator.Rsh(tempValue, 1).Cmp(zero) == 1 {
		result += 1
	}

	return result
}
