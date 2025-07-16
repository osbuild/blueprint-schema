package conv

import (
	"errors"
	"fmt"
)

// warnings is a simple logger for warning messages as well as it is a Go error.
type warnings struct {
	msgs []error
}

type Logger interface {
	Printf(format string, args ...any)
	Println(args ...any)
}

var _ Logger = (*warnings)(nil)

// Printf formats its arguments according to the format and appends the result to the error collector.
func (c *warnings) Printf(format string, args ...any) {
	c.msgs = append(c.msgs, fmt.Errorf(format, args...))
}

// Println formats its arguments using fmt.Sprint and appends the result to the error collector.
// There is a slight difference from Go log Println which joins the arguments with a space, this
// function joins multiple args as separate errors.
func (c *warnings) Println(args ...any) {
	for _, arg := range args {
		c.msgs = append(c.msgs, fmt.Errorf("%v", arg))
	}
}

func (c *warnings) Error() error {
	if len(c.msgs) == 0 {
		return nil
	}

	return errors.Join(c.msgs...)
}

func (c *warnings) Unwrap() []error {
	return c.msgs
}
