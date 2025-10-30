//go:build !386 && !arm

package safecast_test

// The tests in conversion_test.go are the ones that are not architecture dependent
// The tests in conversion_64bit_test.go complete them for 64-bit systems
//
// This architecture dependent file covers the fact, you can reach a higher value with int and uint
// on 64-bit systems, but you will get a compile error on 32-bit.
// This is why it needs to be tested in an architecture dependent way.

import (
	"math"
	"testing"

	"github.com/ccoveille/go-safecast/v2"
)

// TestConvert_64bit completes the [TestConvert] tests in conversion_test.go
// it contains the tests that can only works on 64-bit systems
func TestConvert_64bit(t *testing.T) {
	t.Run("to uint32", func(t *testing.T) {
		for name, tt := range map[string]struct {
			input uint64
			want  uint32
		}{
			"positive out of range": {input: uint64(math.MaxUint32 + 1), want: 0},
		} {
			t.Run(name, func(t *testing.T) {
				got, err := safecast.Convert[uint32](tt.input)
				assertEqual(t, tt.want, got)
				requireErrorIs(t, err, safecast.ErrConversionIssue)
			})
		}
	})
}
