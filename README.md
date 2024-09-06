# Go SafeCast

[![Go Report Card](https://goreportcard.com/badge/github.com/ccoveille/go-safecast)](https://goreportcard.com/report/github.com/ccoveille/go-safecast)
[![GoDoc](https://godoc.org/github.com/ccoVeille/go-safecast?status.svg)](https://godoc.org/github.com/ccoVeille/go-safecast)
[![codecov](https://codecov.io/gh/ccoVeille/go-safecast/graph/badge.svg?token=VW0VO503U6)](https://codecov.io/gh/ccoVeille/go-safecast)
[![Code Climate](https://codeclimate.com/github/ccoVeille/go-safecast.png)](https://codeclimate.com/github/ccoVeille/go-safecast)

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
  fmt.Println(b) // -2147483648 integer overflow

  c := -1
  d := uint32(c)
  fmt.Println(d) // 4294967295
}
```

GoPlay: [https://go.dev/play/p/9PRWI7e0x1T](https://go.dev/play/p/9PRWI7e0x1T)

### What is the problem ?

[CWE-190](https://cwe.mitre.org/data/definitions/190.html) explains in detail.

But to sum it up, you can face:

- infinite loop
- access to wrong resource by id
- grant access to someone who exhausted their quota

## Motivation

The gosec G115 will now report issues in a lot of project.

## Alternatives

Some libraries existed, but they were not able to cover all the use cases.

- [github.com/rung/go-safecast](https://github.com/rung/go-safecast):
  Unmaintained, not architecture agnostic, do not support uint -> int conversion

- [github.com/cybergarage/go-safecast](https://github.com/cybergarage/go-safecast)
  Work with pointer like json.Marshall
