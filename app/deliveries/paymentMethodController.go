package deliveries

import (
	// "net/http"

	// "github.com/coroo/go-starter/models"
	usecases "github.com/coroo/go-starter/app/usecases"
	"github.com/coroo/go-starter/app/middlewares"
	
	"github.com/gin-gonic/gin"
)

type PaymentMethodController interface {
	GetPaymentMethod(ctx *gin.Context)
	// 	Save(ctx *gin.Context) error
}

type paymentMethodController struct {
	usecases usecases.PaymentMethodService
}

// var validate *validator.Validate

func NewPaymentMethodController(router *gin.Engine, apiPrefix string, paymentMethodService usecases.PaymentMethodService) {
	// validate = validator.New()
	// validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	handlerPaymentMethod := &paymentMethodController{
		usecases: paymentMethodService,
	}
	paymentMethodsGroup := router.Group(apiPrefix + "paymentMethod")
	{
		paymentMethodsGroup.GET("index", handlerPaymentMethod.GetAllPaymentMethod)
	}
}

func (deliveries *paymentMethodController) GetAllPaymentMethod(ctx *gin.Context) {
}