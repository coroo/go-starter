package usecases

import (
	"encoding/json"
	"os"
	entity "github.com/coroo/go-starter/app/entity"
	repositories "github.com/coroo/go-starter/app/repositories"
	"github.com/coroo/go-starter/app/utils"

	_ "github.com/go-sql-driver/mysql"
)

type CheckSyUserInvoice struct {
	ID string `json:"id"`
}

type SyUserInvoiceService interface {
	GetAllSyUserInvoices() []entity.SyUserInvoice
	SyMapEtlLatestPayment() []entity.SyUserInvoice
}

type syUserInvoiceService struct {
	repositories repositories.SyUserInvoiceRepository
}

func NewSyUserInvoiceService(syUserInvoiceRepository repositories.SyUserInvoiceRepository) SyUserInvoiceService {
	return &syUserInvoiceService{
		repositories: syUserInvoiceRepository,
	}
}

func (service *syUserInvoiceService) GetAllSyUserInvoices() []entity.SyUserInvoice {
	lumpSumPayment := service.repositories.GetAllPaidUserInvoices()
	return lumpSumPayment
}

// Migrate Super You DB
func (service *syUserInvoiceService) SyMapEtlLatestPayment() []entity.SyUserInvoice {
	lumpSumPayment := service.repositories.GetAllPaidUserInvoices()
	for _, s := range lumpSumPayment {
		// Populate Data to SY ETL Payment
		jsonData := new(entity.SyEtlPayment)
		jsonData.OdsPolicyNumber = s.PolicyNumber
		jsonData.ProposalNumber = s.ProposalNumber
		jsonData.PaymentMethodName = s.PaymentMethodName
		jsonData.TotalPremium = s.TotalPremium
		jsonData.PolicyStatus = s.SyUserPolicy.Status
		jsonData.PaidDate = s.PaidAt

		jsonValue, _ := json.Marshal(jsonData)

		_, err := utils.CreateHttpRequest("POST", os.Getenv("MAIN_SCHEMES")+"://"+os.Getenv("MAIN_URL")+"/"+os.Getenv("API_PREFIX")+"syEtl/payment/create", jsonValue)
		if err != nil {
			utils.ErrorLogger(err)
		}
	}
	return lumpSumPayment
}
