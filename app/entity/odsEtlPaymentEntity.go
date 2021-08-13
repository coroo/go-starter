package entity

import (
	"time"
)

type OdsEtlPayment struct {
	ID               int       `json:"id"`
	CollectionId     string    `json:"collection_id"`
	ProposalNumber   string    `json:"proposal_number"`
	PolicyNumber     string    `gorm:"column:policy_number;unique;" json:"policy_number"`
	FirstPaymentDate time.Time `json:"fist_payment_date"`
	PaymentDate      time.Time `json:"payment_date"`
	PaymentMethod    string    `json:"payment_method"`
	TotalAmount      float64   `json:"total_amount"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func (odsEtlPayment *OdsEtlPayment) TableName() string {
	return "etl_ods_payments"
}
