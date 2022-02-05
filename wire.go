//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

func InitializeHandler(dsn string) *Routing {
	wire.Build(SuperSet, NewRouting)
	return nil
}
