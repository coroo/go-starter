package usecases

import (
	entity "github.com/coroo/go-starter/app/entity"
	repositories "github.com/coroo/go-starter/app/repositories"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type SyOdsEtlPaymentService interface {
	CreateSyOdsEtlPayment(entity.SyOdsEtlPayment) error
	GetAllSyOdsEtlPayments() []entity.SyOdsEtlPayment
	GetSyOdsEtlPaymentByPolicyNumber(ctx *gin.Context) []entity.SyOdsEtlPayment
	GetSyOdsEtlPaymentByStatus(status string) []entity.SyOdsEtlPayment
	GetSyOdsEtlPaymentDailyByStatus(status string) []entity.SyOdsEtlPayment
	CancelOutstandingSyOdsEtlPayments() []entity.SyOdsEtlPayment
}

type syOdsEtlPaymentService struct {
	paymentRepository repositories.LumpSumPaymentRepository
	repositories        repositories.SyOdsEtlPaymentRepository
}

func NewSyOdsEtlPaymentService(syOdsEtlPaymentRepository repositories.SyOdsEtlPaymentRepository) SyOdsEtlPaymentService {
	// return &syOdsEtlPaymentService{
	// 	models: models,
	// }
	return &syOdsEtlPaymentService{
		repositories: syOdsEtlPaymentRepository,
	}
}

// func (service *syOdsEtlPaymentService) GetAllSyOdsEtlPayments() []entity.SyOdsEtlPayment {
// 	return service.repositories.GetAllSyOdsEtlPayments()
// }

func (service *syOdsEtlPaymentService) CreateSyOdsEtlPayment(syOdsEtlPayment entity.SyOdsEtlPayment) error {
	service.repositories.CreateSyOdsEtlPayment(syOdsEtlPayment)
	return nil
}

func (service *syOdsEtlPaymentService) GetAllSyOdsEtlPayments() []entity.SyOdsEtlPayment {
	syOdsEtlPayment := service.repositories.GetAllSyOdsEtlPayments()
	return syOdsEtlPayment
}

func (service *syOdsEtlPaymentService) GetSyOdsEtlPaymentByPolicyNumber(ctx *gin.Context) []entity.SyOdsEtlPayment {
	syOdsEtlPayment := service.repositories.GetSyOdsEtlPaymentByPolicyNumber(ctx)
	return syOdsEtlPayment
}

func (service *syOdsEtlPaymentService) GetSyOdsEtlPaymentByStatus(status string) []entity.SyOdsEtlPayment {
	syOdsEtlPayment := service.repositories.GetSyOdsEtlPaymentByStatus(status)
	return syOdsEtlPayment
}

func (service *syOdsEtlPaymentService) GetSyOdsEtlPaymentDailyByStatus(status string) []entity.SyOdsEtlPayment {
	syOdsEtlPayment := service.repositories.GetSyOdsEtlPaymentDailyByStatus(status)
	return syOdsEtlPayment
}

func (service *syOdsEtlPaymentService) CancelOutstandingSyOdsEtlPayments() []entity.SyOdsEtlPayment {
	syOdsEtlPayment := service.repositories.CancelOutstandingSyOdsEtlPayments()
	return syOdsEtlPayment
}

// func (service *syOdsEtlPaymentService) GetAllSyOdsEtlPayments(ctx *gin.Context) ([]entity.SyOdsEtlPayment, error) {
// 	db, _ := ctx.Get("db")
// 	conn := db.(sql.DB)
// 	syOdsEtlPayments, err := models.GetAllSyOdsEtlPayments(&conn)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return syOdsEtlPayments, err
// }
