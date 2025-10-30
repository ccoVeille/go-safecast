package safecast_test

import (
	"fmt"
	"math"

	"github.com/ccoveille/go-safecast"
)

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
	// 0 conversion issue: cannot convert from `abc` to uint (base auto-detection)
	// 0 conversion issue: -1.1 (float64) is less than 0 (uint): minimum value for this type exceeded
	// 12345 <nil>
}
