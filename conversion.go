package safecast

import (
	"errors"
	"fmt"
	"math"
)

type Integer interface {
	~int | ~uint | ~int8 | ~uint8 | ~int16 | ~uint16 | ~int32 | ~uint32 | ~int64 | ~uint64
}

var ErrOutOfRange = errors.New("out of range")

func ToInt[T Integer](i T) (int, error) {
	if i > 0 && uint64(i) > math.MaxInt {
		return 0, fmt.Errorf("%w: %d is greater than math.MaxInt", ErrOutOfRange, i)
	}

	if i < 0 && uint64(-i) > -math.MinInt {
		return 0, fmt.Errorf("%w: %d is less than math.MinInt", ErrOutOfRange, i)
	}

	return int(i), nil
}

func ToUint[T Integer](i T) (uint, error) {
	if i < 0 {
		return 0, fmt.Errorf("%w: %d is negative", ErrOutOfRange, i)
	}

	if uint64(i) > math.MaxUint {
		return 0, fmt.Errorf("%w: %d is greater than math.MaxUint", ErrOutOfRange, i)
	}

	return uint(i), nil
}

func ToInt8[T Integer](i T) (int8, error) {
	if i > math.MaxInt8 {
		return 0, fmt.Errorf("%w: %d is greater than math.MaxInt8", ErrOutOfRange, i)
	}

	if i < 0 && uint64(-i) > -math.MinInt8 {
		return 0, fmt.Errorf("%w: %d is less than math.MinInt8", ErrOutOfRange, i)
	}

	return int8(i), nil
}

func ToUint8[T Integer](i T) (uint8, error) {
	if i < 0 {
		return 0, fmt.Errorf("%w: %d is negative", ErrOutOfRange, i)
	}

	if uint64(i) > math.MaxUint8 {
		return 0, fmt.Errorf("%w: %d is greater than math.MaxUint8", ErrOutOfRange, i)
	}

	return uint8(i), nil
}

func ToInt16[T Integer](i T) (int16, error) {
	if i > 0 && uint64(i) > math.MaxInt16 {
		return 0, fmt.Errorf("%w: %d is greater than math.MaxInt16", ErrOutOfRange, i)
	}

	if i < 0 && uint64(-i) > -math.MinInt16 {
		return 0, fmt.Errorf("%w: %d is less than math.MinInt16", ErrOutOfRange, i)
	}

	return int16(i), nil
}

func ToUint16[T Integer](i T) (uint16, error) {
	if i < 0 {
		return 0, fmt.Errorf("%w: %d is negative", ErrOutOfRange, i)
	}

	if uint64(i) > math.MaxUint16 {
		return 0, fmt.Errorf("%w: %d is greater than math.MaxUint16", ErrOutOfRange, i)
	}

	return uint16(i), nil
}

func ToInt32[T Integer](i T) (int32, error) {
	if i > 0 && uint64(i) > math.MaxInt32 {
		return 0, fmt.Errorf("%w: %d is greater than math.MaxInt32", ErrOutOfRange, i)
	}

	if i < 0 && uint64(-i) > -math.MinInt32 {
		return 0, fmt.Errorf("%w: %d is less than math.MinInt32", ErrOutOfRange, i)
	}

	return int32(i), nil
}

func ToUint32[T Integer](i T) (uint32, error) {
	if i < 0 {
		return 0, fmt.Errorf("%w: %d is negative", ErrOutOfRange, i)
	}

	if uint64(i) > math.MaxUint32 {
		return 0, fmt.Errorf("%w: %d is greater than math.MaxUint32", ErrOutOfRange, i)
	}

	return uint32(i), nil
}

func ToInt64[T Integer](i T) (int64, error) {
	if i > 0 && uint64(i) > math.MaxInt64 {
		return 0, fmt.Errorf("%w: %d is greater than math.MaxInt64", ErrOutOfRange, i)
	}

	if i < 0 && uint64(-i) > -math.MinInt64 {
		return 0, fmt.Errorf("%w: %d is less than math.MinInt64", ErrOutOfRange, i)
	}

	return int64(i), nil
}

func ToUint64[T Integer](i T) (uint64, error) {
	if i < 0 {
		return 0, fmt.Errorf("%w: %d is negative", ErrOutOfRange, i)
	}

	return uint64(i), nil
}
