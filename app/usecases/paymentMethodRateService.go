package usecases

import (
	entity "github.com/coroo/go-starter/app/entity"
	repositories "github.com/coroo/go-starter/app/repositories"

	// "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type PaymentMethodRateService interface {
	SavePaymentMethodRate(entity.PaymentMethodRate) (int, error)
	UpdatePaymentMethodRate(entity.PaymentMethodRate) error
	DeletePaymentMethodRate(entity.PaymentMethodRate) error
	GetAllPaymentMethodRates() []entity.PaymentMethodRate
	GetPaymentMethodRate(id string) []entity.PaymentMethodRate
}

type paymentMethodRateService struct {
	repositories repositories.PaymentMethodRateRepository
}

func NewPaymentMethodRateService(repository repositories.PaymentMethodRateRepository) PaymentMethodRateService {
	return &paymentMethodRateService{
		repositories: repository,
	}
}

func (usecases *paymentMethodRateService) GetAllPaymentMethodRates() []entity.PaymentMethodRate {
	return usecases.repositories.GetAllPaymentMethodRates()
}

func (usecases *paymentMethodRateService) GetPaymentMethodRate(id string) []entity.PaymentMethodRate {
	return usecases.repositories.GetPaymentMethodRate(id)
}

func (usecases *paymentMethodRateService) SavePaymentMethodRate(paymentMethodRate entity.PaymentMethodRate) (int, error) {
	return usecases.repositories.SavePaymentMethodRate(paymentMethodRate)
}

func (usecases *paymentMethodRateService) UpdatePaymentMethodRate(paymentMethodRate entity.PaymentMethodRate) error {
	return usecases.repositories.UpdatePaymentMethodRate(paymentMethodRate)
}

func (usecases *paymentMethodRateService) DeletePaymentMethodRate(paymentMethodRate entity.PaymentMethodRate) error {
	return usecases.repositories.DeletePaymentMethodRate(paymentMethodRate)
}