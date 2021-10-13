package entity

import (
	"time"
)

type Status string

const (
	Inactive Status = "inactive"
	Active Status = "active"
)

type PaymentMethod struct {
	ID               		int            		`gorm:"type:BIGINT UNSIGNED NOT NULL AUTO_INCREMENT" json:"id"`
	Code 					string	    		`gorm:"unique;type:varchar(191) NOT NULL" json:"code"`
	InitPaymentCode  		string    			`gorm:"type:varchar(191) NULL default NULL;default:null" json:"init_payment_code"`
	RenewalPaymentCode 		string    			`gorm:"type:varchar(191) NULL default NULL;default:null" json:"renewal_payment_code"`
	FastpayCode 		 	string    			`gorm:"type:varchar(191) NOT NULL" json:"fastpay_code"`
	BankCode 		 		string    			`gorm:"type:varchar(191) NULL default NULL;default:null" json:"bank_code"`
	Name 		 			string    			`gorm:"type:varchar(191) NOT NULL" json:"name"`
	PaymentLogo 		 	string    			`gorm:"type:TEXT NOT NULL" json:"payment_logo"`
	// RateApplied 		 	bool    			`gorm:"type:tinyint(1) NOT NULL" json:"rate_applied"`
	// UserId 		 		int    				`gorm:"type:INT UNSIGNED NOT NULL" json:"user_id"`
	Status 		 			string    			`gorm:"type:enum('active','inactive') NOT NULL DEFAULT 'inactive';default:'inactive'" json:"status"`
	Spec 		 			string    			`gorm:"type:LONGTEXT" json:"spec"`
	PaymentMethodRate 		[]PaymentMethodRate   `gorm:"foreignKey:PaymentMethodCode;references:Code"`
	CreatedAt        		time.Time      		`gorm:"type:TIMESTAMP DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt        		time.Time      		`gorm:"type:TIMESTAMP DEFAULT CURRENT_TIMESTAMP" json:"updated_at"`
}

type PaymentMethodWithPremium struct {
	ID               		int            	`gorm:"type:BIGINT UNSIGNED NOT NULL AUTO_INCREMENT" json:"id"`
	Code 					string	    	`gorm:"unique;type:varchar(191) NOT NULL" json:"code"`
	InitPaymentCode  		string    		`gorm:"type:varchar(191) NULL default NULL;default:null" json:"init_payment_code"`
	RenewalPaymentCode 		string    		`gorm:"type:varchar(191) NULL default NULL;default:null" json:"renewal_payment_code"`
	FastpayCode 		 	string    		`gorm:"type:varchar(191) NOT NULL" json:"fastpay_code"`
	BankCode 		 		string    		`gorm:"type:varchar(191) NULL default NULL;default:null" json:"bank_code"`
	Name 		 			string    		`gorm:"type:varchar(191) NOT NULL" json:"name"`
	PaymentLogo 		 	string    		`gorm:"type:TEXT NOT NULL" json:"payment_logo"`
	// RateApplied 		 	bool    		`gorm:"type:tinyint(1) NOT NULL" json:"rate_applied"`
	// UserId 		 		int    			`gorm:"type:INT UNSIGNED NOT NULL" json:"user_id"`
	Status 		 			string    		`gorm:"type:enum('active','inactive') NOT NULL DEFAULT 'inactive';default:'inactive'" json:"status"`
	Spec 		 			string    		`gorm:"type:LONGTEXT" json:"spec"`
	TotalPremium 		 	int    			`json:"total_premium"`
	Fee 		 			int    			`json:"fee"`
	TotalPayment		 	int    			`json:"total_payment"`
	CreatedAt        		time.Time      	`gorm:"type:TIMESTAMP DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt        		time.Time      	`gorm:"type:TIMESTAMP DEFAULT CURRENT_TIMESTAMP" json:"updated_at"`
}
