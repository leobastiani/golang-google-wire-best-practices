package main

import (
	"fmt"

	"github.com/leobastiani/golang-google-wire-best-practices/src"
)

func main() {
	// Create a "real" greeter.
	// Greet() will include the real current time, so elide it for repeatable
	// tests.
	fmt.Printf("Real time greeting: %s [current time elided]\n", InitApp().Greet())

	// There are two approaches for creating an app with mocks.

	// Approach A: create the mocks manually, and pass them to an injector.
	// This approach is useful if you need to prime the mocks beforehand.
	fmt.Println("Approach A")
	mt := src.NewMockTimer()
	mockedApp := InitMockedAppFromArgs(mt)
	fmt.Println(mockedApp.Greet()) // prints greeting with time = zero time
	mt.T = mt.T.AddDate(1999, 0, 0)
	fmt.Println(mockedApp.Greet()) // prints greeting with time = year 2000

	// Approach B: allow the injector to create the mocks, and return a struct
	// that includes the resulting app plus the mocks.
	fmt.Println("Approach B")
	appWithMocks := InitMockedApp()
	fmt.Println(appWithMocks.App.Greet()) // prints greeting with time = zero time
	appWithMocks.Mt.T = appWithMocks.Mt.T.AddDate(999, 0, 0)
	fmt.Println(appWithMocks.App.Greet()) // prints greeting with time = year 1000
}
