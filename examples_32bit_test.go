//go:build i386 || arm

package safecast_test

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

import (
	"fmt"
	"math"

	"github.com/ccoveille/go-safecast"
)

func ExampleToInt() {
	a := uint64(42)
	i, err := safecast.ToInt(a)
	fmt.Println(i, err)

	b := float32(math.MaxFloat32)
	i, err = safecast.ToInt(b)
	fmt.Println(i, err)

	// Output:
	// 42 <nil>
	// 0 conversion issue: 3.4028235e+38 (float32) is greater than 2147483647 (int): maximum value for this type exceeded
}
