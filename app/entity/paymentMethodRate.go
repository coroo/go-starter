package entity

import (
	"time"
)

type PaymentMethodRate struct {
	ID               		int            	`gorm:"type:BIGINT UNSIGNED NOT NULL AUTO_INCREMENT" json:"id"`
	PaymentMethodId 		int	    	`gorm:"type:INT UNSIGNED NOT NULL" json:"payment_method_id"`
	MinTransaction  		int    		`gorm:"type:INT UNSIGNED NOT NULL" json:"min_transaction"`
	MaxTransaction  		int    		`gorm:"type:INT UNSIGNED NOT NULL" json:"max_transaction"`
	TransactionFee  		int    		`gorm:"type:DOUBLE(8,2) NOT NULL" json:"transaction_fee"`
	Premi  					int    		`gorm:"type:DOUBLE(8,2) NOT NULL" json:"premi"`
	CreatedAt        		time.Time      	`gorm:"type:TIMESTAMP DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt        		time.Time      	`gorm:"type:TIMESTAMP DEFAULT CURRENT_TIMESTAMP" json:"updated_at"`
}
