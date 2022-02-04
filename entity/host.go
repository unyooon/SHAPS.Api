package entity

import "github.com/jinzhu/gorm"

type Host struct {
	gorm.Model
	UserId         uint `column:"user_id"`
	SubscriptionId uint `column:"subscription_id"`
}
