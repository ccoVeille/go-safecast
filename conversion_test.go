package safecast_test

// The tests in conversion_test.go are the ones that are not architecture dependent
// The tests in conversion_64bit_test.go complete them for 64-bit systems
//
// The architecture dependent file covers the fact, you can reach a higher value with int and uint
// on 64-bit systems, but you will get a compile error on 32-bit.
// This is why it needs to be tested in an architecture dependent way.

import (
	"math"
	"testing"
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
