package usecases

import (
	repositories "github.com/coroo/go-starter/app/repositories"

	// "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type PaymentMethodService interface {
	CreatePaymentMethod()
}

type paymentMethodService struct {
	repositories repositories.PaymentMethodRepository
}

func NewPaymentMethodService(repository repositories.PaymentMethodRepository) PaymentMethodService {
	return &paymentMethodService{
		repositories: repository,
	}
}

func (usecases *paymentMethodService) CreatePaymentMethod(){

}