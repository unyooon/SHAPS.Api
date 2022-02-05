package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"shaps.api/controller"
)

type Routing struct {
	Gin  *gin.Engine
	Port string
	sc   *controller.SubscriptionController
}

func NewRouting(
	sc *controller.SubscriptionController,
) *Routing {
	r := &Routing{
		Gin:  gin.Default(),
		Port: os.Getenv("PORT"),
		sc:   sc,
	}

	r.setRouting()

	return r
}

func (r *Routing) setRouting() {
	r.Gin.POST("/subscriptions", func(c *gin.Context) { r.sc.Post(c) })
}

func (r *Routing) Run() {
	r.Gin.Run(r.Port)
}
