package usecases

import (
	entity "github.com/coroo/go-starter/app/entity"
	repositories "github.com/coroo/go-starter/app/repositories"

	// "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type OdsEtlPaymentService interface {
	CreateOdsEtlPayment(entity.OdsEtlPayment) error
	GetAllOdsEtlPayments() []entity.OdsEtlPayment
	GetOdsEtlPayment(id string) []entity.OdsEtlPayment
	TruncateTableOdsEtlPayments() error
}

type odsEtlPaymentService struct {
	paymentRepository repositories.LumpSumPaymentRepository
	repository        repositories.OdsEtlPaymentRepository
}

func NewOdsEtlPaymentService(odsEtlPaymentRepository repositories.OdsEtlPaymentRepository) OdsEtlPaymentService {
	// return &odsEtlPaymentService{
	// 	models: models,
	// }
	return &odsEtlPaymentService{
		repository: odsEtlPaymentRepository,
	}
}

// func (service *odsEtlPaymentService) GetAllOdsEtlPayments() []entity.OdsEtlPayment {
// 	return service.repository.GetAllOdsEtlPayments()
// }

func (service *odsEtlPaymentService) CreateOdsEtlPayment(odsEtlPayment entity.OdsEtlPayment) error {
	mappingData := odsEtlPayment
	if mappingData.PaymentMethod == "Credit Card" {
		mappingData.PaymentMethod = "Visa Master"
	}
	service.repository.CreateOdsEtlPayment(mappingData)
	return nil
}

func (service *odsEtlPaymentService) GetAllOdsEtlPayments() []entity.OdsEtlPayment {
	odsEtlPayment := service.repository.GetAllOdsEtlPayments()
	return odsEtlPayment
}

func (service *odsEtlPaymentService) GetOdsEtlPayment(id string) []entity.OdsEtlPayment {
	odsEtlPayment := service.repository.GetOdsEtlPayment(id)
	return odsEtlPayment
}

func (service *odsEtlPaymentService) TruncateTableOdsEtlPayments() error {
	service.repository.TruncateTableOdsEtlPayments()
	return nil
}

// func (service *odsEtlPaymentService) GetAllOdsEtlPayments(ctx *gin.Context) ([]entity.OdsEtlPayment, error) {
// 	db, _ := ctx.Get("db")
// 	conn := db.(sql.DB)
// 	odsEtlPayments, err := models.GetAllOdsEtlPayments(&conn)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return odsEtlPayments, err
// }
