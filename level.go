package opinlog

import "github.com/sirupsen/logrus"

// LogLevel is the type for log level
// this is used so that users do not have to expose logrus
type LogLevel int

const (
	Trace LogLevel = 1
	Debug LogLevel = 2
	Info  LogLevel = 3
	Warn  LogLevel = 4
	Error LogLevel = 5
)

// SetLevel sets the log level
func SetLevel(level LogLevel) {
	switch level {
	case Trace:
		logrus.SetLevel(logrus.TraceLevel)
	case Debug:
		logrus.SetLevel(logrus.DebugLevel)
	case Info:
		logrus.SetLevel(logrus.InfoLevel)
	case Warn:
		logrus.SetLevel(logrus.WarnLevel)
	case Error:
		logrus.SetLevel(logrus.ErrorLevel)
	}
}
