package external

import (
	"github.com/stripe/stripe-go/v72/client"
	"shaps.api/domain/setting"
)

type StripeClient struct {
	*client.API
}

func NewStripeClient(setting setting.Setting) *StripeClient {
	sc := &client.API{}
	sc.Init(setting.StripeApiKey, nil)
	return &StripeClient{
		sc,
	}
}
