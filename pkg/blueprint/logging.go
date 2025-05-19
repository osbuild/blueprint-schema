package blueprint

type Logger interface {
	Printf(format string, args ...any)
	Println(args ...any)
}

type discardLogger struct{}

func (d *discardLogger) Printf(format string, args ...any) {}
func (d *discardLogger) Println(args ...any)               {}

var _ Logger = (*discardLogger)(nil)

var log Logger = &discardLogger{}

// SetLogger sets the logger to be used by the package.
// If not set, a discard logger is used which ignores all log messages.
// The function is not thread-safe and should be called before any other functions in the package.
func SetLogger(l Logger) {
	log = l
}
