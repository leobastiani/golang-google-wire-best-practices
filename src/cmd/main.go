package main

import (
	"fmt"

	"github.com/leobastiani/golang-google-wire-best-practices/src"
)

func main() {
	app1 := Init(src.Options{
		Timer: src.RealTime{},
	})
	fmt.Println(app1.Greeter.Greet())

	app2 := Init(src.Options{
		Timer: &src.MockTimer{},
	})
	fmt.Println(app2.Greeter.Greet())
}
