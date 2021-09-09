package entity

import (
	"time"
)

type LumpSumPayment struct {
	ID                 int       `json:"id"`
	FeeId              string    `json:"fee_id"`
	CollectionId       string    `json:"collection_id"`
	ProposalNumber     string    `json:"proposal_number"`
	PolicyNumber       string    `json:"policy_number"`
	// FirstEffectiveDate time.Time `json:"first_effective_date"`
	EffectiveDate      time.Time `json:"effective_date"`
	SettledDate        time.Time `json:"settled_date"`
	PaymentMethod      string    `json:"payment_method"`
	BankName           string    `json:"bank_name"`
	TotalAmount        float64   `json:"total_amount"`
}
