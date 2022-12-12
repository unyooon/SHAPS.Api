package entity

import "github.com/jinzhu/gorm"

type Construct struct {
	gorm.Model
	SubscriptionID uint       `column:"subscriptionId"`
	Subscription Subscription
	UserID string       			`column:"userId"`
}
