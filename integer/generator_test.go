package integer

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTenPow(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		decimals uint8
		expected *big.Int
	}{
		{
			name:     "it should generate correct number",
			decimals: 9,
			expected: big.NewInt(1e9),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := TenPow(tc.decimals)

			assert.Zero(t, tc.expected.Cmp(actual))
		})
	}
}
