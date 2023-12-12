//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
)

// InitApp returns a real app.
func InitApp() *App {
	wire.Build(AppSet)
	return nil
}

// InitMockedAppFromArgs returns an app with mocked dependencies provided via
// arguments (Approach A). Note that the argument's type is the interface
// type (timer), but the concrete mock type should be passed.
func InitMockedAppFromArgs(mt Timer) *App {
	wire.Build(AppSetWithoutMocks)
	return nil
}

// InitMockedApp returns an app with its mocked dependencies, created
// via providers (Approach B).
func InitMockedApp() *AppWithMocks {
	wire.Build(MockAppSet)
	return nil
}
