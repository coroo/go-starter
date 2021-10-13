package usecases

import (
	entity "github.com/coroo/go-starter/app/entity"
	repositories "github.com/coroo/go-starter/app/repositories"
	// "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type PaymentMethodLinkService interface {
	SavePaymentMethodLink(entity.PaymentMethodLink) (int, error)
	UpdatePaymentMethodLink(entity.PaymentMethodLink) error
	DeletePaymentMethodLink(entity.PaymentMethodLink) error
	GetAllPaymentMethodLinks() []entity.PaymentMethodLink
	GetPaymentMethodLink(id string) []entity.PaymentMethodLink
	GetPaymentMethodLinkByCode(code string) []entity.PaymentMethodLink
}

type paymentMethodLinkService struct {
	repositories repositories.PaymentMethodLinkRepository
}

func NewPaymentMethodLinkService(repository repositories.PaymentMethodLinkRepository) PaymentMethodLinkService {
	return &paymentMethodLinkService{
		repositories: repository,
	}
}

func (usecases *paymentMethodLinkService) GetAllPaymentMethodLinks() []entity.PaymentMethodLink {
	return usecases.repositories.GetAllPaymentMethodLinks()
}

func (usecases *paymentMethodLinkService) GetPaymentMethodLink(id string) []entity.PaymentMethodLink {
	return usecases.repositories.GetPaymentMethodLink(id)
}

func (usecases *paymentMethodLinkService) GetPaymentMethodLinkByCode(code string) []entity.PaymentMethodLink {
	return usecases.repositories.GetPaymentMethodLinkByCode(code)
}

func (usecases *paymentMethodLinkService) SavePaymentMethodLink(paymentMethodLink entity.PaymentMethodLink) (int, error) {
	return usecases.repositories.SavePaymentMethodLink(paymentMethodLink)
}

func (usecases *paymentMethodLinkService) UpdatePaymentMethodLink(paymentMethodLink entity.PaymentMethodLink) error {
	return usecases.repositories.UpdatePaymentMethodLink(paymentMethodLink)
}

func (usecases *paymentMethodLinkService) DeletePaymentMethodLink(paymentMethodLink entity.PaymentMethodLink) error {
	return usecases.repositories.DeletePaymentMethodLink(paymentMethodLink)
}