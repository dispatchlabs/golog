package modifiers

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/dispatchlabs/golog/contracts"
)

type SimpleFormatterLogger struct {
	loggerToSendTo contracts.Logger
}

func NewSimpleFormatterLogger(logger contracts.Logger) contracts.Logger {
	return &SimpleFormatterLogger{
		loggerToSendTo: logger,
	}
}

func (thisRef *SimpleFormatterLogger) Log(logEntry contracts.LogEntry) {
	var formattedTime = logEntry.Time.UTC().Format(time.RFC3339Nano)
	var formattedTimeLen = len(formattedTime)
	if formattedTimeLen < 30 {
		var spacesCount = 30 - formattedTimeLen

		var newV = fmt.Sprintf("%"+strconv.Itoa(spacesCount+1)+"v", "Z")
		newV = strings.Replace(newV, " ", "0", spacesCount)

		formattedTime = strings.Replace(
			formattedTime,
			"Z",
			newV,
			1,
		)
	}

	logEntry.Message = fmt.Sprintf(
		"%s %s %s %"+strconv.Itoa(logEntry.Level*4)+"v %s",
		formattedTime,
		logEntry.ErrorType,
		logEntry.Id,
		"",
		logEntry.Message,
	)

	thisRef.loggerToSendTo.Log(logEntry)
}

func (thisRef *SimpleFormatterLogger) LogInfo(id string, level int, message string)    {}
func (thisRef *SimpleFormatterLogger) LogWarning(id string, level int, message string) {}
func (thisRef *SimpleFormatterLogger) LogError(id string, level int, message string)   {}
