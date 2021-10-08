package deliveries

import (
	"net/http"

	// "github.com/coroo/go-starter/models"
	usecases "github.com/coroo/go-starter/app/usecases"
	entity "github.com/coroo/go-starter/app/entity"
	// "github.com/coroo/go-starter/app/middlewares"
	
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
		// paymentMethodsGroup.GET("index", handlerPaymentMethod.GetAllPaymentMethod)
		paymentMethodsGroup.GET("index", handlerPaymentMethod.PaymentMethodsIndex)
		paymentMethodsGroup.GET("detail/:id", handlerPaymentMethod.PaymentMethodsDetail)
		paymentMethodsGroup.POST("create", handlerPaymentMethod.PaymentMethodCreate)
		paymentMethodsGroup.PUT("update", handlerPaymentMethod.PaymentMethodUpdate)
		paymentMethodsGroup.DELETE("delete", handlerPaymentMethod.PaymentMethodDelete)
	}
}

// GetPaymentMethodsIndex godoc
// @Security basicAuth
// @Summary Show all existing PaymentMethods
// @Description Get all existing PaymentMethods
// @Tags PaymentMethods
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.PaymentMethod
// @Failure 401 {object} dto.Response
// @Router /paymentMethod/index [get]
func (deliveries *paymentMethodController) PaymentMethodsIndex(c *gin.Context) {
	paymentMethods := deliveries.usecases.GetAllPaymentMethods()
	c.JSON(http.StatusOK, gin.H{"data": paymentMethods})
}

// GetPaymentMethodsDetail godoc
// @Security basicAuth
// @Summary Show an existing PaymentMethods
// @Description Get detail the existing PaymentMethods
// @Tags PaymentMethods
// @Accept  json
// @Produce  json
// @Param  id path int true "Master Question ID"
// @Success 200 {array} entity.PaymentMethod
// @Failure 401 {object} dto.Response
// @Router /paymentMethod/detail/{id} [get]
func (deliveries *paymentMethodController) PaymentMethodsDetail(c *gin.Context) {
	paymentMethod := deliveries.usecases.GetPaymentMethod(c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"data": paymentMethod})
}

// CreatePaymentMethods godoc
// @Security basicAuth
// @Summary Create new PaymentMethods
// @Description Create a new PaymentMethods
// @Tags PaymentMethods
// @Accept  json
// @Produce  json
// @Param paymentMethod body entity.PaymentMethod true "Create paymentMethod"
// @Success 200 {object} entity.PaymentMethod
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /paymentMethod/create [post]
func (deliveries *paymentMethodController) PaymentMethodCreate(c *gin.Context) {
	var paymentMethodEntity entity.PaymentMethod
	c.ShouldBindJSON(&paymentMethodEntity)
	paymentMethodPK, err := deliveries.usecases.SavePaymentMethod(paymentMethodEntity)
	if(err!=nil){
		c.JSON(http.StatusConflict, err)
	} else {
		paymentMethodEntity.ID = paymentMethodPK
		c.JSON(http.StatusOK, paymentMethodEntity)
	}
}

// UpdatePaymentMethods godoc
// @Security basicAuth
// @Summary Update PaymentMethods
// @Description Update a PaymentMethods
// @Tags PaymentMethods
// @Accept  json
// @Produce  json
// @Param paymentMethod body entity.PaymentMethod true "Update paymentMethod"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /paymentMethod/update [put]
func (deliveries *paymentMethodController) PaymentMethodUpdate(c *gin.Context) {
	var paymentMethodEntity entity.PaymentMethod
	c.ShouldBindJSON(&paymentMethodEntity)
	paymentMethod := deliveries.usecases.UpdatePaymentMethod(paymentMethodEntity)
	c.JSON(http.StatusOK, paymentMethod)
}

// DeletePaymentMethods godoc
// @Security basicAuth
// @Summary Delete PaymentMethods
// @Description Delete a PaymentMethods
// @Tags PaymentMethods
// @Accept  json
// @Produce  json
// @Param paymentMethod body entity.PaymentMethod true "Delete paymentMethod"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /paymentMethod/delete [delete]
func (deliveries *paymentMethodController) PaymentMethodDelete(c *gin.Context) {
	var paymentMethodEntity entity.PaymentMethod
	c.ShouldBindJSON(&paymentMethodEntity)
	paymentMethod := deliveries.usecases.DeletePaymentMethod(paymentMethodEntity)
	c.JSON(http.StatusOK, paymentMethod)
}