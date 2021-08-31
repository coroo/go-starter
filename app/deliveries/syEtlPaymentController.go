package deliveries

import (
	"net/http"
	// "github.com/Sequis-Digital-Channel/superyou-ods/models"

	entity "github.com/coroo/go-starter/app/entity"
	usecases "github.com/coroo/go-starter/app/usecases"
	"github.com/coroo/go-starter/app/middlewares"

	"github.com/gin-gonic/gin"
)

type syEtlPaymentController struct {
	usecases usecases.SyEtlPaymentService
}

func NewSyEtlPaymentController(router *gin.Engine, apiPrefix string, syEtlPaymentService usecases.SyEtlPaymentService) {
	handlerSyEtlPayment := &syEtlPaymentController{
		usecases: syEtlPaymentService,
	}
	syETLGroup := router.Group(apiPrefix + "syEtl", middlewares.Auth)
	{
		syETLGroup.GET("payment/index", handlerSyEtlPayment.GetAllSyEtlPayments)
		syETLGroup.GET("payment/map-etl-payment", handlerSyEtlPayment.SyOdsMapEtlLatestPayment)
		syETLGroup.GET("payment/detail/:policyNumber", handlerSyEtlPayment.GetSyEtlPayment)
		syETLGroup.POST("payment/create", handlerSyEtlPayment.CreateSyEtlPayment)
		syETLGroup.GET("payment/remove-before-map", handlerSyEtlPayment.TruncateTableSyEtlPayments)
	}
}

// type SyEtlPaymentController interface {
// 	Save(ctx *gin.Context) error
// 	GetAllSyEtlPayments() 
// 	SyOdsMapEtlLatestPayment() 
// 	GetSyEtlPayment(ctx *gin.Context) 
// 	TruncateTableSyEtlPayments() error
// }

// type syEtlPaymentController struct {
// 	service service.SyEtlPaymentService
// }

// // var validate *validator.Validate

// func NewSyEtlPayment(service service.SyEtlPaymentService) SyEtlPaymentController {
// 	// validate = validator.New()
// 	// validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
// 	return &syEtlPaymentController{
// 		service: service,
// 	}
// }

// GetAllSyEtlPayments godoc
// @Param Authorization header string true "Bearer"
// @Security basicAuth
// @Summary List existing syEtlPayments
// @Description Get all the existing syEtlPayments
// @Tags syEtlPayments
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.SyEtlPayment
// @Failure 401 {object} dto.Response
// @Router /syEtl/payment/index [get]
func (deliveries *syEtlPaymentController) GetAllSyEtlPayments(ctx *gin.Context){
	syEtlPayments :=  deliveries.usecases.GetAllSyEtlPayments()
	ctx.JSON(http.StatusOK, gin.H{"data": syEtlPayments})
}

// SyOdsMapEtlLatestPayment godoc
// @Param Authorization header string true "Bearer"
// @Security basicAuth
// @Summary Map existing syEtlPayments to syOdsEtlPayments
// @Description Map all the existing syEtlPayments to syOdsEtlPayments
// @Tags syEtlPayments
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.SyEtlPayment
// @Failure 401 {object} dto.Response
// @Router /syEtl/payment/map-etl-payment [get]
func (deliveries *syEtlPaymentController) SyOdsMapEtlLatestPayment(ctx *gin.Context){
	latestSyEtlPayments :=  deliveries.usecases.SyOdsMapEtlLatestPayment()
	ctx.JSON(http.StatusOK, gin.H{"data": latestSyEtlPayments})
}

// GetSyEtlPaymentsDetail godoc
// @Param Authorization header string true "Bearer"
// @Security basicAuth
// @Summary Show an existing syEtlPayments
// @Description Get detail the existing syEtlPayments
// @Tags syEtlPayments
// @Accept  json
// @Produce  json
// @Param  policyNumber path int true "Lump Sum Policy Number"
// @Success 200 {array} entity.SyEtlPayment
// @Failure 401 {object} dto.Response
// @Router /syEtl/payment/detail/{policyNumber} [get]
func (deliveries *syEtlPaymentController) GetSyEtlPayment(ctx *gin.Context){
	syEtlPayment :=  deliveries.usecases.GetSyEtlPayment(ctx.Param("policyNumber"))
	ctx.JSON(http.StatusOK, gin.H{"data": syEtlPayment})
}

// SyEtlPaymentsCreate godoc
// @Param Authorization header string true "Bearer"
// @Security basicAuth
// @Summary Create new syEtlPayments
// @Description Create a new syEtlPayment
// @Tags syEtlPayments
// @Accept  json
// @Produce  json
// @Param syEtlPayment body entity.SyEtlPayment true "Create syEtlPayment"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /syEtl/payment/create [post]
func (deliveries *syEtlPaymentController) CreateSyEtlPayment(ctx *gin.Context){
	var syEtlPayment entity.SyEtlPayment
	err := ctx.ShouldBindJSON(&syEtlPayment)
	if err != nil {
		ctx.JSON(http.StatusConflict, err)
	}
	// err = validate.Struct(syEtlPayment)
	// if err != nil {
	// 	return err
	// }
	deliveries.usecases.CreateSyEtlPayment(syEtlPayment)
	ctx.JSON(http.StatusOK, syEtlPayment)
}

// RemoveSyEtlBeforeMap godoc
// @Param Authorization header string true "Bearer"
// @Security basicAuth
// @Summary Truncate existing syEtlPayments
// @Description Remove all the existing syEtlPayments
// @Tags syEtlPayments
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.SyEtlPayment
// @Failure 401 {object} dto.Response
// @Router /syEtl/payment/remove-before-map [get]
func (deliveries *syEtlPaymentController) TruncateTableSyEtlPayments(ctx *gin.Context){
	truncateSyEtlPayment :=  deliveries.usecases.TruncateTableSyEtlPayments()
	ctx.JSON(http.StatusOK, gin.H{"data": truncateSyEtlPayment})
}
