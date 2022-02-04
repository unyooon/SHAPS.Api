package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

type Routing struct {
	gin  *gin.Engine
	port string
}

func NewRouting() *Routing {
	r := &Routing{
		gin:  gin.Default(),
		port: os.Getenv("PORT"),
	}

	r.setRouting()

	return r
}

func (r *Routing) setRouting() {

}

func (r *Routing) Run() {
	r.gin.Run(r.port)
}
