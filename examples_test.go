package safecast_test

import (
	"fmt"
	"math"

	"github.com/ccoveille/go-safecast"
)

// The tests in examples_test.go are the ones that are not architecture dependent
// The tests in examples_32bit_test.go are for 32-bit systems
// The tests in examples_64bit_test.go are for 64-bit systems
//
// The architecture dependent files cover the difference when dealing with [safecast.ToInt]
// The error message is different on 32-bit and 64-bit systems
// The max is 9223372036854775807 on 64-bit and 2147483647 on 32-bit
//
// The examples could have been skipped for 32-bit systems,
// but I wanted the Examples to be launched on this architecture.

func ExampleToInt8() {
	a := float64(42.42)
	i, err := safecast.ToInt8(a)
	fmt.Println(i, err)

	a = float64(200.42)
	i, err = safecast.ToInt8(a)
	fmt.Println(i, err)

	// Output:
	// 42 <nil>
	// 0 conversion issue: 200.42 (float64) is greater than 127 (int8): maximum value for this type exceeded
}

func ExampleToUint8() {
	a := int64(42)
	i, err := safecast.ToUint8(a)
	fmt.Println(i, err)

	a = int64(-1)
	i, err = safecast.ToUint8(a)
	fmt.Println(i, err)

	// Output:
	// 42 <nil>
	// 0 conversion issue: -1 (int64) is less than 0 (uint8): minimum value for this type exceeded
}

func ExampleToInt16() {
	a := int32(42)
	i, err := safecast.ToInt16(a)
	fmt.Println(i, err)

	a = int32(40000)
	i, err = safecast.ToInt16(a)
	fmt.Println(i, err)
	// Output:
	// 42 <nil>
	// 0 conversion issue: 40000 (int32) is greater than 32767 (int16): maximum value for this type exceeded
}

func ExampleToUint16() {
	a := int64(42)
	i, err := safecast.ToUint16(a)
	fmt.Println(i, err)

	a = int64(-1)
	i, err = safecast.ToUint16(a)
	fmt.Println(i, err)
	// Output:
	// 42 <nil>
	// 0 conversion issue: -1 (int64) is less than 0 (uint16): minimum value for this type exceeded
}

func ExampleToInt32() {
	a := int(42)
	i, err := safecast.ToInt32(a)
	fmt.Println(i, err)

	b := uint32(math.MaxInt32) + 1
	i, err = safecast.ToInt32(b)
	fmt.Println(i, err)

	// Output:
	// 42 <nil>
	// 0 conversion issue: 2147483648 (uint32) is greater than 2147483647 (int32): maximum value for this type exceeded
}

func ExampleToUint32() {
	a := int16(42)
	i, err := safecast.ToUint32(a)
	fmt.Println(i, err)

	a = int16(-1)
	i, err = safecast.ToUint32(a)
	fmt.Println(i, err)

	// Output:
	// 42 <nil>
	// 0 conversion issue: -1 (int16) is less than 0 (uint32): minimum value for this type exceeded
}

func ExampleToInt64() {
	a := uint64(42)
	i, err := safecast.ToInt64(a)
	fmt.Println(i, err)

	a = uint64(math.MaxInt64) + 1
	i, err = safecast.ToInt64(a)
	fmt.Println(i, err)

	// Output:
	// 42 <nil>
	// 0 conversion issue: 9223372036854775808 (uint64) is greater than 9223372036854775807 (int64): maximum value for this type exceeded
}

func ExampleToUint64() {
	a := int8(42)
	i, err := safecast.ToUint64(a)
	fmt.Println(i, err)

	a = int8(-1)
	i, err = safecast.ToUint64(a)
	fmt.Println(i, err)

	// Output:
	// 42 <nil>
	// 0 conversion issue: -1 (int8) is less than 0 (uint64): minimum value for this type exceeded
}

func ExampleToUint() {
	a := int8(42)
	i, err := safecast.ToUint(a)
	fmt.Println(i, err)

	a = int8(-1)
	i, err = safecast.ToUint(a)
	fmt.Println(i, err)

	// Output:
	// 42 <nil>
	// 0 conversion issue: -1 (int8) is less than 0 (uint): minimum value for this type exceeded
}

func ExampleToFloat32() {
	a := int8(42)
	i, err := safecast.ToFloat32(a)
	fmt.Println(i, err)

	b := math.MaxFloat64
	i, err = safecast.ToFloat32(b)
	fmt.Println(i, err)

	// Output:
	// 42 <nil>
	// 0 conversion issue: 1.7976931348623157e+308 (float64) is greater than 3.4028235e+38 (float32): maximum value for this type exceeded
}

func ExampleToFloat64() {
	a := int8(42)
	i, err := safecast.ToFloat64(a)
	fmt.Println(i, err)

	b := math.MaxFloat64
	i, err = safecast.ToFloat64(b)
	fmt.Println(i, err)

	// Output:
	// 42 <nil>
	// 1.7976931348623157e+308 <nil>
}

func ExampleConvert() {
	b, err := safecast.Convert[int8](true)
	fmt.Println(b, err)

	b, err = safecast.Convert[int8]("true")
	fmt.Println(b, err)

	c, err := safecast.Convert[int16](17.1)
	fmt.Println(c, err)

	c, err = safecast.Convert[int16](int64(17))
	fmt.Println(c, err)

	d, err := safecast.Convert[int32]("17.1")
	fmt.Println(d, err)

	i, err := safecast.Convert[uint]("100_000")
	fmt.Println(i, err)

	i, err = safecast.Convert[uint]("0b11")
	fmt.Println(i, err)

	i, err = safecast.Convert[uint]("0x11")
	fmt.Println(i, err)

	a := int8(-1)
	i, err = safecast.Convert[uint](a)
	fmt.Println(i, err)

	i, err = safecast.Convert[uint]("-1")
	fmt.Println(i, err)

	i, err = safecast.Convert[uint]("abc")
	fmt.Println(i, err)

	i, err = safecast.Convert[uint]("-1.1")
	fmt.Println(i, err)

	i, err = safecast.Convert[uint](".12345E+5")
	fmt.Println(i, err)

	// Output:
	// 1 <nil>
	// 1 <nil>
	// 17 <nil>
	// 17 <nil>
	// 17 <nil>
	// 100000 <nil>
	// 3 <nil>
	// 17 <nil>
	// 0 conversion issue: -1 (int8) is less than 0 (uint): minimum value for this type exceeded
	// 0 conversion issue: -1 (int64) is less than 0 (uint): minimum value for this type exceeded
	// 0 conversion issue: cannot convert from string abc to uint
	// 0 conversion issue: -1.1 (float64) is less than 0 (uint): minimum value for this type exceeded
	// 12345 <nil>
}
