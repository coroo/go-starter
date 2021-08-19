package usecases

import (
	"encoding/json"
	"os"
	"github.com/coroo/go-starter/app/utils"
	entity "github.com/coroo/go-starter/app/entity"
	repositories "github.com/coroo/go-starter/app/repositories"
	// entity "github.com/Sequis-Digital-Channel/superyou-ods/entity/ods"
	// repositories "github.com/Sequis-Digital-Channel/superyou-ods/repositories/ods"
	// syrepositories "github.com/Sequis-Digital-Channel/superyou-ods/repositories/sy"

	// "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type CheckInvoice struct {
	ID string `json:"id"`
}

type LumpSumPaymentService interface {
	GetAllLumpSumPayments() []entity.LumpSumPayment
	OdsMapEtlLatestPayment() []entity.LumpSumPayment
	GetLumpSumPayment(policyNumber string) []entity.LumpSumPayment
}

type lumpSumPaymentService struct {
	repositories   repositories.LumpSumPaymentRepository
}

func NewLumpSumPaymentService(lumpSumPaymentRepository repositories.LumpSumPaymentRepository) LumpSumPaymentService {
	return &lumpSumPaymentService{
		repositories: lumpSumPaymentRepository,
	}
}

func (service *lumpSumPaymentService) GetAllLumpSumPayments() []entity.LumpSumPayment {
	lumpSumPayment := service.repositories.GetAllLumpSumPayments()
	return lumpSumPayment
}

func (service *lumpSumPaymentService) OdsMapEtlLatestPayment() []entity.LumpSumPayment {
	lumpSumPayment := service.repositories.GetAllLatestGroupLumpSumPayments()

	for _, s := range lumpSumPayment {
		// Populate Data to SY ETL Payment
		jsonData := new(entity.OdsEtlPayment)
		jsonData.CollectionId = s.CollectionId
		jsonData.PolicyNumber = s.PolicyNumber
		jsonData.ProposalNumber = s.ProposalNumber
		jsonData.PaymentMethod = s.PaymentMethod
		jsonData.TotalAmount = s.TotalAmount
		jsonData.FirstPaymentDate = s.FirstEffectiveDate
		jsonData.PaymentDate = s.EffectiveDate

		jsonValue, _ := json.Marshal(jsonData)
		res, err := utils.CreateHttpRequest("POST", os.Getenv("MAIN_SCHEMES")+"://"+os.Getenv("MAIN_URL")+"/"+os.Getenv("API_PREFIX")+"odsEtl/payment/create", jsonValue)
		defer res.Body.Close()
		if err != nil {
			// utils.ErrorLogger(err)
		}
	}
	return lumpSumPayment
}

func (service *lumpSumPaymentService) GetLumpSumPayment(policyNumber string) []entity.LumpSumPayment {
	return service.repositories.GetLumpSumPayment(policyNumber)
}
