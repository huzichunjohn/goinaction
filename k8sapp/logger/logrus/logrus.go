package logrus

import (
	"github.com/Sirupsen/logrus"

	"goinaction/k8sapp/logger"
)

func New(config *logger.Config) logger.Logger {
	logger := logrus.New()
	logger.Level = logrusLevelConverter(config.Level)
	logger.WithFields(logrus.Fields(config.Fields))
	return logger
}

func logrusLevelConverter(level logger.Level) logrus.Level {
	switch level {
	case logger.LevelDebug:
		return logrus.DebugLevel
	case logger.LevelInfo:
		return logrus.InfoLevel
	case logger.LevelWarn:
		return logrus.WarnLevel
	case logger.LevelError:
		return logrus.ErrorLevel
	case logger.LevelFatal:
		return logrus.FatalLevel
	default:
		return logrus.InfoLevel
	}
}
