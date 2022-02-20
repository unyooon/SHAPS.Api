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
}

func NewRouting(
	sc *controller.SubscriptionController,
	uc *controller.UserController,
	setting setting.Setting,
) *Routing {
	r := &Routing{
		Gin:     gin.New(),
		Setting: setting,
		sc:      sc,
		uc:      uc,
	}

	r.Gin.Use(gin.Recovery())
	r.Gin.Use(middleware.Logging)
	r.setRouting()

	return r
}

func (r *Routing) setRouting() {
	r.Gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := r.Gin.Group("/api/v1", middleware.ValidateToken(r.Setting))
	{
		v1.POST("/subscriptions", func(c *gin.Context) { r.sc.Create(c) })
		v1.POST("/users", func(c *gin.Context) { r.uc.Create(c) })
		v1.GET("/users", func(c *gin.Context) { r.uc.Read(c) })
	}
}

func (r *Routing) Run() {
	r.Gin.Run(r.Setting.Port)
}
