package src

import (
	"fmt"
	"time"

	"github.com/google/wire"
)

var AppSet = wire.NewSet(
	wire.Struct(new(App), "*"),
	wire.Struct(new(Greeter), "*"),
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
