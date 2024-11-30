package safecast

import "math"

func negative[T Number](t T) bool {
	return t < 0
}

func sameSign[T1, T2 Number](a T1, b T2) bool {
	return negative(a) == negative(b)
}

func getUpperBoundary(value any) any {
	switch value.(type) {
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
	case int64:
		return int64(math.MinInt64)
	case int32:
		return int32(math.MinInt32)
	case int16:
		return int16(math.MinInt16)
	case int8:
		return int8(math.MinInt8)
	case int:
		return int(math.MinInt)
	case uint:
		return uint(0)
	case uint8:
		return uint8(0)
	case uint16:
		return uint16(0)
	case uint32:
		return uint32(0)
	case uint64:
		return uint64(0)
	}

	return 0
}
