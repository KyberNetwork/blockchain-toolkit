package number_test

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/KyberNetwork/blockchain-toolkit/number"
	"github.com/holiman/uint256"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWrapper(t *testing.T) {
	for i := 0; i < 200; i++ {
		xStr := number.RandNumberHexString(64)
		yStr := number.RandNumberHexString(64)
		var err error

		t.Run(fmt.Sprintf("test %s %s", xStr, yStr), func(t *testing.T) {
			xBI, ok := new(big.Int).SetString(xStr, 16)
			require.True(t, ok)
			yBI, ok := new(big.Int).SetString(yStr, 16)
			require.True(t, ok)

			// known correct values to compare against
			var x1, y1 uint256.Int
			err = x1.SetFromHex("0x" + xStr)
			require.Nil(t, err)
			err = y1.SetFromHex("0x" + yStr)
			require.Nil(t, err)

			x2 := number.SetFromBig(xBI)
			y2 := number.SetFromBig(yBI)
			assert.Equal(t, x1, *x2)
			assert.Equal(t, y1, *y2)

			x3 := number.Set(&x1)
			y3 := number.Set(&y1)
			assert.Equal(t, x1, *x3)
			assert.Equal(t, y1, *y3)

			assert.Equal(t, new(uint256.Int).Add(&x1, &y1), number.Add(&x1, &y1))
			assert.Equal(t, new(uint256.Int).Sub(&x1, &y1), number.Sub(&x1, &y1))
			assert.Equal(t, new(uint256.Int).Mul(&x1, &y1), number.Mul(&x1, &y1))
			assert.Equal(t, new(uint256.Int).Div(&x1, &y1), number.Div(&x1, &y1))
		})
	}
}

func SampleUsage(x, y *uint256.Int) bool {
	res := number.Add(
		number.Mul(
			number.Div(x, number.SetUint64(10)),
			number.AddUint64(y, 1),
		),
		number.Sub(
			number.SubUint64(x, 1),
			number.Set(y),
		),
	)
	return res.IsZero()
}

func BenchmarkWrapper(b *testing.B) {
	x := uint256.MustFromHex("0xfffffffffffffffffffffff6")
	y := uint256.MustFromHex("0xeffffffffffff")

	// this should has `0 allocs/op`
	for i := 0; i < b.N; i++ {
		_ = SampleUsage(x, y)
	}
}
