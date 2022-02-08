package controller

import (
	"github.com/gin-gonic/gin"
	"shaps.api/usecase"
)

type MeController struct {
	create usecase.UserCreater
}

func NewMeController(uc usecase.UserCreater) *MeController {
	return &MeController{
		create: uc,
	}
}

func (mc *MeController) Post(c *gin.Context) {
	mc.create.Excecute(c)
	c.JSON(200, nil)
}
