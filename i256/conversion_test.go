package i256_test

import (
	"testing"

	"github.com/KyberNetwork/blockchain-toolkit/i256"
	"github.com/KyberNetwork/blockchain-toolkit/number"
	"github.com/KyberNetwork/int256"
	"github.com/holiman/uint256"
	"github.com/stretchr/testify/assert"
)

func TestUnsafeToUInt256(t *testing.T) {
	assert.Equal(t, uint256.NewInt(1), i256.UnsafeToUInt256(int256.NewInt(1)))
	assert.Equal(t, uint256.NewInt(0), i256.UnsafeToUInt256(int256.NewInt(0)))
	assert.Equal(t, number.MaxU256, i256.UnsafeToUInt256(int256.NewInt(-1)))
	assert.Equal(t, uint256.MustFromHex("0x7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"), i256.UnsafeToUInt256(int256.MaxI256))
	assert.Equal(t, uint256.MustFromHex("0x8000000000000000000000000000000000000000000000000000000000000000"), i256.UnsafeToUInt256(int256.MinI256))
}

func TestSafeConvertToUInt256(t *testing.T) {
	assert.Equal(t, uint256.NewInt(1), i256.SafeConvertToUInt256(int256.NewInt(1)))
	assert.Equal(t, uint256.NewInt(0), i256.SafeConvertToUInt256(int256.NewInt(0)))
	assert.Equal(t, uint256.MustFromHex("0x7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"), i256.UnsafeToUInt256(int256.MaxI256))

	assert.Panics(t, func() { i256.SafeConvertToUInt256(int256.NewInt(-1)) })
	assert.Panics(t, func() { i256.SafeConvertToUInt256(int256.MinI256) })
}

func TestSafeToInt256(t *testing.T) {
	assert.Equal(t, int256.NewInt(1), i256.SafeToInt256(uint256.NewInt(1)))
	assert.Equal(t, int256.NewInt(0), i256.SafeToInt256(uint256.NewInt(0)))
	assert.Equal(t, int256.MaxI256, i256.SafeToInt256(uint256.MustFromHex("0x7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff")))

	assert.Panics(t, func() {
		i256.SafeToInt256(uint256.MustFromHex("0x8000000000000000000000000000000000000000000000000000000000000000"))
	})
	assert.Panics(t, func() { i256.SafeToInt256(number.MaxU256) })
}
