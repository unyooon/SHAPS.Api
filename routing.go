package main

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"shaps.api/controller"
	"shaps.api/domain/setting"
	"shaps.api/middleware"
)

type Routing struct {
	Gin     *gin.Engine
	Setting setting.Setting
	sc      *controller.SubscriptionController
	uc      *controller.UserController
	mc      *controller.MeController
}

func NewRouting(
	sc *controller.SubscriptionController,
	uc *controller.UserController,
	mc *controller.MeController,
	setting setting.Setting,
) *Routing {
	r := &Routing{
		Gin:     gin.New(),
		Setting: setting,
		sc:      sc,
		uc:      uc,
		mc:      mc,
	}

	r.Gin.Use(gin.Recovery())
	r.Gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Gin.Use(middleware.ValidateToken(r.Setting))
	r.Gin.Use(middleware.HttpLogger)
	r.setRouting()

	return r
}

func (r *Routing) setRouting() {
	v1 := r.Gin.Group("/api/v1")
	{
		v1.GET("/me", func(c *gin.Context) { r.mc.ReadMe(c) })
		v1.POST("/me/stripe-connect", func(c *gin.Context) { r.mc.CreateStripeConnect(c) })

		v1.GET("/users", func(c *gin.Context) { r.uc.Read(c) })
		v1.POST("/users", func(c *gin.Context) { r.uc.Create(c) })

		v1.POST("/subscriptions", func(c *gin.Context) { r.sc.Create(c) })
	}
}

func (r *Routing) Run() {
	r.Gin.Run(r.Setting.Port)
}
