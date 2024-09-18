package safecast_test

import (
	"fmt"
	"math"

	"github.com/ccoveille/go-safecast"
)

func ExampleToInt8() {
	a := float64(42.42)
	i, err := safecast.ToInt8(a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)

	a = float64(200.42)
	_, err = safecast.ToInt8(a)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// 42
	// conversion issue: 200.42 (float64) is greater than 127 (int8): maximum value for this type exceeded
}

func ExampleToUint8() {
	a := int64(42)
	i, err := safecast.ToUint8(a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)

	a = int64(-1)
	_, err = safecast.ToUint8(a)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// 42
	// conversion issue: -1 (int64) is less than 0 (uint8): minimum value for this type exceeded
}

func ExampleToInt16() {
	a := int32(42)
	i, err := safecast.ToInt16(a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)

	a = int32(40000)
	_, err = safecast.ToInt16(a)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// 42
	// conversion issue: 40000 (int32) is greater than 32767 (int16): maximum value for this type exceeded
}

func ExampleToUint16() {
	a := int64(42)
	i, err := safecast.ToUint16(a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)

	a = int64(-1)
	_, err = safecast.ToUint16(a)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// 42
	// conversion issue: -1 (int64) is less than 0 (uint16): minimum value for this type exceeded
}

func ExampleToInt32() {
	a := int(42)
	i, err := safecast.ToInt32(a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)

	a = int(math.MaxUint32 + 1)
	_, err = safecast.ToInt32(a)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// 42
	// conversion issue: 4294967296 (int) is greater than 2147483647 (int32): maximum value for this type exceeded
}

func ExampleToUint32() {
	a := int16(42)
	i, err := safecast.ToUint32(a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)

	a = int16(-1)
	_, err = safecast.ToUint32(a)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// 42
	// conversion issue: -1 (int16) is less than 0 (uint32): minimum value for this type exceeded
}

func ExampleToInt64() {
	a := uint64(42)
	i, err := safecast.ToInt64(a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)

	a = uint64(math.MaxInt64) + 1
	_, err = safecast.ToInt64(a)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// 42
	// conversion issue: 9223372036854775808 (uint64) is greater than 9223372036854775807 (int64): maximum value for this type exceeded
}

func ExampleToUint64() {
	a := int8(42)
	i, err := safecast.ToUint64(a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)

	a = int8(-1)
	_, err = safecast.ToUint64(a)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// 42
	// conversion issue: -1 (int8) is less than 0 (uint64): minimum value for this type exceeded
}

func ExampleToInt() {
	a := uint64(42)
	i, err := safecast.ToInt(a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)

	a = uint64(math.MaxInt64) + 1
	_, err = safecast.ToInt(a)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// 42
	// conversion issue: 9223372036854775808 (uint64) is greater than 9223372036854775807 (int): maximum value for this type exceeded
}

func ExampleToUint() {
	a := int8(42)
	i, err := safecast.ToUint(a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)

	a = int8(-1)
	_, err = safecast.ToUint(a)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// 42
	// conversion issue: -1 (int8) is less than 0 (uint): minimum value for this type exceeded
}

func ExampleConvert() {
	a := int(42)
	i, err := safecast.Convert[int32](a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)

	a = int(math.MaxUint32 + 1)
	_, err = safecast.ToInt32(a)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// 42
	// conversion issue: 4294967296 (int) is greater than 2147483647 (int32): maximum value for this type exceeded
}
