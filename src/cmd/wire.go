//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/leobastiani/golang-google-wire-best-practices/src"
)

// InitApp returns a real app.
func InitApp() *src.App {
	wire.Build(src.AppSet)
	return nil
}

// InitMockedAppFromArgs returns an app with mocked dependencies provided via
// arguments (Approach A). Note that the argument's type is the interface
// type (timer), but the concrete mock type should be passed.
func InitMockedAppFromArgs(mt src.Timer) *src.App {
	wire.Build(src.AppSetWithoutMocks)
	return nil
}

// InitMockedApp returns an app with its mocked dependencies, created
// via providers (Approach B).
func InitMockedApp() *src.AppWithMocks {
	wire.Build(src.MockAppSet)
	return nil
}
