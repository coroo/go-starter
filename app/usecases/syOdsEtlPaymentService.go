package usecases

import (
	entity "github.com/coroo/go-starter/app/entity"
	repositories "github.com/coroo/go-starter/app/repositories"
	
	_ "github.com/go-sql-driver/mysql"
)

type SyOdsEtlPaymentService interface {
	CreateSyOdsEtlPayment(entity.SyOdsEtlPayment) error
	GetAllSyOdsEtlPayments() []entity.SyOdsEtlPayment
	GetSyOdsEtlPaymentByPolicyNumber(policyNumber string) []entity.SyOdsEtlPayment
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

func (service *syOdsEtlPaymentService) CreateSyOdsEtlPayment(syOdsEtlPayment entity.SyOdsEtlPayment) error {
	service.repositories.CreateSyOdsEtlPayment(syOdsEtlPayment)
	return nil
}

func (service *syOdsEtlPaymentService) GetAllSyOdsEtlPayments() []entity.SyOdsEtlPayment {
	syOdsEtlPayment := service.repositories.GetAllSyOdsEtlPayments()
	return syOdsEtlPayment
}

func (service *syOdsEtlPaymentService) GetSyOdsEtlPaymentByPolicyNumber(policyNumber string) []entity.SyOdsEtlPayment {
	syOdsEtlPayment := service.repositories.GetSyOdsEtlPaymentByPolicyNumber(policyNumber)
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
