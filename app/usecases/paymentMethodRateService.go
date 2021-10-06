package usecases

import (
	repositories "github.com/coroo/go-starter/app/repositories"

	// "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type PaymentMethodRateService interface {
	CreatePaymentMethodRate()
}

type paymentMethodRateService struct {
	repositories repositories.PaymentMethodRateRepository
}

func NewPaymentMethodRateService(repository repositories.PaymentMethodRateRepository) PaymentMethodRateService {
	return &paymentMethodRateService{
		repositories: repository,
	}
}

func (usecases *paymentMethodRateService) CreatePaymentMethodRate(){

}