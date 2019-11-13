package logger

import (
	"fmt"
	"io"
	"path"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func init() {
	logger = logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	logger.SetReportCaller(true)
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:          true,
		DisableLevelTruncation: true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			pcs := make([]uintptr, 25)
			depth := runtime.Callers(1, pcs)
			frames := runtime.CallersFrames(pcs[:depth])

			for f, again := frames.Next(); again; f, again = frames.Next() {
				skip := strings.Contains(f.Function, "logrus") || strings.Contains(f.Function, "logger")
				if !skip {
					filename := path.Base(f.File)
					return fmt.Sprintf("%s()", path.Base(f.Function)), fmt.Sprintf("%s:%d", filename, f.Line)
				}
			}

			return "", ""
		},
	})
}

func Debugf(msg string, args ...interface{}) {
	logger.Debugf(msg, args...)
}

func Errorf(msg string, args ...interface{}) {
	logger.Errorf(msg, args...)
}

func Fatalf(msg string, args ...interface{}) {
	logger.Fatalf(msg, args...)
}

func Infof(msg string, args ...interface{}) {
	logger.Infof(msg, args...)
}

func Panicf(msg string, args ...interface{}) {
	logger.Panicf(msg, args...)
}

func SetOutput(w io.Writer) {
	logger.SetOutput(w)
}

func Warnf(msg string, args ...interface{}) {
	logger.Warnf(msg, args...)
}
