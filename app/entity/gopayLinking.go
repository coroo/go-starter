package entity

import (
	"time"
)

type GopayLinking struct {
	ID               		int            	`gorm:"type:BIGINT UNSIGNED NOT NULL AUTO_INCREMENT" json:"id"`
	PhoneNumber				string			`gorm:"type:VARCHAR(191) NOT NULL" json:"phone_number"`
	AccountId				string			`gorm:"type:VARCHAR(191) NOT NULL" json:"account_id"`
	PaymentOptionToken		string			`gorm:"type:VARCHAR(191) NOT NULL" json:"payment_option_token"`
	CreatedAt        		time.Time      	`gorm:"type:TIMESTAMP DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt        		time.Time      	`gorm:"type:TIMESTAMP DEFAULT CURRENT_TIMESTAMP" json:"updated_at"`
}
type GopayLinkingTesting struct {
	ID               		int            	`gorm:"type:BIGINT UNSIGNED NOT NULL" json:"id"`
	PhoneNumber				string			`gorm:"type:VARCHAR(191) NOT NULL" json:"phone_number"`
	AccountId				string			`gorm:"type:VARCHAR(191) NOT NULL" json:"account_id"`
	PaymentOptionToken		string			`gorm:"type:VARCHAR(191) NOT NULL" json:"payment_option_token"`
	CreatedAt        		time.Time      	`gorm:"type:TIMESTAMP DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt        		time.Time      	`gorm:"type:TIMESTAMP DEFAULT CURRENT_TIMESTAMP" json:"updated_at"`
}

func (GopayLinking *GopayLinkingTesting) TableName() string {
	return "gopay_linkings"
}
