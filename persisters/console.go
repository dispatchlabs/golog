package persisters

import (
	"fmt"

	"github.com/nic0lae/golog/contracts"
)

type ConsoleLogger struct {
}

func NewConsoleLogger() contracts.Logger {
	return &ConsoleLogger{}
}

func (thisRef *ConsoleLogger) Log(logEntry contracts.LogEntry) {
	fmt.Println(logEntry.Message)
}

func (thisRef *ConsoleLogger) LogInfo(id string, level int, message string)    {}
func (thisRef *ConsoleLogger) LogWarning(id string, level int, message string) {}
func (thisRef *ConsoleLogger) LogError(id string, level int, message string)   {}
