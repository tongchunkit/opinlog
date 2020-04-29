package opinlog

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

// LogFormat is the type for log format
// this is used so that users do not have to expose logrus
type LogFormat int

const (
	FormatAsNone LogFormat = 0
	FormatAsText LogFormat = 1
	FormatAsJSON LogFormat = 2
)

// SetLogFormat sets the log format
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
