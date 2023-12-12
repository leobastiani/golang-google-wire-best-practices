//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/leobastiani/golang-google-wire-best-practices/src"
)

func Init(mt src.Timer) *src.App {
	wire.Build(src.AppSet)
	return nil
}
