package safecast

import (
	"math"
)

func Negative[T Type](t T) bool {
	return t < 0
}

func SameSign[T1, T2 Type](a T1, b T2) bool {
	return Negative(a) == Negative(b)
}

// idea: @ccoVeille (me)
// Suggestion of improvement by @ldemailly
// Fixed floats issues by @ccoveille (me)

func Convert[NumOut Type, NumIn Type](orig NumIn) (converted NumOut, err error) {
	converted = NumOut(orig)

	if !SameSign(orig, converted) {
		err = ErrConversionIssue
		return
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

	errBoundary := ErrExceedMaximumValue
	boundary := getUpperBoundary(converted)
	if Negative(orig) {
		errBoundary = ErrExceedMinimumValue
		boundary = getLowerBoundary(converted)
	}

	return converted, Error{
		value:    orig,
		err:      errBoundary,
		boundary: boundary,
	}
}

func getUpperBoundary(value any) any {
	switch value.(type) {
	case float64:
		return float64(math.MaxFloat64)
	case float32:
		return float32(math.MaxFloat32)
	case int64:
		return int64(math.MaxInt64)
	case int32:
		return int32(math.MaxInt32)
	case int16:
		return int16(math.MaxInt16)
	case int8:
		return int8(math.MaxInt8)
	case int:
		return int(math.MaxInt)
	case uint64:
		return uint64(math.MaxUint64)
	case uint32:
		return uint32(math.MaxUint32)
	case uint16:
		return uint16(math.MaxUint16)
	case uint8:
		return uint8(math.MaxUint8)
	case uint:
		return uint(math.MaxUint)
	}

	return nil
}

func getLowerBoundary(value any) any {
	switch value.(type) {
	case float64:
		return -math.MaxFloat64
	case float32:
		return -math.MaxFloat32
	case int64:
		return math.MinInt64
	case int32:
		return math.MinInt32
	case int16:
		return math.MinInt16
	case int8:
		return math.MinInt8
	case int:
		return math.MinInt
	}

	return 0
}
