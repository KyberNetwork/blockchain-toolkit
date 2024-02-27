package number_test

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/KyberNetwork/blockchain-toolkit/integer"
	"github.com/KyberNetwork/blockchain-toolkit/number"
	"github.com/holiman/uint256"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFillBig(t *testing.T) {
	var err error
	var xU256 uint256.Int

	for i := 0; i < 200; i++ {
		xStr := number.RandNumberHexString(64)

		t.Run(fmt.Sprintf("test %s", xStr), func(t *testing.T) {
			err = xU256.SetFromHex("0x" + xStr)
			require.Nil(t, err)

			// known correct value to compare against
			xBI, ok := new(big.Int).SetString(xStr, 16)
			require.True(t, ok)

			// should resize and fill small int
			smallBI := big.NewInt(12)
			number.FillBig(&xU256, smallBI)
			assert.Equal(t, xBI, smallBI)

			// should resize and fill big int
			largeBI := integer.TenPow(200)
			number.FillBig(&xU256, largeBI)
			assert.Equal(t, xBI, largeBI)

		})
	}
}

func BenchmarkFillBig(b *testing.B) {
	xU256 := number.Number_1e18
	xBI := big.NewInt(0)
	for i := 0; i < b.N; i++ {
		number.FillBig(xU256, xBI)
	}
}
