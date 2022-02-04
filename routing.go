package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"shaps.api/infrastructure"
)

type Routing struct {
	Db   *infrastructure.Db
	Gin  *gin.Engine
	Port string
}

func NewRouting(d *infrastructure.Db) *Routing {
	r := &Routing{
		Db:   d,
		Gin:  gin.Default(),
		Port: os.Getenv("PORT"),
	}

	r.setRouting()

	return r
}

func (r *Routing) setRouting() {

}

func (r *Routing) Run() {
	r.Gin.Run(r.Port)
}
