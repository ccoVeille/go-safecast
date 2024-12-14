package safecast

import "math"

func negative[T Number](t T) bool {
	return t < 0
}

func sameSign[T1, T2 Number](a T1, b T2) bool {
	return negative(a) == negative(b)
}

func getUpperBoundary(value any) any {
	var upper any = math.Inf(1)
	switch value.(type) {
	case int8:
		upper = int8(math.MaxInt8)
	case int16:
		upper = int16(math.MaxInt16)
	case int32:
		upper = int32(math.MaxInt32)
	case int64:
		upper = int64(math.MaxInt64)
	case int:
		upper = int(math.MaxInt)
	case uint8:
		upper = uint8(math.MaxUint8)
	case uint32:
		upper = uint32(math.MaxUint32)
	case uint16:
		upper = uint16(math.MaxUint16)
	case uint64:
		upper = uint64(math.MaxUint64)
	case uint:
		upper = uint(math.MaxUint)

	// Note: there is no float64 boundary
	// because float64 cannot overflow
	case float32:
		upper = float32(math.MaxFloat32)
	}

	return upper
}

func getLowerBoundary(value any) any {
	var lower any = math.Inf(-1)
	switch value.(type) {
	case int64:
		lower = int64(math.MinInt64)
	case int32:
		lower = int32(math.MinInt32)
	case int16:
		lower = int16(math.MinInt16)
	case int8:
		lower = int8(math.MinInt8)
	case int:
		lower = int(math.MinInt)
	case uint:
		lower = uint(0)
	case uint8:
		lower = uint8(0)
	case uint16:
		lower = uint16(0)
	case uint32:
		lower = uint32(0)
	case uint64:
		lower = uint64(0)

	// Note: there is no float64 boundary
	// because float64 cannot overflow
	case float32:
		lower = float32(-math.MaxFloat32)

	}

	return lower
}
