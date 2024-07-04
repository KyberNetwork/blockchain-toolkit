package account

import "fmt"

var (
	ErrInvalidAddress         = fmt.Errorf("invalid Ethereum address")
	ErrInvalidNumVisibleChars = fmt.Errorf("number of visible characters must be even")
	ErrInvalidMaskedChars     = fmt.Errorf("invalid number of masked characters")
)

// MaskAddress masks the middle part of an Ethereum address with a number of masked characters and returns the masked address.
// The number of visible characters must be even.
// The number of masked characters must be greater than 0.
// The masked character can be any character.
// @param address: The Ethereum address to mask
// @param numVisibleChars: The number of visible characters to keep at the start and end of the address
// @param maskedCharacter: The character to use for masking the middle part of the address
// @return The masked address and error if any
func MaskAddress(address string, numVisibleChars int, maskedCharacter rune) (string, error) {
	if !IsValidAddress(address) {
		return "", ErrInvalidAddress
	}

	if numVisibleChars%2 != 0 {
		return "", ErrInvalidNumVisibleChars
	}

	prefixLength := len("0x") + numVisibleChars/2 // The number of characters to keep at the start (0x + first some hex digits)
	suffixLength := numVisibleChars / 2           // The number of characters to keep at the end (last some hex digits)

	addressLength := len(address)
	maskedCharacterCount := addressLength - prefixLength - suffixLength

	if maskedCharacterCount <= 0 {
		return "", ErrInvalidMaskedChars
	}

	// Create the masked part
	maskedPart := make([]rune, maskedCharacterCount)
	for i := range maskedPart {
		maskedPart[i] = maskedCharacter
	}

	// Construct the masked address
	maskedAddress := address[:prefixLength] + string(maskedPart) + address[addressLength-suffixLength:]

	return maskedAddress, nil
}
