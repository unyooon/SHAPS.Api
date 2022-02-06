package main

import (
	"github.com/gin-gonic/gin"
	"shaps.api/controller"
	"shaps.api/domain/setting"
	"shaps.api/middleware"
)

type Routing struct {
	Gin     *gin.Engine
	Setting setting.Setting
	sc      *controller.SubscriptionController
}

func NewRouting(
	sc *controller.SubscriptionController,
	setting setting.Setting,
) *Routing {
	r := &Routing{
		Gin:     gin.Default(),
		Setting: setting,
		sc:      sc,
	}

	r.Gin.Use(middleware.ValidateToken(r.Setting))
	r.setRouting()

	return r
}

func (r *Routing) setRouting() {
	r.Gin.POST("/subscriptions", func(c *gin.Context) { r.sc.Post(c) })
}

func (r *Routing) Run() {
	r.Gin.Run(r.Setting.Port)
}
