package entity

import (
	"time"
)

type SyOdsEtlPayment struct {
	ID                int       `json:"id" ;gorm:"AUTO_INCREMENT"`
	ProposalNumber    string    `json:"proposal_number"`
	PolicyNumber      string    `gorm:"column:policy_number;" json:"policy_number"`
	OdsFirstPaidDate  time.Time `json:"ods_first_paid_date"`
	OdsPaidDate       time.Time `json:"ods_paid_date"`
	SyPaidDate        time.Time `json:"sy_paid_date"`
	PaymentMethodName string    `json:"payment_method_name"`
	CollectionId      string    `json:"collection_id"`
	SyTotalAmount     float64   `json:"sy_total_amount"`
	OdsTotalAmount    float64   `json:"ods_total_amount"`
	PolicyStatus      string    `json:"policy_status"`
	Status            string    `json:"status"`
	StatusDescription string    `gorm:"type:text" json:"status_description"`
	UpdatedAt         time.Time `json:"updated_at"`
	// LumpSumPayment    LumpSumPayment `gorm:"foreignkey:PolicyNumber;AssociationForeignKey:PolicyNumber"`
	// UserPolicy        UserPolicy `gorm:"foreignkey:PolicyNumber;AssociationForeignKey:PolicyNumber"`
}

func (syOdsEtlPayment *SyOdsEtlPayment) TableName() string {
	return "etl_sy_ods_payments"
}
