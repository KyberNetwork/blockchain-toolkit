package number

import (
	"fmt"
	"math/rand"
)

// RandNumberString generate random string of number with length <= maxLen
func RandNumberString(maxLen int) string {
	sLen := rand.Intn(maxLen) + 1
	var s string
	for i := 0; i < sLen; i++ {
		var c int
		if i == 0 {
			c = rand.Intn(9) + 1
		} else {
			c = rand.Intn(10)
		}
		s = fmt.Sprintf("%s%d", s, c)
	}
	return s
}

// RandNumberHexString generate random string of hex number with length <= maxLen (without 0x prefix)
func RandNumberHexString(maxLen int) string {
	sLen := rand.Intn(maxLen) + 1
	var s string
	for i := 0; i < sLen; i++ {
		var c int
		if i == 0 {
			c = rand.Intn(15) + 1
		} else {
			c = rand.Intn(16)
		}
		s = fmt.Sprintf("%s%x", s, c)
	}
	return s
}
