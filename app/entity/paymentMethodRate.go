package entity

import (
	"time"
)

type PaymentMethodRate struct {
	ID               		int            	`gorm:"type:BIGINT UNSIGNED NOT NULL AUTO_INCREMENT" json:"id"`
	PaymentMethodCode 		string	    	`gorm:"type:VARCHAR(191) NOT NULL" json:"payment_method_code"`
	MinTransaction  		int    			`gorm:"type:INT UNSIGNED NOT NULL" json:"min_transaction"`
	MaxTransaction  		int    			`gorm:"type:INT UNSIGNED NOT NULL" json:"max_transaction"`
	// TransactionFee  		int    			`gorm:"type:DOUBLE(8,2) NOT NULL" json:"transaction_fee"`
	FormulaFee 		 		string    		`gorm:"type:VARCHAR(191) NOT NULL" json:"formula_fee"`
	Premi  					int    			`gorm:"type:DOUBLE(8,2) NOT NULL" json:"premi"`
	CreatedAt        		time.Time      	`gorm:"type:TIMESTAMP DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt        		time.Time      	`gorm:"type:TIMESTAMP DEFAULT CURRENT_TIMESTAMP" json:"updated_at"`
}
type PaymentMethodRateTesting struct {
	ID               		int            	`gorm:"type:BIGINT UNSIGNED NOT NULL" json:"id"`
	PaymentMethodCode 		string	    	`gorm:"type:VARCHAR(191) NOT NULL" json:"payment_method_code"`
	MinTransaction  		int    			`gorm:"type:INT UNSIGNED NOT NULL" json:"min_transaction"`
	MaxTransaction  		int    			`gorm:"type:INT UNSIGNED NOT NULL" json:"max_transaction"`
	// TransactionFee  		int    			`gorm:"type:DOUBLE(8,2) NOT NULL" json:"transaction_fee"`
	FormulaFee 		 		string    		`gorm:"type:VARCHAR(191) NOT NULL" json:"formula_fee"`
	Premi  					int    			`gorm:"type:DOUBLE(8,2) NOT NULL" json:"premi"`
	CreatedAt        		time.Time      	`gorm:"type:TIMESTAMP DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt        		time.Time      	`gorm:"type:TIMESTAMP DEFAULT CURRENT_TIMESTAMP" json:"updated_at"`
}

func (PaymentMethodRateTesting *PaymentMethodRateTesting) TableName() string {
	return "payment_method_rates"
}
