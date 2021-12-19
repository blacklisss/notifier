package services

import (
	log "github.com/sirupsen/logrus"
)

type Logger struct {
	logger *log.Logger
}

func NewLogger(logger *log.Logger) *Logger {
	return &Logger{logger: logger}
}
