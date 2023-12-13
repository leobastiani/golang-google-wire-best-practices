//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/leobastiani/golang-google-wire-best-practices/internal/src"
)

func Init(options src.Options) *src.App {
	wire.Build(src.AppSet)
	return nil
}
