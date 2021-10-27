package entity

import (
	"time"
)

type UserInvoiceLog struct {
	ID               		int            		`gorm:"type:BIGINT UNSIGNED NOT NULL AUTO_INCREMENT" json:"id"`
	PaymentMethodCode 		string	    		`gorm:"type:VARCHAR(191) NOT NULL" json:"payment_method_code"`
	PolicyGroupNumber  		string    			`gorm:"type:VARCHAR(191) NOT NULL" json:"policy_group_number"`
	PaymentType 			string    			`gorm:"type:ENUM('first_payment','recurring','insufficient_payment') NULL;default:'first_payment'" json:"payment_type"`
	InvoiceCode 		 	string    			`gorm:"type:ENUM('FP#','REC#','IP#') NULL; default:'FP#'" json:"invoice_code"`
	SummaryToken 		 	string    			`gorm:"type:TEXT NOT NULL" json:"summary_token"`
	VirtualAccount 		 	string    			`gorm:"type:VARCHAR(191) NULL" json:"virtual_account"`
	InvoiceNumber 		 	string    			`gorm:"type:varchar(191) NOT NULL" json:"invoice_number"`
	MemberId 		 		string    			`gorm:"type:VARCHAR(191) NULL" json:"member_id"`
	PaymentCycle 		 	string    			`gorm:"type:ENUM('MONTHLY','YEARLY') NOT NULL" json:"payment_cycle"`
	TransactionFee 		 	string    			`gorm:"type:VARCHAR(191) NOT NULL; default:0" json:"transaction_fee"`
	AgentFee 		 		string    			`gorm:"type:VARCHAR(191) NOT NULL; default:0" json:"agent_fee"`
	TotalPremium 		 	string    			`gorm:"type:VARCHAR(191) NOT NULL" json:"total_premium"`
	TotalPayment		 	string    			`gorm:"type:VARCHAR(191) NOT NULL" json:"total_payment"`
	PromoCode			 	string    			`gorm:"type:VARCHAR(191) NULL" json:"promocode"`
	Status 		 			string    			`gorm:"type:VARCHAR(191) NOT NULL;default:'init'" json:"status"`
	LogEncryptedId 		 	string    			`gorm:"type:VARCHAR(191) NULL" json:"log_encrypted_id"`
	TrxFaspayId 		 	string    			`gorm:"type:VARCHAR(191) NULL" json:"trx_faspay_id"`
	Type 		 			string    			`gorm:"type:ENUM('script','show','redirect') NULL" json:"type"`
	RequestMessage	 		string    			`gorm:"type:TEXT NULL" json:"request_message"`
	ResponseMessage 		string    			`gorm:"type:TEXT NULL" json:"response_message"`
	PaidAt 		 			time.Time    		`gorm:"type:TIMESTAMP NULL" json:"paid_at"`
	ExpiredAt	 		 	time.Time    		`gorm:"type:TIMESTAMP NULL" json:"expired_at"`
	CreatedAt        		time.Time      		`gorm:"type:TIMESTAMP DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt        		time.Time      		`gorm:"type:TIMESTAMP DEFAULT CURRENT_TIMESTAMP" json:"updated_at"`
}

type UserInvoiceLogTesting struct {
	ID               		int            		`gorm:"type:BIGINT UNSIGNED NOT NULL" json:"id"`
	PaymentMethodCode 		string	    		`gorm:"type:VARCHAR(191) NOT NULL" json:"payment_method_code"`
	PolicyGroupNumber  		string    			`gorm:"type:VARCHAR(191) NOT NULL" json:"policy_group_number"`
	PaymentType 			string    			`gorm:"type:VARCHAR(191) NULL;default:'first_payment'" json:"payment_type"`
	InvoiceCode 		 	string    			`gorm:"type:VARCHAR(191) NULL; default:'FP#'" json:"invoice_code"`
	SummaryToken 		 	string    			`gorm:"type:TEXT NOT NULL" json:"summary_token"`
	VirtualAccount 		 	string    			`gorm:"type:VARCHAR(191) NULL" json:"virtual_account"`
	InvoiceNumber 		 	string    			`gorm:"type:varchar(191) NOT NULL" json:"invoice_number"`
	MemberId 		 		string    			`gorm:"type:VARCHAR(191) NULL" json:"member_id"`
	PaymentCycle 		 	string    			`gorm:"type:VARCHAR(191) NOT NULL" json:"payment_cycle"`
	TransactionFee 		 	string    			`gorm:"type:VARCHAR(191) NOT NULL; default:0" json:"transaction_fee"`
	AgentFee 		 		string    			`gorm:"type:VARCHAR(191) NOT NULL; default:0" json:"agent_fee"`
	TotalPremium 		 	string    			`gorm:"type:VARCHAR(191) NOT NULL" json:"total_premium"`
	TotalPayment		 	string    			`gorm:"type:VARCHAR(191) NOT NULL" json:"total_payment"`
	PromoCode			 	string    			`gorm:"type:VARCHAR(191) NULL" json:"promocode"`
	Status 		 			string    			`gorm:"type:VARCHAR(191) NOT NULL;default:'init'" json:"status"`
	LogEncryptedId 		 	string    			`gorm:"type:VARCHAR(191) NULL" json:"log_encrypted_id"`
	TrxFaspayId 		 	string    			`gorm:"type:VARCHAR(191) NULL" json:"trx_faspay_id"`
	Type 		 			string    			`gorm:"type:VARCHAR(191) NULL" json:"type"`
	RequestMessage	 		string    			`gorm:"type:TEXT NULL" json:"request_message"`
	ResponseMessage 		string    			`gorm:"type:TEXT NULL" json:"response_message"`
	PaidAt 		 			time.Time    		`gorm:"type:TIMESTAMP NULL" json:"paid_at"`
	ExpiredAt	 		 	time.Time    		`gorm:"type:TIMESTAMP NULL" json:"expired_at"`
	CreatedAt        		time.Time      		`gorm:"type:TIMESTAMP DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt        		time.Time      		`gorm:"type:TIMESTAMP DEFAULT CURRENT_TIMESTAMP" json:"updated_at"`
}

func (UserInvoiceLogTesting *UserInvoiceLogTesting) TableName() string {
	return "user_invoice_logs"
}