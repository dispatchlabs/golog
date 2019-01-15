package maing

import (
	"github.com/dispatchlabs/golog/contracts"
	"github.com/dispatchlabs/golog/modifiers"
)

func NewLogger(logger contracts.Logger) contracts.Logger {
	return modifiers.NewDefaultLogger(logger)
}

var instance contracts.Logger

func StoreSingleton(logger contracts.Logger) {
	instance = logger
}

func Instance() contracts.Logger {
	return instance
}
