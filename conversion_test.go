package safecast_test

import (
	"errors"
	"math"
	"strings"
	"testing"

	"github.com/ccoveille/go-safecast"
)

func assertEqual[V comparable](t *testing.T, expected, got V) {
	t.Helper()

	if expected == got {
		return
	}

	t.Errorf("Not equal: \n"+
		"expected: %v (%T)\n"+
		"actual  : %v (%T)", expected, expected, got, got)
}

func requireError(t *testing.T, err error) {
	t.Helper()

	if err == nil {
		t.Fatal("expected error")
	}
}

func requireErrorIs(t *testing.T, err error, expected error) {
	t.Helper()
	requireError(t, err)

	if !errors.Is(err, expected) {
		t.Fatalf("unexpected error got %v, expected %v", err, expected)
	}
}

func requireErrorContains(t *testing.T, err error, text string) {
	t.Helper()
	requireErrorIs(t, err, safecast.ErrConversionIssue)

	errMessage := err.Error()
	if !strings.Contains(errMessage, text) {
		t.Fatalf("error message should contain %q: %q", text, errMessage)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

type caseInt8[in safecast.Type] struct {
	name  string
	input in
	want  int8
}

func assertInt8OK[in safecast.Type](t *testing.T, tests []caseInt8[in]) {
	t.Helper()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := safecast.ToInt8(tt.input)
			assertNoError(t, err)
			assertEqual(t, tt.want, got)
		})
	}
}

func assertInt8Error[in safecast.Type](t *testing.T, tests []caseInt8[in]) {
	t.Helper()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := safecast.ToInt8(tt.input)
			requireErrorIs(t, err, safecast.ErrConversionIssue)
			assertEqual(t, tt.want, got)
		})
	}
}

func TestToInt8(t *testing.T) {
	t.Run("from int", func(t *testing.T) {
		assertInt8OK(t, []caseInt8[int]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
			{name: "negative within range", input: -100, want: -100},
		})

		assertInt8Error(t, []caseInt8[int]{
			{name: "positive out of range", input: math.MaxInt8 + 1},
			{name: "negative out of range", input: math.MinInt8 - 1},
		})
	})

	t.Run("from int8", func(t *testing.T) {
		assertInt8OK(t, []caseInt8[int8]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
			{name: "negative within range", input: -100, want: -100},
		})
	})

	t.Run("from int16", func(t *testing.T) {
		assertInt8OK(t, []caseInt8[int16]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
			{name: "negative within range", input: -100, want: -100},
		})

		assertInt8Error(t, []caseInt8[int16]{
			{name: "positive out of range", input: math.MaxInt8 + 1},
			{name: "negative out of range", input: math.MinInt8 - 1},
		})
	})

	t.Run("from int32", func(t *testing.T) {
		assertInt8OK(t, []caseInt8[int32]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
			{name: "negative within range", input: -100, want: -100},
		})

		assertInt8Error(t, []caseInt8[int32]{
			{name: "positive out of range", input: math.MaxInt8 + 1},
			{name: "negative out of range", input: math.MinInt8 - 1},
		})
	})

	t.Run("from int64", func(t *testing.T) {
		assertInt8OK(t, []caseInt8[int64]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
			{name: "negative within range", input: -100, want: -100},
		})

		assertInt8Error(t, []caseInt8[int64]{
			{name: "positive out of range", input: math.MaxInt8 + 1},
			{name: "negative out of range", input: math.MinInt8 - 1},
		})
	})

	t.Run("from uint", func(t *testing.T) {
		assertInt8OK(t, []caseInt8[uint]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertInt8Error(t, []caseInt8[uint]{
			{name: "positive out of range", input: math.MaxInt8 + 1},
		})
	})

	t.Run("from uint8", func(t *testing.T) {
		assertInt8OK(t, []caseInt8[uint8]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})
	})

	t.Run("from uint16", func(t *testing.T) {
		assertInt8OK(t, []caseInt8[uint16]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertInt8Error(t, []caseInt8[uint16]{
			{name: "positive out of range", input: math.MaxInt8 + 1},
		})
	})

	t.Run("from uint32", func(t *testing.T) {
		assertInt8OK(t, []caseInt8[uint32]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertInt8Error(t, []caseInt8[uint32]{
			{name: "positive out of range", input: math.MaxInt8 + 1},
		})
	})

	t.Run("from uint64", func(t *testing.T) {
		assertInt8OK(t, []caseInt8[uint64]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertInt8Error(t, []caseInt8[uint64]{
			{name: "positive out of range", input: math.MaxInt8 + 1},
		})
	})

	t.Run("from float32", func(t *testing.T) {
		assertInt8OK(t, []caseInt8[float32]{
			{name: "zero", input: 0.0, want: 0},
			{name: "rounded value", input: 1.1, want: 1},
		})

		assertInt8Error(t, []caseInt8[float32]{
			{name: "positive out of range", input: math.MaxInt8 + 1},
		})
	})

	t.Run("from float64", func(t *testing.T) {
		assertInt8OK(t, []caseInt8[float64]{
			{name: "zero", input: 0.0, want: 0},
			{name: "rounded value", input: 1.1, want: 1},
		})

		assertInt8Error(t, []caseInt8[float64]{
			{name: "positive out of range", input: math.MaxInt8 + 1},
		})
	})
}

func TestErrorMessage(t *testing.T) {
	_, err := safecast.ToUint8(-1)
	requireErrorIs(t, err, safecast.ErrConversionIssue)
	requireErrorIs(t, err, safecast.ErrExceedMinimumValue)
	requireErrorContains(t, err, "than 0 (uint8)")

	_, err = safecast.ToUint8(math.MaxInt16)
	requireErrorIs(t, err, safecast.ErrConversionIssue)
	requireErrorIs(t, err, safecast.ErrExceedMaximumValue)
	requireErrorContains(t, err, "than 255 (uint8)")

	_, err = safecast.ToInt8(-math.MaxInt16)
	requireErrorIs(t, err, safecast.ErrConversionIssue)
	requireErrorIs(t, err, safecast.ErrExceedMinimumValue)
	requireErrorContains(t, err, "than -128 (int8)")
}

type caseUint8[in safecast.Type] struct {
	name  string
	input in
	want  uint8
}

func assertUint8OK[in safecast.Type](t *testing.T, tests []caseUint8[in]) {
	t.Helper()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := safecast.ToUint8(tt.input)
			assertNoError(t, err)
			assertEqual(t, tt.want, got)
		})
	}
}

func assertUint8Error[in safecast.Type](t *testing.T, tests []caseUint8[in]) {
	t.Helper()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Helper()

			got, err := safecast.ToUint8(tt.input)
			requireErrorIs(t, err, safecast.ErrConversionIssue)
			assertEqual(t, tt.want, got)
		})
	}
}

func TestToUint8(t *testing.T) {
	t.Run("from int", func(t *testing.T) {
		assertUint8OK(t, []caseUint8[int]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertUint8Error(t, []caseUint8[int]{
			{name: "positive out of range", input: math.MaxUint8 + 1},
			{name: "negative value", input: -1},
		})
	})

	t.Run("from int8", func(t *testing.T) {
		assertUint8OK(t, []caseUint8[int8]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertUint8Error(t, []caseUint8[int8]{
			{name: "negative value", input: -1},
		})
	})

	t.Run("from int16", func(t *testing.T) {
		assertUint8OK(t, []caseUint8[int16]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertUint8Error(t, []caseUint8[int16]{
			{name: "positive out of range", input: 10000},
			{name: "negative value", input: -1},
		})
	})

	t.Run("from int32", func(t *testing.T) {
		assertUint8OK(t, []caseUint8[int32]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertUint8Error(t, []caseUint8[int32]{
			{name: "positive out of range", input: 100000},
			{name: "negative value", input: -1},
		})
	})

	t.Run("from int64", func(t *testing.T) {
		assertUint8OK(t, []caseUint8[int64]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertUint8Error(t, []caseUint8[int64]{
			{name: "positive out of range", input: 100000},
			{name: "negative value", input: -1},
		})
	})

	t.Run("from uint", func(t *testing.T) {
		assertUint8OK(t, []caseUint8[uint]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertUint8Error(t, []caseUint8[uint]{
			{name: "positive out of range", input: math.MaxUint8 + 1},
		})
	})

	t.Run("from uint8", func(t *testing.T) {
		assertUint8OK(t, []caseUint8[uint8]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})
	})

	t.Run("from uint16", func(t *testing.T) {
		assertUint8OK(t, []caseUint8[uint16]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertUint8Error(t, []caseUint8[uint]{
			{name: "positive out of range", input: math.MaxUint8 + 1},
		})
	})

	t.Run("from uint32", func(t *testing.T) {
		assertUint8OK(t, []caseUint8[uint32]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertUint8Error(t, []caseUint8[uint]{
			{name: "positive out of range", input: math.MaxUint8 + 1},
		})
	})

	t.Run("from uint64", func(t *testing.T) {
		assertUint8OK(t, []caseUint8[uint64]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertUint8Error(t, []caseUint8[uint]{
			{name: "positive out of range", input: math.MaxUint8 + 1},
		})
	})

	t.Run("from float32", func(t *testing.T) {
		assertUint8OK(t, []caseUint8[float32]{
			{name: "zero", input: 0.0, want: 0},
			{name: "rounded value", input: 1.1, want: 1},
			{name: "positive within range", input: 100.9, want: 100},
		})

		assertUint8Error(t, []caseUint8[float32]{
			{name: "positive out of range", input: math.MaxUint8 + 1},
		})
	})

	t.Run("from float64", func(t *testing.T) {
		assertUint8OK(t, []caseUint8[float64]{
			{name: "zero", input: 0.0, want: 0},
			{name: "rounded value", input: 1.1, want: 1},
			{name: "positive within range", input: 100.9, want: 100},
		})

		assertUint8Error(t, []caseUint8[float64]{
			{name: "positive out of range", input: math.MaxUint8 + 1},
		})
	})
}

type caseInt16[in safecast.Type] struct {
	name  string
	input in
	want  int16
}

func assertInt16OK[in safecast.Type](t *testing.T, tests []caseInt16[in]) {
	t.Helper()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := safecast.ToInt16(tt.input)
			assertNoError(t, err)
			assertEqual(t, tt.want, got)
		})
	}
}

func assertInt16Error[in safecast.Type](t *testing.T, tests []caseInt16[in]) {
	t.Helper()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := safecast.ToInt16(tt.input)
			requireErrorIs(t, err, safecast.ErrConversionIssue)
			assertEqual(t, tt.want, got)
		})
	}
}

func TestToInt16(t *testing.T) {
	t.Run("from int", func(t *testing.T) {
		assertInt16OK(t, []caseInt16[int]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
			{name: "negative within range", input: -10000, want: -10000},
		})

		assertInt16Error(t, []caseInt16[int]{
			{name: "positive out of range", input: math.MaxInt16 + 1},
			{name: "negative out of range", input: math.MinInt16 - 1},
		})
	})

	t.Run("from int8", func(t *testing.T) {
		assertInt16OK(t, []caseInt16[int8]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
			{name: "negative within range", input: -100, want: -100},
		})
	})

	t.Run("from int16", func(t *testing.T) {
		assertInt16OK(t, []caseInt16[int16]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
			{name: "negative within range", input: -10000, want: -10000},
		})
	})

	t.Run("from int32", func(t *testing.T) {
		assertInt16OK(t, []caseInt16[int32]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
			{name: "negative within range", input: -10000, want: -10000},
		})

		assertInt16Error(t, []caseInt16[int32]{
			{name: "positive out of range", input: 100000},
			{name: "negative out of range", input: -100000},
		})
	})

	t.Run("from int64", func(t *testing.T) {
		assertInt16OK(t, []caseInt16[int64]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
			{name: "negative within range", input: -10000, want: -10000},
		})

		assertInt16Error(t, []caseInt16[int64]{
			{name: "positive out of range", input: 100000},
			{name: "negative out of range", input: -100000},
		})
	})

	t.Run("from uint", func(t *testing.T) {
		assertInt16OK(t, []caseInt16[uint]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertInt16Error(t, []caseInt16[uint]{
			{name: "positive out of range", input: math.MaxInt16 + 1},
		})
	})

	t.Run("from uint8", func(t *testing.T) {
		assertInt16OK(t, []caseInt16[uint8]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})
	})

	t.Run("from uint16", func(t *testing.T) {
		assertInt16OK(t, []caseInt16[uint16]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertInt16Error(t, []caseInt16[uint16]{
			{name: "positive out of range", input: math.MaxInt16 + 1},
		})
	})

	t.Run("from uint32", func(t *testing.T) {
		assertInt16OK(t, []caseInt16[uint32]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertInt16Error(t, []caseInt16[uint32]{
			{name: "positive out of range", input: math.MaxInt16 + 1},
		})
	})

	t.Run("from uint64", func(t *testing.T) {
		assertInt16OK(t, []caseInt16[uint64]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertInt16Error(t, []caseInt16[uint64]{
			{name: "positive out of range", input: math.MaxInt16 + 1},
		})
	})

	t.Run("from float32", func(t *testing.T) {
		assertInt16OK(t, []caseInt16[float32]{
			{name: "zero", input: 0.0, want: 0},
			{name: "rounded value", input: 1.1, want: 1},
			{name: "positive within range", input: 10000, want: 10000},
		})

		assertInt16Error(t, []caseInt16[float32]{
			{name: "positive out of range", input: math.MaxInt16 + 1},
		})
	})

	t.Run("from float64", func(t *testing.T) {
		assertInt16OK(t, []caseInt16[float64]{
			{name: "zero", input: 0.0, want: 0},
			{name: "rounded value", input: 1.1, want: 1},
			{name: "positive within range", input: 10000, want: 10000},
		})

		assertInt16Error(t, []caseInt16[float64]{
			{name: "positive out of range", input: math.MaxInt16 + 1},
		})
	})
}

type caseUint16[in safecast.Type] struct {
	name  string
	input in
	want  uint16
}

func assertUint16OK[in safecast.Type](t *testing.T, tests []caseUint16[in]) {
	t.Helper()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := safecast.ToUint16(tt.input)
			assertNoError(t, err)
			assertEqual(t, tt.want, got)
		})
	}
}

func assertUint16Error[in safecast.Type](t *testing.T, tests []caseUint16[in]) {
	t.Helper()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Helper()

			got, err := safecast.ToUint16(tt.input)
			requireErrorIs(t, err, safecast.ErrConversionIssue)
			assertEqual(t, tt.want, got)
		})
	}
}

func TestToUint16(t *testing.T) {
	t.Run("from int", func(t *testing.T) {
		assertUint16OK(t, []caseUint16[int]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
		})

		assertUint16Error(t, []caseUint16[int]{
			{name: "positive out of range", input: math.MaxUint16 + 1},
			{name: "negative value", input: -1},
		})
	})

	t.Run("from int8", func(t *testing.T) {
		assertUint16OK(t, []caseUint16[int8]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})
	})

	t.Run("from int16", func(t *testing.T) {
		assertUint16OK(t, []caseUint16[int16]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
		})
	})

	t.Run("from int32", func(t *testing.T) {
		assertUint16OK(t, []caseUint16[int32]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
		})

		assertUint16Error(t, []caseUint16[int32]{
			{name: "positive out of range", input: 100000},
			{name: "negative value", input: -1},
		})
	})

	t.Run("from int64", func(t *testing.T) {
		assertUint16OK(t, []caseUint16[int64]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
		})

		assertUint16Error(t, []caseUint16[int64]{
			{name: "positive out of range", input: 100000},
			{name: "negative value", input: -1},
		})
	})

	t.Run("from uint", func(t *testing.T) {
		assertUint16OK(t, []caseUint16[uint]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertUint16Error(t, []caseUint16[uint]{
			{name: "positive out of range", input: math.MaxUint16 + 1},
		})
	})

	t.Run("from uint8", func(t *testing.T) {
		assertUint16OK(t, []caseUint16[uint8]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})
	})

	t.Run("from uint16", func(t *testing.T) {
		assertUint16OK(t, []caseUint16[uint16]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})
	})

	t.Run("from uint32", func(t *testing.T) {
		assertUint16OK(t, []caseUint16[uint32]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertUint16Error(t, []caseUint16[uint32]{
			{name: "positive out of range", input: math.MaxUint16 + 1},
		})
	})

	t.Run("from uint64", func(t *testing.T) {
		assertUint16OK(t, []caseUint16[uint64]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertUint16Error(t, []caseUint16[uint64]{
			{name: "positive out of range", input: math.MaxUint16 + 1},
		})
	})

	t.Run("from float32", func(t *testing.T) {
		assertUint16OK(t, []caseUint16[float32]{
			{name: "zero", input: 0.0, want: 0},
			{name: "rounded value", input: 1.1, want: 1},
			{name: "positive within range", input: 10000, want: 10000},
		})

		assertUint16Error(t, []caseUint16[float32]{
			{name: "positive out of range", input: math.MaxUint16 + 1},
		})
	})

	t.Run("from float64", func(t *testing.T) {
		assertUint16OK(t, []caseUint16[float64]{
			{name: "zero", input: 0.0, want: 0},
			{name: "rounded value", input: 1.1, want: 1},
			{name: "positive within range", input: 10000, want: 10000},
		})

		assertUint16Error(t, []caseUint16[float32]{
			{name: "positive out of range", input: math.MaxUint16 + 1},
		})
	})
}

type caseInt32[in safecast.Type] struct {
	name  string
	input in
	want  int32
}

func assertInt32OK[in safecast.Type](t *testing.T, tests []caseInt32[in]) {
	t.Helper()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := safecast.ToInt32(tt.input)
			assertNoError(t, err)
			assertEqual(t, tt.want, got)
		})
	}
}

func assertInt32Error[in safecast.Type](t *testing.T, tests []caseInt32[in]) {
	t.Helper()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := safecast.ToInt32(tt.input)
			requireErrorIs(t, err, safecast.ErrConversionIssue)
			assertEqual(t, tt.want, got)
		})
	}
}

func TestToInt32(t *testing.T) {
	t.Run("from int", func(t *testing.T) {
		assertInt32OK(t, []caseInt32[int]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
			{name: "negative within range", input: -10000, want: -10000},
		})

		assertInt32Error(t, []caseInt32[int]{
			{name: "positive out of range", input: math.MaxInt32 + 1},
			{name: "negative out of range", input: math.MinInt32 - 1},
		})
	})

	t.Run("from int8", func(t *testing.T) {
		assertInt32OK(t, []caseInt32[int8]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
			{name: "negative within range", input: -100, want: -100},
		})
	})

	t.Run("from int16", func(t *testing.T) {
		assertInt32OK(t, []caseInt32[int16]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
			{name: "negative within range", input: -10000, want: -10000},
		})
	})

	t.Run("from int32", func(t *testing.T) {
		assertInt32OK(t, []caseInt32[int32]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
			{name: "negative within range", input: -10000, want: -10000},
		})
	})

	t.Run("from int64", func(t *testing.T) {
		assertInt32OK(t, []caseInt32[int64]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
			{name: "negative within range", input: -10000, want: -10000},
		})

		assertInt32Error(t, []caseInt32[int64]{
			{name: "positive out of range", input: math.MaxInt32 + 1},
			{name: "negative out of range", input: math.MinInt32 - 1},
		})
	})

	t.Run("from uint", func(t *testing.T) {
		assertInt32OK(t, []caseInt32[uint]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertInt32Error(t, []caseInt32[uint]{
			{name: "positive out of range", input: math.MaxInt32 + 1},
		})
	})

	t.Run("from uint8", func(t *testing.T) {
		assertInt32OK(t, []caseInt32[uint8]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})
	})

	t.Run("from uint16", func(t *testing.T) {
		assertInt32OK(t, []caseInt32[uint16]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: math.MaxUint16, want: math.MaxUint16},
		})
	})

	t.Run("from uint32", func(t *testing.T) {
		assertInt32OK(t, []caseInt32[uint32]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertInt32Error(t, []caseInt32[uint32]{
			{name: "positive out of range", input: math.MaxInt32 + 1},
		})
	})

	t.Run("from uint64", func(t *testing.T) {
		assertInt32OK(t, []caseInt32[uint64]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertInt32Error(t, []caseInt32[uint64]{
			{name: "positive out of range", input: math.MaxInt32 + 1},
		})
	})

	t.Run("from float32", func(t *testing.T) {
		assertInt32OK(t, []caseInt32[float32]{
			{name: "zero", input: 0.0, want: 0},
			{name: "rounded value", input: 1.1, want: 1},
			{name: "positive within range", input: 10000, want: 10000},
		})

		assertInt32Error(t, []caseInt32[float32]{
			{name: "positive out of range", input: math.MaxInt32 + 1},
		})
	})

	t.Run("from float64", func(t *testing.T) {
		assertInt32OK(t, []caseInt32[float64]{
			{name: "zero", input: 0.0, want: 0},
			{name: "rounded value", input: 1.1, want: 1},
			{name: "positive within range", input: 10000, want: 10000},
		})

		assertInt32Error(t, []caseInt32[float64]{
			{name: "positive out of range", input: math.MaxInt32 + 1},
		})
	})
}

type caseUint32[in safecast.Type] struct {
	name  string
	input in
	want  uint32
}

func assertUint32OK[in safecast.Type](t *testing.T, tests []caseUint32[in]) {
	t.Helper()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := safecast.ToUint32(tt.input)
			assertNoError(t, err)
			assertEqual(t, tt.want, got)
		})
	}
}

func assertUint32Error[in safecast.Type](t *testing.T, tests []caseUint32[in]) {
	t.Helper()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Helper()

			got, err := safecast.ToUint32(tt.input)
			requireErrorIs(t, err, safecast.ErrConversionIssue)
			assertEqual(t, tt.want, got)
		})
	}
}

func TestToUint32(t *testing.T) {
	t.Run("from int", func(t *testing.T) {
		assertUint32OK(t, []caseUint32[int]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
		})

		assertUint32Error(t, []caseUint32[int]{
			{name: "positive out of range", input: math.MaxUint32 + 1},
			{name: "negative value", input: -1},
		})
	})

	t.Run("from int8", func(t *testing.T) {
		assertUint32OK(t, []caseUint32[int8]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertUint32Error(t, []caseUint32[int8]{
			{name: "negative value", input: -1},
		})
	})

	t.Run("from int16", func(t *testing.T) {
		assertUint32OK(t, []caseUint32[int16]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
		})

		assertUint32Error(t, []caseUint32[int16]{
			{name: "negative value", input: -1},
		})
	})

	t.Run("from int32", func(t *testing.T) {
		assertUint32OK(t, []caseUint32[int32]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
		})

		assertUint32Error(t, []caseUint32[int32]{
			{name: "negative value", input: -1},
		})
	})

	t.Run("from int64", func(t *testing.T) {
		assertUint32OK(t, []caseUint32[int64]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
		})

		assertUint32Error(t, []caseUint32[int64]{
			{name: "positive out of range", input: math.MaxUint32 + 1},
			{name: "negative value", input: -1},
		})
	})

	t.Run("from uint", func(t *testing.T) {
		assertUint32OK(t, []caseUint32[uint]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertUint32Error(t, []caseUint32[uint]{
			{name: "positive out of range", input: math.MaxUint32 + 1},
		})
	})

	t.Run("from uint8", func(t *testing.T) {
		assertUint32OK(t, []caseUint32[uint8]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})
	})

	t.Run("from uint16", func(t *testing.T) {
		assertUint32OK(t, []caseUint32[uint16]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})
	})

	t.Run("from uint32", func(t *testing.T) {
		assertUint32OK(t, []caseUint32[uint32]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})
	})

	t.Run("from uint64", func(t *testing.T) {
		assertUint32OK(t, []caseUint32[uint64]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertUint32Error(t, []caseUint32[uint64]{
			{name: "positive out of range", input: math.MaxUint32 + 1},
		})
	})

	t.Run("from float32", func(t *testing.T) {
		assertUint32OK(t, []caseUint32[float32]{
			{name: "zero", input: 0.0, want: 0},
			{name: "rounded value", input: 1.1, want: 1},
			{name: "positive within range", input: 10000.9, want: 10000},
		})

		assertUint32Error(t, []caseUint32[float32]{
			{name: "positive out of range", input: math.MaxUint32 + 1},
		})
	})

	t.Run("from float64", func(t *testing.T) {
		assertUint32OK(t, []caseUint32[float64]{
			{name: "zero", input: 0.0, want: 0},
			{name: "rounded value", input: 1.1, want: 1},
			{name: "positive within range", input: 10000.9, want: 10000},
		})

		assertUint32Error(t, []caseUint32[float64]{
			{name: "positive out of range", input: math.MaxUint32 + 1},
			{name: "negative value", input: -1},
		})
	})
}

type caseInt64[in safecast.Type] struct {
	name  string
	input in
	want  int64
}

func assertInt64OK[in safecast.Type](t *testing.T, tests []caseInt64[in]) {
	t.Helper()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := safecast.ToInt64(tt.input)
			assertNoError(t, err)
			assertEqual(t, tt.want, got)
		})
	}
}

func assertInt64Error[in safecast.Type](t *testing.T, tests []caseInt64[in]) {
	t.Helper()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := safecast.ToInt64(tt.input)
			requireErrorIs(t, err, safecast.ErrConversionIssue)
			assertEqual(t, tt.want, got)
		})
	}
}

func TestToInt64(t *testing.T) {
	t.Run("from int", func(t *testing.T) {
		assertInt64OK(t, []caseInt64[int]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
			{name: "negative within range", input: -10000, want: -10000},
		})
	})

	t.Run("from int8", func(t *testing.T) {
		assertInt64OK(t, []caseInt64[int8]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
			{name: "negative within range", input: -100, want: -100},
		})
	})

	t.Run("from int16", func(t *testing.T) {
		assertInt64OK(t, []caseInt64[int16]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
			{name: "negative within range", input: -10000, want: -10000},
		})
	})

	t.Run("from int32", func(t *testing.T) {
		assertInt64OK(t, []caseInt64[int32]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
			{name: "negative within range", input: -10000, want: -10000},
		})
	})

	t.Run("from int64", func(t *testing.T) {
		assertInt64OK(t, []caseInt64[int64]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
			{name: "negative within range", input: -10000, want: -10000},
		})
	})

	t.Run("from uint", func(t *testing.T) {
		assertInt64OK(t, []caseInt64[uint]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertInt64Error(t, []caseInt64[uint]{
			{name: "positive out of range", input: math.MaxInt64 + 1},
		})
	})

	t.Run("from uint8", func(t *testing.T) {
		assertInt64OK(t, []caseInt64[uint8]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})
	})

	t.Run("from uint16", func(t *testing.T) {
		assertInt64OK(t, []caseInt64[uint16]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})
	})

	t.Run("from uint32", func(t *testing.T) {
		assertInt64OK(t, []caseInt64[uint32]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})
	})

	t.Run("from uint64", func(t *testing.T) {
		assertInt64OK(t, []caseInt64[uint64]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertInt64Error(t, []caseInt64[uint64]{
			{name: "positive out of range", input: math.MaxInt64 + 1},
		})
	})

	t.Run("from float32", func(t *testing.T) {
		assertInt64OK(t, []caseInt64[float32]{
			{name: "zero", input: 0.0, want: 0},
			{name: "rounded value", input: 1.1, want: 1},
			{name: "positive within range", input: 10000.9, want: 10000},
			{name: "max int16", input: math.MaxInt16, want: math.MaxInt16},
			{name: "min int16", input: math.MinInt16, want: math.MinInt16},
			{name: "max int32", input: math.MaxInt32, want: 2147483648}, // number differs due to float imprecision
			{name: "min int32", input: math.MinInt32, want: math.MinInt32},
		})

		assertInt64Error(t, []caseInt64[float32]{
			{name: "out of range math.MaxInt64 + 1", input: math.MaxInt64 + 1},
			{name: "out of range math.MaxUint64", input: math.MaxUint64},
			{name: "out of range math.MinInt64", input: math.MinInt64 * 1.02}, // because of float imprecision, we have to exceed the min int64 to trigger the error
			{name: "out of range math.MaxFloat32", input: math.MaxFloat32},
			{name: "out of range -math.MaxFloat32", input: -math.MaxFloat32},
		})
	})

	t.Run("from float64", func(t *testing.T) {
		assertInt64OK(t, []caseInt64[float64]{
			{name: "zero", input: 0.0, want: 0},
			{name: "rounded value", input: 1.1, want: 1},
			{name: "positive within range", input: 10000.9, want: 10000},
			{name: "max int16", input: math.MaxInt16, want: math.MaxInt16},
			{name: "min int16", input: math.MinInt16, want: math.MinInt16},
			{name: "max int32", input: math.MaxInt32, want: math.MaxInt32},
			{name: "min int32", input: math.MinInt32, want: math.MinInt32},
		})

		assertInt64Error(t, []caseInt64[float64]{
			{name: "out of range math.MaxInt64 + 1", input: math.MaxInt64 + 1},
			{name: "out of range math.MaxUint64", input: math.MaxUint64},
			{name: "out of range math.MinInt64", input: math.MinInt64 * 1.02}, // because of float imprecision, we have to exceed the min int64 to trigger the error
			{name: "out of range math.MaxFloat32", input: math.MaxFloat32},
			{name: "out of range -math.MaxFloat32", input: -math.MaxFloat32},
			{name: "out of range math.MaxFloat64", input: math.MaxFloat64},
			{name: "out of range -math.MaxFloat64", input: -math.MaxFloat64},
		})
	})
}

type caseUint64[in safecast.Type] struct {
	name  string
	input in
	want  uint64
}

func assertUint64OK[in safecast.Type](t *testing.T, tests []caseUint64[in]) {
	t.Helper()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := safecast.ToUint64(tt.input)
			assertNoError(t, err)
			assertEqual(t, tt.want, got)
		})
	}
}

func assertUint64Error[in safecast.Type](t *testing.T, tests []caseUint64[in]) {
	t.Helper()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := safecast.ToUint64(tt.input)
			requireErrorIs(t, err, safecast.ErrConversionIssue)
			assertEqual(t, tt.want, got)
		})
	}
}

func TestToUint64(t *testing.T) {
	t.Run("from int", func(t *testing.T) {
		assertUint64OK(t, []caseUint64[int]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
		})

		assertUint64Error(t, []caseUint64[int]{
			{name: "negative value", input: -1},
		})
	})

	t.Run("from int8", func(t *testing.T) {
		assertUint64OK(t, []caseUint64[int8]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertUint64Error(t, []caseUint64[int]{
			{name: "negative value", input: -1},
		})
	})

	t.Run("from int16", func(t *testing.T) {
		assertUint64OK(t, []caseUint64[int16]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
		})

		assertUint64Error(t, []caseUint64[int]{
			{name: "negative value", input: -1},
		})
	})

	t.Run("from int32", func(t *testing.T) {
		assertUint64OK(t, []caseUint64[int32]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
		})

		assertUint64Error(t, []caseUint64[int]{
			{name: "negative value", input: -1},
		})
	})

	t.Run("from int64", func(t *testing.T) {
		assertUint64OK(t, []caseUint64[int64]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
		})

		assertUint64Error(t, []caseUint64[int]{
			{name: "negative value", input: -1},
		})
	})

	t.Run("from uint", func(t *testing.T) {
		assertUint64OK(t, []caseUint64[uint]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})
	})

	t.Run("from uint8", func(t *testing.T) {
		assertUint64OK(t, []caseUint64[uint8]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})
	})

	t.Run("from uint16", func(t *testing.T) {
		assertUint64OK(t, []caseUint64[uint16]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})
	})

	t.Run("from uint32", func(t *testing.T) {
		assertUint64OK(t, []caseUint64[uint32]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})
	})

	t.Run("from uint64", func(t *testing.T) {
		assertUint64OK(t, []caseUint64[uint64]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})
	})

	t.Run("from float32", func(t *testing.T) {
		assertUint64OK(t, []caseUint64[float32]{
			{name: "zero", input: 0.0, want: 0},
			{name: "rounded value", input: 1.1, want: 1},
			{name: "positive within range", input: 10000.9, want: 10000},
		})

		assertUint64Error(t, []caseUint64[float32]{
			{name: "negative value", input: -1},
			{name: "out of range max uint64", input: math.MaxUint64},
		})
	})

	t.Run("from float64", func(t *testing.T) {
		assertUint64OK(t, []caseUint64[float64]{
			{name: "zero", input: 0.0, want: 0},
			{name: "rounded value", input: 1.1, want: 1},
			{name: "positive within range", input: 10000.9, want: 10000},
		})

		assertUint64Error(t, []caseUint64[float64]{
			{name: "negative value", input: -1},
			{name: "out of range max uint64", input: math.MaxUint64},
			{name: "out of range max float32", input: math.MaxFloat32},
			{name: "out of range max float64", input: math.MaxFloat64},
		})
	})
}

type caseInt[in safecast.Type] struct {
	name  string
	input in
	want  int
}

func assertIntOK[in safecast.Type](t *testing.T, tests []caseInt[in]) {
	t.Helper()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := safecast.ToInt(tt.input)
			assertNoError(t, err)
			assertEqual(t, tt.want, got)
		})
	}
}

func assertIntError[in safecast.Type](t *testing.T, tests []caseInt[in]) {
	t.Helper()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := safecast.ToInt(tt.input)
			requireErrorIs(t, err, safecast.ErrConversionIssue)
			assertEqual(t, tt.want, got)
		})
	}
}

func TestToInt(t *testing.T) {
	t.Run("from int", func(t *testing.T) {
		assertIntOK(t, []caseInt[int]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
			{name: "negative within range", input: -10000, want: -10000},
		})
	})

	t.Run("from int8", func(t *testing.T) {
		assertIntOK(t, []caseInt[int8]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
			{name: "negative within range", input: -100, want: -100},
		})
	})

	t.Run("from int16", func(t *testing.T) {
		assertIntOK(t, []caseInt[int16]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
			{name: "negative within range", input: -10000, want: -10000},
		})
	})

	t.Run("from int32", func(t *testing.T) {
		assertIntOK(t, []caseInt[int32]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
			{name: "negative within range", input: -10000, want: -10000},
		})
	})

	t.Run("from int64", func(t *testing.T) {
		assertIntOK(t, []caseInt[int64]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
			{name: "negative within range", input: -10000, want: -10000},
		})
	})

	t.Run("from uint", func(t *testing.T) {
		assertIntOK(t, []caseInt[uint]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertIntError(t, []caseInt[uint]{
			{name: "positive out of range", input: math.MaxInt64 + 1},
		})
	})

	t.Run("from uint8", func(t *testing.T) {
		assertIntOK(t, []caseInt[uint8]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})
	})

	t.Run("from uint16", func(t *testing.T) {
		assertIntOK(t, []caseInt[uint16]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})
	})

	t.Run("from uint32", func(t *testing.T) {
		assertIntOK(t, []caseInt[uint32]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})
	})

	t.Run("from uint64", func(t *testing.T) {
		assertIntOK(t, []caseInt[uint64]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertIntError(t, []caseInt[uint64]{
			{name: "positive out of range", input: math.MaxInt64 + 1},
		})
	})

	t.Run("from float32", func(t *testing.T) {
		assertIntOK(t, []caseInt[float32]{
			{name: "zero", input: 0.0, want: 0},
			{name: "rounded value", input: 1.1, want: 1},
			{name: "positive within range", input: 10000.9, want: 10000},
		})

		assertIntError(t, []caseInt[float32]{
			{name: "out of range math.MaxInt64 + 1", input: math.MaxInt64 + 1},
			{name: "out of range math.MaxUint64", input: math.MaxUint64},
			{name: "out of range math.MinInt64", input: math.MinInt64 * 1.02}, // because of float imprecision, we have to exceed the min int64 to trigger the error
			{name: "out of range math.MaxFloat32", input: math.MaxFloat32},
			{name: "out of range -math.MaxFloat32", input: -math.MaxFloat32},
		})
	})

	t.Run("from float64", func(t *testing.T) {
		assertIntOK(t, []caseInt[float64]{
			{name: "zero", input: 0.0, want: 0},
			{name: "rounded value", input: 1.1, want: 1},
			{name: "positive within range", input: 10000.9, want: 10000},
			{name: "math.MinInt64", input: math.MinInt64, want: math.MinInt64}, // pass because of float imprecision
		})

		assertIntError(t, []caseInt[float64]{
			{name: "out of range math.MaxInt64 + 1", input: math.MaxInt64 + 1},
			{name: "out of range math.MaxUint64", input: math.MaxUint64},
			{name: "out of range math.MinInt64", input: math.MinInt64 * 1.02}, // because of float imprecision, we have to exceed the min int64 to trigger the error
			{name: "out of range math.MaxFloat32", input: math.MaxFloat32},
			{name: "out of range -math.MaxFloat32", input: -math.MaxFloat32},
			{name: "out of range math.MaxFloat64", input: math.MaxFloat64},
			{name: "out of range -math.MaxFloat64", input: -math.MaxFloat64},
		})
	})
}

type caseUint[in safecast.Type] struct {
	name  string
	input in
	want  uint
}

func assertUintOK[in safecast.Type](t *testing.T, tests []caseUint[in]) {
	t.Helper()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := safecast.ToUint(tt.input)
			assertNoError(t, err)
			assertEqual(t, tt.want, got)
		})
	}
}

func assertUintError[in safecast.Type](t *testing.T, tests []caseUint[in]) {
	t.Helper()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := safecast.ToUint(tt.input)
			requireErrorIs(t, err, safecast.ErrConversionIssue)
			assertEqual(t, tt.want, got)
		})
	}
}

func TestToUint(t *testing.T) {
	t.Run("from int", func(t *testing.T) {
		assertUintOK(t, []caseUint[int]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
		})

		assertUintError(t, []caseUint[int]{
			{name: "negative value", input: -1},
		})
	})

	t.Run("from int8", func(t *testing.T) {
		assertUintOK(t, []caseUint[int8]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})

		assertUintError(t, []caseUint[int8]{
			{name: "negative value", input: -1},
		})
	})

	t.Run("from int16", func(t *testing.T) {
		assertUintOK(t, []caseUint[int16]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
		})

		assertUintError(t, []caseUint[int16]{
			{name: "negative value", input: -1},
		})
	})

	t.Run("from int32", func(t *testing.T) {
		assertUintOK(t, []caseUint[int32]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
		})

		assertUintError(t, []caseUint[int32]{
			{name: "negative value", input: -1},
		})
	})

	t.Run("from int64", func(t *testing.T) {
		assertUintOK(t, []caseUint[int64]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 10000, want: 10000},
		})

		assertUintError(t, []caseUint[int64]{
			{name: "negative value", input: -1},
		})
	})

	t.Run("from uint", func(t *testing.T) {
		assertUintOK(t, []caseUint[uint]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})
	})

	t.Run("from uint8", func(t *testing.T) {
		assertUintOK(t, []caseUint[uint8]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})
	})

	t.Run("from uint16", func(t *testing.T) {
		assertUintOK(t, []caseUint[uint16]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})
	})

	t.Run("from uint32", func(t *testing.T) {
		assertUintOK(t, []caseUint[uint32]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})
	})

	t.Run("from uint64", func(t *testing.T) {
		assertUintOK(t, []caseUint[uint64]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})
	})

	t.Run("from float32", func(t *testing.T) {
		assertUintOK(t, []caseUint[float32]{
			{name: "zero", input: 0.0, want: 0},
			{name: "rounded value", input: 1.1, want: 1},
			{name: "positive within range", input: 10000.9, want: 10000},
		})

		assertUint64Error(t, []caseUint64[float32]{
			{name: "negative value", input: -1},
			{name: "out of range max uint64", input: math.MaxUint64},
			{name: "out of range max float32", input: math.MaxFloat32},
		})
	})

	t.Run("from float64", func(t *testing.T) {
		assertUintOK(t, []caseUint[float64]{
			{name: "zero", input: 0.0, want: 0},
			{name: "rounded value", input: 1.1, want: 1},
			{name: "positive within range", input: 10000.9, want: 10000},
		})

		assertUintError(t, []caseUint[float64]{
			{name: "negative value", input: -1},
			{name: "out of range max uint64", input: math.MaxUint64},
			{name: "out of range max float32", input: math.MaxFloat32},
			{name: "out of range max float64", input: math.MaxFloat64},
		})
	})

	type UintAlias uint
	t.Run("from type alias", func(t *testing.T) {
		assertUintOK(t, []caseUint[UintAlias]{
			{name: "zero", input: 0, want: 0},
			{name: "positive within range", input: 100, want: 100},
		})
	})
}
