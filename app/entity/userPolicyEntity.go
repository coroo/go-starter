package entity

import (
	"time"
)

type UserPolicy struct {
	// gorm.Model
	// LumpSumPayment         LumpSumPayment `gorm:"foreignkey:PolicyNumber";association_foreignkey:PolicyNumber`
	ID                     int       `json:"id"`
	PolicyNumber           string    `json:"policy_number"`
	ProposalNumber         string    `json:"proposal_number"`
	ProposalDate           time.Time `json:"proposal_date"`
	CommDate               time.Time `json:"comm_date"`
	IssuedDate             time.Time `json:"issued_date"`
	ProductCode            string    `json:"product_code"`
	ProductName            string    `json:"product_name"`
	BenefitLevel           string    `json:"benefit_level"`
	PhName                 string    `json:"ph_name"`
	PhGender               string    `json:"ph_gender"`
	PhDob                  time.Time `json:"ph_dob"`
	PhProvince             string    `json:"ph_province"`
	PhCitizenId            string    `json:"ph_citizen_id"`
	PhEmail                string    `json:"ph_email"`
	PhMobile               string    `json:"ph_mobile"`
	PhHomeTel              string    `json:"ph_home_tel"`
	Number                 int       `json:"number"`
	LifeAssured            string    `json:"life_assured"`
	LaDob                  time.Time `json:"la_dob"`
	LaGender               string    `json:"la_gender"`
	LaIdNumber             string    `json:"la_id_number"`
	LaRelation             string    `json:"la_relation"`
	TotalPremium           string    `json:"total_premium"`
	Beneficiary            string    `json:"beneficiary"`
	BeneDob                time.Time `json:"bene_dob"`
	BeneGender             string    `json:"bene_gender"`
	BeneIdNumber           string    `json:"bene_id_number"`
	Frequency              string    `json:"frequency"`
	PolicyAmount           float64   `json:"policy_amount"`
	LastPaidDate           time.Time `json:"last_paid_date"`
	Ndd                    string    `json:"ndd"`
	OverduePremium         int       `json:"overdue_premium"`
	RenewalSuspense        int       `json:"renewal_suspense"`
	Mop                    string    `json:"mop"`
	PolicyStatus           string    `json:"policy_status"`
	AdditionalStatus       string    `json:"additional_status"`
	StatusDate             time.Time `json:"status_date"`
	TerminationLapseReasos string    `json:"termination_lapse_reasos"`
	UpdatedAt              time.Time `json:"updated_at"`
}
