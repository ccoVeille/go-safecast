// Package go-safecast solves the type conversion issues in Go
//
// In Go, integer type conversion can lead to unexpected behavior and errors if not handled carefully.
// Issues can happen when converting between signed and unsigned integers, or when converting to a smaller integer type.

package safecast

import (
	"math"
)

func convertFromNumber[NumOut Number, NumIn Number](orig NumIn) (converted NumOut, err error) {
	converted = NumOut(orig)

	errBoundary := ErrExceedMaximumValue
	boundary := getUpperBoundary(converted)
	if negative(orig) {
		errBoundary = ErrExceedMinimumValue
		boundary = getLowerBoundary(converted)
	}

	if !sameSign(orig, converted) {
		return 0, Error{
			value:    orig,
			err:      errBoundary,
			boundary: boundary,
		}
	}

	base := orig
	switch f := any(orig).(type) {
	case float64:
		base = NumIn(math.Trunc(f))
	case float32:
		base = NumIn(math.Trunc(float64(f)))
	}

	if NumIn(converted) == base {
		return converted, nil
	}

	return 0, Error{
		value:    orig,
		err:      errBoundary,
		boundary: boundary,
	}
}

// ToInt attempts to convert any [Number] value to an int.
// If the conversion results in a value outside the range of an int,
// an [ErrConversionIssue] error is returned.
func ToInt[T Number](i T) (int, error) {
	return convertFromNumber[int](i)
}

// ToUint attempts to convert any [Number] value to an uint.
// If the conversion results in a value outside the range of an uint,
// an [ErrConversionIssue] error is returned.
func ToUint[T Number](i T) (uint, error) {
	return convertFromNumber[uint](i)
}

// ToInt8 attempts to convert any [Number] value to an int8.
// If the conversion results in a value outside the range of an int8,
// an [ErrConversionIssue] error is returned.
func ToInt8[T Number](i T) (int8, error) {
	return convertFromNumber[int8](i)
}

// ToUint8 attempts to convert any [Number] value to an uint8.
// If the conversion results in a value outside the range of an uint8,
// an [ErrConversionIssue] error is returned.
func ToUint8[T Number](i T) (uint8, error) {
	return convertFromNumber[uint8](i)
}

// ToInt16 attempts to convert any [Number] value to an int16.
// If the conversion results in a value outside the range of an int16,
// an [ErrConversionIssue] error is returned.
func ToInt16[T Number](i T) (int16, error) {
	return convertFromNumber[int16](i)
}

// ToUint16 attempts to convert any [Number] value to an uint16.
// If the conversion results in a value outside the range of an uint16,
// an [ErrConversionIssue] error is returned.
func ToUint16[T Number](i T) (uint16, error) {
	return convertFromNumber[uint16](i)
}

// ToInt32 attempts to convert any [Number] value to an int32.
// If the conversion results in a value outside the range of an int32,
// an [ErrConversionIssue] error is returned.
func ToInt32[T Number](i T) (int32, error) {
	return convertFromNumber[int32](i)
}

// ToUint32 attempts to convert any [Number] value to an uint32.
// If the conversion results in a value outside the range of an uint32,
// an [ErrConversionIssue] error is returned.
func ToUint32[T Number](i T) (uint32, error) {
	return convertFromNumber[uint32](i)
}

// ToInt64 attempts to convert any [Number] value to an int64.
// If the conversion results in a value outside the range of an int64,
// an [ErrConversionIssue] error is returned.
func ToInt64[T Number](i T) (int64, error) {
	return convertFromNumber[int64](i)
}

// ToUint64 attempts to convert any [Number] value to an uint64.
// If the conversion results in a value outside the range of an uint64,
// an [ErrConversionIssue] error is returned.
func ToUint64[T Number](i T) (uint64, error) {
	return convertFromNumber[uint64](i)
}
