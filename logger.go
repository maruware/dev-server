package httpdump

import (
	"log"
	"os"
)

type Logger struct {
	out *log.Logger
	err *log.Logger
}

var logger *Logger

func init() {
	logger = &Logger{
		out: log.New(os.Stdout, "", log.Ldate|log.Ltime),
		err: log.New(os.Stderr, "[Error]", log.Ldate|log.Ltime),
	}
}

func (l *Logger) Printf(format string, v ...interface{}) {
	l.out.Printf(format, v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.err.Printf(format, v...)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.err.Fatalf(format, v...)
}
