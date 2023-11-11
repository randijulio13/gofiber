package logger

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var customLogger *logrus.Logger

func InitializeLogger() {
	customLogger = logrus.New()
	customLogger.SetFormatter(&logrus.JSONFormatter{})
	level, err := logrus.ParseLevel(viper.GetString("LOG_LEVEL"))
	if err != nil {
		level = logrus.InfoLevel
	}

	customLogger.SetLevel(level)
}

func GetLogger() *logrus.Logger {
	return customLogger
}
