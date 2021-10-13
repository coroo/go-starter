package entity

import (
	"time"
)

type PaymentMethodLink struct {
	ID               		int            		`gorm:"type:BIGINT UNSIGNED NOT NULL AUTO_INCREMENT" json:"id"`
	PaymentMethodCode 		string	    		`gorm:"type:VARCHAR(191) NOT NULL" json:"payment_method_code"`
	ProcessType	 			string	    		`gorm:"type:VARCHAR(191) NOT NULL" json:"process_type"`
	Url 					string	    		`gorm:"type:TEXT NOT NULL" json:"url"`
	CreatedAt        		time.Time      		`gorm:"type:TIMESTAMP DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt        		time.Time      		`gorm:"type:TIMESTAMP DEFAULT CURRENT_TIMESTAMP" json:"updated_at"`
}
