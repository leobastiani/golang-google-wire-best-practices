package src

import (
	"fmt"
	"time"

	"github.com/google/wire"
)

type Options struct {
	Timer Timer
}

var AppSet = wire.NewSet(
	wire.FieldsOf(new(Options), "Timer"),
	wire.Struct(new(App), "*"),
	wire.Struct(new(Greeter), "*"),
)

type Timer interface {
	Now() time.Time
}

type RealTime struct{}

func (RealTime) Now() time.Time { return time.Now() }

type MockTimer struct {
	timer time.Time
}

func (m *MockTimer) Now() time.Time { return m.timer }

// Greeter issues greetings with the time provided by T.
type Greeter struct {
	Timer Timer
}

func (g Greeter) Greet() string {
	return fmt.Sprintf("Good day! It is %v", g.Timer.Now())
}

type App struct {
	Greeter Greeter
}
