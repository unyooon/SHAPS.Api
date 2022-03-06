package controller

import (
	"github.com/gin-gonic/gin"
	"shaps.api/usecase"
)

type UserController struct {
	create usecase.UserCreater
	read   usecase.UserReader
}

func NewUserController(
	uc usecase.UserCreater,
	ur usecase.UserReader,
) *UserController {
	return &UserController{
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
func (uc *UserController) Create(c *gin.Context) {
	err := uc.create.Excecute(c)
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
func (uc *UserController) Read(c *gin.Context) {
	u, err := uc.read.Excecute(c)
	Handler(c, u, err)
}
