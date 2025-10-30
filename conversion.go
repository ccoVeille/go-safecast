package safecast

import (
	"math"
	"reflect"
	"strconv"
	"strings"
)

// Convert attempts to convert any value to the desired type
//
// # Behavior
//
//   - If the conversion is possible, the converted value is returned.
//
// # Errors when conversion exceeds range of the desired type, the following errors are wrapped in the returned error:
//
//   - [ErrRangeOverflow] when the value is outside the range of the desired type. (example: 1000 or -1 to uint8).
//   - [ErrExceedMaximumValue] when the value exceeds the maximum value of the desired type (example: 1000 to uint8).
//   - [ErrExceedMinimumValue] when the value is less than the minimum value of the desired type (example: -1 to uint16).
//
// # Errors when conversion is not possible, the following errors are wrapped in the returned error:
//
//   - [ErrorUnsupportedConversion] when the conversion is not possible for the desired type (example: NaN to int).
//   - [ErrStringConversion] when the conversion from string fails (example: "abc" to int).
//
// # General errors wrapped on conversion failure:
//
//   - [ErrConversionIssue] is always wrapped in the returned error when [Convert] fails (example "abc", -1, or 1000 to uint8).
//
// # ⚠️ Note about string support
//
// ️️Please consider using [Parse] instead, which provides more options for string parsing.
// The support of string input in this function will be deprecated in future major release.
func Convert[NumOut Number, NumIn Input](orig NumIn) (converted NumOut, err error) {
	v := reflect.ValueOf(orig)
	switch v.Kind() {
	case reflect.Int:
		return convertFromNumber[NumOut](int(v.Int()))
	case reflect.Uint:
		return convertFromNumber[NumOut](uint(v.Uint()))
	case reflect.Int8:
		//nolint:gosec // the int8 is confirmed
		return convertFromNumber[NumOut](int8(v.Int()))
	case reflect.Uint8:
		//nolint:gosec // the uint8 is confirmed
		return convertFromNumber[NumOut](uint8(v.Uint()))
	case reflect.Int16:
		//nolint:gosec // the int16 is confirmed
		return convertFromNumber[NumOut](int16(v.Int()))
	case reflect.Uint16:
		//nolint:gosec // the uint16 is confirmed
		return convertFromNumber[NumOut](uint16(v.Uint()))
	case reflect.Int32:
		//nolint:gosec // the int32 is confirmed
		return convertFromNumber[NumOut](int32(v.Int()))
	case reflect.Uint32:
		//nolint:gosec // the uint32 is confirmed
		return convertFromNumber[NumOut](uint32(v.Uint()))
	case reflect.Int64:
		return convertFromNumber[NumOut](int64(v.Int()))
	case reflect.Uint64:
		return convertFromNumber[NumOut](uint64(v.Uint()))
	case reflect.Float32:
		return convertFromNumber[NumOut](float32(v.Float()))
	case reflect.Float64:
		return convertFromNumber[NumOut](float64(v.Float()))
	case reflect.Bool:
		if v.Bool() {
			return NumOut(1), nil
		}
		return NumOut(0), nil
	case reflect.String:
		converted, err = convertFromString[NumOut](v.String())
		// this falls through to default statement is a deliberate hack for increasing the code coverage.
		// without this, the default case would be impossible to reach in tests.
		fallthrough
	default:
		return converted, err
	}
}

// MustConvert calls [Convert] to convert the value to the desired type, and panics if the conversion fails.
func MustConvert[NumOut Number, NumIn Input](orig NumIn) NumOut {
	converted, err := Convert[NumOut](orig)
	if err != nil {
		panic(err)
	}
	return converted
}

// TestingT is an interface wrapper used by [RequireConvert] that we need for testing purposes.
//
// Only the methods used by [RequireConvert] are expected to be implemented.
//
// [*testing.T], [*testing.B], or [*testing.F] types satisfy this interface.
type TestingT interface {
	Helper()
	Fatal(args ...any)
}

// RequireConvert is a test helper that calls [Convert] that converts the value to the desired type,
// and fails the test if the conversion fails.
func RequireConvert[NumOut Number, NumIn Input](t TestingT, orig NumIn) (converted NumOut) {
	t.Helper()

	converted, err := Convert[NumOut](orig)
	if err != nil {
		t.Fatal(err)
	}
	return converted
}

func convertFromNumber[NumOut Number, NumIn Number](orig NumIn) (NumOut, error) {
	converted := NumOut(orig)
	if isFloat64[NumIn]() {
		floatOrig := float64(orig)
		if math.IsInf(floatOrig, 1) || math.IsInf(floatOrig, -1) {
			return 0, getRangeError[NumOut](orig)
		}
		if math.IsNaN(floatOrig) {
			return 0, errorHelper[NumOut]{
				value: orig,
				err:   ErrUnsupportedConversion,
			}
		}
	}

	if isFloat64[NumOut]() {
		// float64 cannot overflow, so we don't have to worry about it
		return converted, nil
	}

	if isFloat32[NumOut]() {
		// check boundary
		if math.Abs(float64(orig)) < math.MaxFloat32 {
			// the value is within float32 range, there is no overflow
			return converted, nil
		}

		// TODO: check for numbers close to math.MaxFloat32

		return 0, getRangeError[NumOut](orig)
	}

	if !sameSign(orig, converted) {
		return 0, getRangeError[NumOut](orig)
	}

	// and compare
	base := orig
	if isFloat[NumIn]() {
		base = NumIn(math.Trunc(float64(orig)))
	}

	// convert back to the original type
	cast := NumIn(converted)
	if cast != base {
		return 0, getRangeError[NumOut](orig)
	}

	return converted, nil
}

func convertFromString[NumOut Number](input string) (converted NumOut, err error) {
	s := strings.TrimSpace(input)

	if b, err := strconv.ParseBool(s); err == nil {
		if b {
			return NumOut(1), nil
		}
		return NumOut(0), nil

	}

	return Parse[NumOut](input, WithBaseAutoDetection(), withLegacyMode())
}

func getRangeError[NumOut Number, NumIn Number](value NumIn) error {
	err := ErrExceedMaximumValue
	if value < 0 {
		err = ErrExceedMinimumValue
	}

	return errorHelper[NumOut]{
		value: value,
		err:   err,
	}
}
