package domain

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v72"
	"shaps.api/core/validatation"
	"shaps.api/domain/dto"
	"shaps.api/domain/exception"
	"shaps.api/infrastructure/external"
	"shaps.api/infrastructure/repository"
)

type CreateStripeConnectInteractor struct {
	UserRepository repository.UserRepositoryInterface
	StripeClient   *external.StripeClient
}

func NewCreateStripeConnectInteractor(
	r repository.UserRepositoryInterface,
	sc *external.StripeClient,
) *CreateStripeConnectInteractor {
	return &CreateStripeConnectInteractor{
		UserRepository: r,
		StripeClient:   sc,
	}
}

func (i *CreateStripeConnectInteractor) Execute(c *gin.Context) *exception.CustomException {
	body, _ := ioutil.ReadAll(c.Request.Body)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	req := new(dto.CreateStripeConnectRequest)
	ve := validatation.RequestValidate(req, c)
	if ve.Code == exception.BadRequestCode {
		return ve
	}
	json.Unmarshal(body, &req)

	ip := c.ClientIP()
	time := time.Now().Unix()

	params := &stripe.AccountParams{
		Capabilities: &stripe.AccountCapabilitiesParams{
			CardPayments: &stripe.AccountCapabilitiesCardPaymentsParams{
				Requested: stripe.Bool(true),
			},
			Transfers: &stripe.AccountCapabilitiesTransfersParams{
				Requested: stripe.Bool(true),
			},
		},
		Type:         stripe.String("custom"),
		BusinessType: stripe.String("individual"),
		BusinessProfile: &stripe.AccountBusinessProfileParams{
			MCC:                &req.Mcc,
			URL:                &req.BuisinessUrl,
			ProductDescription: &req.ProductionDescription,
		},
		TOSAcceptance: &stripe.AccountTOSAcceptanceParams{
			IP:   &ip,
			Date: &time,
		},
		Individual: &stripe.PersonParams{
			AddressKanji: &stripe.AccountAddressParams{
				PostalCode: &req.PostalCode,
				Line1:      &req.Line1,
				Line2:      &req.Line2,
			},
			AddressKana: &stripe.AccountAddressParams{
				PostalCode: &req.PostalCode,
				Line1:      &req.Line1,
				Line2:      &req.Line2,
			},
			FirstNameKanji: &req.FirstName,
			FirstNameKana:  &req.FirstNameKana,
			LastNameKanji:  &req.LastName,
			LastNameKana:   &req.LastNameKana,
			DOB: &stripe.DOBParams{
				Day:   &req.DobDay,
				Month: &req.DobMonth,
				Year:  &req.DobYear,
			},
			Email: &req.Email,
			Phone: &req.Phone,
		},
	}

	_, err := i.StripeClient.Account.New(params)
	if err != nil {
		e := &exception.CustomException{
			Code:    exception.InternalServerErrorCode,
			Message: exception.StripeError,
			Err:     err,
		}
		return e
	}

	return nil
}
