# Go SafeCast

## Origin of this project

In Go, integer type conversion can lead to unexpected behavior and errors if not handled carefully.

Issues can happen when converting between signed and unsigned integers, or when converting to a smaller integer type.

The gosec project raised this to my attention when the gosec [G115 rule was added](https://github.com/securego/gosec/pull/1149)

> G115: Potential integer overflow when converting between integer types

This issue was way more complex than expected, and required multiple fixes.

## Example

This code seems OK

```go
package main

import (
  "fmt"
)

func main() {
  var a uint64
  a = 42
  b := int32(a)
  fmt.Println(b) // 42
}
```

But the conversion to int32 will behave differently depending on the value

```go
package main

import (
  "fmt"
)

func main() {
  var a uint64
  a = 2147483647
  b := int32(a)
  fmt.Println(b) // 2147483647

  a = 2147483647 + 1
  b = int32(a)
  fmt.Println(b) // -2147483648 Stack overflow

  c := -1
  d := uint32(c)
  fmt.Println(d) // 4294967295
}
```

GoPlay: [https://go.dev/play/p/9PRWI7e0x1T](https://go.dev/play/p/9PRWI7e0x1T)

## Motivation

The gosec G115 will now report issues in a lot of project.

Some libraries existed (See [alternatives](#alternatives) section), but they were not able to cover all the use cases.

## Alternatives

- [github.com/rung/go-safecast](https://github.com/rung/go-safecast):
  Unmaintained, not architecture agnostic, do not support uint -> int conversion

- [github.com/cybergarage/go-safecast](https://github.com/cybergarage/go-safecast)
  Work with pointer like json.Marshall