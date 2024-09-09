package safecast_test

import (
	"fmt"

	"github.com/ccoveille/go-safecast"
)

func Example() {
	var a int

	a = 42
	b, err := safecast.ToUint8(a) // everything is fine
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)

	a = 255 + 1
	_, err = safecast.ToUint8(a) // 256 is greater than uint8 maximum value
	if err != nil {
		fmt.Println(err)
	}

	a = -1
	_, err = safecast.ToUint8(a) // -1 cannot fit in uint8
	if err != nil {
		fmt.Println(err)
	}

	str := "\x99" // ASCII code 153 for Trademark symbol
	e := str[0]
	_, err = safecast.ToInt8(e)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// 42
	// conversion issue: 256 (int) is greater than 255 (uint8): maximum value for this type exceeded
	// conversion issue: -1 (int) is less than 0 (uint8): minimum value for this type exceeded
	// conversion issue: 153 (uint8) is greater than 127 (int8): maximum value for this type exceeded
}
