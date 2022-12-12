package entity

import "github.com/jinzhu/gorm"

type Contract struct {
	gorm.Model
	SubscriptionID uint       `column:"subscriptionId"`
	Subscription Subscription
	UserID string       			`column:"userId"`
}
