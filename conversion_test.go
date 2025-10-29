package safecast_test

// The tests in conversion_test.go are the ones that are not architecture dependent
// The tests in conversion_64bit_test.go complete them for 64-bit systems
//
// The architecture dependent file covers the fact, you can reach a higher value with int and uint
// on 64-bit systems, but you will get a compile error on 32-bit.
// This is why it needs to be tested in an architecture dependent way.

import (
	"errors"
	"fmt"
	"math"
	"testing"

	"github.com/ccoveille/go-safecast"
)

func TestToInt8(t *testing.T) {
	t.Run("from int", func(t *testing.T) {
		assertInt8OK(t, []caseInt8[int]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
			{name: "negative within range", input: -100, want: -100},
		})

		assertInt8Error(t, []caseInt8[int]{
			{name: "positive out of range", input: math.MaxInt8 + 1},
			{name: "negative out of range", input: math.MinInt8 - 1},
		})
	})

	t.Run("from int8", func(t *testing.T) {
		assertInt8OK(t, []caseInt8[int8]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
			{name: "negative within range", input: -100, want: -100},
		})
	})

	t.Run("from int16", func(t *testing.T) {
		assertInt8OK(t, []caseInt8[int16]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
			{name: "negative within range", input: -100, want: -100},
		})

		assertInt8Error(t, []caseInt8[int16]{
			{name: "positive out of range", input: math.MaxInt8 + 1},
			{name: "negative out of range", input: math.MinInt8 - 1},
		})
	})

	t.Run("from int32", func(t *testing.T) {
		assertInt8OK(t, []caseInt8[int32]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
			{name: "negative within range", input: -100, want: -100},
		})

		assertInt8Error(t, []caseInt8[int32]{
			{name: "positive out of range", input: math.MaxInt8 + 1},
			{name: "negative out of range", input: math.MinInt8 - 1},
		})
	})

	t.Run("from int64", func(t *testing.T) {
		assertInt8OK(t, []caseInt8[int64]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
			{name: "negative within range", input: -100, want: -100},
		})

		assertInt8Error(t, []caseInt8[int64]{
			{name: "positive out of range", input: math.MaxInt8 + 1},
			{name: "negative out of range", input: math.MinInt8 - 1},
		})
	})

	t.Run("from uint", func(t *testing.T) {
		assertInt8OK(t, []caseInt8[uint]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertInt8Error(t, []caseInt8[uint]{
			{name: "positive out of range", input: math.MaxInt8 + 1},
		})
	})

	t.Run("from uint8", func(t *testing.T) {
		assertInt8OK(t, []caseInt8[uint8]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})
	})

	t.Run("from uint16", func(t *testing.T) {
		assertInt8OK(t, []caseInt8[uint16]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertInt8Error(t, []caseInt8[uint16]{
			{name: "positive out of range", input: math.MaxInt8 + 1},
		})
	})

	t.Run("from uint32", func(t *testing.T) {
		assertInt8OK(t, []caseInt8[uint32]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertInt8Error(t, []caseInt8[uint32]{
			{name: "positive out of range", input: math.MaxInt8 + 1},
		})
	})

	t.Run("from uint64", func(t *testing.T) {
		assertInt8OK(t, []caseInt8[uint64]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertInt8Error(t, []caseInt8[uint64]{
			{name: "positive out of range", input: math.MaxInt8 + 1},
		})
	})

	t.Run("from float32", func(t *testing.T) {
		assertInt8OK(t, []caseInt8[float32]{
			{name: "zero", input: 0.0, want: 0},
			{name: "rounded value", input: 1.1, want: 1},
		})

		assertInt8Error(t, []caseInt8[float32]{
			{name: "positive out of range", input: math.MaxInt8 + 1},
		})
	})

	t.Run("from float64", func(t *testing.T) {
		assertInt8OK(t, []caseInt8[float64]{
			{name: "zero", input: 0.0, want: 0},
			{name: "rounded value", input: 1.1, want: 1},
		})

		assertInt8Error(t, []caseInt8[float64]{
			{name: "positive out of range", input: math.MaxInt8 + 1},
		})
	})
}

func TestToUint8(t *testing.T) {
	t.Run("from int", func(t *testing.T) {
		assertUint8OK(t, []caseUint8[int]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertUint8Error(t, []caseUint8[int]{
			{name: "positive out of range", input: math.MaxUint8 + 1},
			{name: "negative value", input: -1},
		})
	})

	t.Run("from int8", func(t *testing.T) {
		assertUint8OK(t, []caseUint8[int8]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertUint8Error(t, []caseUint8[int8]{
			{name: "negative value", input: -1},
		})
	})

	t.Run("from int16", func(t *testing.T) {
		assertUint8OK(t, []caseUint8[int16]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertUint8Error(t, []caseUint8[int16]{
			{name: "positive out of range", input: 10000},
			{name: "negative value", input: -1},
		})
	})

	t.Run("from int32", func(t *testing.T) {
		assertUint8OK(t, []caseUint8[int32]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertUint8Error(t, []caseUint8[int32]{
			{name: "positive out of range", input: 100000},
			{name: "negative value", input: -1},
		})
	})

	t.Run("from int64", func(t *testing.T) {
		assertUint8OK(t, []caseUint8[int64]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertUint8Error(t, []caseUint8[int64]{
			{name: "positive out of range", input: 100000},
			{name: "negative value", input: -1},
		})
	})

	t.Run("from uint", func(t *testing.T) {
		assertUint8OK(t, []caseUint8[uint]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertUint8Error(t, []caseUint8[uint]{
			{name: "positive out of range", input: math.MaxUint8 + 1},
		})
	})

	t.Run("from uint8", func(t *testing.T) {
		assertUint8OK(t, []caseUint8[uint8]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})
	})

	t.Run("from uint16", func(t *testing.T) {
		assertUint8OK(t, []caseUint8[uint16]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertUint8Error(t, []caseUint8[uint]{
			{name: "positive out of range", input: math.MaxUint8 + 1},
		})
	})

	t.Run("from uint32", func(t *testing.T) {
		assertUint8OK(t, []caseUint8[uint32]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertUint8Error(t, []caseUint8[uint]{
			{name: "positive out of range", input: math.MaxUint8 + 1},
		})
	})

	t.Run("from uint64", func(t *testing.T) {
		assertUint8OK(t, []caseUint8[uint64]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertUint8Error(t, []caseUint8[uint]{
			{name: "positive out of range", input: math.MaxUint8 + 1},
		})
	})

	t.Run("from float32", func(t *testing.T) {
		assertUint8OK(t, []caseUint8[float32]{
			{name: "zero", input: 0.0, want: 0},
			{name: "rounded value", input: 1.1, want: 1},
			{name: "positive within range", input: 100.9, want: 100},
		})

		assertUint8Error(t, []caseUint8[float32]{
			{name: "positive out of range", input: math.MaxUint8 + 1},
		})
	})

	t.Run("from float64", func(t *testing.T) {
		assertUint8OK(t, []caseUint8[float64]{
			{name: "zero", input: 0.0, want: 0},
			{name: "rounded value", input: 1.1, want: 1},
			{name: "positive within range", input: 100.9, want: 100},
		})

		assertUint8Error(t, []caseUint8[float64]{
			{name: "positive out of range", input: math.MaxUint8 + 1},
		})
	})
}

func TestToInt16(t *testing.T) {
	t.Run("from int", func(t *testing.T) {
		assertInt16OK(t, []caseInt16[int]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
			{name: "negative within range", input: -10000, want: -10000},
		})

		assertInt16Error(t, []caseInt16[int]{
			{name: "positive out of range", input: math.MaxInt16 + 1},
			{name: "negative out of range", input: math.MinInt16 - 1},
		})
	})

	t.Run("from int8", func(t *testing.T) {
		assertInt16OK(t, []caseInt16[int8]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
			{name: "negative within range", input: -100, want: -100},
		})
	})

	t.Run("from int16", func(t *testing.T) {
		assertInt16OK(t, []caseInt16[int16]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
			{name: "negative within range", input: -10000, want: -10000},
		})
	})

	t.Run("from int32", func(t *testing.T) {
		assertInt16OK(t, []caseInt16[int32]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
			{name: "negative within range", input: -10000, want: -10000},
		})

		assertInt16Error(t, []caseInt16[int32]{
			{name: "positive out of range", input: 100000},
			{name: "negative out of range", input: -100000},
		})
	})

	t.Run("from int64", func(t *testing.T) {
		assertInt16OK(t, []caseInt16[int64]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
			{name: "negative within range", input: -10000, want: -10000},
		})

		assertInt16Error(t, []caseInt16[int64]{
			{name: "positive out of range", input: 100000},
			{name: "negative out of range", input: -100000},
		})
	})

	t.Run("from uint", func(t *testing.T) {
		assertInt16OK(t, []caseInt16[uint]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertInt16Error(t, []caseInt16[uint]{
			{name: "positive out of range", input: math.MaxInt16 + 1},
		})
	})

	t.Run("from uint8", func(t *testing.T) {
		assertInt16OK(t, []caseInt16[uint8]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})
	})

	t.Run("from uint16", func(t *testing.T) {
		assertInt16OK(t, []caseInt16[uint16]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertInt16Error(t, []caseInt16[uint16]{
			{name: "positive out of range", input: math.MaxInt16 + 1},
		})
	})

	t.Run("from uint32", func(t *testing.T) {
		assertInt16OK(t, []caseInt16[uint32]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertInt16Error(t, []caseInt16[uint32]{
			{name: "positive out of range", input: math.MaxInt16 + 1},
		})
	})

	t.Run("from uint64", func(t *testing.T) {
		assertInt16OK(t, []caseInt16[uint64]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertInt16Error(t, []caseInt16[uint64]{
			{name: "positive out of range", input: math.MaxInt16 + 1},
		})
	})

	t.Run("from float32", func(t *testing.T) {
		assertInt16OK(t, []caseInt16[float32]{
			{name: "zero", input: 0.0, want: 0},
			{name: "rounded value", input: 1.1, want: 1},
			{name: "positive within range", input: 10000, want: 10000},
		})

		assertInt16Error(t, []caseInt16[float32]{
			{name: "positive out of range", input: math.MaxInt16 + 1},
		})
	})

	t.Run("from float64", func(t *testing.T) {
		assertInt16OK(t, []caseInt16[float64]{
			{name: "zero", input: 0.0, want: 0},
			{name: "rounded value", input: 1.1, want: 1},
			{name: "positive within range", input: 10000, want: 10000},
		})

		assertInt16Error(t, []caseInt16[float64]{
			{name: "positive out of range", input: math.MaxInt16 + 1},
		})
	})
}

func TestToUint16(t *testing.T) {
	t.Run("from int", func(t *testing.T) {
		assertUint16OK(t, []caseUint16[int]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
		})

		assertUint16Error(t, []caseUint16[int]{
			{name: "positive out of range", input: math.MaxUint16 + 1},
			{name: "negative value", input: -1},
		})
	})

	t.Run("from int8", func(t *testing.T) {
		assertUint16OK(t, []caseUint16[int8]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})
	})

	t.Run("from int16", func(t *testing.T) {
		assertUint16OK(t, []caseUint16[int16]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
		})
	})

	t.Run("from int32", func(t *testing.T) {
		assertUint16OK(t, []caseUint16[int32]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
		})

		assertUint16Error(t, []caseUint16[int32]{
			{name: "positive out of range", input: 100000},
			{name: "negative value", input: -1},
		})
	})

	t.Run("from int64", func(t *testing.T) {
		assertUint16OK(t, []caseUint16[int64]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
		})

		assertUint16Error(t, []caseUint16[int64]{
			{name: "positive out of range", input: 100000},
			{name: "negative value", input: -1},
		})
	})

	t.Run("from uint", func(t *testing.T) {
		assertUint16OK(t, []caseUint16[uint]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertUint16Error(t, []caseUint16[uint]{
			{name: "positive out of range", input: math.MaxUint16 + 1},
		})
	})

	t.Run("from uint8", func(t *testing.T) {
		assertUint16OK(t, []caseUint16[uint8]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})
	})

	t.Run("from uint16", func(t *testing.T) {
		assertUint16OK(t, []caseUint16[uint16]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})
	})

	t.Run("from uint32", func(t *testing.T) {
		assertUint16OK(t, []caseUint16[uint32]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertUint16Error(t, []caseUint16[uint32]{
			{name: "positive out of range", input: math.MaxUint16 + 1},
		})
	})

	t.Run("from uint64", func(t *testing.T) {
		assertUint16OK(t, []caseUint16[uint64]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertUint16Error(t, []caseUint16[uint64]{
			{name: "positive out of range", input: math.MaxUint16 + 1},
		})
	})

	t.Run("from float32", func(t *testing.T) {
		assertUint16OK(t, []caseUint16[float32]{
			{name: "zero", input: 0.0, want: 0},
			{name: "rounded value", input: 1.1, want: 1},
			{name: "positive within range", input: 10000, want: 10000},
		})

		assertUint16Error(t, []caseUint16[float32]{
			{name: "positive out of range", input: math.MaxUint16 + 1},
		})
	})

	t.Run("from float64", func(t *testing.T) {
		assertUint16OK(t, []caseUint16[float64]{
			{name: "zero", input: 0.0, want: 0},
			{name: "rounded value", input: 1.1, want: 1},
			{name: "positive within range", input: 10000, want: 10000},
		})

		assertUint16Error(t, []caseUint16[float32]{
			{name: "positive out of range", input: math.MaxUint16 + 1},
		})
	})
}

func TestToInt32(t *testing.T) {
	t.Run("from int", func(t *testing.T) {
		assertInt32OK(t, []caseInt32[int]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
			{name: "negative within range", input: -10000, want: -10000},
		})

		// There are extra checks in [TestToInt32_64bit]
		// the tests are separated because they cannot work on i386
	})

	t.Run("from int8", func(t *testing.T) {
		assertInt32OK(t, []caseInt32[int8]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
			{name: "negative within range", input: -100, want: -100},
		})
	})

	t.Run("from int16", func(t *testing.T) {
		assertInt32OK(t, []caseInt32[int16]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
			{name: "negative within range", input: -10000, want: -10000},
		})
	})

	t.Run("from int32", func(t *testing.T) {
		assertInt32OK(t, []caseInt32[int32]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
			{name: "negative within range", input: -10000, want: -10000},
		})
	})

	t.Run("from int64", func(t *testing.T) {
		assertInt32OK(t, []caseInt32[int64]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
			{name: "negative within range", input: -10000, want: -10000},
		})

		assertInt32Error(t, []caseInt32[int64]{
			{name: "positive out of range", input: math.MaxInt32 + 1},
			{name: "negative out of range", input: math.MinInt32 - 1},
		})
	})

	t.Run("from uint", func(t *testing.T) {
		assertInt32OK(t, []caseInt32[uint]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertInt32Error(t, []caseInt32[uint]{
			{name: "positive out of range", input: math.MaxInt32 + 1},
		})
	})

	t.Run("from uint8", func(t *testing.T) {
		assertInt32OK(t, []caseInt32[uint8]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})
	})

	t.Run("from uint16", func(t *testing.T) {
		assertInt32OK(t, []caseInt32[uint16]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: math.MaxUint16, want: math.MaxUint16},
		})
	})

	t.Run("from uint32", func(t *testing.T) {
		assertInt32OK(t, []caseInt32[uint32]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertInt32Error(t, []caseInt32[uint32]{
			{name: "positive out of range", input: math.MaxInt32 + 1},
		})
	})

	t.Run("from uint64", func(t *testing.T) {
		assertInt32OK(t, []caseInt32[uint64]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertInt32Error(t, []caseInt32[uint64]{
			{name: "positive out of range", input: math.MaxInt32 + 1},
		})
	})

	t.Run("from float32", func(t *testing.T) {
		assertInt32OK(t, []caseInt32[float32]{
			{name: "zero", input: 0.0, want: 0},
			{name: "rounded value", input: 1.1, want: 1},
			{name: "positive within range", input: 10000, want: 10000},
		})

		assertInt32Error(t, []caseInt32[float32]{
			{name: "positive out of range", input: math.MaxInt32 * 1.02}, // because of float imprecision, we have to exceed the min int64 to trigger the error
		})
	})

	t.Run("from float64", func(t *testing.T) {
		assertInt32OK(t, []caseInt32[float64]{
			{name: "zero", input: 0.0, want: 0},
			{name: "rounded value", input: 1.1, want: 1},
			{name: "positive within range", input: 10000, want: 10000},
		})

		assertInt32Error(t, []caseInt32[float64]{
			{name: "positive out of range", input: math.MaxInt32 * 1.02}, // because of float imprecision, we have to exceed the min int64 to trigger the error
		})
	})
}

func TestToUint32(t *testing.T) {
	t.Run("from int", func(t *testing.T) {
		assertUint32OK(t, []caseUint32[int]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
		})

		assertUint32Error(t, []caseUint32[int]{
			// There are extra checks in [TestToUint32_64bit]
			// the tests are separated because they cannot work on i386
			{name: "negative value", input: -1},
		})
	})

	t.Run("from int8", func(t *testing.T) {
		assertUint32OK(t, []caseUint32[int8]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertUint32Error(t, []caseUint32[int8]{
			{name: "negative value", input: -1},
		})
	})

	t.Run("from int16", func(t *testing.T) {
		assertUint32OK(t, []caseUint32[int16]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
		})

		assertUint32Error(t, []caseUint32[int16]{
			{name: "negative value", input: -1},
		})
	})

	t.Run("from int32", func(t *testing.T) {
		assertUint32OK(t, []caseUint32[int32]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
		})

		assertUint32Error(t, []caseUint32[int32]{
			{name: "negative value", input: -1},
		})
	})

	t.Run("from int64", func(t *testing.T) {
		assertUint32OK(t, []caseUint32[int64]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
		})

		assertUint32Error(t, []caseUint32[int64]{
			{name: "positive out of range", input: math.MaxUint32 + 1},
			{name: "negative value", input: -1},
		})
	})

	t.Run("from uint", func(t *testing.T) {
		assertUint32OK(t, []caseUint32[uint]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		// There are extra checks in [TestToUint32_64bit]
		// the tests are separated because they cannot work on i386
	})

	t.Run("from uint8", func(t *testing.T) {
		assertUint32OK(t, []caseUint32[uint8]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})
	})

	t.Run("from uint16", func(t *testing.T) {
		assertUint32OK(t, []caseUint32[uint16]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})
	})

	t.Run("from uint32", func(t *testing.T) {
		assertUint32OK(t, []caseUint32[uint32]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})
	})

	t.Run("from uint64", func(t *testing.T) {
		assertUint32OK(t, []caseUint32[uint64]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertUint32Error(t, []caseUint32[uint64]{
			{name: "positive out of range", input: math.MaxUint32 + 1},
		})
	})

	t.Run("from float32", func(t *testing.T) {
		assertUint32OK(t, []caseUint32[float32]{
			{name: "zero", input: 0.0, want: 0},
			{name: "rounded value", input: 1.1, want: 1},
			{name: "positive within range", input: 10000.9, want: 10000},
		})

		assertUint32Error(t, []caseUint32[float32]{
			{name: "positive out of range", input: math.MaxUint32 * 1.02}, // because of float imprecision, we have to exceed the min int64 to trigger the error
		})
	})

	t.Run("from float64", func(t *testing.T) {
		assertUint32OK(t, []caseUint32[float64]{
			{name: "zero", input: 0.0, want: 0},
			{name: "rounded value", input: 1.1, want: 1},
			{name: "positive within range", input: 10000.9, want: 10000},
		})

		assertUint32Error(t, []caseUint32[float64]{
			{name: "positive out of range", input: math.MaxUint32 * 1.02}, // because of float imprecision, we have to exceed the min int64 to trigger the error
			{name: "negative value", input: -1},
		})
	})
}

func TestToInt64(t *testing.T) {
	t.Run("from int", func(t *testing.T) {
		assertInt64OK(t, []caseInt64[int]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
			{name: "negative within range", input: -10000, want: -10000},
		})
	})

	t.Run("from int8", func(t *testing.T) {
		assertInt64OK(t, []caseInt64[int8]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
			{name: "negative within range", input: -100, want: -100},
		})
	})

	t.Run("from int16", func(t *testing.T) {
		assertInt64OK(t, []caseInt64[int16]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
			{name: "negative within range", input: -10000, want: -10000},
		})
	})

	t.Run("from int32", func(t *testing.T) {
		assertInt64OK(t, []caseInt64[int32]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
			{name: "negative within range", input: -10000, want: -10000},
		})
	})

	t.Run("from int64", func(t *testing.T) {
		assertInt64OK(t, []caseInt64[int64]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
			{name: "negative within range", input: -10000, want: -10000},
		})
	})

	t.Run("from uint", func(t *testing.T) {
		assertInt64OK(t, []caseInt64[uint]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		// There are extra checks in [TestToInt64_64bit]
		// the tests are separated because they cannot work on i386
	})

	t.Run("from uint8", func(t *testing.T) {
		assertInt64OK(t, []caseInt64[uint8]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})
	})

	t.Run("from uint16", func(t *testing.T) {
		assertInt64OK(t, []caseInt64[uint16]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})
	})

	t.Run("from uint32", func(t *testing.T) {
		assertInt64OK(t, []caseInt64[uint32]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})
	})

	t.Run("from uint64", func(t *testing.T) {
		assertInt64OK(t, []caseInt64[uint64]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertInt64Error(t, []caseInt64[uint64]{
			{name: "positive out of range", input: math.MaxInt64 + 1},
		})
	})

	t.Run("from float32", func(t *testing.T) {
		assertInt64OK(t, []caseInt64[float32]{
			{name: "zero", input: 0.0, want: 0},
			{name: "rounded value", input: 1.1, want: 1},
			{name: "positive within range", input: 10000.9, want: 10000},
			{name: "max int16", input: math.MaxInt16, want: math.MaxInt16},
			{name: "min int16", input: math.MinInt16, want: math.MinInt16},
			{name: "max int32", input: math.MaxInt32, want: 2147483648}, // number differs due to float imprecision
			{name: "min int32", input: math.MinInt32, want: math.MinInt32},
		})

		assertInt64Error(t, []caseInt64[float32]{
			{name: "out of range math.MaxInt64", input: math.MaxInt64 * 1.02}, // because of float imprecision, we have to exceed the min int64 to trigger the error
			{name: "out of range math.MaxUint64", input: math.MaxUint64},
			{name: "out of range math.MinInt64", input: math.MinInt64 * 1.02}, // because of float imprecision, we have to exceed the min int64 to trigger the error
			{name: "out of range math.MaxFloat32", input: math.MaxFloat32},
			{name: "out of range -math.MaxFloat32", input: -math.MaxFloat32},
		})
	})

	t.Run("from float64", func(t *testing.T) {
		assertInt64OK(t, []caseInt64[float64]{
			{name: "zero", input: 0.0, want: 0},
			{name: "rounded value", input: 1.1, want: 1},
			{name: "positive within range", input: 10000.9, want: 10000},
			{name: "max int16", input: math.MaxInt16, want: math.MaxInt16},
			{name: "min int16", input: math.MinInt16, want: math.MinInt16},
			{name: "max int32", input: math.MaxInt32, want: math.MaxInt32},
			{name: "min int32", input: math.MinInt32, want: math.MinInt32},
		})

		assertInt64Error(t, []caseInt64[float64]{
			{name: "out of range math.MaxInt64 + 1", input: math.MaxInt64 * 1.02}, // because of float imprecision, we have to exceed the min int64 to trigger the error
			{name: "out of range math.MaxUint64", input: math.MaxUint64},
			{name: "out of range math.MinInt64", input: math.MinInt64 * 1.02}, // because of float imprecision, we have to exceed the min int64 to trigger the error
			{name: "out of range math.MaxFloat32", input: math.MaxFloat32},
			{name: "out of range -math.MaxFloat32", input: -math.MaxFloat32},
			{name: "out of range math.MaxFloat64", input: math.MaxFloat64},
			{name: "out of range -math.MaxFloat64", input: -math.MaxFloat64},
		})
	})
}

func TestToUint64(t *testing.T) {
	t.Run("from int", func(t *testing.T) {
		assertUint64OK(t, []caseUint64[int]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
		})

		assertUint64Error(t, []caseUint64[int]{
			{name: "negative value", input: -1},
		})
	})

	t.Run("from int8", func(t *testing.T) {
		assertUint64OK(t, []caseUint64[int8]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertUint64Error(t, []caseUint64[int]{
			{name: "negative value", input: -1},
		})
	})

	t.Run("from int16", func(t *testing.T) {
		assertUint64OK(t, []caseUint64[int16]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
		})

		assertUint64Error(t, []caseUint64[int]{
			{name: "negative value", input: -1},
		})
	})

	t.Run("from int32", func(t *testing.T) {
		assertUint64OK(t, []caseUint64[int32]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
		})

		assertUint64Error(t, []caseUint64[int]{
			{name: "negative value", input: -1},
		})
	})

	t.Run("from int64", func(t *testing.T) {
		assertUint64OK(t, []caseUint64[int64]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
		})

		assertUint64Error(t, []caseUint64[int]{
			{name: "negative value", input: -1},
		})
	})

	t.Run("from uint", func(t *testing.T) {
		assertUint64OK(t, []caseUint64[uint]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})
	})

	t.Run("from uint8", func(t *testing.T) {
		assertUint64OK(t, []caseUint64[uint8]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})
	})

	t.Run("from uint16", func(t *testing.T) {
		assertUint64OK(t, []caseUint64[uint16]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})
	})

	t.Run("from uint32", func(t *testing.T) {
		assertUint64OK(t, []caseUint64[uint32]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})
	})

	t.Run("from uint64", func(t *testing.T) {
		assertUint64OK(t, []caseUint64[uint64]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})
	})

	t.Run("from float32", func(t *testing.T) {
		assertUint64OK(t, []caseUint64[float32]{
			{name: "zero", input: 0.0, want: 0},
			{name: "rounded value", input: 1.1, want: 1},
			{name: "positive within range", input: 10000.9, want: 10000},
		})

		assertUint64Error(t, []caseUint64[float32]{
			{name: "negative value", input: -1},
			{name: "out of range max uint64", input: math.MaxUint64 * 1.02}, // because of float imprecision, we have to exceed the min int64 to trigger the error
		})
	})

	t.Run("from float64", func(t *testing.T) {
		assertUint64OK(t, []caseUint64[float64]{
			{name: "zero", input: 0.0, want: 0},
			{name: "rounded value", input: 1.1, want: 1},
			{name: "positive within range", input: 10000.9, want: 10000},
		})

		assertUint64Error(t, []caseUint64[float64]{
			{name: "negative value", input: -1},
			{name: "out of range max uint64", input: math.MaxUint64 * 1.02}, // because of float imprecision, we have to exceed the min int64 to trigger the error
			{name: "out of range max float32", input: math.MaxFloat32},
			{name: "out of range max float64", input: math.MaxFloat64},
		})
	})
}

func TestToInt(t *testing.T) {
	t.Run("from int", func(t *testing.T) {
		assertIntOK(t, []caseInt[int]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
			{name: "negative within range", input: -10000, want: -10000},
		})
	})

	t.Run("from int8", func(t *testing.T) {
		assertIntOK(t, []caseInt[int8]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
			{name: "negative within range", input: -100, want: -100},
		})
	})

	t.Run("from int16", func(t *testing.T) {
		assertIntOK(t, []caseInt[int16]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
			{name: "negative within range", input: -10000, want: -10000},
		})
	})

	t.Run("from int32", func(t *testing.T) {
		assertIntOK(t, []caseInt[int32]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
			{name: "negative within range", input: -10000, want: -10000},
		})
	})

	t.Run("from int64", func(t *testing.T) {
		assertIntOK(t, []caseInt[int64]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
			{name: "negative within range", input: -10000, want: -10000},
		})
	})

	t.Run("from uint", func(t *testing.T) {
		assertIntOK(t, []caseInt[uint]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		// There are extra checks in [TestToInt_64bit]
		// the tests are separated because they cannot work on i386
	})

	t.Run("from uint8", func(t *testing.T) {
		assertIntOK(t, []caseInt[uint8]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})
	})

	t.Run("from uint16", func(t *testing.T) {
		assertIntOK(t, []caseInt[uint16]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})
	})

	t.Run("from uint32", func(t *testing.T) {
		assertIntOK(t, []caseInt[uint32]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})
	})

	t.Run("from uint64", func(t *testing.T) {
		assertIntOK(t, []caseInt[uint64]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertIntError(t, []caseInt[uint64]{
			{name: "positive out of range", input: math.MaxInt64 + 1},
		})
	})

	t.Run("from float32", func(t *testing.T) {
		assertIntOK(t, []caseInt[float32]{
			{name: "zero", input: 0.0, want: 0},
			{name: "rounded value", input: 1.1, want: 1},
			{name: "positive within range", input: 10000.9, want: 10000},
		})

		assertIntError(t, []caseInt[float32]{
			{name: "out of range math.MaxInt64 + 1", input: math.MaxInt64 * 1.02}, // because of float imprecision, we have to exceed the min int64 to trigger the error
			{name: "out of range math.MaxUint64", input: math.MaxUint64},
			{name: "out of range math.MinInt64", input: math.MinInt64 * 1.02}, // because of float imprecision, we have to exceed the min int64 to trigger the error
			{name: "out of range math.MaxFloat32", input: math.MaxFloat32},
			{name: "out of range -math.MaxFloat32", input: -math.MaxFloat32},
		})
	})

	t.Run("from float64", func(t *testing.T) {
		assertIntOK(t, []caseInt[float64]{
			{name: "zero", input: 0.0, want: 0},
			{name: "rounded value", input: 1.1, want: 1},
			{name: "positive within range", input: 10000.9, want: 10000},
			// There are extra checks in [TestToInt_64bit]
			// the tests are separated because they cannot work on i386
		})

		assertIntError(t, []caseInt[float64]{
			{name: "out of range math.MaxInt64 + 1", input: math.MaxInt64 * 1.02}, // because of float imprecision, we have to exceed the min int64 to trigger the error
			{name: "out of range math.MaxUint64", input: math.MaxUint64},
			{name: "out of range math.MinInt64", input: math.MinInt64 * 1.02}, // because of float imprecision, we have to exceed the min int64 to trigger the error
			{name: "out of range math.MaxFloat32", input: math.MaxFloat32},
			{name: "out of range -math.MaxFloat32", input: -math.MaxFloat32},
			{name: "out of range math.MaxFloat64", input: math.MaxFloat64},
			{name: "out of range -math.MaxFloat64", input: -math.MaxFloat64},
		})
	})
}

func TestToUint(t *testing.T) {
	t.Run("from int", func(t *testing.T) {
		assertUintOK(t, []caseUint[int]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
		})

		assertUintError(t, []caseUint[int]{
			{name: "negative value", input: -1},
		})
	})

	t.Run("from int8", func(t *testing.T) {
		assertUintOK(t, []caseUint[int8]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertUintError(t, []caseUint[int8]{
			{name: "negative value", input: -1},
		})
	})

	t.Run("from int16", func(t *testing.T) {
		assertUintOK(t, []caseUint[int16]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
		})

		assertUintError(t, []caseUint[int16]{
			{name: "negative value", input: -1},
		})
	})

	t.Run("from int32", func(t *testing.T) {
		assertUintOK(t, []caseUint[int32]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
		})

		assertUintError(t, []caseUint[int32]{
			{name: "negative value", input: -1},
		})
	})

	t.Run("from int64", func(t *testing.T) {
		assertUintOK(t, []caseUint[int64]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
		})

		assertUintError(t, []caseUint[int64]{
			{name: "negative value", input: -1},
		})
	})

	t.Run("from uint", func(t *testing.T) {
		assertUintOK(t, []caseUint[uint]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})
	})

	t.Run("from uint8", func(t *testing.T) {
		assertUintOK(t, []caseUint[uint8]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})
	})

	t.Run("from uint16", func(t *testing.T) {
		assertUintOK(t, []caseUint[uint16]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})
	})

	t.Run("from uint32", func(t *testing.T) {
		assertUintOK(t, []caseUint[uint32]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})
	})

	t.Run("from uint64", func(t *testing.T) {
		assertUintOK(t, []caseUint[uint64]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})
	})

	t.Run("from float32", func(t *testing.T) {
		assertUintOK(t, []caseUint[float32]{
			{name: "zero", input: 0.0, want: 0},
			{name: "rounded value", input: 1.1, want: 1},
			{name: "positive within range", input: 10000.9, want: 10000},
		})

		assertUint64Error(t, []caseUint64[float32]{
			{name: "negative value", input: -1},
			{name: "out of range max uint64", input: math.MaxUint64 * 1.02}, // because of float imprecision, we have to exceed the min int64 to trigger the error
			{name: "out of range max float32", input: math.MaxFloat32},
		})
	})

	t.Run("from float64", func(t *testing.T) {
		assertUintOK(t, []caseUint[float64]{
			{name: "zero", input: 0.0, want: 0},
			{name: "rounded value", input: 1.1, want: 1},
			{name: "positive within range", input: 10000.9, want: 10000},
		})

		assertUintError(t, []caseUint[float64]{
			{name: "negative value", input: -1},
			{name: "out of range max uint64", input: math.MaxUint64 * 1.02}, // because of float imprecision, we have to exceed the min int64 to trigger the error
			{name: "out of range max float32", input: math.MaxFloat32},
			{name: "out of range max float64", input: math.MaxFloat64},
		})
	})

	type UintAlias uint
	t.Run("from type alias", func(t *testing.T) {
		assertUintOK(t, []caseUint[UintAlias]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})
	})
}

func TestToFloat32(t *testing.T) {
	t.Run("from int", func(t *testing.T) {
		assertFloat32OK(t, []caseFloat32[int]{
			{name: "zero", input: 0, want: 0.0},
			{name: "negative within range", input: -100, want: -100.0},
			{name: "positive within range", input: 100, want: 100.0},
		})
	})

	t.Run("from int8", func(t *testing.T) {
		assertFloat32OK(t, []caseFloat32[int8]{
			{name: "zero", input: 0, want: 0.0},
			{name: "negative within range", input: -100, want: -100.0},
			{name: "positive within range", input: 100, want: 100.0},
		})
	})

	t.Run("from int16", func(t *testing.T) {
		assertFloat32OK(t, []caseFloat32[int16]{
			{name: "zero", input: 0, want: 0.0},
			{name: "negative within range", input: -100, want: -100.0},
			{name: "positive within range", input: 100, want: 100.0},
		})
	})

	t.Run("from int32", func(t *testing.T) {
		assertFloat32OK(t, []caseFloat32[int32]{
			{name: "zero", input: 0, want: 0.0},
			{name: "negative within range", input: -100, want: -100.0},
			{name: "positive within range", input: 100, want: 100.0},
		})
	})

	t.Run("from int64", func(t *testing.T) {
		assertFloat32OK(t, []caseFloat32[int64]{
			{name: "zero", input: 0, want: 0.0},
			{name: "negative within range", input: -100, want: -100.0},
			{name: "positive within range", input: 100, want: 100.0},
		})
	})

	t.Run("from uint", func(t *testing.T) {
		assertFloat32OK(t, []caseFloat32[uint]{
			{name: "zero", input: 0, want: 0.0},
			{name: "positive within range", input: 100, want: 100.0},
		})
	})

	t.Run("from uint8", func(t *testing.T) {
		assertFloat32OK(t, []caseFloat32[uint8]{
			{name: "zero", input: 0, want: 0.0},
			{name: "positive within range", input: 100, want: 100.0},
		})
	})

	t.Run("from uint16", func(t *testing.T) {
		assertFloat32OK(t, []caseFloat32[uint16]{
			{name: "zero", input: 0, want: 0.0},
			{name: "positive within range", input: 100, want: 100.0},
		})
	})

	t.Run("from uint32", func(t *testing.T) {
		assertFloat32OK(t, []caseFloat32[uint32]{
			{name: "zero", input: 0, want: 0.0},
			{name: "positive within range", input: 100, want: 100.0},
		})
	})

	t.Run("from uint64", func(t *testing.T) {
		assertFloat32OK(t, []caseFloat32[uint64]{
			{name: "zero", input: 0, want: 0.0},
			{name: "positive within range", input: 100, want: 100.0},
		})
	})

	t.Run("from float32", func(t *testing.T) {
		assertFloat32OK(t, []caseFloat32[float32]{
			{name: "zero", input: 0.0, want: 0.0},
			{name: "rounded value", input: 1.1, want: 1.1},
			{name: "negative within range", input: -100.9, want: -100.9},
			{name: "positive within range", input: 100.9, want: 100.9},
		})
	})

	t.Run("from float64", func(t *testing.T) {
		assertFloat32OK(t, []caseFloat32[float64]{
			{name: "zero", input: 0.0, want: 0.0},
			{name: "negative zero", input: math.Copysign(0, -1), want: -0},
			{name: "almost zero", input: math.SmallestNonzeroFloat32, want: 1e-45},
			{name: "almost negative zero", input: -math.SmallestNonzeroFloat32, want: -1e-45},
			{name: "negative within range", input: -100.9, want: -100.9},
			{name: "positive within range", input: 100.9, want: 100.9},
			{name: "with imprecision due to conversion", input: 2.67428e+28, want: 2.67428e+28},
		})

		assertFloat32Error(t, []caseFloat32[float64]{
			{name: "out of range max float32", input: math.MaxFloat32 * 1.02},  // because of float imprecision, we have to exceed the min int64 to trigger the error
			{name: "out of range min float32", input: -math.MaxFloat32 * 1.02}, // because of float imprecision, we have to exceed the min int64 to trigger the error
		})
	})
}

func TestToFloat64(t *testing.T) {
	t.Run("from int", func(t *testing.T) {
		assertFloat64OK(t, []caseFloat64[int]{
			{name: "zero", input: 0, want: 0.0},
			{name: "negative within range", input: -100, want: -100.0},
			{name: "positive within range", input: 100, want: 100.0},
		})
	})

	t.Run("from int8", func(t *testing.T) {
		assertFloat64OK(t, []caseFloat64[int8]{
			{name: "zero", input: 0, want: 0.0},
			{name: "negative within range", input: -100, want: -100.0},
			{name: "positive within range", input: 100, want: 100.0},
		})
	})

	t.Run("from int16", func(t *testing.T) {
		assertFloat64OK(t, []caseFloat64[int16]{
			{name: "zero", input: 0, want: 0.0},
			{name: "negative within range", input: -100, want: -100.0},
			{name: "positive within range", input: 100, want: 100.0},
		})
	})

	t.Run("from int32", func(t *testing.T) {
		assertFloat64OK(t, []caseFloat64[int32]{
			{name: "zero", input: 0, want: 0.0},
			{name: "negative within range", input: -100, want: -100.0},
			{name: "positive within range", input: 100, want: 100.0},
		})
	})

	t.Run("from int64", func(t *testing.T) {
		assertFloat64OK(t, []caseFloat64[int64]{
			{name: "zero", input: 0, want: 0.0},
			{name: "negative within range", input: -100, want: -100.0},
			{name: "positive within range", input: 100, want: 100.0},
		})
	})

	t.Run("from uint", func(t *testing.T) {
		assertFloat64OK(t, []caseFloat64[uint]{
			{name: "zero", input: 0, want: 0.0},
			{name: "positive within range", input: 100, want: 100.0},
		})
	})

	t.Run("from uint8", func(t *testing.T) {
		assertFloat64OK(t, []caseFloat64[uint8]{
			{name: "zero", input: 0, want: 0.0},
			{name: "positive within range", input: 100, want: 100.0},
		})
	})

	t.Run("from uint16", func(t *testing.T) {
		assertFloat64OK(t, []caseFloat64[uint16]{
			{name: "zero", input: 0, want: 0.0},
			{name: "positive within range", input: 100, want: 100.0},
		})
	})

	t.Run("from uint32", func(t *testing.T) {
		assertFloat64OK(t, []caseFloat64[uint32]{
			{name: "zero", input: 0, want: 0.0},
			{name: "positive within range", input: 100, want: 100.0},
		})
	})

	t.Run("from uint64", func(t *testing.T) {
		assertFloat64OK(t, []caseFloat64[uint64]{
			{name: "zero", input: 0, want: 0.0},
			{name: "positive within range", input: 100, want: 100.0},
		})
	})

	t.Run("from float32", func(t *testing.T) {
		assertFloat64OK(t, []caseFloat64[float64]{
			{name: "zero", input: 0.0, want: 0.0},
			{name: "rounded value", input: 1.1, want: 1.1},
			{name: "negative within range", input: -100.9, want: -100.9},
			{name: "positive within range", input: 100.9, want: 100.9},
		})
	})

	t.Run("from float64", func(t *testing.T) {
		assertFloat64OK(t, []caseFloat64[float64]{
			{name: "zero", input: 0.0, want: 0.0},
			{name: "negative within range", input: -100.9, want: -100.9},
			{name: "positive within range", input: 100.0, want: 100.0},
		})
	})
}

type MapTest[TypeInput safecast.Input, TypeOutput safecast.Number] struct {
	Input          TypeInput
	ExpectedOutput TypeOutput
	ExpectedError  error
	ErrorContains  string
}

func (mt MapTest[I, O]) Run(t *testing.T) {
	t.Helper()

	// configure a helper to validate there is no panic
	defer func(t *testing.T) {
		t.Helper()

		err := recover()
		if err != nil {
			t.Fatalf("panic with %v", err)
		}
	}(t)

	out, err := safecast.Convert[O](mt.Input)
	if mt.ExpectedError != nil {
		requireErrorIs(t, err, safecast.ErrConversionIssue)
		requireErrorIs(t, err, mt.ExpectedError)

		if mt.ErrorContains != "" {
			requireErrorContains(t, err, mt.ErrorContains)
		}

		return
	}

	assertNoError(t, err)
	assertEqual(t, mt.ExpectedOutput, out)
}

func TestConvert(t *testing.T) {
	t.Run("untyped integer", func(t *testing.T) {
		out, err := safecast.Convert[uint](42)
		assertNoError(t, err)
		assertEqual(t, uint(42), out)
	})

	for name, c := range map[string]TestRunner{
		"int to float32":     MapTest[int, float32]{Input: 42, ExpectedOutput: 42},
		"int to float64":     MapTest[int, float64]{Input: 42, ExpectedOutput: 42},
		"int to int":         MapTest[int, int]{Input: 42, ExpectedOutput: 42},
		"int to int16":       MapTest[int, int16]{Input: 42, ExpectedOutput: 42},
		"int to int32":       MapTest[int, int32]{Input: 42, ExpectedOutput: 42},
		"int to int64":       MapTest[int, int64]{Input: 42, ExpectedOutput: 42},
		"int to int8":        MapTest[int, int8]{Input: 42, ExpectedOutput: 42},
		"int to uint":        MapTest[int, uint]{Input: 42, ExpectedOutput: 42},
		"int to uint16":      MapTest[int, uint16]{Input: 42, ExpectedOutput: 42},
		"int to uint32":      MapTest[int, uint32]{Input: 42, ExpectedOutput: 42},
		"int to uint64":      MapTest[int, uint64]{Input: 42, ExpectedOutput: 42},
		"int to uint8":       MapTest[int, uint8]{Input: 42, ExpectedOutput: 42},
		"int8 to float32":    MapTest[int8, float32]{Input: 42, ExpectedOutput: 42},
		"int8 to float64":    MapTest[int8, float64]{Input: 42, ExpectedOutput: 42},
		"int8 to int":        MapTest[int8, int]{Input: 42, ExpectedOutput: 42},
		"int8 to int16":      MapTest[int8, int16]{Input: 42, ExpectedOutput: 42},
		"int8 to int32":      MapTest[int8, int32]{Input: 42, ExpectedOutput: 42},
		"int8 to int64":      MapTest[int8, int64]{Input: 42, ExpectedOutput: 42},
		"int8 to int8":       MapTest[int8, int8]{Input: 42, ExpectedOutput: 42},
		"int8 to uint":       MapTest[int8, uint]{Input: 42, ExpectedOutput: 42},
		"int8 to uint16":     MapTest[int8, uint16]{Input: 42, ExpectedOutput: 42},
		"int8 to uint32":     MapTest[int8, uint32]{Input: 42, ExpectedOutput: 42},
		"int8 to uint64":     MapTest[int8, uint64]{Input: 42, ExpectedOutput: 42},
		"int8 to uint8":      MapTest[int8, uint8]{Input: 42, ExpectedOutput: 42},
		"int16 to float32":   MapTest[int16, float32]{Input: 42, ExpectedOutput: 42},
		"int16 to float64":   MapTest[int16, float64]{Input: 42, ExpectedOutput: 42},
		"int16 to int":       MapTest[int16, int]{Input: 42, ExpectedOutput: 42},
		"int16 to int16":     MapTest[int16, int16]{Input: 42, ExpectedOutput: 42},
		"int16 to int32":     MapTest[int16, int32]{Input: 42, ExpectedOutput: 42},
		"int16 to int64":     MapTest[int16, int64]{Input: 42, ExpectedOutput: 42},
		"int16 to int8":      MapTest[int16, int8]{Input: 42, ExpectedOutput: 42},
		"int16 to uint":      MapTest[int16, uint]{Input: 42, ExpectedOutput: 42},
		"int16 to uint16":    MapTest[int16, uint16]{Input: 42, ExpectedOutput: 42},
		"int16 to uint32":    MapTest[int16, uint32]{Input: 42, ExpectedOutput: 42},
		"int16 to uint64":    MapTest[int16, uint64]{Input: 42, ExpectedOutput: 42},
		"int16 to uint8":     MapTest[int16, uint8]{Input: 42, ExpectedOutput: 42},
		"int32 to float32":   MapTest[int32, float32]{Input: 42, ExpectedOutput: 42},
		"int32 to float64":   MapTest[int32, float64]{Input: 42, ExpectedOutput: 42},
		"int32 to int":       MapTest[int32, int]{Input: 42, ExpectedOutput: 42},
		"int32 to int16":     MapTest[int32, int16]{Input: 42, ExpectedOutput: 42},
		"int32 to int32":     MapTest[int32, int32]{Input: 42, ExpectedOutput: 42},
		"int32 to int64":     MapTest[int32, int64]{Input: 42, ExpectedOutput: 42},
		"int32 to int8":      MapTest[int32, int8]{Input: 42, ExpectedOutput: 42},
		"int32 to uint":      MapTest[int32, uint]{Input: 42, ExpectedOutput: 42},
		"int32 to uint16":    MapTest[int32, uint16]{Input: 42, ExpectedOutput: 42},
		"int32 to uint32":    MapTest[int32, uint32]{Input: 42, ExpectedOutput: 42},
		"int32 to uint64":    MapTest[int32, uint64]{Input: 42, ExpectedOutput: 42},
		"int32 to uint8":     MapTest[int32, uint8]{Input: 42, ExpectedOutput: 42},
		"int64 to float32":   MapTest[int64, float32]{Input: 42, ExpectedOutput: 42},
		"int64 to float64":   MapTest[int64, float64]{Input: 42, ExpectedOutput: 42},
		"int64 to int":       MapTest[int64, int]{Input: 42, ExpectedOutput: 42},
		"int64 to int16":     MapTest[int64, int16]{Input: 42, ExpectedOutput: 42},
		"int64 to int32":     MapTest[int64, int32]{Input: 42, ExpectedOutput: 42},
		"int64 to int64":     MapTest[int64, int64]{Input: 42, ExpectedOutput: 42},
		"int64 to int8":      MapTest[int64, int8]{Input: 42, ExpectedOutput: 42},
		"int64 to uint":      MapTest[int64, uint]{Input: 42, ExpectedOutput: 42},
		"int64 to uint16":    MapTest[int64, uint16]{Input: 42, ExpectedOutput: 42},
		"int64 to uint32":    MapTest[int64, uint32]{Input: 42, ExpectedOutput: 42},
		"int64 to uint64":    MapTest[int64, uint64]{Input: 42, ExpectedOutput: 42},
		"int64 to uint8":     MapTest[int64, uint8]{Input: 42, ExpectedOutput: 42},
		"uint to float32":    MapTest[uint, float32]{Input: 42, ExpectedOutput: 42},
		"uint to float64":    MapTest[uint, float64]{Input: 42, ExpectedOutput: 42},
		"uint to int":        MapTest[uint, int]{Input: 42, ExpectedOutput: 42},
		"uint to int16":      MapTest[uint, int16]{Input: 42, ExpectedOutput: 42},
		"uint to int32":      MapTest[uint, int32]{Input: 42, ExpectedOutput: 42},
		"uint to int64":      MapTest[uint, int64]{Input: 42, ExpectedOutput: 42},
		"uint to int8":       MapTest[uint, int8]{Input: 42, ExpectedOutput: 42},
		"uint to uint":       MapTest[uint, uint]{Input: 42, ExpectedOutput: 42},
		"uint to uint16":     MapTest[uint, uint16]{Input: 42, ExpectedOutput: 42},
		"uint to uint32":     MapTest[uint, uint32]{Input: 42, ExpectedOutput: 42},
		"uint to uint64":     MapTest[uint, uint64]{Input: 42, ExpectedOutput: 42},
		"uint to uint8":      MapTest[uint, uint8]{Input: 42, ExpectedOutput: 42},
		"uint8 to float32":   MapTest[uint8, float32]{Input: 42, ExpectedOutput: 42},
		"uint8 to float64":   MapTest[uint8, float64]{Input: 42, ExpectedOutput: 42},
		"uint8 to int":       MapTest[uint8, int]{Input: 42, ExpectedOutput: 42},
		"uint8 to int16":     MapTest[uint8, int16]{Input: 42, ExpectedOutput: 42},
		"uint8 to int32":     MapTest[uint8, int32]{Input: 42, ExpectedOutput: 42},
		"uint8 to int64":     MapTest[uint8, int64]{Input: 42, ExpectedOutput: 42},
		"uint8 to int8":      MapTest[uint8, int8]{Input: 42, ExpectedOutput: 42},
		"uint8 to uint":      MapTest[uint8, uint]{Input: 42, ExpectedOutput: 42},
		"uint8 to uint16":    MapTest[uint8, uint16]{Input: 42, ExpectedOutput: 42},
		"uint8 to uint32":    MapTest[uint8, uint32]{Input: 42, ExpectedOutput: 42},
		"uint8 to uint64":    MapTest[uint8, uint64]{Input: 42, ExpectedOutput: 42},
		"uint8 to uint8":     MapTest[uint8, uint8]{Input: 42, ExpectedOutput: 42},
		"uint16 to float32":  MapTest[uint16, float32]{Input: 42, ExpectedOutput: 42},
		"uint16 to float64":  MapTest[uint16, float64]{Input: 42, ExpectedOutput: 42},
		"uint16 to int":      MapTest[uint16, int]{Input: 42, ExpectedOutput: 42},
		"uint16 to int16":    MapTest[uint16, int16]{Input: 42, ExpectedOutput: 42},
		"uint16 to int32":    MapTest[uint16, int32]{Input: 42, ExpectedOutput: 42},
		"uint16 to int64":    MapTest[uint16, int64]{Input: 42, ExpectedOutput: 42},
		"uint16 to int8":     MapTest[uint16, int8]{Input: 42, ExpectedOutput: 42},
		"uint16 to uint":     MapTest[uint16, uint]{Input: 42, ExpectedOutput: 42},
		"uint16 to uint16":   MapTest[uint16, uint16]{Input: 42, ExpectedOutput: 42},
		"uint16 to uint32":   MapTest[uint16, uint32]{Input: 42, ExpectedOutput: 42},
		"uint16 to uint64":   MapTest[uint16, uint64]{Input: 42, ExpectedOutput: 42},
		"uint16 to uint8":    MapTest[uint16, uint8]{Input: 42, ExpectedOutput: 42},
		"uint32 to float32":  MapTest[uint32, float32]{Input: 42, ExpectedOutput: 42},
		"uint32 to float64":  MapTest[uint32, float64]{Input: 42, ExpectedOutput: 42},
		"uint32 to int":      MapTest[uint32, int]{Input: 42, ExpectedOutput: 42},
		"uint32 to int16":    MapTest[uint32, int16]{Input: 42, ExpectedOutput: 42},
		"uint32 to int32":    MapTest[uint32, int32]{Input: 42, ExpectedOutput: 42},
		"uint32 to int64":    MapTest[uint32, int64]{Input: 42, ExpectedOutput: 42},
		"uint32 to int8":     MapTest[uint32, int8]{Input: 42, ExpectedOutput: 42},
		"uint32 to uint":     MapTest[uint32, uint]{Input: 42, ExpectedOutput: 42},
		"uint32 to uint16":   MapTest[uint32, uint16]{Input: 42, ExpectedOutput: 42},
		"uint32 to uint32":   MapTest[uint32, uint32]{Input: 42, ExpectedOutput: 42},
		"uint32 to uint64":   MapTest[uint32, uint64]{Input: 42, ExpectedOutput: 42},
		"uint32 to uint8":    MapTest[uint32, uint8]{Input: 42, ExpectedOutput: 42},
		"uint64 to float32":  MapTest[uint64, float32]{Input: 42, ExpectedOutput: 42},
		"uint64 to float64":  MapTest[uint64, float64]{Input: 42, ExpectedOutput: 42},
		"uint64 to int":      MapTest[uint64, int]{Input: 42, ExpectedOutput: 42},
		"uint64 to int16":    MapTest[uint64, int16]{Input: 42, ExpectedOutput: 42},
		"uint64 to int32":    MapTest[uint64, int32]{Input: 42, ExpectedOutput: 42},
		"uint64 to int64":    MapTest[uint64, int64]{Input: 42, ExpectedOutput: 42},
		"uint64 to int8":     MapTest[uint64, int8]{Input: 42, ExpectedOutput: 42},
		"uint64 to uint":     MapTest[uint64, uint]{Input: 42, ExpectedOutput: 42},
		"uint64 to uint16":   MapTest[uint64, uint16]{Input: 42, ExpectedOutput: 42},
		"uint64 to uint32":   MapTest[uint64, uint32]{Input: 42, ExpectedOutput: 42},
		"uint64 to uint64":   MapTest[uint64, uint64]{Input: 42, ExpectedOutput: 42},
		"uint64 to uint8":    MapTest[uint64, uint8]{Input: 42, ExpectedOutput: 42},
		"float32 to int":     MapTest[float32, int]{Input: 42, ExpectedOutput: 42},
		"float32 to int8":    MapTest[float32, int8]{Input: 42, ExpectedOutput: 42},
		"float32 to int16":   MapTest[float32, int16]{Input: 42, ExpectedOutput: 42},
		"float32 to int32":   MapTest[float32, int32]{Input: 42, ExpectedOutput: 42},
		"float32 to int64":   MapTest[float32, int64]{Input: 42, ExpectedOutput: 42},
		"float32 to uint":    MapTest[float32, uint]{Input: 42, ExpectedOutput: 42},
		"float32 to uint8":   MapTest[float32, uint8]{Input: 42, ExpectedOutput: 42},
		"float32 to uint16":  MapTest[float32, uint16]{Input: 42, ExpectedOutput: 42},
		"float32 to uint32":  MapTest[float32, uint32]{Input: 42, ExpectedOutput: 42},
		"float32 to uint64":  MapTest[float32, uint64]{Input: 42, ExpectedOutput: 42},
		"float32 to float32": MapTest[float32, float32]{Input: 42, ExpectedOutput: 42},
		"float32 to float64": MapTest[float32, float64]{Input: 42, ExpectedOutput: 42},
		"float64 to int":     MapTest[float64, int]{Input: 42, ExpectedOutput: 42},
		"float64 to int8":    MapTest[float64, int8]{Input: 42, ExpectedOutput: 42},
		"float64 to int16":   MapTest[float64, int16]{Input: 42, ExpectedOutput: 42},
		"float64 to int32":   MapTest[float64, int32]{Input: 42, ExpectedOutput: 42},
		"float64 to int64":   MapTest[float64, int64]{Input: 42, ExpectedOutput: 42},
		"float64 to uint":    MapTest[float64, uint]{Input: 42, ExpectedOutput: 42},
		"float64 to uint8":   MapTest[float64, uint8]{Input: 42, ExpectedOutput: 42},
		"float64 to uint16":  MapTest[float64, uint16]{Input: 42, ExpectedOutput: 42},
		"float64 to uint32":  MapTest[float64, uint32]{Input: 42, ExpectedOutput: 42},
		"float64 to uint64":  MapTest[float64, uint64]{Input: 42, ExpectedOutput: 42},
		"float64 to float32": MapTest[float64, float32]{Input: 42, ExpectedOutput: 42},
		"float64 to float64": MapTest[float64, float64]{Input: 42, ExpectedOutput: 42},
	} {
		t.Run(name, func(t *testing.T) {
			c.Run(t)
		})
	}

	for name, c := range map[string]TestRunner{
		"string integer":              MapTest[string, uint]{Input: "42", ExpectedOutput: 42},
		"string with spaces":          MapTest[string, uint]{Input: "42 ", ExpectedOutput: 42},
		"string float":                MapTest[string, uint]{Input: "42.0", ExpectedOutput: 42},
		"string true":                 MapTest[string, uint]{Input: "true", ExpectedOutput: 1},
		"string false":                MapTest[string, uint]{Input: "false", ExpectedOutput: 0},
		"string 10_0":                 MapTest[string, uint]{Input: "10_0", ExpectedOutput: 100},
		"string binary":               MapTest[string, uint]{Input: "0b101010", ExpectedOutput: 42},
		"string short octal notation": MapTest[string, uint]{Input: "042", ExpectedOutput: 34},
		"string octal":                MapTest[string, uint]{Input: "0o42", ExpectedOutput: 34},
		"string hexadecimal":          MapTest[string, uint]{Input: "0x42", ExpectedOutput: 66},
		"boolean true":                MapTest[bool, uint]{Input: true, ExpectedOutput: 1},
		"boolean false":               MapTest[bool, uint]{Input: false, ExpectedOutput: 0},

		"empty string": MapTest[string, uint]{
			Input:         "",
			ExpectedError: safecast.ErrStringConversion,
			ErrorContains: "cannot convert from string  to uint",
		},
		"simple space": MapTest[string, uint]{
			Input:         " ",
			ExpectedError: safecast.ErrStringConversion,
			ErrorContains: "cannot convert from string   to uint",
		},
		"simple dot": MapTest[string, uint]{
			Input:         ".",
			ExpectedError: safecast.ErrStringConversion,
			ErrorContains: "cannot convert from string . to uint"},
		"simple dash": MapTest[string, uint]{
			Input:         "-",
			ExpectedError: safecast.ErrStringConversion,
			ErrorContains: "cannot convert from string - to uint"},
		"invalid string": MapTest[string, uint]{
			Input:         "abc",
			ExpectedError: safecast.ErrStringConversion,
			ErrorContains: `cannot convert from string abc to uint`},
		"invalid string with dot": MapTest[string, uint]{
			Input:         "ab.c",
			ExpectedError: safecast.ErrStringConversion,
			ErrorContains: `cannot convert from string ab.c to uint`},
		"strings with leading +": MapTest[string, uint]{
			Input:         "+42",
			ExpectedError: safecast.ErrStringConversion,
			ErrorContains: "cannot convert from string +42 to uint",
		},
		"invalid string multiple leading dashes": MapTest[string, uint]{Input: "--42", ExpectedError: safecast.ErrStringConversion},
		"invalid string with dash":               MapTest[string, uint]{Input: "-abc", ExpectedError: safecast.ErrStringConversion},
		"invalid string with dash and dot":       MapTest[string, uint]{Input: "-ab.c", ExpectedError: safecast.ErrStringConversion},
	} {
		t.Run(name, func(t *testing.T) {
			c.Run(t)
		})
	}

	negativeZero := math.Copysign(0, -1)
	t.Run("convert to float32 near zero", func(t *testing.T) {
		for name, tt := range map[string]TestRunner{
			"negative untyped zero":              MapTest[float64, float32]{Input: negativeZero, ExpectedOutput: float32(negativeZero)},
			"smallest positive non-zero float32": MapTest[float64, float32]{Input: math.SmallestNonzeroFloat32, ExpectedOutput: 1e-45},
			"smallest negative non-zero float32": MapTest[float64, float32]{Input: -math.SmallestNonzeroFloat32, ExpectedOutput: -1e-45},
			"smallest positive non-zero float64": MapTest[float64, float32]{Input: math.SmallestNonzeroFloat64, ExpectedOutput: 4.9e-324},
			"smallest negative non-zero float64": MapTest[float64, float32]{Input: -math.SmallestNonzeroFloat64, ExpectedOutput: -4.9e-324},
		} {
			t.Run(name, func(t *testing.T) {
				tt.Run(t)
			})
		}
	})

	t.Run("convert to float64 near zero", func(t *testing.T) {
		for name, tt := range map[string]TestRunner{
			"negative untyped zero":              MapTest[float64, float64]{Input: negativeZero, ExpectedOutput: negativeZero},
			"smallest positive non-zero float32": MapTest[float64, float64]{Input: math.SmallestNonzeroFloat32, ExpectedOutput: 1.401298464324817e-45},
			"smallest negative non-zero float32": MapTest[float64, float64]{Input: -math.SmallestNonzeroFloat32, ExpectedOutput: -1.401298464324817e-45},
			"smallest positive non-zero float64": MapTest[float64, float64]{Input: math.SmallestNonzeroFloat64, ExpectedOutput: 4.9e-324},
			"smallest negative non-zero float64": MapTest[float64, float64]{Input: -math.SmallestNonzeroFloat64, ExpectedOutput: -4.9e-324},
		} {
			t.Run(name, func(t *testing.T) {
				tt.Run(t)
			})
		}
	})

	for name, c := range map[string]TestRunner{
		"math.Inf to float64": MapTest[float64, float64]{
			Input:         math.Inf(1),
			ExpectedError: safecast.ErrExceedMaximumValue,
			ErrorContains: "+Inf (float64) is greater than 1.7976931348623157e+308",
		},
		"-math.Inf to float64": MapTest[float64, float64]{
			Input:         math.Inf(-1),
			ExpectedError: safecast.ErrExceedMinimumValue,
			ErrorContains: "-Inf (float64) is less than -1.7976931348623157e+308",
		},
		"math.NaN to float64": MapTest[float64, float64]{
			Input:         math.NaN(),
			ExpectedError: safecast.ErrUnsupportedConversion,
			ErrorContains: "NaN (float64) is not supported",
		},
		"math.Inf to float32": MapTest[float64, float32]{
			Input:         math.Inf(1),
			ExpectedError: safecast.ErrExceedMaximumValue,
			ErrorContains: "+Inf (float64) is greater than 3.4028235e+38",
		},
		"-math.Inf to float32": MapTest[float64, float32]{
			Input:         math.Inf(-1),
			ExpectedError: safecast.ErrExceedMinimumValue,
			ErrorContains: "-Inf (float64) is less than -3.4028235e+38",
		},
		"math.NaN to float32": MapTest[float64, float32]{
			Input:         math.NaN(),
			ExpectedError: safecast.ErrUnsupportedConversion,
			ErrorContains: "NaN (float64) is not supported",
		},
		"upper bound overflows for int": MapTest[uint, int]{
			Input:         uint(math.MaxInt + 1),
			ExpectedError: safecast.ErrExceedMaximumValue,
		},
		"upper bound overflows for int8": MapTest[uint, int8]{
			Input:         uint(math.MaxInt8 + 1),
			ExpectedError: safecast.ErrExceedMaximumValue,
		},
		"upper bound overflows for int16": MapTest[uint, int16]{
			Input:         uint(math.MaxInt16 + 1),
			ExpectedError: safecast.ErrExceedMaximumValue,
		},
		"upper bound overflows for int32": MapTest[uint, int32]{
			Input:         uint(math.MaxInt32 + 1),
			ExpectedError: safecast.ErrExceedMaximumValue,
		},
		"upper bound overflows for int64": MapTest[float64, int64]{
			Input:         float64(math.MaxInt64 * 1.01), // using float64 here avoid issue when testing on 32-bit systems
			ExpectedError: safecast.ErrExceedMaximumValue,
		},
		"upper bound overflows for uint": MapTest[float64, uint]{
			Input:         float64(math.MaxUint * 1.01), // using float64 here avoid issue when testing on 32-bit systems
			ExpectedError: safecast.ErrExceedMaximumValue,
		},
		"upper bound overflows for uint8": MapTest[uint, uint8]{
			Input:         uint(math.MaxUint8 + 1),
			ExpectedError: safecast.ErrExceedMaximumValue,
		},
		"upper bound overflows for uint16": MapTest[uint, uint16]{
			Input:         uint(math.MaxUint16 + 1),
			ExpectedError: safecast.ErrExceedMaximumValue,
		},
		"upper bound overflows for uint32": MapTest[float64, uint32]{
			Input:         float64(math.MaxUint32 * 1.01),
			ExpectedError: safecast.ErrExceedMaximumValue,
		},
		"upper bound overflows for uint64": MapTest[float64, uint64]{
			Input:         float64(math.MaxUint64 * 1.01),
			ExpectedError: safecast.ErrExceedMaximumValue,
		},

		"upper bound overflows for int string": MapTest[string, int]{
			Input:         "9223372036854775808", // math.MaxInt64 + 1
			ExpectedError: safecast.ErrExceedMaximumValue,
		},
		"upper bound overflows for int8 string": MapTest[string, int8]{
			Input:         "129", // math.MaxInt8 + 1
			ExpectedError: safecast.ErrExceedMaximumValue,
		},
		"upper bound overflows for int16 string": MapTest[string, int16]{
			Input:         "32769", // math.MaxInt16 + 1
			ExpectedError: safecast.ErrExceedMaximumValue,
		},
		"upper bound overflows for int32 string": MapTest[string, int32]{
			Input:         "2147483648", // math.MaxInt32 + 1
			ExpectedError: safecast.ErrExceedMaximumValue,
		},
		"upper bound overflows for int64 string": MapTest[string, int64]{
			Input:         "9223372036854775808", // math.MaxInt64 + 1
			ExpectedError: safecast.ErrExceedMaximumValue,
		},

		"upper bound overflows for int64 string overflow": MapTest[string, int64]{
			Input:         "123456789012345678901234567890", // more characters than math.MaxInt64 represented as string
			ExpectedError: safecast.ErrExceedMaximumValue,
		},
	} {
		t.Run(name, func(t *testing.T) {
			c.Run(t)
		})
	}

	for name, c := range map[string]TestRunner{
		"lower bound overflows for int": MapTest[float64, int]{
			Input:         float64(math.MinInt * 1.01), // the float64 conversion is used to avoid overflow on 32-bit
			ExpectedError: safecast.ErrExceedMinimumValue,
		},
		"lower bound overflows for int8": MapTest[int, int8]{
			Input:         math.MinInt8 - 1,
			ExpectedError: safecast.ErrExceedMinimumValue,
		},
		"lower bound overflows for int16": MapTest[int, int16]{
			Input:         math.MinInt16 - 1,
			ExpectedError: safecast.ErrExceedMinimumValue,
		},
		"lower bound overflows for int32": MapTest[float64, int32]{
			Input:         float64(math.MinInt32 - 1), // the float64 conversion is used to avoid overflow on 32-bit,
			ExpectedError: safecast.ErrExceedMinimumValue,
		},
		"lower bound overflows for int64": MapTest[float64, int64]{
			Input:         float64(math.MinInt64 * 1.01), // the float64 conversion is used to avoid overflow on 32-bit
			ExpectedError: safecast.ErrExceedMinimumValue,
		},

		"lower bound overflows for float32": MapTest[float64, float32]{
			Input:         -float64(math.MaxFloat32 * 1.01),
			ExpectedError: safecast.ErrExceedMinimumValue,
			ErrorContains: "float32", // this increases the test coverage of the specific error message for float32
		},

		// Note: float64 cannot overflow

		"negative overflows uint": MapTest[int, uint]{
			Input:         -42,
			ExpectedError: safecast.ErrExceedMinimumValue,
		},

		"negative overflows uint8": MapTest[int, uint8]{
			Input:         -42,
			ExpectedError: safecast.ErrExceedMinimumValue,
		},

		"negative overflows uint16": MapTest[int, uint16]{
			Input:         -42,
			ExpectedError: safecast.ErrExceedMinimumValue,
		},

		"negative overflows uint32": MapTest[int, uint32]{
			Input:         -42,
			ExpectedError: safecast.ErrExceedMinimumValue,
		},

		"negative overflows uint64": MapTest[int, uint64]{
			Input:         -42,
			ExpectedError: safecast.ErrExceedMinimumValue,
		},

		"lower bound overflows int from string": MapTest[string, int]{
			Input:         "-9223372036854775809", // math.MinInt64 - 1
			ExpectedError: safecast.ErrExceedMinimumValue,
		},

		"lower bound overflows int8 from string": MapTest[string, int8]{
			Input:         "-129", // math.MinInt8 - 1
			ExpectedError: safecast.ErrExceedMinimumValue,
		},

		"lower bound overflows int16 from string": MapTest[string, int16]{
			Input:         "-32769", // math.MinInt16 - 1
			ExpectedError: safecast.ErrExceedMinimumValue,
		},

		"lower bound overflows int32 from string": MapTest[string, int32]{
			Input:         "-2147483649", // math.MinInt32 - 1
			ExpectedError: safecast.ErrExceedMinimumValue,
		},

		"lower bound overflows int64 from string": MapTest[string, int64]{
			Input:         "-9223372036854775809", // math.MinInt64 - 1
			ExpectedError: safecast.ErrExceedMinimumValue,
		},

		"lower bound overflows int64 from string overflow": MapTest[string, int64]{
			Input:         "-123456789012345678901234567890", // more characters than math.MinInt64 represented as string
			ExpectedError: safecast.ErrExceedMinimumValue,
		},

		"negative string overflows uint": MapTest[string, uint]{
			Input:         "-1",
			ExpectedError: safecast.ErrExceedMinimumValue,
		},

		"negative string overflows uint8": MapTest[string, uint8]{
			Input:         "-1",
			ExpectedError: safecast.ErrExceedMinimumValue,
		},

		"negative string overflows uint16": MapTest[string, uint16]{
			Input:         "-1",
			ExpectedError: safecast.ErrExceedMinimumValue,
		},

		"negative string overflows uint32": MapTest[string, uint32]{
			Input:         "-1",
			ExpectedError: safecast.ErrExceedMinimumValue,
		},

		"negative string overflows uint64": MapTest[string, uint64]{
			Input:         "-1",
			ExpectedError: safecast.ErrExceedMinimumValue,
		},
	} {
		t.Run(name, func(t *testing.T) {
			c.Run(t)
		})
	}

	t.Run("aliases", func(t *testing.T) {
		// Type aliases are handled separately

		type (
			// UintSimpleAlias is an alias
			UintSimpleAlias = uint

			// UintTypeAlias is a type alias
			UintTypeAlias uint

			// StringAlias is an alias
			StringAlias = string

			// StringTypeAlias is a type alias
			StringTypeAlias string

			// BoolAlias is an alias
			BoolAlias = bool

			// BoolTypeAlias is a type alias
			BoolTypeAlias bool
		)

		for name, c := range map[string]TestRunner{
			"integer simple alias": MapTest[UintSimpleAlias, int8]{
				Input:          UintSimpleAlias(42),
				ExpectedOutput: int8(42),
			},

			"integer type alias": MapTest[UintTypeAlias, int8]{
				Input:          UintTypeAlias(42),
				ExpectedOutput: int8(42),
			},

			"string simple alias": MapTest[StringAlias, int8]{
				Input:          StringAlias("42"),
				ExpectedOutput: int8(42),
			},

			"string type alias": MapTest[StringTypeAlias, int8]{
				Input:          StringTypeAlias("42"),
				ExpectedOutput: int8(42),
			},

			"bool simple alias": MapTest[BoolAlias, int8]{
				Input:          BoolAlias(true),
				ExpectedOutput: int8(1),
			},

			"bool type alias": MapTest[BoolTypeAlias, int8]{
				Input:          BoolTypeAlias(true),
				ExpectedOutput: int8(1),
			},

			"simple alias overflows for int8": MapTest[UintSimpleAlias, int8]{
				Input:         UintSimpleAlias(255),
				ExpectedError: safecast.ErrExceedMaximumValue,
				ErrorContains: "255 (uint) is greater than 127 (int8)",
			},

			"type alias overflows for int8": MapTest[UintTypeAlias, int8]{
				Input:         UintTypeAlias(255),
				ExpectedError: safecast.ErrExceedMaximumValue,
				ErrorContains: "255 (uint) is greater than 127 (int8)",
			},
		} {
			t.Run(name, func(t *testing.T) {
				c.Run(t)
			})
		}
	})
}

type MapMustConvertTest[TypeInput safecast.Input, TypeOutput safecast.Number] struct {
	Input          TypeInput
	ExpectedOutput TypeOutput
	ExpectedError  error
}

func (mt MapMustConvertTest[I, O]) Run(t *testing.T) {
	t.Helper()

	var out O
	defer func(t *testing.T) {
		t.Helper()

		assertEqual(t, mt.ExpectedOutput, out)

		r := recover()

		if mt.ExpectedError == nil && r == nil {
			return
		}

		if r == nil {
			t.Fatal("did not panic")
		}

		err, ok := r.(error)
		if !ok {
			t.Fatalf("panic value is not an error: %v", r)
		}

		if !errors.Is(err, safecast.ErrConversionIssue) {
			t.Fatalf("panic with unexpected error: %v", err)
		}

		if !errors.Is(err, mt.ExpectedError) {
			t.Fatalf("panic with unexpected error: %v", err)
		}
	}(t)

	out = safecast.MustConvert[O](mt.Input)
}

func TestMustConvert(t *testing.T) {
	// [TestConvert] tested all the cases
	// here we are simply checking that the function panic on errors

	t.Run("panic on error", func(t *testing.T) {
		for name, tt := range map[string]TestRunner{
			"negative": MapMustConvertTest[int, uint8]{Input: -1, ExpectedError: safecast.ErrExceedMinimumValue},
			"overflow": MapMustConvertTest[int, uint8]{Input: math.MaxInt, ExpectedError: safecast.ErrExceedMaximumValue},
			"string":   MapMustConvertTest[string, uint8]{Input: "cats", ExpectedError: safecast.ErrStringConversion},
		} {
			t.Run(name, func(t *testing.T) {
				tt.Run(t)
			})
		}
	})

	t.Run("no panic", func(t *testing.T) {
		for name, tt := range map[string]TestRunner{
			"number": MapMustConvertTest[int, uint8]{Input: 42, ExpectedOutput: 42},
			"string": MapMustConvertTest[string, uint8]{Input: "42", ExpectedOutput: 42},
			"float":  MapMustConvertTest[float64, uint8]{Input: 42.0, ExpectedOutput: 42},
			"octal":  MapMustConvertTest[string, uint8]{Input: "0o52", ExpectedOutput: 42},
		} {
			t.Run(name, func(t *testing.T) {
				tt.Run(t)
			})
		}
	})
}

type MapRequireConvertTest[TypeInput safecast.Input, TypeOutput safecast.Number] struct {
	Input               TypeInput
	ExpectedOutput      TypeOutput
	ExpectedTestFailure bool
}

func (mt MapRequireConvertTest[I, O]) Run(t *testing.T) {
	t.Helper()

	// We need to use a fake testing.T to avoid the test failing when we expect a failure
	fakeT := new(mockTestingT)
	out := safecast.RequireConvert[O](fakeT, mt.Input)
	assertEqual(t, mt.ExpectedOutput, out)

	if mt.ExpectedTestFailure == fakeT.Failed() {
		return // test passed as expected
	}

	if mt.ExpectedTestFailure {
		t.Fatal("test should have failed, but it did not")
	}
	t.Error("test should not have failed, but it did")
	// we know that the test failed, so we use the real testing.T to report the error
	_ = safecast.RequireConvert[O](t, mt.Input)
}

func TestRequireConvert(t *testing.T) {
	// [TestConvert] tested all the cases
	// here we are simply checking that the test fails on errors

	t.Run("no conversion error", func(t *testing.T) {
		for name, tt := range map[string]TestRunner{
			"number": MapRequireConvertTest[int, uint8]{Input: 42, ExpectedOutput: 42},
			"string": MapRequireConvertTest[string, uint8]{Input: "42", ExpectedOutput: 42},
			"float":  MapRequireConvertTest[float64, uint8]{Input: 42.0, ExpectedOutput: 42},
			"octal":  MapRequireConvertTest[string, uint8]{Input: "0o52", ExpectedOutput: 42},
		} {
			t.Run(name, func(t *testing.T) {
				tt.Run(t)
			})
		}
	})

	t.Run("test fail on error", func(t *testing.T) {
		for name, tt := range map[string]TestRunner{
			"negative": MapRequireConvertTest[int, uint8]{Input: -1, ExpectedTestFailure: true},
			"overflow": MapRequireConvertTest[int, uint8]{Input: math.MaxInt, ExpectedTestFailure: true},
			"string":   MapRequireConvertTest[string, uint8]{Input: "cats", ExpectedTestFailure: true},
		} {
			t.Run(name, func(t *testing.T) {
				tt.Run(t)
			})
		}
	})
}

// This example shows how to use [safecast.RequireConvert] to convert a value in a test
//
// Here t is a mock implementation of the [safecast.TestingT] interface
// you can use any type that implements this interface, such as [*testing.T], [*testing.B], or [*testing.F]
func ExampleRequireConvert_success() {
	t := new(mockTestingT)

	out := safecast.RequireConvert[uint8](t, 42.0)
	fmt.Println(out, t.Failed())

	// Output:
	// 42 false
}

// mockTestingTExample is a mock implementation of the [safecast.TestingT] interface
// its purpose is to display the error message in the [ExampleRequireConvert_failure].
type mockTestingTExample struct {
	failed bool
}

func (m *mockTestingTExample) Helper() {}

func (m *mockTestingTExample) Fatal(args ...any) {
	// here we simulate the behavior of [testing.T.Fatal] so it can be shown in the example output
	fmt.Printf("--- FAIL:\n\t")
	fmt.Println(args...)
	m.failed = true
}

func (m mockTestingTExample) Failed() bool {
	return m.failed
}

// This example shows how [safecast.RequireConvert] reacts when the conversion fails
//
// An error is raised via [testing.T.Fatal] if the conversion fails.
func ExampleRequireConvert_failure() {
	t := new(mockTestingTExample) // use *testing.T, *testing.B, or *testing.F here

	_ = safecast.RequireConvert[uint8](t, "foo")
	// here an error is raised via [testing.T.Fatal]

	// Output:
	// --- FAIL:
	// 	conversion issue: cannot convert from string foo to uint8 (base auto-detection)
}
