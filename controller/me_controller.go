package controller

import (
	"github.com/gin-gonic/gin"
	"shaps.api/usecase"
)

type MeController struct {
	create usecase.UserCreater
	read   usecase.UserReader
}

func NewMeController(
	uc usecase.UserCreater,
	ur usecase.UserReader,
) *MeController {
	return &MeController{
		create: uc,
		read:   ur,
	}
}

// CreateMe godoc
// @Summary      Create User
// @Description  create user
// @Tags         user
// @Accept       json
// @Produce      json
// @Success      200
// @Router       /me [post]
func (mc *MeController) Create(c *gin.Context) {
	err := mc.create.Excecute(c)
	Handler(c, nil, err)
}

// ReadMe godoc
// @Summary      Read User
// @Description  read user
// @Tags         user
// @Accept       json
// @Produce      json
// @Success      200
// @Router       /me [get]
func (mc *MeController) Read(c *gin.Context) {
	u, err := mc.read.Excecute(c)
	Handler(c, u, err)
}
