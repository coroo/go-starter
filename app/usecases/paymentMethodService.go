package usecases

import (
	entity "github.com/coroo/go-starter/app/entity"
	repositories "github.com/coroo/go-starter/app/repositories"

	// "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type PaymentMethodService interface {
	SavePaymentMethod(entity.PaymentMethod) (int, error)
	UpdatePaymentMethod(entity.PaymentMethod) error
	DeletePaymentMethod(entity.PaymentMethod) error
	GetAllPaymentMethods() []entity.PaymentMethod
	GetPaymentMethod(id string) []entity.PaymentMethod
}

type paymentMethodService struct {
	repositories repositories.PaymentMethodRepository
}

func NewPaymentMethodService(repository repositories.PaymentMethodRepository) PaymentMethodService {
	return &paymentMethodService{
		repositories: repository,
	}
}

func (usecases *paymentMethodService) GetAllPaymentMethods() []entity.PaymentMethod {
	return usecases.repositories.GetAllPaymentMethods()
}

func (usecases *paymentMethodService) GetPaymentMethod(id string) []entity.PaymentMethod {
	return usecases.repositories.GetPaymentMethod(id)
}

func (usecases *paymentMethodService) SavePaymentMethod(paymentMethod entity.PaymentMethod) (int, error) {
	return usecases.repositories.SavePaymentMethod(paymentMethod)
}

func (usecases *paymentMethodService) UpdatePaymentMethod(paymentMethod entity.PaymentMethod) error {
	return usecases.repositories.UpdatePaymentMethod(paymentMethod)
}

func (usecases *paymentMethodService) DeletePaymentMethod(paymentMethod entity.PaymentMethod) error {
	return usecases.repositories.DeletePaymentMethod(paymentMethod)
}