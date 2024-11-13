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
)

func TestToInt32_64bit(t *testing.T) {
	t.Run("from int", func(t *testing.T) {
		assertInt32Error(t, []caseInt32[int]{
			{name: "positive out of range", input: math.MaxInt32 + 1},
			{name: "negative out of range", input: math.MinInt32 - 1},
		})
	})
}

func TestToUint32_64bit(t *testing.T) {
	t.Run("from int", func(t *testing.T) {
		assertUint32Error(t, []caseUint32[int]{
			{name: "positive out of range", input: math.MaxUint32 + 1},
			{name: "negative value", input: -1},
		})
	})
}

func TestToInt64_64bit(t *testing.T) {
	t.Run("from uint", func(t *testing.T) {
		assertInt64Error(t, []caseInt64[uint]{
			{name: "positive out of range", input: math.MaxInt64 + 1},
		})
	})
}

func TestToInt_64bit(t *testing.T) {
	t.Run("from uint", func(t *testing.T) {
		assertIntError(t, []caseInt[uint]{
			{name: "positive out of range", input: math.MaxInt64 + 1},
		})
	})

	t.Run("from float64", func(t *testing.T) {
		assertIntOK(t, []caseInt[float64]{
			{name: "math.MinInt64", input: math.MinInt64, want: math.MinInt64}, // pass because of float imprecision
		})
	})
}
