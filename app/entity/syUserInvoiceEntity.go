package entity

import (
	"time"
)

type SyUserInvoice struct {
	ID                int          `json:"id"`
	PolicyNumber      string       `json:"policy_number"`
	PolicyGroupNumber string       `json:"policy_group_number"`
	ProposalNumber    string       `json:"proposal_number"`
	PaymentMethodName string       `json:"payment_method_name"`
	TotalPremium      float64      `json:"total_premium"`
	Status            string       `json:"status"`
	PaidAt            time.Time    `json:"paid_at"`
	SyUserPolicy      SyUserPolicy `gorm:"foreignKey:PolicyNumber;references:PolicyNumber"`
	// ProposalNumber string    `json:"proposal_number"`
	// PolicyNumber   string    `json:"policy_number";sql:"index"`
	// PaymentMethod  string    `json:"payment_method"`
	// TotalAmount    float64   `json:"Total_amount"`
	// UserPolicyID   uint64    `json:"-"`
}

func (syUserInvoice *SyUserInvoice) TableName() string {
	return "user_invoices"
}
