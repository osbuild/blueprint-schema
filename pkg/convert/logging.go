package convert

type Logger interface {
	Printf(format string, args ...any)
	Println(args ...any)
}

var log Logger

func SetLogger(l Logger) {
	log = l
}
