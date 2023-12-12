//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/leobastiani/golang-google-wire-best-practices/src"
)

func Init(options src.Options) *src.App {
	wire.Build(src.AppSet)
	return nil
}
