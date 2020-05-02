package opinlog

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

// TextFormatter embeds the logrus struct so that users do not have to import logrus
type TextFormatter struct {
	logrus.TextFormatter
}

// JSONFormatter embeds the logrus struct so that users do not have to import logrus
type JSONFormatter struct {
	logrus.JSONFormatter
}

// LogFormat is the type for log format
// this is used so that users do not have to expose logrus
type LogFormat int

const (
	FormatAsNone LogFormat = 0
	FormatAsText LogFormat = 1
	FormatAsJSON LogFormat = 2
)

// SetLogFormat sets default log formats
func SetLogFormat(format LogFormat) {
	switch format {
	case FormatAsNone:
		logrus.SetOutput(ioutil.Discard)
	case FormatAsText:
		logrus.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
	case FormatAsJSON:
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}
}

// SetLogFormatManual allows users to customize the options manually
func SetLogFormatManual(format logrus.Formatter) {
	logrus.SetFormatter(format)
}
