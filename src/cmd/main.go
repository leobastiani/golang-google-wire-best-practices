package main

import (
	"fmt"

	"github.com/leobastiani/golang-google-wire-best-practices/src"
)

func main() {
	mt := &src.MockTimer{}
	app := Init(mt)

	fmt.Println(app.G.Greet())

	mt.T = mt.T.AddDate(1999, 0, 0)
	fmt.Println(app.G.Greet())
}
