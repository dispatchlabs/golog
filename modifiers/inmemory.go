package modifiers

import (
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/dispatchlabs/golog/contracts"
)

type InmemoryLogger struct {
	logLinesByKey         map[string][]contracts.LogEntry
	logLinesByKeyWithTime map[string]time.Time
	loggerToSendTo        contracts.Logger
	rwMutex               sync.RWMutex
}

func NewInmemoryLogger(logger contracts.Logger) contracts.Logger {
	return &InmemoryLogger{
		logLinesByKey:         map[string][]contracts.LogEntry{},
		logLinesByKeyWithTime: map[string]time.Time{},
		loggerToSendTo:        logger,
		rwMutex:               sync.RWMutex{},
	}
}

func (thisRef *InmemoryLogger) Log(logEntry contracts.LogEntry) {
	thisRef.rwMutex.Lock()
	defer thisRef.rwMutex.Unlock()

	thisRef.logLinesByKey[logEntry.Id] = append(thisRef.logLinesByKey[logEntry.Id], logEntry)

	// Remember the earliest log-line
	var ok bool
	if _, ok = thisRef.logLinesByKeyWithTime[logEntry.Id]; !ok {
		thisRef.logLinesByKeyWithTime[logEntry.Id] = logEntry.Time
	} else {
		var storedTime = thisRef.logLinesByKeyWithTime[logEntry.Id]
		if storedTime.After(logEntry.Time) {
			thisRef.logLinesByKeyWithTime[logEntry.Id] = logEntry.Time
		}
	}
}

func (thisRef *InmemoryLogger) Flush() {
	thisRef.rwMutex.RLock()
	defer thisRef.rwMutex.RUnlock()

	// Sort by time
	var allTimes = []time.Time{}
	var timeToLogEntryId = map[int64]string{}
	for key, value := range thisRef.logLinesByKeyWithTime {
		allTimes = append(allTimes, value)
		timeToLogEntryId[value.UnixNano()] = key
	}
	sort.Slice(
		allTimes,
		func(i, j int) bool {
			return allTimes[i].Before(allTimes[j])
		},
	)

	for index, _ := range allTimes {
		var logEntryId = timeToLogEntryId[allTimes[index].UnixNano()]
		var arrayOfLogEntries = thisRef.logLinesByKey[logEntryId]

		sort.Slice(
			arrayOfLogEntries,
			func(i, j int) bool {
				return arrayOfLogEntries[i].Time.Before(arrayOfLogEntries[j].Time)
			},
		)

		for i, _ := range arrayOfLogEntries {
			// if i == 0 {
			arrayOfLogEntries[i].Id = fmt.Sprintf("%10s", arrayOfLogEntries[i].Id)
			// } else {
			// arrayOfLogEntries[i].Id = fmt.Sprintf("%10s", "")
			// }

			thisRef.loggerToSendTo.Log(arrayOfLogEntries[i])
		}
	}
}

func (thisRef *InmemoryLogger) LogInfo(id string, level int, message string)    {}
func (thisRef *InmemoryLogger) LogWarning(id string, level int, message string) {}
func (thisRef *InmemoryLogger) LogError(id string, level int, message string)   {}
