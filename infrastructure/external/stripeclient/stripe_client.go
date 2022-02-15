package stripeclient

import (
	"github.com/stripe/stripe-go/v72/client"
)

type Client struct {
	*client.API
}

func NewStripeClietn(sk string) *Client {
	sc := &client.API{}
	sc.Init(sk, nil)
	return &Client{
		sc,
	}
}
