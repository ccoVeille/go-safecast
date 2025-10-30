package safecast_test

import (
	"fmt"

	"github.com/ccoveille/go-safecast/v2"
)

func ExampleConvert() {
	c, err := safecast.Convert[int16](17.1)
	fmt.Println(c, err)

	c, err = safecast.Convert[int16](int64(17))
	fmt.Println(c, err)

	a := int8(-1)
	i, err := safecast.Convert[uint16](a)
	fmt.Println(i, err)

	// Output:
	// 17 <nil>
	// 17 <nil>
	// 65535 conversion issue: -1 (int8) is less than 0 (uint16): minimum value for this type exceeded
}
