package account

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestIsZeroAddress(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		iAddress any
		expected bool
	}{
		{
			name:     "it should return false when it is a valid address",
			iAddress: common.HexToAddress("0x323b5d4c32345ced77393b3530b1eed0f346429d"),
			expected: false,
		},
		{
			name:     "it should return true when it is a zero address",
			iAddress: common.HexToAddress("0x0000000000000000000000000000000000000000"),
			expected: true,
		},
		{
			name:     "it should return false when it is a valid address in str",
			iAddress: "0x323b5d4c32345ced77393b3530b1eed0f346429d",
			expected: false,
		},
		{
			name:     "it should return true when it is a zero address in str",
			iAddress: "0x0000000000000000000000000000000000000000",
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := IsZeroAddress(tc.iAddress)

			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestIsValidAddress(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		iAddress any
		expected bool
	}{
		{
			name:     "it should return true when it is a valid address",
			iAddress: "0x323b5d4c32345ced77393b3530b1eed0f346429d",
			expected: true,
		},
		{
			name:     "it should return false when it an invalid address",
			iAddress: "0xabc",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := IsValidAddress(tc.iAddress)

			assert.Equal(t, tc.expected, actual)
		})
	}
}
