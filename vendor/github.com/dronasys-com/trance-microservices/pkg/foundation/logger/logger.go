package logger

// binu@dronasys.com

import (
	logrus "github.com/sirupsen/logrus"
)

var Log *logrus.Logger

type Logger struct {
	*logrus.Logger
}

func NewLogger() {

	if Log == nil {

		customFormatter := new(logrus.TextFormatter)
		customFormatter.TimestampFormat = "2006-01-02 15:04:05"
		// Log.Logrus.SetFormatter(customFormatter)
		logrus.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})

		Log = &logrus.Logger{}
	}

}

// func (l *Logger) Info(args ...interface{}) {
// 	if config.Config.LogVerbosity > 3 {
// 		l.Info(args)
// 	}
// }

// func (l *Logger) Debug(args ...interface{}) {
// 	if config.Config.LogVerbosity > 7 {
// 		l.Debug(args)
// 	}
// }

// func (l *Logger) Error(args ...interface{}) {
// 	if config.Config.LogVerbosity > 0 {
// 		l.Error(args)
// 	}
// }

// func (l *Logger) Fatal(args ...interface{}) {
// 	if config.Config.LogVerbosity > 0 {
// 		l.Fatal(args)
// 	}
// }
