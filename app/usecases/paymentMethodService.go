package usecases

import (
	"fmt"
	"strconv"
	entity "github.com/coroo/go-starter/app/entity"
	utils "github.com/coroo/go-starter/app/utils"
	repositories "github.com/coroo/go-starter/app/repositories"
	// "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type PaymentMethodService interface {
	SavePaymentMethod(entity.PaymentMethod) (int, error)
	UpdatePaymentMethod(entity.PaymentMethod) error
	DeletePaymentMethod(entity.PaymentMethod) error
	GetAllPaymentMethods(total_premium string) []entity.PaymentMethodWithPremium
	GetPaymentMethod(id string) []entity.PaymentMethod
	GetPaymentMethodByCode(code string) []entity.PaymentMethod
}

type paymentMethodService struct {
	repositories repositories.PaymentMethodRepository
}

func NewPaymentMethodService(repository repositories.PaymentMethodRepository) PaymentMethodService {
	return &paymentMethodService{
		repositories: repository,
	}
}

func (usecases *paymentMethodService) GetAllPaymentMethods(total_premium string) []entity.PaymentMethodWithPremium {
	// get all payment method + rates
	allPaymentMethod := usecases.repositories.GetAllPaymentMethods()
	AllPaymentMethodWithTotalPremium := []entity.PaymentMethodWithPremium{}
	// change total premium query from string (A) into integer (i)
	tempPremium, err := strconv.Atoi(total_premium)
	if err != nil {
		fmt.Println(err)
	}
	// loop all payment to add the total payment and fee
	for _, value := range allPaymentMethod{
		// adding data to schema
		data := entity.PaymentMethodWithPremium{}
		data.ID = value.ID
		data.Code = value.Code
		data.InitPaymentCode = value.InitPaymentCode
		data.RenewalPaymentCode = value.RenewalPaymentCode
		data.FastpayCode = value.FastpayCode
		data.BankCode = value.BankCode
		data.Name = value.Name
		data.PaymentLogo = value.PaymentLogo
		data.Status = value.Status
		data.Spec = value.Spec
		data.CreatedAt = value.CreatedAt
		data.UpdatedAt = value.UpdatedAt
		// total payment calculation
		if (total_premium != ""){
			for _, valueRate := range value.PaymentMethodRate {
				// checking transaction range
				if valueRate.MinTransaction <= tempPremium && valueRate.MaxTransaction >= tempPremium {
					// total premium
					data.TotalPremium = tempPremium
					// calculate total fee
					data.Fee = utils.EvaluateStringToFormula(valueRate.FormulaFee, tempPremium)
					// total payment
					data.TotalPayment = data.TotalPremium + data.Fee
				}
			}
		}
		AllPaymentMethodWithTotalPremium = append(AllPaymentMethodWithTotalPremium, data) 
	}
	return AllPaymentMethodWithTotalPremium
}

func (usecases *paymentMethodService) GetPaymentMethod(id string) []entity.PaymentMethod {
	return usecases.repositories.GetPaymentMethod(id)
}

func (usecases *paymentMethodService) GetPaymentMethodByCode(code string) []entity.PaymentMethod {
	return usecases.repositories.GetPaymentMethodByCode(code)
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