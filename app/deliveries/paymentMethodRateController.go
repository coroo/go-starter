package deliveries

import (
	"net/http"

	usecases "github.com/coroo/go-starter/app/usecases"
	entity "github.com/coroo/go-starter/app/entity"
	// "github.com/coroo/go-starter/app/middlewares"
	
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
		paymentMethodRatesGroup.GET("index", handlerPaymentMethodRate.PaymentMethodRatesIndex)
		paymentMethodRatesGroup.GET("detail/:id", handlerPaymentMethodRate.PaymentMethodRatesDetail)
		paymentMethodRatesGroup.POST("create", handlerPaymentMethodRate.PaymentMethodRateCreate)
		paymentMethodRatesGroup.PUT("update", handlerPaymentMethodRate.PaymentMethodRateUpdate)
		paymentMethodRatesGroup.DELETE("delete", handlerPaymentMethodRate.PaymentMethodRateDelete)
	}
}

// GetPaymentMethodRatesIndex godoc
// @Security basicAuth
// @Summary Show all existing PaymentMethodRates
// @Description Get all existing PaymentMethodRates
// @Tags PaymentMethodRates
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.PaymentMethodRate
// @Failure 401 {object} dto.Response
// @Router /paymentMethodRate/index [get]
func (deliveries *paymentMethodRateController) PaymentMethodRatesIndex(c *gin.Context) {
	paymentMethodRates := deliveries.usecases.GetAllPaymentMethodRates()
	c.JSON(http.StatusOK, gin.H{"data": paymentMethodRates})
}

// GetPaymentMethodRatesDetail godoc
// @Security basicAuth
// @Summary Show an existing PaymentMethodRates
// @Description Get detail the existing PaymentMethodRates
// @Tags PaymentMethodRates
// @Accept  json
// @Produce  json
// @Param  id path int true "Master Question ID"
// @Success 200 {array} entity.PaymentMethodRate
// @Failure 401 {object} dto.Response
// @Router /paymentMethodRate/detail/{id} [get]
func (deliveries *paymentMethodRateController) PaymentMethodRatesDetail(c *gin.Context) {
	paymentMethodRate := deliveries.usecases.GetPaymentMethodRate(c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"data": paymentMethodRate})
}

// CreatePaymentMethodRates godoc
// @Security basicAuth
// @Summary Create new PaymentMethodRates
// @Description Create a new PaymentMethodRates
// @Tags PaymentMethodRates
// @Accept  json
// @Produce  json
// @Param paymentMethodRate body entity.PaymentMethodRate true "Create paymentMethodRate"
// @Success 200 {object} entity.PaymentMethodRate
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /paymentMethodRate/create [post]
func (deliveries *paymentMethodRateController) PaymentMethodRateCreate(c *gin.Context) {
	var paymentMethodRateEntity entity.PaymentMethodRate
	c.ShouldBindJSON(&paymentMethodRateEntity)
	paymentMethodRatePK, err := deliveries.usecases.SavePaymentMethodRate(paymentMethodRateEntity)
	if(err!=nil){
		c.JSON(http.StatusConflict, err)
	} else {
		paymentMethodRateEntity.ID = paymentMethodRatePK
		c.JSON(http.StatusOK, paymentMethodRateEntity)
	}
}

// UpdatePaymentMethodRates godoc
// @Security basicAuth
// @Summary Update PaymentMethodRates
// @Description Update a PaymentMethodRates
// @Tags PaymentMethodRates
// @Accept  json
// @Produce  json
// @Param paymentMethodRate body entity.PaymentMethodRate true "Update paymentMethodRate"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /paymentMethodRate/update [put]
func (deliveries *paymentMethodRateController) PaymentMethodRateUpdate(c *gin.Context) {
	var paymentMethodRateEntity entity.PaymentMethodRate
	c.ShouldBindJSON(&paymentMethodRateEntity)
	paymentMethodRate := deliveries.usecases.UpdatePaymentMethodRate(paymentMethodRateEntity)
	c.JSON(http.StatusOK, paymentMethodRate)
}

// DeletePaymentMethodRates godoc
// @Security basicAuth
// @Summary Delete PaymentMethodRates
// @Description Delete a PaymentMethodRates
// @Tags PaymentMethodRates
// @Accept  json
// @Produce  json
// @Param paymentMethodRate body entity.PaymentMethodRate true "Delete paymentMethodRate"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /paymentMethodRate/delete [delete]
func (deliveries *paymentMethodRateController) PaymentMethodRateDelete(c *gin.Context) {
	var paymentMethodRateEntity entity.PaymentMethodRate
	c.ShouldBindJSON(&paymentMethodRateEntity)
	paymentMethodRate := deliveries.usecases.DeletePaymentMethodRate(paymentMethodRateEntity)
	c.JSON(http.StatusOK, paymentMethodRate)
}