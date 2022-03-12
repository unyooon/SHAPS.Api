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

// CreateStripeConnect godoc
// @Summary      Create StripeConnect
// @Description  create stripeconnect
// @Tags         me
// @Accept       json
// @Produce      json
// @Param        body body dto.CreateStripeConnectRequest true "body"
// @Success      200
// @Router       /me/stripe-connect [post]
func (mc *MeController) CreateStripeConnect(c *gin.Context) {
	err := mc.createStripeConnect.Execute(c)
	Handler(c, nil, err)
}
