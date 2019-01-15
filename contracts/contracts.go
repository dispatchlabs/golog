package contracts

import "time"

type ErrorType string

const (
	TypeError   ErrorType = "EROR"
	TypeWarning           = "WARN"
	TypeInfo              = "INFO"
)

type LogEntry struct {
	Time      time.Time
	ErrorType ErrorType
	Id        string
	Level     int
	Message   string
}

type Logger interface {
	Log(logEntry LogEntry)

	LogInfo(id string, level int, message string)
	LogWarning(id string, level int, message string)
	LogError(id string, level int, message string)
}
