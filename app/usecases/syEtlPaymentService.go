package usecases

import (
	"encoding/json"
	"os"
	entity "github.com/coroo/go-starter/app/entity"
	repositories "github.com/coroo/go-starter/app/repositories"
	"github.com/coroo/go-starter/app/utils"
	"time"

	// "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type SyEtlPaymentService interface {
	CreateSyEtlPayment(entity.SyEtlPayment) error
	SyOdsMapEtlLatestPayment() []entity.SyEtlPayment
	GetAllSyEtlPayments() []entity.SyEtlPayment
	GetSyEtlPayment(policyNumber string) []entity.SyEtlPayment
	TruncateTableSyEtlPayments() error
}

type syEtlPaymentService struct {
	paymentRepository repositories.LumpSumPaymentRepository
	repositories        repositories.SyEtlPaymentRepository
}

func NewSyEtlPaymentService(syEtlPaymentRepository repositories.SyEtlPaymentRepository) SyEtlPaymentService {
	return &syEtlPaymentService{
		repositories: syEtlPaymentRepository,
	}
}

func (service *syEtlPaymentService) CreateSyEtlPayment(syEtlPayment entity.SyEtlPayment) error {
	service.repositories.CreateSyEtlPayment(syEtlPayment)
	return nil
}

func (service *syEtlPaymentService) GetAllSyEtlPayments() []entity.SyEtlPayment {
	syEtlPayment := service.repositories.GetAllSyEtlPayments()
	return syEtlPayment
}

func (service *syEtlPaymentService) SyOdsMapEtlLatestPayment() []entity.SyEtlPayment {
	syEtlPayment := service.repositories.GetAllSyEtlPayments()

	for _, s := range syEtlPayment {
		if s.OdsEtlPayment.CollectionId != "" {
			// Populate Data to SY ETL Payment
			jsonData := new(entity.SyOdsEtlPayment)
			jsonData.CollectionId = s.OdsEtlPayment.CollectionId
			jsonData.PolicyNumber = s.PolicyNumber
			jsonData.ProposalNumber = s.ProposalNumber
			jsonData.PaymentMethodName = s.PaymentMethodName
			jsonData.SyTotalAmount = s.TotalPremium
			jsonData.OdsTotalAmount = s.OdsEtlPayment.TotalAmount
			jsonData.SyPaidDate = s.PaidDate
			jsonData.PolicyStatus = s.PolicyStatus
			jsonData.OdsFirstPaidDate = s.OdsEtlPayment.FirstPaymentDate
			jsonData.OdsPaidDate = s.OdsEtlPayment.PaymentDate

			if s.PaidDate.Truncate(24 * time.Hour).Before(s.OdsEtlPayment.PaymentDate.Truncate(24 * time.Hour)) {
				if s.PolicyStatus != "success" {
					jsonData.Status = "closed"
					jsonData.StatusDescription = "Policy is not in success status"
				} else if s.OdsEtlPayment.FirstPaymentDate == s.OdsEtlPayment.PaymentDate {
					jsonData.Status = "closed"
					jsonData.StatusDescription = "Stated as First Payment, use SY data for First Payment"
				} else if s.TotalPremium != s.OdsEtlPayment.TotalAmount {
					jsonData.Status = "conflict"
					jsonData.StatusDescription = "SY-ODS has different amount"
				} else {
					jsonData.Status = "queue"
					jsonData.StatusDescription = s.PaidDate.Format("2006-01-02") + " updated with " + s.OdsEtlPayment.PaymentDate.Format("2006-01-02")
				}
			} else if s.TotalPremium != s.OdsEtlPayment.TotalAmount {
				jsonData.Status = "conflict"
				jsonData.StatusDescription = "SY-ODS has different amount"
			} else {
				jsonData.Status = "closed"
				jsonData.StatusDescription = "SY-ODS no need to update"
			}

			jsonValue, _ := json.Marshal(jsonData)

			res, err := utils.CreateHttpRequest("POST", os.Getenv("MAIN_SCHEMES")+"://"+os.Getenv("MAIN_URL")+"/"+os.Getenv("API_PREFIX")+"syOdsEtl/payment/create", jsonValue)
			defer res.Body.Close()
			if err != nil {
				// utils.ErrorLogger(err)
			}
		}
	}
	return syEtlPayment
}

func (service *syEtlPaymentService) GetSyEtlPayment(policyNumber string) []entity.SyEtlPayment {
	syEtlPayment := service.repositories.GetSyEtlPayment(policyNumber)
	return syEtlPayment
}

func (service *syEtlPaymentService) TruncateTableSyEtlPayments() error {
	service.repositories.TruncateTableSyEtlPayments()
	return nil
}
