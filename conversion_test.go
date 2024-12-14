package safecast_test

// The tests in conversion_test.go are the ones that are not architecture dependent
// The tests in conversion_64bit_test.go complete them for 64-bit systems
//
// The architecture dependent file covers the fact, you can reach a higher value with int and uint
// on 64-bit systems, but you will get a compile error on 32-bit.
// This is why it needs to be tested in an architecture dependent way.

import (
	"fmt"
	"math"
	"testing"

	"github.com/ccoveille/go-safecast"
)

type anyStringer struct {
	string
}

func (anyStringer) String() string {
	return "42"
}

type anyError struct {
	string
}

func (anyError) Error() string {
	return "42"
}

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

func TestConvert(t *testing.T) {
	negativeZero := math.Copysign(0, -1)

	type helper func(input any) (any, error)

	convertUint := func(input any) (any, error) {
		return safecast.Convert[uint](input)
	}

	convertUint8 := func(input any) (any, error) {
		return safecast.Convert[uint8](input)
	}

	convertUint16 := func(input any) (any, error) {
		return safecast.Convert[uint16](input)
	}

	convertUint32 := func(input any) (any, error) {
		return safecast.Convert[uint32](input)
	}

	convertUint64 := func(input any) (any, error) {
		return safecast.Convert[uint64](input)
	}

	convertInt := func(input any) (any, error) {
		return safecast.Convert[int](input)
	}

	convertInt8 := func(input any) (any, error) {
		return safecast.Convert[int8](input)
	}

	convertInt16 := func(input any) (any, error) {
		return safecast.Convert[int16](input)
	}

	convertInt32 := func(input any) (any, error) {
		return safecast.Convert[int32](input)
	}

	convertInt64 := func(input any) (any, error) {
		return safecast.Convert[int64](input)
	}

	convertFloat32 := func(input any) (any, error) {
		return safecast.Convert[float32](input)
	}

	convertFloat64 := func(input any) (any, error) {
		return safecast.Convert[float64](input)
	}

	unsignedConverters := map[string]helper{
		"uint":   convertUint,
		"uint8":  convertUint8,
		"uint16": convertUint16,
		"uint32": convertUint32,
		"uint64": convertUint64,
	}

	allConverters := map[string]helper{
		"uint":    convertUint,
		"uint8":   convertUint8,
		"uint16":  convertUint16,
		"uint32":  convertUint32,
		"uint64":  convertUint64,
		"int":     convertInt,
		"int8":    convertInt8,
		"int16":   convertInt16,
		"int32":   convertInt32,
		"int64":   convertInt64,
		"float32": convertFloat32,
		"float64": convertFloat64,
	}

	for name, converter := range allConverters {
		t.Run(fmt.Sprintf("convert to %s", name), func(t *testing.T) {
			for name, tt := range map[string]struct {
				input any
				want  any
			}{
				"untyped int zero": {input: 0, want: 0},

				"positive untyped int within range": {input: 42, want: 42},

				"int zero":                  {input: int(0), want: 0},
				"positive int within range": {input: int(42), want: 42},

				"int8 zero":                  {input: int8(0), want: 0},
				"positive int8 within range": {input: int8(42), want: 42},

				"int16 zero":                  {input: int16(0), want: 0},
				"positive int16 within range": {input: int16(42), want: 42},

				"int32 zero":                  {input: int32(0), want: 0},
				"positive int32 within range": {input: int32(42), want: 42},

				"int64 zero":                  {input: int64(0), want: 0},
				"positive int64 within range": {input: int64(42), want: 42},

				"uint zero":         {input: uint(0), want: 0},
				"uint within range": {input: uint(42), want: 42},

				"uint8 zero":         {input: uint8(0), want: 0},
				"uint8 within range": {input: uint8(42), want: 42},

				"uint16 zero":         {input: uint16(0), want: 0},
				"uint16 within range": {input: uint16(42), want: 42},

				"uint32 zero":         {input: uint32(0), want: 0},
				"uint32 within range": {input: uint32(42), want: 42},

				"uint64 zero":         {input: uint64(0), want: 0},
				"uint64 within range": {input: uint64(42), want: 42},

				"float32 zero":                  {input: float32(0), want: 0},
				"positive float32 within range": {input: float32(42.0), want: 42.0},

				"float64 zero":                  {input: float64(0), want: 0},
				"positive float64 within range": {input: float64(42.0), want: 42.0},

				"string integer":              {input: "42", want: 42},
				"string with spaces":          {input: "42 ", want: 42},
				"string float":                {input: "42.0", want: 42},
				"string true":                 {input: "true", want: 1},
				"string false":                {input: "false", want: 0},
				"string 10_0":                 {input: "10_0", want: 100},
				"string binary":               {input: "0b101010", want: 42},
				"string short octal notation": {input: "042", want: 34},
				"string octal":                {input: "0o42", want: 34},
				"string hexadecimal":          {input: "0x42", want: 66},

				"boolean true":  {input: true, want: 1},
				"boolean false": {input: false, want: 0},

				"error":    {input: anyError{"42"}, want: 42},
				"stringer": {input: anyStringer{"42"}, want: 42},
			} {
				t.Run(fmt.Sprintf("from %s", name), func(t *testing.T) {
					got, err := converter(tt.input)
					assertNoError(t, err)

					if fmt.Sprint(got) != fmt.Sprint(tt.want) {
						t.Fatalf("unexpected result %+v != %+v", tt.want, got)
					}
				})
			}

			for name, tt := range map[string]struct {
				input       any
				errExpected error
			}{
				"nil": {
					input:       nil,
					errExpected: safecast.ErrUnsupportedConversion,
				},
				"unexpected type": {
					input:       struct{}{},
					errExpected: safecast.ErrUnsupportedConversion,
				},
				"empty string": {
					input:       "",
					errExpected: safecast.ErrStringConversion,
				},
				"simple space": {
					input:       " ",
					errExpected: safecast.ErrStringConversion,
				},
				"simple dot": {
					input:       ".",
					errExpected: safecast.ErrStringConversion,
				},
				"simple dash": {
					input:       "-",
					errExpected: safecast.ErrStringConversion,
				},
				"invalid string": {
					input:       "abc",
					errExpected: safecast.ErrStringConversion,
				},
				"invalid string with dot": {
					input:       "ab.c",
					errExpected: safecast.ErrStringConversion,
				},
				"string with leading +": {
					input:       "+42",
					errExpected: safecast.ErrStringConversion,
				},
				"invalid string multiple leading dashes": {
					input:       "--42",
					errExpected: safecast.ErrStringConversion,
				},
				"invalid string with leading dash": {
					input:       "-abc",
					errExpected: safecast.ErrStringConversion,
				},
				"invalid string with leading dash and dot": {
					input:       "-ab.c",
					errExpected: safecast.ErrStringConversion,
				},
			} {
				t.Run(name, func(t *testing.T) {
					_, err := converter(tt.input)
					requireErrorIs(t, err, safecast.ErrConversionIssue)
					requireErrorIs(t, err, tt.errExpected)
				})
			}
		})
	}

	t.Run("convert to float32 near zero", func(t *testing.T) {
		for name, tt := range map[string]struct {
			input any
			want  float32
		}{
			"negative untyped zero":              {input: negativeZero, want: float32(negativeZero)},
			"smallest positive non-zero float32": {input: math.SmallestNonzeroFloat32, want: 1e-45},
			"smallest negative non-zero float32": {input: -math.SmallestNonzeroFloat32, want: -1e-45},
			"smallest positive non-zero float64": {input: math.SmallestNonzeroFloat64, want: 4.9e-324},
			"smallest negative non-zero float64": {input: -math.SmallestNonzeroFloat64, want: -4.9e-324},
		} {
			t.Run(fmt.Sprintf("from %s", name), func(t *testing.T) {
				got, err := convertFloat32(tt.input)
				assertNoError(t, err)

				if got != tt.want {
					t.Fatalf("unexpected result want:%+v got:%+v", tt.want, got)
				}
			})
		}
	})

	t.Run("convert to float64 near zero", func(t *testing.T) {
		for name, tt := range map[string]struct {
			input any
			want  float64
		}{
			"negative untyped zero":              {input: negativeZero, want: negativeZero},
			"smallest positive non-zero float32": {input: math.SmallestNonzeroFloat32, want: 1.401298464324817e-45},
			"smallest negative non-zero float32": {input: -math.SmallestNonzeroFloat32, want: -1.401298464324817e-45},
			"smallest positive non-zero float64": {input: math.SmallestNonzeroFloat64, want: 4.9e-324},
			"smallest negative non-zero float64": {input: -math.SmallestNonzeroFloat64, want: -4.9e-324},
		} {
			t.Run(fmt.Sprintf("from %s", name), func(t *testing.T) {
				got, err := convertFloat64(tt.input)
				assertNoError(t, err)

				if got != tt.want {
					t.Fatalf("unexpected result want:%+v got:%+v", tt.want, got)
				}
			})
		}
	})

	t.Run("upper bound overflows", func(t *testing.T) {
		for name, tt := range map[string]struct {
			converter helper
			value     any
		}{
			"int": {
				converter: convertInt,
				value:     uint(math.MaxInt + 1),
			},
			"int8": {
				converter: convertInt8,
				value:     math.MaxInt8 + 1,
			},
			"int16": {
				converter: convertInt16,
				value:     math.MaxInt16 + 1,
			},
			"int32": {
				converter: convertInt32,
				value:     uint(math.MaxInt32 + 1),
			},
			"int64": {
				converter: convertInt64,
				value:     float64(math.MaxInt64 + 1), // the float64 conversion is used to avoid overflow on 32-bit
			},
			"uint": {
				converter: convertUint,
				value:     float64(math.MaxUint * 1.01),
			},
			"uint8": {
				converter: convertUint8,
				value:     math.MaxUint8 + 1,
			},
			"uint16": {
				converter: convertUint16,
				value:     math.MaxUint16 + 1,
			},
			"uint32": {
				converter: convertUint32,
				value:     float64(math.MaxUint32 * 1.01), // the float64 conversion is used to avoid overflow on 32-bit
			},
			"uint64": {
				converter: convertUint64,
				value:     float64(math.MaxUint64 * 1.01),
			},
			"float32": {
				converter: convertFloat32,
				value:     math.MaxFloat32 * 1.01,
			},
			"int string": {
				converter: convertInt,
				value:     "9223372036854775808", // math.MaxInt64 + 1
			},
			"int8 string": {
				converter: convertInt8,
				value:     "129", // math.MaxInt8 + 1
			},
			"int16 string": {
				converter: convertInt16,
				value:     "32769", // math.MaxInt16 + 1
			},
			"int32 string": {
				converter: convertInt32,
				value:     "2147483648", // math.MaxInt32 + 1
			},
			"int64 string": {
				converter: convertInt64,
				value:     "9223372036854775808", // math.MaxInt64 + 1
			},

			"int64 string overflow": {
				converter: convertInt64,
				value:     "123456789012345678901234567890", // more string than math.MaxInt64 represented as string
			},
		} {
			t.Run("for "+name, func(t *testing.T) {
				_, err := tt.converter(tt.value)
				requireErrorIs(t, err, safecast.ErrConversionIssue)
				requireErrorIs(t, err, safecast.ErrExceedMaximumValue)
			})
		}
	})

	t.Run("lower bound overflows", func(t *testing.T) {
		for name, tt := range map[string]struct {
			converter helper
			value     any
		}{
			"int": {
				converter: convertInt,
				value:     float32(math.MinInt * 1.01),
			},
			"int8": {
				converter: convertInt8,
				value:     math.MinInt8 - 1,
			},
			"int16": {
				converter: convertInt16,
				value:     math.MinInt16 - 1,
			},
			"int32": {
				converter: convertInt32,
				value:     float64(math.MinInt32 - 1), // the float64 conversion is used to avoid overflow on 32-bit
			},
			"int64": {
				converter: convertInt64,
				value:     float32(math.MinInt64 * 1.01),
			},
			"float32": {
				converter: convertFloat32,
				value:     -float64(math.MaxFloat32 * 1.01),
			},

			// Note: float64 cannot overflow

			"int from string": {
				converter: convertInt,
				value:     "-9223372036854775809", // math.MinInt64 - 1
			},

			"int8 from string": {
				converter: convertInt8,
				value:     "-129", // math.MinInt8 - 1
			},

			"int16 from string": {
				converter: convertInt16,
				value:     "-32769", // math.MinInt16 - 1
			},

			"int32 from string": {
				converter: convertInt32,
				value:     "-2147483649", // math.MinInt32 - 1
			},

			"int64 from string": {
				converter: convertInt64,
				value:     "-9223372036854775809", // math.MinInt64 - 1
			},

			"int64 string overflow": {
				converter: convertInt64,
				value:     "-123456789012345678901234567890", // more characters than math.MinInt64 represented as string
			},
		} {
			t.Run("for "+name, func(t *testing.T) {
				_, err := tt.converter(tt.value)
				requireErrorIs(t, err, safecast.ErrConversionIssue)
				requireErrorIs(t, err, safecast.ErrExceedMinimumValue)
			})
		}

		for name, converter := range unsignedConverters {
			t.Run(fmt.Sprintf("for %s", name), func(t *testing.T) {
				for name, input := range map[string]any{
					"untyped int within range": -42,
					"int":                      int(-42),
					"int8":                     int8(-42),
					"int16":                    int16(-42),
					"int32":                    int32(-42),
					"int64":                    int64(-42),
					"float32":                  float32(-42),
					"float64":                  float64(-42),

					"negative string":          "-42",
					"negative string with dot": "-42.0",
				} {
					t.Run(fmt.Sprintf("from %s", name), func(t *testing.T) {
						_, err := converter(input)
						requireErrorIs(t, err, safecast.ErrConversionIssue)
						requireErrorIs(t, err, safecast.ErrExceedMinimumValue)
					})
				}
			})
		}
	})
}
