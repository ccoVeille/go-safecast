package safecast_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/ccoveille/go-safecast"
)

func ExampleParse() {
	d, err := safecast.Parse[int32]("17.1")
	fmt.Println(d, err)

	_, err = safecast.Parse[uint64]("-1")
	fmt.Println(err)

	_, err = safecast.Parse[uint64]("abc")
	fmt.Println(err)

	_, err = safecast.Parse[uint64]("-1.1")
	fmt.Println(err)

	i, err := safecast.Parse[uint64](".12345E+5")
	fmt.Println(i, err)

	// Output:
	// 17 <nil>
	// conversion issue: -1 (int64) is less than 0 (uint64): minimum value for this type exceeded
	// conversion issue: cannot convert from `abc` to uint64
	// conversion issue: -1.1 (float64) is less than 0 (uint64): minimum value for this type exceeded
	// 12345 <nil>
}

func ExampleParse_with_options() {
	i, err := safecast.Parse[uint64]("100_000", safecast.WithBaseAutoDetection())
	fmt.Println(i, err)

	i, err = safecast.Parse[uint64]("11")
	fmt.Println(i, err)

	i, err = safecast.Parse[uint64]("0b11", safecast.WithBaseAutoDetection())
	fmt.Println(i, err)

	i, err = safecast.Parse[uint64]("11", safecast.WithBaseBinary())
	fmt.Println(i, err)

	i, err = safecast.Parse[uint64]("0x11", safecast.WithBaseAutoDetection())
	fmt.Println(i, err)

	i, err = safecast.Parse[uint64]("11", safecast.WithBaseHexadecimal())
	fmt.Println(i, err)

	// Output:
	// 100000 <nil>
	// 11 <nil>
	// 3 <nil>
	// 3 <nil>
	// 17 <nil>
	// 17 <nil>
}

type MapTestParse[TypeOutput safecast.Number] struct {
	Input          string
	ParseOptions   []safecast.ParseOption
	ExpectedOutput TypeOutput
	ExpectedError  error
	ErrorContains  string
}

func (mt MapTestParse[O]) Run(t *testing.T) {
	t.Helper()

	// configure a helper to validate there is no panic
	defer func(t *testing.T) {
		t.Helper()

		err := recover()
		if err != nil {
			t.Fatalf("panic with %v", err)
		}
	}(t)

	out, err := safecast.Parse[O](mt.Input, mt.ParseOptions...)
	if mt.ExpectedError != nil {
		requireErrorIs(t, err, safecast.ErrConversionIssue)
		requireErrorIs(t, err, mt.ExpectedError)

		if mt.ErrorContains != "" {
			requireErrorContains(t, err, mt.ErrorContains)
		}

		return
	}

	assertNoError(t, err)
	assertEqual(t, mt.ExpectedOutput, out)
}

func TestParse(t *testing.T) {
	for name, c := range map[string]TestRunner{
		"integer":                   MapTestParse[uint]{Input: "42", ExpectedOutput: 42},
		"integer with decimal base": MapTestParse[uint]{Input: "42", ExpectedOutput: 42, ParseOptions: []safecast.ParseOption{safecast.WithBaseDecimal()}},
		"float":                     MapTestParse[uint]{Input: "42.0", ExpectedOutput: 42},

		"decimal separator without options": MapTestParse[uint]{Input: "10_0", ExpectedError: safecast.ErrStringConversion},
		"decimal separator with base auto":  MapTestParse[uint]{Input: "10_0", ExpectedOutput: 100, ParseOptions: []safecast.ParseOption{safecast.WithBaseAutoDetection()}},

		"binary prefix without options":            MapTestParse[uint]{Input: "0b101010", ExpectedError: safecast.ErrStringConversion},
		"binary prefix with base auto":             MapTestParse[uint]{Input: "0b101010", ExpectedOutput: 42, ParseOptions: []safecast.ParseOption{safecast.WithBaseAutoDetection()}},
		"binary without prefix with binary option": MapTestParse[uint]{Input: "101010", ExpectedOutput: 42, ParseOptions: []safecast.ParseOption{safecast.WithBaseBinary()}},

		"decimal with leading zero without option": MapTestParse[uint]{Input: "042", ExpectedOutput: 42},
		"short octal notation with base auto":      MapTestParse[uint]{Input: "042", ExpectedOutput: 34, ParseOptions: []safecast.ParseOption{safecast.WithBaseAutoDetection()}},
		"octal with base auto":                     MapTestParse[uint]{Input: "0o42", ExpectedOutput: 34, ParseOptions: []safecast.ParseOption{safecast.WithBaseAutoDetection()}},
		"octal without prefix with octal option":   MapTestParse[uint]{Input: "42", ExpectedOutput: 34, ParseOptions: []safecast.ParseOption{safecast.WithBaseOctal()}},

		"hexadecimal prefix without options":         MapTestParse[uint]{Input: "0x42", ExpectedError: safecast.ErrStringConversion},
		"hexadecimal prefix with base auto":          MapTestParse[uint]{Input: "0x42", ExpectedOutput: 66, ParseOptions: []safecast.ParseOption{safecast.WithBaseAutoDetection()}},
		"hexadecimal without prefix with hex option": MapTestParse[uint]{Input: "42", ExpectedOutput: 66, ParseOptions: []safecast.ParseOption{safecast.WithBaseHexadecimal()}},

		"boolean string":       MapTestParse[uint]{Input: "true", ExpectedError: safecast.ErrStringConversion},
		"short boolean string": MapTestParse[uint]{Input: "t", ExpectedError: safecast.ErrStringConversion},

		"integer trailing spaces": MapTestParse[uint]{
			Input:         "42 ",
			ExpectedError: safecast.ErrStringConversion,
			ErrorContains: "cannot convert from `42 ` to uint",
		},

		"empty string": MapTestParse[uint]{
			Input:         "",
			ExpectedError: safecast.ErrStringConversion,
			ErrorContains: "cannot convert from `` to uint",
		},
		"simple space": MapTestParse[uint]{
			Input:         " ",
			ExpectedError: safecast.ErrStringConversion,
			ErrorContains: "cannot convert from ` ` to uint",
		},
		"simple dot": MapTestParse[uint]{
			Input:         ".",
			ExpectedError: safecast.ErrStringConversion,
			ErrorContains: "cannot convert from `.` to uint"},
		"simple dash": MapTestParse[uint]{
			Input:         "-",
			ExpectedError: safecast.ErrStringConversion,
			ErrorContains: "cannot convert from `-` to uint"},
		"invalid string": MapTestParse[uint]{
			Input:         "abc",
			ExpectedError: safecast.ErrStringConversion,
			ErrorContains: "cannot convert from `abc` to uint"},
		"invalid string with dot": MapTestParse[uint]{
			Input:         "ab.c",
			ExpectedError: safecast.ErrStringConversion,
			ErrorContains: "cannot convert from `ab.c` to uint"},
		"strings with leading +": MapTestParse[uint]{
			Input:         "+42",
			ExpectedError: safecast.ErrStringConversion,
			ErrorContains: "cannot convert from `+42` to uint",
		},

		"integer with leading +":                 MapTestParse[uint]{Input: "+42", ExpectedError: safecast.ErrStringConversion},
		"float with leading plus":                MapTestParse[uint8]{Input: "+42.0", ExpectedOutput: 42},
		"invalid string multiple leading dashes": MapTestParse[uint]{Input: "--42", ExpectedError: safecast.ErrStringConversion},
		"invalid string with dash":               MapTestParse[uint]{Input: "-abc", ExpectedError: safecast.ErrStringConversion},
		"invalid string with dash and dot":       MapTestParse[uint]{Input: "-ab.c", ExpectedError: safecast.ErrStringConversion},
	} {
		t.Run(name, func(t *testing.T) {
			c.Run(t)
		})
	}

	for name, c := range map[string]TestRunner{
		"upper bound overflows for int string": MapTestParse[int]{
			Input:         "9223372036854775808", // math.MaxInt64 + 1
			ExpectedError: safecast.ErrExceedMaximumValue,
		},
		"upper bound overflows for int8 string": MapTestParse[int8]{
			Input:         "129", // math.MaxInt8 + 1
			ExpectedError: safecast.ErrExceedMaximumValue,
		},
		"upper bound overflows for int16 string": MapTestParse[int16]{
			Input:         "32769", // math.MaxInt16 + 1
			ExpectedError: safecast.ErrExceedMaximumValue,
		},
		"upper bound overflows for int32 string": MapTestParse[int32]{
			Input:         "2147483648", // math.MaxInt32 + 1
			ExpectedError: safecast.ErrExceedMaximumValue,
		},
		"upper bound overflows for int64 string": MapTestParse[int64]{
			Input:         "9223372036854775808", // math.MaxInt64 + 1
			ExpectedError: safecast.ErrExceedMaximumValue,
		},
		"upper bound overflows for int64 string overflow": MapTestParse[int64]{
			Input:         "123456789012345678901234567890", // more characters than math.MaxInt64 represented as string
			ExpectedError: safecast.ErrExceedMaximumValue,
		},
		"upper bound overflows with huge float": MapTestParse[float64]{
			Input:         "1.8446744073709552e+400", // much larger than [math.MaxFloat64]
			ExpectedError: safecast.ErrExceedMaximumValue,
		},
	} {
		t.Run(name, func(t *testing.T) {
			c.Run(t)
		})
	}

	for name, c := range map[string]TestRunner{
		"lower bound overflows int": MapTestParse[int]{
			Input:         "-9223372036854775809", // math.MinInt64 - 1
			ExpectedError: safecast.ErrExceedMinimumValue,
		},

		"lower bound overflows int8": MapTestParse[int8]{
			Input:         "-129", // math.MinInt8 - 1
			ExpectedError: safecast.ErrExceedMinimumValue,
		},

		"lower bound overflows int16": MapTestParse[int16]{
			Input:         "-32769", // math.MinInt16 - 1
			ExpectedError: safecast.ErrExceedMinimumValue,
		},

		"lower bound overflows int32": MapTestParse[int32]{
			Input:         "-2147483649", // math.MinInt32 - 1
			ExpectedError: safecast.ErrExceedMinimumValue,
		},

		"lower bound overflows int64": MapTestParse[int64]{
			Input:         "-9223372036854775809", // math.MinInt64 - 1
			ExpectedError: safecast.ErrExceedMinimumValue,
		},

		"lower bound overflows int64 overflow": MapTestParse[int64]{
			Input:         "-123456789012345678901234567890", // more characters than math.MinInt64 represented as string
			ExpectedError: safecast.ErrExceedMinimumValue,
		},

		"lower bound overflows with huge float": MapTestParse[float64]{
			Input:         "-1.8446744073709552e+400", // much larger than -[math.MaxFloat64]
			ExpectedError: safecast.ErrExceedMinimumValue,
		},

		"negative string overflows uint": MapTestParse[uint]{
			Input:         "-1",
			ExpectedError: safecast.ErrExceedMinimumValue,
		},

		"negative string overflows uint8": MapTestParse[uint8]{
			Input:         "-1",
			ExpectedError: safecast.ErrExceedMinimumValue,
		},

		"negative string overflows uint16": MapTestParse[uint16]{
			Input:         "-1",
			ExpectedError: safecast.ErrExceedMinimumValue,
		},

		"negative string overflows uint32": MapTestParse[uint32]{
			Input:         "-1",
			ExpectedError: safecast.ErrExceedMinimumValue,
		},

		"negative string overflows uint64": MapTestParse[uint64]{
			Input:         "-1",
			ExpectedError: safecast.ErrExceedMinimumValue,
		},
	} {
		t.Run(name, func(t *testing.T) {
			c.Run(t)
		})
	}

	for name, c := range map[string]TestRunner{
		"not a number with no options": MapTestParse[int8]{
			Input:         "foo",
			ExpectedError: safecast.ErrStringConversion,
		},
		"not a number with base decimal": MapTestParse[int8]{
			Input:         "foo",
			ParseOptions:  []safecast.ParseOption{safecast.WithBaseDecimal()},
			ExpectedError: safecast.ErrStringConversion,
		},
		"not a number with base binary": MapTestParse[int8]{
			Input:         "foo",
			ParseOptions:  []safecast.ParseOption{safecast.WithBaseBinary()},
			ExpectedError: safecast.ErrStringConversion,
			ErrorContains: "base binary",
		},
		"not a number with base octal": MapTestParse[int8]{
			Input:         "foo",
			ParseOptions:  []safecast.ParseOption{safecast.WithBaseOctal()},
			ExpectedError: safecast.ErrStringConversion,
			ErrorContains: "base octal",
		},
		"not a number with base hexadecimal": MapTestParse[int8]{
			Input:         "foo",
			ParseOptions:  []safecast.ParseOption{safecast.WithBaseHexadecimal()},
			ExpectedError: safecast.ErrStringConversion,
			ErrorContains: "base hexadecimal",
		},
		"not a number with base auto detection": MapTestParse[int8]{
			Input:         "foo",
			ParseOptions:  []safecast.ParseOption{safecast.WithBaseAutoDetection()},
			ExpectedError: safecast.ErrStringConversion,
			ErrorContains: "base auto-detection",
		},
	} {
		t.Run(name, func(t *testing.T) {
			c.Run(t)
		})
	}

	for name, c := range map[string]TestRunner{
		"binary overflows int8 with base 2": MapTestParse[int8]{
			Input:         "10000000", // 128 in base 2
			ParseOptions:  []safecast.ParseOption{safecast.WithBaseBinary()},
			ExpectedError: safecast.ErrExceedMaximumValue,
		},
		"binary overflows int8 with base auto detection": MapTestParse[int8]{
			Input:         "0b10000000", // 128 in base 2
			ParseOptions:  []safecast.ParseOption{safecast.WithBaseAutoDetection()},
			ExpectedError: safecast.ErrExceedMaximumValue,
		},
		"octal overflows int8 with base 8": MapTestParse[int8]{
			Input:         "200", // 128 in base 8
			ParseOptions:  []safecast.ParseOption{safecast.WithBaseOctal()},
			ExpectedError: safecast.ErrExceedMaximumValue,
		},
		"octal overflows int8 with base auto detection": MapTestParse[int8]{
			Input:         "0o200", // 128 in base 8
			ParseOptions:  []safecast.ParseOption{safecast.WithBaseAutoDetection()},
			ExpectedError: safecast.ErrExceedMaximumValue,
		},
		"short octal notation overflows int8 with base auto detection": MapTestParse[int8]{
			Input:         "0200", // 128 in base 8
			ParseOptions:  []safecast.ParseOption{safecast.WithBaseAutoDetection()},
			ExpectedError: safecast.ErrExceedMaximumValue,
		},
		"hexadecimal overflows int8 with base 16": MapTestParse[int8]{
			Input:         "80", // 128 in base 16
			ParseOptions:  []safecast.ParseOption{safecast.WithBaseHexadecimal()},
			ExpectedError: safecast.ErrExceedMaximumValue,
		},
		"hexadecimal overflows int8 with base auto detection": MapTestParse[int8]{
			Input:         "0x80", // 128 in base 16
			ParseOptions:  []safecast.ParseOption{safecast.WithBaseAutoDetection()},
			ExpectedError: safecast.ErrExceedMaximumValue,
		},
		"decimal overflows int8 with base 10": MapTestParse[int8]{
			Input:         "128", // 128 in base 10
			ParseOptions:  []safecast.ParseOption{safecast.WithBaseDecimal()},
			ExpectedError: safecast.ErrExceedMaximumValue,
		},
		"decimal overflows int8 with base auto detection": MapTestParse[int8]{
			Input:         "128", // 128 in base 10
			ParseOptions:  []safecast.ParseOption{safecast.WithBaseAutoDetection()},
			ExpectedError: safecast.ErrExceedMaximumValue,
		},
	} {
		t.Run(name, func(t *testing.T) {
			c.Run(t)
		})
	}

	t.Run("aliases", func(t *testing.T) {
		// Type aliases are handled separately

		type (
			// StringAlias is an alias
			StringAlias = string
		)

		for name, c := range map[string]TestRunner{
			"string simple alias": MapTestParse[int8]{
				Input:          StringAlias("42"),
				ExpectedOutput: int8(42),
			},
		} {
			t.Run(name, func(t *testing.T) {
				c.Run(t)
			})
		}
	})
}

type MapMustParseTest[TypeOutput safecast.Number] struct {
	Input          string
	ParseOptions   []safecast.ParseOption
	ExpectedOutput TypeOutput
	ExpectedError  error
}

func (mt MapMustParseTest[O]) Run(t *testing.T) {
	t.Helper()

	var out O
	defer func(t *testing.T) {
		t.Helper()

		assertEqual(t, mt.ExpectedOutput, out)

		r := recover()

		if mt.ExpectedError == nil && r == nil {

			return
		}

		if r == nil {
			t.Fatal("did not panic")
		}

		err, ok := r.(error)
		if !ok {
			t.Fatalf("panic value is not an error: %v", r)
		}

		if !errors.Is(err, safecast.ErrConversionIssue) {
			t.Fatalf("panic with unexpected error: %v", err)
		}

		if !errors.Is(err, mt.ExpectedError) {
			t.Fatalf("panic with unexpected error: %v", err)
		}
	}(t)

	out = safecast.MustParse[O](mt.Input, mt.ParseOptions...)
}

func TestMustParse(t *testing.T) {
	t.Run("panic on error", func(t *testing.T) {
		for name, tt := range map[string]TestRunner{
			"string": MapMustParseTest[uint8]{Input: "cats", ExpectedError: safecast.ErrStringConversion},
			"string with leading plus without options":          MapMustParseTest[uint8]{Input: "+1000", ExpectedError: safecast.ErrStringConversion},
			"string with leading plus with base auto detection": MapMustParseTest[uint8]{Input: "+1000", ExpectedError: safecast.ErrStringConversion, ParseOptions: []safecast.ParseOption{safecast.WithBaseAutoDetection()}},
			"binary_prefix_without_options":                     MapMustParseTest[uint8]{Input: "0b101010", ExpectedError: safecast.ErrStringConversion},
			"hexadecimal_prefix_without_options":                MapMustParseTest[uint8]{Input: "0x2A", ExpectedError: safecast.ErrStringConversion},
			"octal_prefix_without_options":                      MapMustParseTest[uint8]{Input: "0o52", ExpectedError: safecast.ErrStringConversion},
			"decimal_with_underscore_without_options":           MapMustParseTest[uint8]{Input: "10_000", ExpectedError: safecast.ErrStringConversion},
			"decimal_with_invalid_underscore_without_options":   MapMustParseTest[uint8]{Input: "100_", ExpectedError: safecast.ErrStringConversion},
		} {
			t.Run(name, func(t *testing.T) {
				tt.Run(t)
			})
		}
	})

	t.Run("panic on overflow", func(t *testing.T) {
		for name, tt := range map[string]TestRunner{
			"string": MapMustParseTest[uint8]{Input: "1000", ExpectedError: safecast.ErrExceedMaximumValue},
			"float with leading plus without options":          MapMustParseTest[uint8]{Input: "+1000.0", ExpectedError: safecast.ErrExceedMaximumValue},
			"float with leading plus with base auto detection": MapMustParseTest[uint8]{Input: "+1000.0", ExpectedError: safecast.ErrExceedMaximumValue, ParseOptions: []safecast.ParseOption{safecast.WithBaseAutoDetection()}},
			"binary_with_base_binary":                          MapMustParseTest[uint8]{Input: "100000000", ExpectedError: safecast.ErrExceedMaximumValue, ParseOptions: []safecast.ParseOption{safecast.WithBaseBinary()}},
			"binary_prefix_with_base_auto_detection":           MapMustParseTest[uint8]{Input: "0b100000000", ExpectedError: safecast.ErrExceedMaximumValue, ParseOptions: []safecast.ParseOption{safecast.WithBaseAutoDetection()}},
			"hexadecimal_with_base_hexadecimal":                MapMustParseTest[uint8]{Input: "1FF", ExpectedError: safecast.ErrExceedMaximumValue, ParseOptions: []safecast.ParseOption{safecast.WithBaseHexadecimal()}},
			"hexadecimal_prefix_with_base_auto_detection":      MapMustParseTest[uint8]{Input: "0x1FF", ExpectedError: safecast.ErrExceedMaximumValue, ParseOptions: []safecast.ParseOption{safecast.WithBaseAutoDetection()}},
			"octal_with_base_octal":                            MapMustParseTest[uint8]{Input: "400", ExpectedError: safecast.ErrExceedMaximumValue, ParseOptions: []safecast.ParseOption{safecast.WithBaseOctal()}},
			"octal_prefix_with_base_auto_detection":            MapMustParseTest[uint8]{Input: "0o400", ExpectedError: safecast.ErrExceedMaximumValue, ParseOptions: []safecast.ParseOption{safecast.WithBaseAutoDetection()}},
			"octal_legacy_prefix_with_base_auto_detection":     MapMustParseTest[uint8]{Input: "0400", ExpectedError: safecast.ErrExceedMaximumValue, ParseOptions: []safecast.ParseOption{safecast.WithBaseAutoDetection()}},
		} {
			t.Run(name, func(t *testing.T) {
				tt.Run(t)
			})
		}
	})

	t.Run("panic on underflow", func(t *testing.T) {
		for name, tt := range map[string]TestRunner{
			"negative integer": MapMustParseTest[uint8]{Input: "-1", ExpectedError: safecast.ErrExceedMinimumValue},
			"negative float":   MapMustParseTest[uint8]{Input: "-1.0", ExpectedError: safecast.ErrExceedMinimumValue},
		} {
			t.Run(name, func(t *testing.T) {
				tt.Run(t)
			})
		}
	})

	t.Run("no panic", func(t *testing.T) {
		for name, tt := range map[string]TestRunner{
			"string":                    MapMustParseTest[uint8]{Input: "42", ExpectedOutput: 42},
			"float":                     MapMustParseTest[uint8]{Input: "42.0", ExpectedOutput: 42},
			"float with leading plus":   MapMustParseTest[uint8]{Input: "+42.0", ExpectedOutput: 42},
			"decimal_with_zero_padding": MapMustParseTest[uint8]{Input: "042", ExpectedOutput: 42},
			"binary":                    MapMustParseTest[uint8]{Input: "101010", ExpectedOutput: 42, ParseOptions: []safecast.ParseOption{safecast.WithBaseBinary()}},
			"binary_with_prefix":        MapMustParseTest[uint8]{Input: "0b101010", ExpectedOutput: 42, ParseOptions: []safecast.ParseOption{safecast.WithBaseAutoDetection()}},
			"hexadecimal":               MapMustParseTest[uint8]{Input: "2A", ExpectedOutput: 42, ParseOptions: []safecast.ParseOption{safecast.WithBaseHexadecimal()}},
			"hexadecimal_with_prefix_and_auto_detection": MapMustParseTest[uint8]{Input: "0x2A", ExpectedOutput: 42, ParseOptions: []safecast.ParseOption{safecast.WithBaseAutoDetection()}},
			"octal":                                       MapMustParseTest[uint8]{Input: "52", ExpectedOutput: 42, ParseOptions: []safecast.ParseOption{safecast.WithBaseOctal()}},
			"octal_with_zero_padding":                     MapMustParseTest[uint8]{Input: "042", ExpectedOutput: 34, ParseOptions: []safecast.ParseOption{safecast.WithBaseOctal()}},
			"octal_with_prefix_and_auto_detection":        MapMustParseTest[uint8]{Input: "0o52", ExpectedOutput: 42, ParseOptions: []safecast.ParseOption{safecast.WithBaseAutoDetection()}},
			"octal_with_legacy_prefix_and_auto_detection": MapMustParseTest[uint8]{Input: "052", ExpectedOutput: 42, ParseOptions: []safecast.ParseOption{safecast.WithBaseAutoDetection()}},
		} {
			t.Run(name, func(t *testing.T) {
				tt.Run(t)
			})
		}
	})
}

type MapRequireParseTest[TypeOutput safecast.Number] struct {
	Input               string
	ParseOptions        []safecast.ParseOption
	ExpectedOutput      TypeOutput
	ExpectedTestFailure bool
}

func (mt MapRequireParseTest[O]) Run(t *testing.T) {
	t.Helper()

	// We need to use a fake testing.T to avoid the test failing when we expect a failure
	fakeT := new(mockTestingT)
	out := safecast.RequireParse[O](fakeT, mt.Input, mt.ParseOptions...)

	if !mt.ExpectedTestFailure {
		// only check output if we do not expect a failure
		assertEqual(t, mt.ExpectedOutput, out)
	}

	if mt.ExpectedTestFailure == fakeT.Failed() {
		return // test passed as expected
	}

	if mt.ExpectedTestFailure {
		t.Fatal("test should have failed, but it did not")
	}
	t.Error("test should not have failed, but it did")
	// we know that the test failed, so we use the real testing.T to report the error
	_ = safecast.RequireParse[O](t, mt.Input, mt.ParseOptions...)
}

func TestRequireParse(t *testing.T) {
	t.Run("no conversion error", func(t *testing.T) {
		for name, tt := range map[string]TestRunner{
			"string":                    MapRequireParseTest[uint8]{Input: "42", ExpectedOutput: 42},
			"float":                     MapRequireParseTest[uint8]{Input: "42.0", ExpectedOutput: 42},
			"float with leading plus":   MapRequireParseTest[uint8]{Input: "+42.0", ExpectedOutput: 42},
			"decimal_with_zero_padding": MapRequireParseTest[uint8]{Input: "042", ExpectedOutput: 42},
			"binary":                    MapRequireParseTest[uint8]{Input: "101010", ExpectedOutput: 42, ParseOptions: []safecast.ParseOption{safecast.WithBaseBinary()}},
			"binary_with_prefix":        MapRequireParseTest[uint8]{Input: "0b101010", ExpectedOutput: 42, ParseOptions: []safecast.ParseOption{safecast.WithBaseAutoDetection()}},
			"hexadecimal":               MapRequireParseTest[uint8]{Input: "2A", ExpectedOutput: 42, ParseOptions: []safecast.ParseOption{safecast.WithBaseHexadecimal()}},
			"hexadecimal_with_prefix_and_auto_detection": MapRequireParseTest[uint8]{Input: "0x2A", ExpectedOutput: 42, ParseOptions: []safecast.ParseOption{safecast.WithBaseAutoDetection()}},
			"octal":                                       MapRequireParseTest[uint8]{Input: "52", ExpectedOutput: 42, ParseOptions: []safecast.ParseOption{safecast.WithBaseOctal()}},
			"octal_with_zero_padding":                     MapRequireParseTest[uint8]{Input: "042", ExpectedOutput: 34, ParseOptions: []safecast.ParseOption{safecast.WithBaseOctal()}},
			"octal_with_prefix_and_auto_detection":        MapRequireParseTest[uint8]{Input: "0o52", ExpectedOutput: 42, ParseOptions: []safecast.ParseOption{safecast.WithBaseAutoDetection()}},
			"octal_with_legacy_prefix_and_auto_detection": MapRequireParseTest[uint8]{Input: "052", ExpectedOutput: 42, ParseOptions: []safecast.ParseOption{safecast.WithBaseAutoDetection()}},
		} {
			t.Run(name, func(t *testing.T) {
				tt.Run(t)
			})
		}
	})

	t.Run("test fail on error", func(t *testing.T) {
		for name, tt := range map[string]TestRunner{
			"string": MapMustParseTest[uint8]{Input: "cats", ExpectedError: safecast.ErrStringConversion},
			"string with leading plus without options":          MapRequireParseTest[uint8]{Input: "+1000", ExpectedTestFailure: true},
			"string with leading plus with base auto detection": MapRequireParseTest[uint8]{Input: "+1000", ExpectedTestFailure: true, ParseOptions: []safecast.ParseOption{safecast.WithBaseAutoDetection()}},
			"binary_prefix_without_options":                     MapRequireParseTest[uint8]{Input: "0b101010", ExpectedTestFailure: true},
			"hexadecimal_prefix_without_options":                MapRequireParseTest[uint8]{Input: "0x2A", ExpectedTestFailure: true},
			"octal_prefix_without_options":                      MapRequireParseTest[uint8]{Input: "0o52", ExpectedTestFailure: true},
			"decimal_with_underscore_without_options":           MapRequireParseTest[uint8]{Input: "10_000", ExpectedTestFailure: true},
			"decimal_with_invalid_underscore_without_options":   MapRequireParseTest[uint8]{Input: "100_", ExpectedTestFailure: true},
		} {
			t.Run(name, func(t *testing.T) {
				tt.Run(t)
			})
		}
	})

	t.Run("test fail on overflow", func(t *testing.T) {
		for name, tt := range map[string]TestRunner{
			"string": MapRequireParseTest[uint8]{Input: "1000", ExpectedTestFailure: true},
			"float with leading plus without options":          MapRequireParseTest[uint8]{Input: "+1000.0", ExpectedTestFailure: true},
			"float with leading plus with base auto detection": MapRequireParseTest[uint8]{Input: "+1000.0", ExpectedTestFailure: true, ParseOptions: []safecast.ParseOption{safecast.WithBaseAutoDetection()}},
			"binary_with_base_binary":                          MapRequireParseTest[uint8]{Input: "100000000", ExpectedTestFailure: true, ParseOptions: []safecast.ParseOption{safecast.WithBaseBinary()}},
			"binary_prefix_with_base_auto_detection":           MapRequireParseTest[uint8]{Input: "0b100000000", ExpectedTestFailure: true, ParseOptions: []safecast.ParseOption{safecast.WithBaseAutoDetection()}},
			"hexadecimal_with_base_hexadecimal":                MapRequireParseTest[uint8]{Input: "1FF", ExpectedTestFailure: true, ParseOptions: []safecast.ParseOption{safecast.WithBaseHexadecimal()}},
			"hexadecimal_prefix_with_base_auto_detection":      MapRequireParseTest[uint8]{Input: "0x1FF", ExpectedTestFailure: true, ParseOptions: []safecast.ParseOption{safecast.WithBaseAutoDetection()}},
			"octal_with_base_octal":                            MapRequireParseTest[uint8]{Input: "400", ExpectedTestFailure: true, ParseOptions: []safecast.ParseOption{safecast.WithBaseOctal()}},
			"octal_prefix_with_base_auto_detection":            MapRequireParseTest[uint8]{Input: "0o400", ExpectedTestFailure: true, ParseOptions: []safecast.ParseOption{safecast.WithBaseAutoDetection()}},
			"octal_legacy_prefix_with_base_auto_detection":     MapRequireParseTest[uint8]{Input: "0400", ExpectedTestFailure: true, ParseOptions: []safecast.ParseOption{safecast.WithBaseAutoDetection()}},
		} {
			t.Run(name, func(t *testing.T) {
				tt.Run(t)
			})
		}
	})

	t.Run("test fail on underflow", func(t *testing.T) {
		for name, tt := range map[string]TestRunner{
			"negative integer": MapRequireParseTest[uint8]{Input: "-1", ExpectedTestFailure: true},
			"negative float":   MapRequireParseTest[uint8]{Input: "-1.0", ExpectedTestFailure: true},
		} {
			t.Run(name, func(t *testing.T) {
				tt.Run(t)
			})
		}
	})
}

func ExampleWithBaseAutoDetection() {

	for _, str := range []string{
		"100_000",
		"0b101010", "101010",
		"0o52", "052", "52",
		"0x2A", "2A",
	} {
		i, err := safecast.Parse[uint64](str, safecast.WithBaseAutoDetection())
		fmt.Printf("%-10s => %d %v\n", str, i, err)
	}

	// Output:
	// 100_000    => 100000 <nil>
	// 0b101010   => 42 <nil>
	// 101010     => 101010 <nil>
	// 0o52       => 42 <nil>
	// 052        => 42 <nil>
	// 52         => 52 <nil>
	// 0x2A       => 42 <nil>
	// 2A         => 0 conversion issue: cannot convert from `2A` to uint64 (base auto-detection)
}

func ExampleWithBaseBinary() {

	for _, str := range []string{
		"101010",
		"0b101010",
	} {
		i, err := safecast.Parse[uint64](str, safecast.WithBaseBinary())
		fmt.Printf("%-10s => %d %v\n", str, i, err)
	}

	// Output:
	// 101010     => 42 <nil>
	// 0b101010   => 0 conversion issue: cannot convert from `0b101010` to uint64 (base binary)
}

func ExampleWithBaseOctal() {

	for _, str := range []string{
		"52",
		"0o52",
		"052",
	} {
		i, err := safecast.Parse[uint64](str, safecast.WithBaseOctal())
		fmt.Printf("%-10s => %d %v\n", str, i, err)
	}

	// Output:
	// 52         => 42 <nil>
	// 0o52       => 0 conversion issue: cannot convert from `0o52` to uint64 (base octal)
	// 052        => 42 <nil>
}

func ExampleWithBaseHexadecimal() {

	for _, str := range []string{
		"2A",
		"0x2A",
	} {
		i, err := safecast.Parse[uint64](str, safecast.WithBaseHexadecimal())
		fmt.Printf("%-10s => %d %v\n", str, i, err)
	}

	// Output:
	// 2A         => 42 <nil>
	// 0x2A       => 0 conversion issue: cannot convert from `0x2A` to uint64 (base hexadecimal)
}

func ExampleWithBaseDecimal() {

	for _, str := range []string{
		"100_000",
		"0b101010",
		"42",
		"0x2A",
	} {
		i, err := safecast.Parse[uint64](str, safecast.WithBaseDecimal())
		fmt.Printf("%-10s => %d %v\n", str, i, err)
	}

	// Output:
	// 100_000    => 0 conversion issue: cannot convert from `100_000` to uint64
	// 0b101010   => 0 conversion issue: cannot convert from `0b101010` to uint64
	// 42         => 42 <nil>
	// 0x2A       => 0 conversion issue: cannot convert from `0x2A` to uint64
}
