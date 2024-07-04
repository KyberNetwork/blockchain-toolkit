package account

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaskAddress(t *testing.T) {
	type args struct {
		address         string
		numVisibleChars int
		maskedCharacter rune
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr error
	}{
		{
			name: "valid address",
			args: args{
				address:         "0x1234567890abcdef1234567890abcdef12345678",
				numVisibleChars: 8,
				maskedCharacter: 'x',
			},
			want:    "0x1234xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx5678",
			wantErr: nil,
		},
		{
			name: "invalid address",
			args: args{
				address:         "0x1234567890abcdef1234567890abcdef1234567",
				numVisibleChars: 8,
				maskedCharacter: 'x',
			},
			want:    "",
			wantErr: ErrInvalidAddress,
		},
		{
			name: "invalid number of visible characters",
			args: args{
				address:         "0x1234567890abcdef1234567890abcdef12345678",
				numVisibleChars: 7,
				maskedCharacter: 'x',
			},
			want:    "",
			wantErr: ErrInvalidNumVisibleChars,
		},
		{
			name: "invalid number of masked characters",
			args: args{
				address:         "0x1234567890abcdef1234567890abcdef12345678",
				numVisibleChars: 40,
				maskedCharacter: 'x',
			},
			want:    "",
			wantErr: ErrInvalidMaskedChars,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MaskAddress(tt.args.address, tt.args.numVisibleChars, tt.args.maskedCharacter)
			if err != nil {
				assert.ErrorIsf(t, err, tt.wantErr, "MaskAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equalf(t, tt.want, got, "MaskAddress() = %v, want %v", got, tt.want)
		})
	}
}
