package i256_test

import (
	"fmt"
	"math/big"
	"math/rand"
	"testing"

	"github.com/KyberNetwork/blockchain-toolkit/i256"
	"github.com/KyberNetwork/blockchain-toolkit/number"
	"github.com/KyberNetwork/int256"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWrapper(t *testing.T) {
	for i := 0; i < 200; i++ {
		xStr := number.RandNumberHexString(63)
		yStr := number.RandNumberHexString(63)

		// 1st digit 0-7 to avoid negative
		{
			c := rand.Intn(8)
			xStr = fmt.Sprintf("%x%s", c, xStr)
			c = rand.Intn(8)
			yStr = fmt.Sprintf("%x%s", c, yStr)
		}

		var err error

		t.Run(fmt.Sprintf("test %s %s", xStr, yStr), func(t *testing.T) {
			xBI, ok := new(big.Int).SetString(xStr, 16)
			require.True(t, ok)
			yBI, ok := new(big.Int).SetString(yStr, 16)
			require.True(t, ok)

			var x1, y1 int256.Int
			err = x1.SetFromDec(xBI.String())
			require.Nil(t, err)
			err = y1.SetFromDec(yBI.String())
			require.Nil(t, err)

			x3 := i256.Set(&x1)
			y3 := i256.Set(&y1)
			assert.Equal(t, x1, *x3)
			assert.Equal(t, y1, *y3)

			assert.Equal(t, new(int256.Int).Add(&x1, &y1), i256.Add(&x1, &y1))
			assert.Equal(t, new(int256.Int).Sub(&x1, &y1), i256.Sub(&x1, &y1))
			assert.Equal(t, new(int256.Int).Mul(&x1, &y1), i256.Mul(&x1, &y1))
			assert.Equal(t, new(int256.Int).Quo(&x1, &y1), i256.Div(&x1, &y1))
			assert.Equal(t, new(int256.Int).Neg(&x1), i256.Neg(&x1))
			assert.Equal(t, new(int256.Int).Rsh(&x1, 10), i256.Rsh(&x1, 10))
			assert.Equal(t, new(int256.Int).Lsh(&x1, 10), i256.Lsh(&x1, 10))
		})
	}
}

func SampleUsage(x, y *int256.Int) bool {
	res := i256.Add(
		i256.Mul(
			i256.Div(x, i256.SetInt64(10)),
			y,
		),
		i256.Sub(
			x,
			i256.Add(
				i256.Div(
					i256.Set(y),
					y,
				),
				i256.Sub(
					x,
					i256.Mul(i256.SetInt64(11), y),
				),
			),
		),
	)
	return res.IsZero()
}

func BenchmarkWrapper(b *testing.B) {
	x := int256.MustFromDec("79228162514264337593543950326")
	y := int256.MustFromDec("4222124650659839")

	// this should has `0 allocs/op`
	for i := 0; i < b.N; i++ {
		_ = SampleUsage(x, y)
	}
}
