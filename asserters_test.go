package safecast_test

import (
	"errors"
	"math"
	"strings"
	"testing"

	"github.com/ccoveille/go-safecast/v2"
)

// mockTestingT is a mock implementation of the safecast.TestingT interface
// that captures the arguments passed to the Fatal method for testing purposes.
type mockTestingT struct {
	failed bool
}

func (m *mockTestingT) Helper() {}

func (m *mockTestingT) Fatal(_ ...any) {
	m.failed = true
}

func (m mockTestingT) Failed() bool {
	return m.failed
}

func Map[T any, U any](fn func(v T) U, input []T) []U {
	var output []U
	for _, v := range input {
		output = append(output, fn(v))
	}
	return output
}

type TestRunner interface {
	Run(t *testing.T)
}

// interfaces validation
// this leads to a compile-time error if there is a mismatch
var _ safecast.TestingT = new(testing.T)
var _ safecast.TestingT = new(testing.B)
var _ safecast.TestingT = new(testing.F)
var _ safecast.TestingT = new(mockTestingT)

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
		t.Fatalf("error message should contain %#q: %#q", text, errMessage)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestErrorMessage(t *testing.T) {
	_, err := safecast.Convert[uint8](-1)
	requireErrorIs(t, err, safecast.ErrConversionIssue)
	requireErrorIs(t, err, safecast.ErrExceedMinimumValue)
	requireErrorContains(t, err, "than 0 (uint8)")

	_, err = safecast.Convert[uint8](math.MaxInt16)
	requireErrorIs(t, err, safecast.ErrConversionIssue)
	requireErrorIs(t, err, safecast.ErrExceedMaximumValue)
	requireErrorContains(t, err, "than 255 (uint8)")

	_, err = safecast.Convert[int8](-math.MaxInt16)
	requireErrorIs(t, err, safecast.ErrConversionIssue)
	requireErrorIs(t, err, safecast.ErrExceedMinimumValue)
	requireErrorContains(t, err, "than -128 (int8)")

	_, err = safecast.Convert[int8](3.14, safecast.WithDecimalLossReport())
	requireErrorIs(t, err, safecast.ErrConversionIssue)
	requireErrorIs(t, err, safecast.ErrDecimalLoss)
	requireErrorContains(t, err, "decimal loss")
}
