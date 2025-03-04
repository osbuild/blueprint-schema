package notes

import (
	"errors"
	"fmt"
)

var ErrConversionWarning = errors.New("warning")
var ErrConversionFatal = errors.New("fatal")

// ConversionNotes is a collection of conversion notes. It can be used to collect
// warnings and errors during the conversion process.
type ConversionNotes struct {
	errors []error
}

// Warn adds a warning to the conversion notes. Warnings are considered recoverable
// and should not stop the conversion process.
func (e *ConversionNotes) Warn(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	e.errors = append(e.errors, fmt.Errorf("%w: %s", ErrConversionWarning, msg))
}

// Fatal adds a fatal error to the conversion notes. Fatal errors are considered
// unrecoverable and should stop the conversion process.
func (e *ConversionNotes) Fatal(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	e.errors = append(e.errors, fmt.Errorf("%w: %s", ErrConversionFatal, msg))
}

// Join returns all the errors as a single error. If there are no errors, it returns nil.
// Errors can be unwrapped using errors.Unwrap and checked using errors.Is individually.
func (e *ConversionNotes) Join() error {
	return errors.Join(e.errors...)
}

// Errors returns all the errors as a slice of errors. Never returns nil.
func (e *ConversionNotes) Errors() []error {
	if e.errors == nil {
		return []error{}
	}
	return e.errors
}
