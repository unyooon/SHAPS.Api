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

// CreateUser godoc
// @Summary      Create User
// @Description  create user
// @Tags         user
// @Accept       json
// @Produce      json
// @Success      200
// @Router       /me [post]
func (mc *MeController) Post(c *gin.Context) {
	err := mc.create.Excecute(c)
	Handler(c, nil, err)
}
