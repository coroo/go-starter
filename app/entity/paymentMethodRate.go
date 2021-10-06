package entity

import (
	"time"
)

type PaymentMethodRate struct {
	ID               		int            	`gorm:"type:BIGINT UNSIGNED NOT NULL AUTO_INCREMENT" json:"id"`
	PaymentMethodId 		string	    	`gorm:"type:INT UNSIGNED NOT NULL" json:"payment_method_id"`
	MinTransaction  		string    		`gorm:"type:INT UNSIGNED NOT NULL" json:"min_transaction"`
	MaxTransaction  		string    		`gorm:"type:INT UNSIGNED NOT NULL" json:"max_transaction"`
	TransactionFee  		string    		`gorm:"type:DOUBLE(8,2) NOT NULL" json:"transaction_fee"`
	Premi  					string    		`gorm:"type:DOUBLE(8,2) NOT NULL" json:"premi"`
	CreatedAt        		time.Time      	`gorm:"type:TIMESTAMP DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt        		time.Time      	`gorm:"type:TIMESTAMP DEFAULT CURRENT_TIMESTAMP" json:"updated_at"`
}
