package persisters

import (
	"os"

	"github.com/dispatchlabs/golog/contracts"
)

type FileLogger struct {
	file *os.File
}

func NewFileLogger(fileName string) contracts.Logger {
	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	return &FileLogger{
		file: f,
	}
}

func (thisRef *FileLogger) Log(logEntry contracts.LogEntry) {
	thisRef.file.WriteString(logEntry.Message + "\n")
}

func (thisRef *FileLogger) LogInfo(id string, level int, message string)    {}
func (thisRef *FileLogger) LogWarning(id string, level int, message string) {}
func (thisRef *FileLogger) LogError(id string, level int, message string)   {}
