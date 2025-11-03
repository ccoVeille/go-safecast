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

	"github.com/ccoveille/go-safecast/v2"
)

type MapTest[TypeInput safecast.Number, TypeOutput safecast.Number] struct {
	Input          TypeInput
	Options        []safecast.ConvertOption
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

	out, err := safecast.Convert[O](mt.Input, mt.Options...)
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

		"uintptr to uint":  MapTest[uintptr, uint]{Input: uintptr(42), ExpectedOutput: uint(42)},
		"uint to uintptr":  MapTest[uint, uintptr]{Input: uint(42), ExpectedOutput: uintptr(42)},
		"int8 to uintptr":  MapTest[int8, uintptr]{Input: int8(42), ExpectedOutput: uintptr(42)},
		"int16 to uintptr": MapTest[int16, uintptr]{Input: int16(42), ExpectedOutput: uintptr(42)},
		"int32 to uintptr": MapTest[int32, uintptr]{Input: int32(42), ExpectedOutput: uintptr(42)},
		"int64 to uintptr": MapTest[int64, uintptr]{Input: int64(42), ExpectedOutput: uintptr(42)},
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

		"upper bound overflows for uintptr": MapTest[float64, uintptr]{
			Input:         float64(math.MaxUint64 * 1.01), // uintptr is at least uint32, but can be uint64
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

		"negative overflows uintptr": MapTest[int, uintptr]{
			Input:         -42,
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

			"simple alias overflows for int8": MapTest[UintSimpleAlias, int8]{
				Input:         UintSimpleAlias(255),
				ExpectedError: safecast.ErrExceedMaximumValue,
				ErrorContains: "255 (uint) is greater than 127 (int8)",
			},

			"type alias overflows for int8": MapTest[UintTypeAlias, int8]{
				Input:         UintTypeAlias(255),
				ExpectedError: safecast.ErrExceedMaximumValue,
				ErrorContains: "255 (safecast_test.UintTypeAlias) is greater than 127 (int8)",
			},
		} {
			t.Run(name, func(t *testing.T) {
				c.Run(t)
			})
		}
	})

	t.Run("float to int rounding", func(t *testing.T) {
		t.Run("without decimal loss", func(t *testing.T) {
			for name, tt := range map[string]TestRunner{
				"float32": MapTest[float32, int]{
					Input:          3.14,
					ExpectedOutput: 3,
				},
				"float64": MapTest[float64, int]{
					Input:          3.14,
					ExpectedOutput: 3,
				},
			} {
				t.Run(name, func(t *testing.T) {
					tt.Run(t)
				})
			}
		})

		t.Run("with decimal loss", func(t *testing.T) {
			for name, tt := range map[string]TestRunner{
				"float32": MapTest[float32, int]{
					Input:          3.14,
					Options:        []safecast.ConvertOption{safecast.WithDecimalLossReport()},
					ExpectedOutput: 3,
					ExpectedError:  safecast.ErrDecimalLoss,
				},
				"float64": MapTest[float64, int]{
					Input:          3.14,
					Options:        []safecast.ConvertOption{safecast.WithDecimalLossReport()},
					ExpectedOutput: 3,
					ExpectedError:  safecast.ErrDecimalLoss,
				},
			} {
				t.Run(name, func(t *testing.T) {
					tt.Run(t)
				})
			}
		})
	})
}

type MapMustConvertTest[TypeInput safecast.Number, TypeOutput safecast.Number] struct {
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
		} {
			t.Run(name, func(t *testing.T) {
				tt.Run(t)
			})
		}
	})

	t.Run("no panic", func(t *testing.T) {
		for name, tt := range map[string]TestRunner{
			"number": MapMustConvertTest[int, uint8]{Input: 42, ExpectedOutput: 42},
			"float":  MapMustConvertTest[float64, uint8]{Input: 42.0, ExpectedOutput: 42},
		} {
			t.Run(name, func(t *testing.T) {
				tt.Run(t)
			})
		}
	})
}

type MapRequireConvertTest[TypeInput safecast.Number, TypeOutput safecast.Number] struct {
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
			"float":  MapRequireConvertTest[float64, uint8]{Input: 42.0, ExpectedOutput: 42},
		} {
			t.Run(name, func(t *testing.T) {
				tt.Run(t)
			})
		}
	})

	t.Run("test fail on error", func(t *testing.T) {
		for name, tt := range map[string]TestRunner{
			"negative": MapRequireConvertTest[int, uint8]{Input: -1, ExpectedTestFailure: true, ExpectedOutput: 255},
			"overflow": MapRequireConvertTest[int, uint8]{Input: math.MaxInt, ExpectedTestFailure: true, ExpectedOutput: 255},
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

	_ = safecast.RequireConvert[uint8](t, -1)
	// here an error is raised via [testing.T.Fatal]

	// Output:
	// --- FAIL:
	// 	conversion issue: -1 (int) is less than 0 (uint8): minimum value for this type exceeded
}

func ExampleWithDecimalLossReport() {
	// By default, converting from float to int does not report decimal loss
	val1, err1 := safecast.Convert[int](3.14)
	fmt.Println(val1, err1)

	// Using the WithDecimalLossReport option, decimal loss is reported as an error
	val2, err2 := safecast.Convert[int](3.14, safecast.WithDecimalLossReport())
	fmt.Println(val2, err2)

	// Output:
	// 3 <nil>
	// 3 conversion issue: decimal loss during conversion
}
