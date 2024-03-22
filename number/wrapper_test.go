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

			// from univ3 safeadd: `require((z = x + y) >= x);`
			if number.Add(&x1, &y1).Cmp(&x1) >= 0 {
				assert.Equal(t, new(uint256.Int).Add(&x1, &y1), number.SafeAdd(&x1, &y1))
			} else {
				assert.PanicsWithError(t, number.ErrOverflow.Error(), func() {
					number.SafeAdd(&x1, &y1)
				})
			}

			// from univ3 safesub: `require((z = x - y) <= x);`
			if number.Sub(&x1, &y1).Cmp(&x1) <= 0 {
				assert.Equal(t, new(uint256.Int).Sub(&x1, &y1), number.SafeSub(&x1, &y1))
			} else {
				assert.PanicsWithError(t, number.ErrUnderflow.Error(), func() {
					number.SafeSub(&x1, &y1)
				})
			}

			// from univ3 safemul: `require(x == 0 || (z = x * y) / x == y);`
			if x1.IsZero() || number.Div(number.Mul(&x1, &y1), &x1).Cmp(&y1) == 0 {
				assert.Equal(t, new(uint256.Int).Mul(&x1, &y1), number.SafeMul(&x1, &y1))
			} else {
				assert.PanicsWithError(t, number.ErrOverflow.Error(), func() {
					number.SafeMul(&x1, &y1)
				})
			}
		})
	}
}

func TestSafeAdd(t *testing.T) {
	for _, tc := range []struct {
		x             *uint256.Int
		y             *uint256.Int
		errorOrResult interface{}
	}{
		{number.MaxU256, number.Number_1, number.ErrOverflow},
		{number.MaxU256, number.Number_2, number.ErrOverflow},
		{number.SubUint64(number.MaxU256, 1), number.Number_2, number.ErrOverflow},
		{number.SubUint64(number.MaxU256, 1), number.Number_1, number.MaxU256},
		{number.Number_1, number.Number_2, number.Number_3},
	} {
		t.Run(fmt.Sprintf("test %v %v", tc.x, tc.y), func(t *testing.T) {
			if err, ok := tc.errorOrResult.(error); ok {
				assert.PanicsWithError(t, err.Error(), func() {
					number.SafeAdd(tc.x, tc.y)
				})
				assert.False(t, number.Add(tc.x, tc.y).Cmp(tc.x) >= 0)
			} else {
				assert.True(t, number.Add(tc.x, tc.y).Cmp(tc.x) >= 0)
				assert.Equal(t, tc.errorOrResult, number.SafeAdd(tc.x, tc.y))
			}
		})
	}
}

func TestSafeSub(t *testing.T) {
	for _, tc := range []struct {
		x             *uint256.Int
		y             *uint256.Int
		errorOrResult interface{}
	}{
		{number.MaxU256, number.Number_1, number.SubUint64(number.MaxU256, 1)},
		{number.MaxU256, number.Number_2, number.SubUint64(number.MaxU256, 2)},
		{number.Number_1, number.Number_2, number.ErrUnderflow},
		{uint256.NewInt(0), number.Number_1, number.ErrUnderflow},
		{number.Number_1, number.MaxU256, number.ErrUnderflow},
		{uint256.NewInt(0), number.MaxU256, number.ErrUnderflow},
		{number.SubUint64(number.MaxU256, 1), number.MaxU256, number.ErrUnderflow},
	} {
		t.Run(fmt.Sprintf("test %v %v", tc.x, tc.y), func(t *testing.T) {
			if err, ok := tc.errorOrResult.(error); ok {
				assert.PanicsWithError(t, err.Error(), func() {
					number.SafeSub(tc.x, tc.y)
				})
				assert.False(t, number.Sub(tc.x, tc.y).Cmp(tc.x) <= 0)
			} else {
				assert.True(t, number.Sub(tc.x, tc.y).Cmp(tc.x) <= 0)
				assert.Equal(t, tc.errorOrResult, number.SafeSub(tc.x, tc.y))
			}
		})
	}
}

func TestSafeMul(t *testing.T) {
	for _, tc := range []struct {
		x             *uint256.Int
		y             *uint256.Int
		errorOrResult interface{}
	}{
		{number.MaxU256, uint256.NewInt(0), uint256.NewInt(0)},
		{number.MaxU256, number.Number_1, number.MaxU256},
		{number.MaxU256, number.Number_2, number.ErrOverflow},
		{number.Div(number.MaxU256, number.Number_2), number.Number_2, number.SubUint64(number.MaxU256, 1)},
		{number.AddUint64(number.Div(number.MaxU256, number.Number_2), 1), number.Number_2, number.ErrOverflow},
		{number.Number_1, number.Number_2, number.Number_2},
	} {
		t.Run(fmt.Sprintf("test %v %v", tc.x, tc.y), func(t *testing.T) {
			if err, ok := tc.errorOrResult.(error); ok {
				assert.PanicsWithError(t, err.Error(), func() {
					number.SafeMul(tc.x, tc.y)
				})
				assert.False(t, tc.x.IsZero() || number.Div(number.Mul(tc.x, tc.y), tc.x).Cmp(tc.y) == 0)
			} else {
				assert.True(t, tc.x.IsZero() || number.Div(number.Mul(tc.x, tc.y), tc.x).Cmp(tc.y) == 0)
				assert.Equal(t, tc.errorOrResult, number.SafeMul(tc.x, tc.y))
			}
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
			number.SafeAdd(
				number.Div(
					number.Set(y),
					number.SafeDivZ(y, number.SetUint64(10), &uint256.Int{}),
				),
				number.SafeSub(
					number.SafeDiv(x, number.SetUint64(10)),
					number.SafeMul(number.SetUint64(11), y),
				),
			),
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
