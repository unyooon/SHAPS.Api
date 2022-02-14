package controller

import (
	"github.com/gin-gonic/gin"
	"shaps.api/usecase"
)

type SubscriptionController struct {
	create usecase.SubscriptionCreater
}

func NewSubscriptionController(uc usecase.SubscriptionCreater) *SubscriptionController {
	return &SubscriptionController{
		create: uc,
	}
}

// CreateSubscription godoc
// @Summary      Create Subscription
// @Description  create subscription
// @Tags         subscription
// @Accept       json
// @Produce      json
// @Param        body body dto.CreateSubscriptionRequest true "body"
// @Success      200
// @Router       /subscriptions [post]
func (sc *SubscriptionController) Post(c *gin.Context) {
	err := sc.create.Excecute(c)
	Handler(c, nil, err)
}
