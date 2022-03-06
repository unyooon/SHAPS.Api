package controller

import (
	"github.com/gin-gonic/gin"
	"shaps.api/usecase"
)

type MeController struct {
	createStripeConnect usecase.StripeConnectCreater
}

func NewMeController(
	scc usecase.StripeConnectCreater,
) *MeController {
	return &MeController{
		createStripeConnect: scc,
	}
}

func (mc *MeController) CreateStripeConnect(c *gin.Context) {
	err := mc.createStripeConnect.Excute(c)
	Handler(c, nil, err)
}
