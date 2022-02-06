//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"shaps.api/domain/setting"
)

func InitializeHandler(s setting.Setting) *Routing {
	wire.Build(SuperSet, NewRouting)
	return nil
}
