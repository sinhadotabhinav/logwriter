package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

// Logger func initialises logrus logger with specified settings
func Logger() *logrus.Logger {
	var log = logrus.New()
	log.Out = os.Stdout
	log.Formatter = &logrus.JSONFormatter{}
	log.Level = logrus.DebugLevel
	return log
}
