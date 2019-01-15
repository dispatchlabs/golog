package modifiers

import (
	"time"

	"github.com/dispatchlabs/golog/contracts"
)

type DefaultLogger struct {
	loggerToSendTo contracts.Logger
}

func NewDefaultLogger(logger contracts.Logger) contracts.Logger {
	return &DefaultLogger{
		loggerToSendTo: logger,
	}
}

func (thisRef *DefaultLogger) Log(logEntry contracts.LogEntry) {
	thisRef.loggerToSendTo.Log(logEntry)
}

func (thisRef *DefaultLogger) LogInfo(id string, level int, message string) {
	thisRef.Log(contracts.LogEntry{
		Time:      time.Now(),
		ErrorType: contracts.TypeInfo,
		Id:        id,
		Level:     level,
		Message:   message,
	})
}

func (thisRef *DefaultLogger) LogWarning(id string, level int, message string) {
	thisRef.Log(contracts.LogEntry{
		Time:      time.Now(),
		ErrorType: contracts.TypeWarning,
		Id:        id,
		Level:     level,
		Message:   message,
	})
}

func (thisRef *DefaultLogger) LogError(id string, level int, message string) {
	thisRef.Log(contracts.LogEntry{
		Time:      time.Now(),
		ErrorType: contracts.TypeError,
		Id:        id,
		Level:     level,
		Message:   message,
	})
}
