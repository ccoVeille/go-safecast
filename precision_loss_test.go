package safecast

import (
	"errors"
	"testing"
)

func Test_checkIntegerToFloatPrecisionLoss(t *testing.T) {
	tests := []struct {
		name      string
		original  int64
		converted float32
		wantErr   error
	}{
		{
			name:      "precision loss",
			original:  16_777_217,
			converted: 16_777_216,
			wantErr:   ErrPrecisionLoss,
		},
		{
			name:      "negative number with precision loss",
			original:  -16_777_217,
			converted: -16_777_216,
			wantErr:   ErrPrecisionLoss,
		},
		{
			name:      "no precision loss",
			original:  42,
			converted: 42,
		},
		{
			name:      "negative number without precision loss",
			original:  -42,
			converted: -42,
		},
		{
			name:      "large integer without precision loss",
			original:  33_554_432,
			converted: 33_554_432,
		},
		{
			name:      "large negative integer without precision loss",
			original:  -33_554_432,
			converted: -33_554_432,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := checkIntegerToFloatPrecisionLoss(tt.original, tt.converted)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("expected error %v, got %v", tt.wantErr, err)
			}

			if result != tt.converted {
				t.Errorf("expected %v, got %v", tt.converted, result)
			}
		})
	}
}
