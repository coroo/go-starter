package deliveries

import (
	"fmt"
	"net/http"

	// "github.com/coroo/go-starter/models"
	usecases "github.com/coroo/go-starter/app/usecases"
	entity "github.com/coroo/go-starter/app/entity"
	// "github.com/coroo/go-starter/app/middlewares"
	
	"github.com/gin-gonic/gin"
)

type PaymentMethodLinkController interface {
	GetPaymentMethodLink(ctx *gin.Context)
	// 	Save(ctx *gin.Context) error
}

type paymentMethodLinkController struct {
	usecases usecases.PaymentMethodLinkService
}

// var validate *validator.Validate

func NewPaymentMethodLinkController(router *gin.Engine, apiPrefix string, paymentMethodLinkService usecases.PaymentMethodLinkService) {
	// validate = validator.New()
	// validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	handlerPaymentMethodLink := &paymentMethodLinkController{
		usecases: paymentMethodLinkService,
	}
	paymentMethodLinksGroup := router.Group(apiPrefix + "paymentMethodLink")
	{
		// paymentMethodLinksGroup.GET("index", handlerPaymentMethodLink.GetAllPaymentMethodLink)
		paymentMethodLinksGroup.GET("index", handlerPaymentMethodLink.PaymentMethodLinksIndex)
		paymentMethodLinksGroup.GET("detail/:id", handlerPaymentMethodLink.PaymentMethodLinksDetail)
		paymentMethodLinksGroup.GET("detail-by-code/:code", handlerPaymentMethodLink.PaymentMethodLinksDetailByCode)
		paymentMethodLinksGroup.GET("get-payment-redirect-link", handlerPaymentMethodLink.GetPaymentRedirectLink)
		paymentMethodLinksGroup.POST("create", handlerPaymentMethodLink.PaymentMethodLinkCreate)
		paymentMethodLinksGroup.PUT("update", handlerPaymentMethodLink.PaymentMethodLinkUpdate)
		paymentMethodLinksGroup.DELETE("delete", handlerPaymentMethodLink.PaymentMethodLinkDelete)
	}
}

// GetPaymentMethodLinksIndex godoc
// @Security basicAuth
// @Summary Show all existing PaymentMethodLinks
// @Description Get all existing PaymentMethodLinks
// @Tags PaymentMethodLinks
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.PaymentMethodLink
// @Failure 401 {object} dto.Response
// @Router /paymentMethodLink/index [get]
func (deliveries *paymentMethodLinkController) PaymentMethodLinksIndex(ctx *gin.Context) {
	paymentMethodLinks := deliveries.usecases.GetAllPaymentMethodLinks()
	ctx.JSON(http.StatusOK, gin.H{"data": paymentMethodLinks})
}

// GetPaymentMethodLinksDetail godoc
// @Security basicAuth
// @Summary Show an existing PaymentMethodLinks
// @Description Get detail the existing PaymentMethodLinks
// @Tags PaymentMethodLinks
// @Accept  json
// @Produce  json
// @Param  id path int true "Master Question ID"
// @Success 200 {array} entity.PaymentMethodLink
// @Failure 401 {object} dto.Response
// @Router /paymentMethodLink/detail/{id} [get]
func (deliveries *paymentMethodLinkController) PaymentMethodLinksDetail(c *gin.Context) {
	paymentMethodLink := deliveries.usecases.GetPaymentMethodLink(c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"data": paymentMethodLink})
}

// GetPaymentMethodLinksDetailByCode godoc
// @Security basicAuth
// @Summary Show an existing PaymentMethodLinks
// @Description Get detail the existing PaymentMethodLinks
// @Tags PaymentMethodLinks
// @Accept  json
// @Produce  json
// @Param  id path int true "Master Question ID"
// @Success 200 {array} entity.PaymentMethodLink
// @Failure 401 {object} dto.Response
// @Router /paymentMethodLink/detail/{id} [get]
func (deliveries *paymentMethodLinkController) PaymentMethodLinksDetailByCode(c *gin.Context) {
	paymentMethodLink := deliveries.usecases.GetPaymentMethodLinkByCode(c.Param("code"))
	c.JSON(http.StatusOK, gin.H{"data": paymentMethodLink})
}

// CreatePaymentMethodLinks godoc
// @Security basicAuth
// @Summary Create new PaymentMethodLinks
// @Description Create a new PaymentMethodLinks
// @Tags PaymentMethodLinks
// @Accept  json
// @Produce  json
// @Param paymentMethodLink body entity.PaymentMethodLink true "Create paymentMethodLink"
// @Success 200 {object} entity.PaymentMethodLink
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /paymentMethodLink/create [post]
func (deliveries *paymentMethodLinkController) PaymentMethodLinkCreate(c *gin.Context) {
	var paymentMethodLinkEntity entity.PaymentMethodLink
	c.ShouldBindJSON(&paymentMethodLinkEntity)
	paymentMethodLinkPK, err := deliveries.usecases.SavePaymentMethodLink(paymentMethodLinkEntity)
	if(err!=nil){
		c.JSON(http.StatusConflict, err)
	} else {
		paymentMethodLinkEntity.ID = paymentMethodLinkPK
		c.JSON(http.StatusOK, paymentMethodLinkEntity)
	}
}

// UpdatePaymentMethodLinks godoc
// @Security basicAuth
// @Summary Update PaymentMethodLinks
// @Description Update a PaymentMethodLinks
// @Tags PaymentMethodLinks
// @Accept  json
// @Produce  json
// @Param paymentMethodLink body entity.PaymentMethodLink true "Update paymentMethodLink"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /paymentMethodLink/update [put]
func (deliveries *paymentMethodLinkController) PaymentMethodLinkUpdate(c *gin.Context) {
	var paymentMethodLinkEntity entity.PaymentMethodLink
	c.ShouldBindJSON(&paymentMethodLinkEntity)
	paymentMethodLink := deliveries.usecases.UpdatePaymentMethodLink(paymentMethodLinkEntity)
	c.JSON(http.StatusOK, paymentMethodLink)
}

// DeletePaymentMethodLinks godoc
// @Security basicAuth
// @Summary Delete PaymentMethodLinks
// @Description Delete a PaymentMethodLinks
// @Tags PaymentMethodLinks
// @Accept  json
// @Produce  json
// @Param paymentMethodLink body entity.PaymentMethodLink true "Delete paymentMethodLink"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /paymentMethodLink/delete [delete]
func (deliveries *paymentMethodLinkController) PaymentMethodLinkDelete(c *gin.Context) {
	var paymentMethodLinkEntity entity.PaymentMethodLink
	c.ShouldBindJSON(&paymentMethodLinkEntity)
	paymentMethodLink := deliveries.usecases.DeletePaymentMethodLink(paymentMethodLinkEntity)
	c.JSON(http.StatusOK, paymentMethodLink)
}

// GetPaymentRedirectLink godoc
// @Security basicAuth
// @Summary Get Payment link for redirection
// @Description Get Payment link for redirection to the next page of flow
// @Tags PaymentMethodLinks
// @Accept  json
// @Produce  json
// @Param policy_group_number query string true "Policy Group Number"
// @Param payment_method_code query string true "Payment Method Code"
// @Param process_type query string true "Process Type"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /paymentMethodLink/get-payment-redirect-link [get]
func (deliveries *paymentMethodLinkController) GetPaymentRedirectLink(ctx *gin.Context) {
	var paymentMethodLink entity.PaymentMethodLink
	paymentMethodLink = deliveries.usecases.GetPaymentMethodLinkByCodeAndProcessType(ctx.Query("payment_method_code"), ctx.Query("process_type"))
	fmt.Println("process 1")
	go deliveries.usecases.ProcessPaymentForSettlement(ctx.Query("proposal_number"), paymentMethodLink)
	fmt.Println("process 3")
	ctx.JSON(http.StatusOK, paymentMethodLink)
}