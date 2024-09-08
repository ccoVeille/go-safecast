// In Go, integer type conversion can lead to unexpected behavior and errors if not handled carefully.
// Issues can happen when converting between signed and unsigned integers, or when converting to a smaller integer type.
// This package aims to solve this issue

package safecast

import (
	"fmt"
	"math"
)

// ToInt attempts to convert any [Number] value to an int.
// If the conversion results in a value outside the range of an int,
// an ErrOutOfRange error is returned.
func ToInt[T Number](i T) (int, error) {
	if i > 0 && uint64(i) > math.MaxInt {
		return 0, fmt.Errorf("%w: %v is greater than math.MaxInt", ErrOutOfRange, i)
	}

	return int(i), nil
}

// ToUint attempts to convert any [Number] value to an uint.
// If the conversion results in a value outside the range of an uint,
// an ErrOutOfRange error is returned.
func ToUint[T Number](i T) (uint, error) {
	if err := assertNotNegative(i); err != nil {
		return 0, err
	}

	if float64(i) > math.MaxUint64 {
		return 0, fmt.Errorf("%w: %v is greater than math.MaxUint64", ErrOutOfRange, i)
	}

	return uint(i), nil
}

// ToInt8 attempts to convert any [Number] value to an int8.
// If the conversion results in a value outside the range of an int8,
// an ErrOutOfRange error is returned.
func ToInt8[T Number](i T) (int8, error) {
	if i > math.MaxInt8 {
		return 0, fmt.Errorf("%w: %v is greater than math.MaxInt8", ErrOutOfRange, i)
	}

	if i < 0 && uint64(-i) > -math.MinInt8 {
		return 0, fmt.Errorf("%w: %v is less than math.MinInt8", ErrOutOfRange, i)
	}

	return int8(i), nil
}

// ToUint8 attempts to convert any [Number] value to an uint8.
// If the conversion results in a value outside the range of an uint8,
// an ErrOutOfRange error is returned.
func ToUint8[T Number](i T) (uint8, error) {
	if err := assertNotNegative(i); err != nil {
		return 0, err
	}

	if uint64(i) > math.MaxUint8 {
		return 0, fmt.Errorf("%w: %v is greater than math.MaxUint8", ErrOutOfRange, i)
	}

	return uint8(i), nil
}

// ToInt16 attempts to convert any [Number] value to an int16.
// If the conversion results in a value outside the range of an int16,
// an ErrOutOfRange error is returned.
func ToInt16[T Number](i T) (int16, error) {
	if i > 0 && uint64(i) > math.MaxInt16 {
		return 0, fmt.Errorf("%w: %v is greater than math.MaxInt16", ErrOutOfRange, i)
	}

	if i < 0 && uint64(-i) > -math.MinInt16 {
		return 0, fmt.Errorf("%w: %v is less than math.MinInt16", ErrOutOfRange, i)
	}

	return int16(i), nil
}

// ToUint16 attempts to convert any [Number] value to an uint16.
// If the conversion results in a value outside the range of an uint16,
// an ErrOutOfRange error is returned.
func ToUint16[T Number](i T) (uint16, error) {
	if err := assertNotNegative(i); err != nil {
		return 0, err
	}

	if uint64(i) > math.MaxUint16 {
		return 0, fmt.Errorf("%w: %v is greater than math.MaxUint16", ErrOutOfRange, i)
	}

	return uint16(i), nil
}

// ToInt32 attempts to convert any [Number] value to an int32.
// If the conversion results in a value outside the range of an int32,
// an ErrOutOfRange error is returned.
func ToInt32[T Number](i T) (int32, error) {
	if i > 0 && uint64(i) > math.MaxInt32 {
		return 0, fmt.Errorf("%w: %v is greater than math.MaxInt32", ErrOutOfRange, i)
	}

	if i < 0 && uint64(-i) > -math.MinInt32 {
		return 0, fmt.Errorf("%w: %v is less than math.MinInt32", ErrOutOfRange, i)
	}

	return int32(i), nil
}

// ToUint32 attempts to convert any [Number] value to an uint32.
// If the conversion results in a value outside the range of an uint32,
// an ErrOutOfRange error is returned.
func ToUint32[T Number](i T) (uint32, error) {
	if err := assertNotNegative(i); err != nil {
		return 0, err
	}

	if uint64(i) > math.MaxUint32 {
		return 0, fmt.Errorf("%w: %v is greater than math.MaxUint32", ErrOutOfRange, i)
	}

	return uint32(i), nil
}

// ToInt64 attempts to convert any [Number] value to an int64.
// If the conversion results in a value outside the range of an int64,
// an ErrOutOfRange error is returned.
func ToInt64[T Number](i T) (int64, error) {
	if i > 0 && uint64(i) > math.MaxInt64 {
		return 0, fmt.Errorf("%w: %v is greater than math.MaxInt64", ErrOutOfRange, i)
	}

	return int64(i), nil
}

// ToUint64 attempts to convert any [Number] value to an uint64.
// If the conversion results in a value outside the range of an uint64,
// an ErrOutOfRange error is returned.
func ToUint64[T Number](i T) (uint64, error) {
	if err := assertNotNegative(i); err != nil {
		return 0, err
	}

	if float64(i) > math.MaxUint64 {
		return 0, fmt.Errorf("%w: %v is greater than math.MaxUint64", ErrOutOfRange, i)
	}

	return uint64(i), nil
}
