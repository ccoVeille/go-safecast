package safecast

// This files is highly inspired from https://pkg.go.dev/golang.org/x/exp/constraints
// I didn't import it as it would have added an unnecessary dependency


// Integer is an alias for the int, uint, int8, uint8, int16, uint16, int32, uint32, int64, and uint64 types.
type Integer interface {
	~int | ~uint | ~int8 | ~uint8 | ~int16 | ~uint16 | ~int32 | ~uint32 | ~int64 | ~uint64
}

// Float is an alias for the float32 and float64 types.
type Float interface {
	~float32 | ~float64
}

// Number is an alias for the int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64, float32, and float64 types.
type Number interface {
	Integer | Float
}