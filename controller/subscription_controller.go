package controller

import (
	"github.com/gin-gonic/gin"
	"shaps.api/usecase"
)

type SubscriptionController struct {
	create usecase.SubscriptionCreater
	join   usecase.SubscriptionJoiner
}

func NewSubscriptionController(uc usecase.SubscriptionCreater, uj usecase.SubscriptionJoiner) *SubscriptionController {
	return &SubscriptionController{
		create: uc,
		join: uj,
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
func (sc *SubscriptionController) Create(c *gin.Context) {
	result := sc.create.Execute(c)
	Handler(c, nil, result)
}

// CreateSubscription godoc
// @Summary      Join Subscription
// @Description  join subscription
// @Tags         subscription
// @Accept       json
// @Produce      json
// @Router       /subscriptions/{subscriptionId}/join [post]
func (sc *SubscriptionController) Join(c *gin.Context) {
	result := sc.join.Execute(c)
	Handler(c, nil, result)
}