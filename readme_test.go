package safecast_test

import (
	"fmt"

	"github.com/ccoveille/go-safecast/v2"
)

func Example() {
	var a int64

	a = 42
	b, err := safecast.Convert[int16](a) // everything is fine because 42 fits in int16
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)

	a = 255 + 1
	_, err = safecast.Convert[uint8](a) // 256 is greater than uint8 maximum value
	if err != nil {
		fmt.Println(err)
	}

	a = -1
	_, err = safecast.Convert[uint64](a) // -1 cannot fit in uint64
	if err != nil {
		fmt.Println(err)
	}

	str := "\x99" // ASCII code 153 for Trademark symbol
	e := str[0]
	_, err = safecast.Convert[int8](e)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// 42
	// conversion issue: 256 (int64) is greater than 255 (uint8): maximum value for this type exceeded
	// conversion issue: -1 (int64) is less than 0 (uint64): minimum value for this type exceeded
	// conversion issue: 153 (uint8) is greater than 127 (int8): maximum value for this type exceeded
}
