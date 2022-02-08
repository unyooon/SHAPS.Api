package controller

import (
	"github.com/gin-gonic/gin"
	"shaps.api/usecase"
)

type SubscriptionController struct {
	create usecase.SubscriptionCreater
}

func NewSubscriptionController(c usecase.SubscriptionCreater) *SubscriptionController {
	return &SubscriptionController{
		create: c,
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
func (s *SubscriptionController) Post(c *gin.Context) {
	s.create.Excecute(c)
	c.JSON(200, nil)
}
