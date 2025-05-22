package blueprint

import (
	"errors"
	"fmt"
	"strings"
)

type Logger interface {
	Printf(format string, args ...any)
	Println(args ...any)
}

type Exporter interface {
	Export(bu BuildOptions) error
}

type logs struct {
	msgs []string
}

var _ Logger = (*logs)(nil)

func newCollector() *logs {
	return &logs{
		msgs: make([]string, 0),
	}
}

func (c *logs) Printf(format string, args ...any) {
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

func (c *logs) Println(args ...any) {
	c.msgs = append(c.msgs, strings.Join(toStringSlice(args), " "))
}

func (c *logs) Errors() error {
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
