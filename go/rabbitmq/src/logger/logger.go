package logger

import "fmt"

type logger interface {
	Debugf(format string, v ...interface{})
}

// Logger l
type Logger struct{}

// Debugf 调试
func (l Logger) Debugf(format string, v ...interface{}) {
	fmt.Printf(format, v)
}
