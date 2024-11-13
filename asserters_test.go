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
