package notes

import (
	"errors"
	"fmt"
	"strings"
)

var ErrConversionWarning = errors.New("warning")
var ErrConversionFatal = errors.New("fatal")

type ConversionNotes struct {
	errors []error
}

func (e *ConversionNotes) Warn(err ...string) {
	msg := strings.Join(err, " ")
	e.errors = append(e.errors, fmt.Errorf("%w: %s", ErrConversionWarning, msg))
}

func (e *ConversionNotes) Fatal(err ...string) {
	msg := strings.Join(err, " ")
	e.errors = append(e.errors, fmt.Errorf("%w: %s", ErrConversionFatal, msg))
}

func (e *ConversionNotes) Join() error {
	return errors.Join(e.errors...)
}
