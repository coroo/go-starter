package deliveries

import (
	// "net/http"

	// "github.com/coroo/go-starter/models"
	usecases "github.com/coroo/go-starter/app/usecases"
	"github.com/coroo/go-starter/app/middlewares"
	
	"github.com/gin-gonic/gin"
)

type PaymentMethodRateController interface {
	GetPaymentMethodRate(ctx *gin.Context)
	// 	Save(ctx *gin.Context) error
}

type paymentMethodRateController struct {
	usecases usecases.PaymentMethodRateService
}

// var validate *validator.Validate

func NewPaymentMethodRateController(router *gin.Engine, apiPrefix string, paymentMethodRateService usecases.PaymentMethodRateService) {
	// validate = validator.New()
	// validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	handlerPaymentMethodRate := &paymentMethodRateController{
		usecases: paymentMethodRateService,
	}
	paymentMethodRatesGroup := router.Group(apiPrefix + "paymentMethodRate")
	{
		paymentMethodRatesGroup.GET("index", handlerPaymentMethodRate.GetAllPaymentMethodRate)
	}
}

func (deliveries *paymentMethodRateController) GetAllPaymentMethodRate(ctx *gin.Context) {
}