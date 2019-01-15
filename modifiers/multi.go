package modifiers

import (
	"github.com/dispatchlabs/golog/contracts"
)

type MultiLogger struct {
	loggers []contracts.Logger
}

func NewMultiLogger(loggers []contracts.Logger) contracts.Logger {
	return &MultiLogger{
		loggers: loggers,
	}
}

func (thisRef *MultiLogger) Log(logEntry contracts.LogEntry) {
	for _, logger := range thisRef.loggers {
		logger.Log(logEntry)
	}
}

func (thisRef *MultiLogger) LogInfo(id string, level int, message string)    {}
func (thisRef *MultiLogger) LogWarning(id string, level int, message string) {}
func (thisRef *MultiLogger) LogError(id string, level int, message string)   {}
