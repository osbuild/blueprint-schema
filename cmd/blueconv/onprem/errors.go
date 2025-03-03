package onprem

import (
	"errors"
	"fmt"
	"strings"
)

var ErrConversionWarning = errors.New("warning")
var ErrConversionFatal = errors.New("fatal")

type Errors struct {
	errors []error
}

func (e *Errors) Warn(err ...string) {
	msg := strings.Join(err, " ")
	e.errors = append(e.errors, fmt.Errorf("%w: %s", ErrConversionWarning, msg))
}

func (e *Errors) Fatal(err ...string) {
	msg := strings.Join(err, " ")
	e.errors = append(e.errors, fmt.Errorf("%w: %s", ErrConversionFatal, msg))
}

func (e *Errors) Join() error {
	return errors.Join(e.errors...)
}
