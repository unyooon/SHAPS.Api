package controller

import (
	"github.com/gin-gonic/gin"
	"shaps.api/usecase"
)

type MeController struct {
	StripeConnectCreater usecase.StripeConnectCreater
	MeReader             usecase.MeReader
	MeUpdater            usecase.MeUpdater
	HostsReader usecase.HostsReader
}

func NewMeController(
	scc usecase.StripeConnectCreater,
	mr usecase.MeReader,
	mu usecase.MeUpdater,
	hr usecase.HostsReader,
) *MeController {
	return &MeController{
		StripeConnectCreater: scc,
		MeReader:             mr,
		MeUpdater:            mu,
		HostsReader: hr,
	}
}

// ReadMe godoc
// @Summary      Read Me
// @Description  read me
// @Tags         me
// @Accept       json
// @Produce      json
// @Success      200 {object} dto.ReadMeResponse
// @Router       /me [get]
func (mc *MeController) ReadMe(c *gin.Context) {
	res, err := mc.MeReader.Execute(c)
	Handler(c, res, err)
}

// UpdateMe godoc
// @Summary      Update Me
// @Description  update me
// @Tags         me
// @Accept       json
// @Produce      json
// @Success      200
// @Router       /me [put]
func (mc *MeController) UpdateMe(c *gin.Context) {
	err := mc.MeUpdater.Execute(c)
	Handler(c, nil, err)
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
	err := mc.StripeConnectCreater.Execute(c)
	Handler(c, nil, err)
}

// ReadHosts godoc
// @Summary      Read Hosts
// @Description  read hosts
// @Tags         me
// @Accept       json
// @Produce      json
// @Param 			 size query int true "Page Size"
// @Param 			 page query int true "Page Number"
// @Success      200 {object} dto.ReadHostsResponse
// @Router       /me/hosts [get]
func (mc *MeController) ReadHosts(c *gin.Context) {
	res, err := mc.HostsReader.Execute(c)
	Handler(c, res, err)
}
