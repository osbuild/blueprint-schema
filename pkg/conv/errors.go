package conv

import (
	"errors"
	"fmt"
	"strings"
)

type errs struct {
	msgs []string
}

type Logger interface {
	Printf(format string, args ...any)
	Println(args ...any)
}

var _ Logger = (*errs)(nil)

func newErrorCollector() *errs {
	return &errs{
		msgs: make([]string, 0),
	}
}

// Printf formats its arguments according to the format and appends the result to the error collector.
func (c *errs) Printf(format string, args ...any) {
	c.msgs = append(c.msgs, fmt.Sprintf(format, args...))
}

func toStringSlice(args []any) []string {
	strs := make([]string, len(args))
	for i, arg := range args {
		switch v := arg.(type) {
		case string:
			strs[i] = v
		case fmt.Stringer:
			strs[i] = v.String()
		default:
			strs[i] = fmt.Sprintf("%v", arg)
		}
	}
	return strs
}

// Println formats its arguments using fmt.Sprint and appends the result to the error collector.
func (c *errs) Println(args ...any) {
	c.msgs = append(c.msgs, strings.Join(toStringSlice(args), " "))
}

// Error returns a single error that contains all collected messages. Individual errors can be unwrapped.
func (c *errs) Errors() error {
	if len(c.msgs) == 0 {
		return nil
	}

	errs := make([]error, len(c.msgs))
	for i, msg := range c.msgs {
		errs[i] = errors.New(msg)
	}

	return &logErrors{
		e: errors.Join(errs...),
	}
}

type logErrors struct {
	e error
}

var _ error = (*logErrors)(nil)

func (c *logErrors) Error() string {
	return c.e.Error()
}

func (c *logErrors) Unwrap() []error {
	if unwrapped, ok := c.e.(interface{ Unwrap() []error }); ok {
		return unwrapped.Unwrap()
	}
	return []error{c.e}
}
