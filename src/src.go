package src

import (
	"fmt"
	"time"

	"github.com/google/wire"
)

// AppSet is a provider set for creating a real app.
var AppSet = wire.NewSet(
	wire.Struct(new(App), "*"),
	wire.Struct(new(Greeter), "*"),
	wire.InterfaceValue(new(Timer), RealTime{}),
)

// AppSetWithoutMocks is a provider set for creating an app with mocked
// dependencies. The mocked dependencies are omitted and must be provided as
// arguments to the injector.
// It is used for Approach A.
var AppSetWithoutMocks = wire.NewSet(
	wire.Struct(new(App), "*"),
	wire.Struct(new(Greeter), "*"),
)

// MockAppSet is a provider set for creating a mocked app, including the mocked
// dependencies.
// It is used for Approach B.
var MockAppSet = wire.NewSet(
	wire.Struct(new(App), "*"),
	wire.Struct(new(Greeter), "*"),
	wire.Struct(new(AppWithMocks), "*"),
	// For each mocked dependency, add a provider and use wire.Bind to bind
	// the concrete type to the relevant interface.
	NewMockTimer,
	wire.Bind(new(Timer), new(*MockTimer)),
)

type Timer interface {
	Now() time.Time
}

// RealTime implements timer with the real time.
type RealTime struct{}

func (RealTime) Now() time.Time { return time.Now() }

// MockTimer implements timer using a mocked time.
type MockTimer struct {
	T time.Time
}

func NewMockTimer() *MockTimer      { return &MockTimer{} }
func (m *MockTimer) Now() time.Time { return m.T }

// Greeter issues greetings with the time provided by T.
type Greeter struct {
	T Timer
}

func (g Greeter) Greet() string {
	return fmt.Sprintf("Good day! It is %v", g.T.Now())
}

type App struct {
	G Greeter
}

func (a App) Greet() string {
	return a.G.Greet()
}

// AppWithMocks is used for Approach B, to return the app plus its mocked
// dependencies.
type AppWithMocks struct {
	App App
	Mt  *MockTimer
}
