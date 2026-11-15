package safecast

import (
	"math/bits"
	"reflect"
)

// IEEE 754 significand precision, including the implicit leading bit.
const (
	float32Precision = 24
	float64Precision = 53
)

func checkIntegerToFloatPrecisionLoss[NumOut, NumIn Number](orig NumIn, converted NumOut) (NumOut, error) {
	out := reflect.ValueOf(converted)

	var precision int

	// Determine the IEEE 754 significand precision.
	switch out.Kind() {
	case reflect.Float32:
		precision = float32Precision
	case reflect.Float64:
		precision = float64Precision
	default:
		// Skip if the target is not a float.
		return converted, nil
	}

	in := reflect.ValueOf(orig)

	var n uint64

	// Get the absolute magnitude of the integer.
	switch in.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		v := in.Int()
		if v < 0 {
			v = -v
		}

		n = uint64(v)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		n = in.Uint()
	default:
		// Skip if the original value is not an integer.
		return converted, nil
	}

	significantBits := bits.Len64(n) - bits.TrailingZeros64(n)

	// Report precision loss if the integer requires more bits
	// than the target floating-point significand supports.
	if significantBits > precision {
		return converted, errorHelper[NumOut]{
			value: orig,
			err:   ErrPrecisionLoss,
		}
	}

	// The integer can be represented exactly.
	return converted, nil
}
