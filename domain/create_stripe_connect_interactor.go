package domain

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v72"
	"shaps.api/core/constants"
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
	if ve := validatation.RequestValidate(req, c); ve != nil {
		return ve
	}

	json.Unmarshal(body, &req)

	uid, exists := c.Get("userId")
	if !exists {
		return &exception.CustomException{
			Code:    constants.NotFoundCode,
			Message: constants.NotFoundUser,
			Err:     errors.New("userId is not exists"),
		}
	}

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

	connect, err := i.StripeClient.Account.New(params)
	if err != nil {
		e := &exception.CustomException{
			Code:    constants.InternalServerErrorCode,
			Message: constants.StripeError,
			Err:     err,
		}
		return e
	}

	user, readErr := i.UserRepository.Read(uid.(string))
	if readErr != nil {
		return readErr
	}

	user.ConnectId = connect.ID

	if _, updateErr := i.UserRepository.Update(user); updateErr != nil {
		return updateErr
	}

	return nil
}
