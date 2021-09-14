package log

import "github.com/sirupsen/logrus"

type Logger interface {
	Infof(string, ...interface{})
	Debugf(string, ...interface{})
	Errorf(string, ...interface{})
	Fatalf(string, ...interface{})
}

func NewLogger(debug bool) Logger {
	l := logrus.New()
	if debug {
		l.SetLevel(logrus.DebugLevel)
	}

	timeFormatter := new(logrus.TextFormatter)
	timeFormatter.ForceColors = true
	timeFormatter.FullTimestamp = true
	timeFormatter.TimestampFormat = "2006-01-02 15:04:05"
	l.SetFormatter(timeFormatter)

	return l
}
