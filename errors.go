package safecast

import (
	"errors"
	"fmt"
)

var (
	ErrConversionIssue    = errors.New("conversion issue")
	ErrExceedMaximumValue = errors.New("maximum value for this type exceeded")
	ErrExceedMinimumValue = errors.New("minimum value for this type exceeded")
)

type Error struct {
	value    any
	boundary any
	err      error
}

func (e Error) Error() string {
	errMessage := ErrConversionIssue.Error()

	switch {
	case errors.Is(e.err, ErrExceedMaximumValue):
		errMessage = fmt.Sprintf("%s: %v (%T) is greater than %v (%T)", errMessage, e.value, e.value, e.boundary, e.boundary)
	case errors.Is(e.err, ErrExceedMinimumValue):
		errMessage = fmt.Sprintf("%s: %v (%T) is less than %v (%T)", errMessage, e.value, e.value, e.boundary, e.boundary)
	}

	if e.err != nil {
		errMessage = fmt.Sprintf("%s: %s", errMessage, e.err.Error())
	}
	return errMessage
}

func (e Error) Unwrap() []error {
	errs := []error{ErrConversionIssue}
	if e.err != nil {
		errs = append(errs, e.err)
	}
	return errs
}
