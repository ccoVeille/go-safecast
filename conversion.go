// Package safecast aims to solve the issue of type conversion
//
// In Go, integer type conversion can lead to unexpected behavior and errors if not handled carefully.
// Issues can happen when converting between signed and unsigned integers, or when converting to a smaller integer type.

package safecast

import "math"

// ToInt attempts to convert any [Type] value to an int.
// If the conversion results in a value outside the range of an int,
// an [ErrConversionIssue] error is returned.
func ToInt[T Type](i T) (int, error) {
	if err := checkUpperBoundary(i, int(math.MaxInt)); err != nil {
		return 0, err
	}

	if err := checkLowerBoundary(i, int(math.MinInt)); err != nil {
		return 0, err
	}

	return int(i), nil
}

// ToUint attempts to convert any [Type] value to an uint.
// If the conversion results in a value outside the range of an uint,
// an [ErrConversionIssue] error is returned.
func ToUint[T Type](i T) (uint, error) {
	if err := assertNotNegative(i, uint(0)); err != nil {
		return 0, err
	}

	if err := checkUpperBoundary(i, uint(math.MaxUint)); err != nil {
		return 0, err
	}

	return uint(i), nil
}

// ToInt8 attempts to convert any [Type] value to an int8.
// If the conversion results in a value outside the range of an int8,
// an [ErrConversionIssue] error is returned.
func ToInt8[T Type](i T) (int8, error) {
	if err := checkUpperBoundary(i, int8(math.MaxInt8)); err != nil {
		return 0, err
	}

	if err := checkLowerBoundary(i, int8(math.MinInt8)); err != nil {
		return 0, err
	}

	return int8(i), nil
}

// ToUint8 attempts to convert any [Type] value to an uint8.
// If the conversion results in a value outside the range of an uint8,
// an [ErrConversionIssue] error is returned.
func ToUint8[T Type](i T) (uint8, error) {
	if err := assertNotNegative(i, uint8(0)); err != nil {
		return 0, err
	}

	if err := checkUpperBoundary(i, uint8(math.MaxUint8)); err != nil {
		return 0, err
	}

	return uint8(i), nil
}

// ToInt16 attempts to convert any [Type] value to an int16.
// If the conversion results in a value outside the range of an int16,
// an [ErrConversionIssue] error is returned.
func ToInt16[T Type](i T) (int16, error) {
	if err := checkUpperBoundary(i, int16(math.MaxInt16)); err != nil {
		return 0, err
	}

	if err := checkLowerBoundary(i, int16(math.MinInt16)); err != nil {
		return 0, err
	}

	return int16(i), nil
}

// ToUint16 attempts to convert any [Type] value to an uint16.
// If the conversion results in a value outside the range of an uint16,
// an [ErrConversionIssue] error is returned.
func ToUint16[T Type](i T) (uint16, error) {
	if err := assertNotNegative(i, uint16(0)); err != nil {
		return 0, err
	}

	if err := checkUpperBoundary(i, uint16(math.MaxUint16)); err != nil {
		return 0, err
	}

	return uint16(i), nil
}

// ToInt32 attempts to convert any [Type] value to an int32.
// If the conversion results in a value outside the range of an int32,
// an [ErrConversionIssue] error is returned.
func ToInt32[T Type](i T) (int32, error) {
	if err := checkUpperBoundary(i, int32(math.MaxInt32)); err != nil {
		return 0, err
	}

	if err := checkLowerBoundary(i, int32(math.MinInt32)); err != nil {
		return 0, err
	}

	return int32(i), nil
}

// ToUint32 attempts to convert any [Type] value to an uint32.
// If the conversion results in a value outside the range of an uint32,
// an [ErrConversionIssue] error is returned.
func ToUint32[T Type](i T) (uint32, error) {
	if err := assertNotNegative(i, uint32(0)); err != nil {
		return 0, err
	}

	if err := checkUpperBoundary(i, uint32(math.MaxUint32)); err != nil {
		return 0, err
	}

	return uint32(i), nil
}

// ToInt64 attempts to convert any [Type] value to an int64.
// If the conversion results in a value outside the range of an int64,
// an [ErrConversionIssue] error is returned.
func ToInt64[T Type](i T) (int64, error) {
	if err := checkUpperBoundary(i, int64(math.MaxInt64)); err != nil {
		return 0, err
	}

	return int64(i), nil
}

// ToUint64 attempts to convert any [Type] value to an uint64.
// If the conversion results in a value outside the range of an uint64,
// an [ErrConversionIssue] error is returned.
func ToUint64[T Type](i T) (uint64, error) {
	if err := assertNotNegative(i, uint64(0)); err != nil {
		return 0, err
	}

	if err := checkUpperBoundary(i, uint64(math.MaxUint64)); err != nil {
		return 0, err
	}

	return uint64(i), nil
}
