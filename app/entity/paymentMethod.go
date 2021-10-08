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
	ID               		int            	`gorm:"type:BIGINT UNSIGNED NOT NULL AUTO_INCREMENT" json:"id"`
	Code 					string	    	`gorm:"unique;type:varchar(191) NOT NULL" json:"code"`
	InitPaymentCode  		string    		`gorm:"type:varchar(191) NULL default NULL;default:null" json:"init_payment_code"`
	RenewalPaymentCode 		string    		`gorm:"type:varchar(191) NULL default NULL;default:null" json:"renewal_payment_code"`
	FastpayCode 		 	string    		`gorm:"type:varchar(191) NOT NULL" json:"fastpay_code"`
	BankCode 		 		string    		`gorm:"type:varchar(191) NULL default NULL;default:null" json:"bank_code"`
	Name 		 			string    		`gorm:"type:varchar(191) NOT NULL" json:"name"`
	PaymentLogo 		 	string    		`gorm:"type:TEXT NOT NULL" json:"payment_logo"`
	RateApplied 		 	bool    		`gorm:"type:tinyint(1) NOT NULL" json:"rate_applied"`
	TransactionFee 		 	int    			`gorm:"type:double(8,2) NOT NULL" json:"transaction_fee"`
	PercentageFee 		 	int    			`gorm:"type:INT NOT NULL" json:"percentage_fee"`
	// UserId 		 			int    			`gorm:"type:INT UNSIGNED NOT NULL" json:"user_id"`
	Status 		 			string    		`gorm:"type:enum('active','inactive') NOT NULL DEFAULT 'inactive';default:'inactive'" json:"status"`
	Spec 		 			string    		`gorm:"type:LONGTEXT" json:"spec"`
	CreatedAt        		time.Time      	`gorm:"type:TIMESTAMP DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt        		time.Time      	`gorm:"type:TIMESTAMP DEFAULT CURRENT_TIMESTAMP" json:"updated_at"`
}
