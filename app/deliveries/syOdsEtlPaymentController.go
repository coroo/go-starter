package deliveries

import (
	"net/http"

	entity "github.com/coroo/go-starter/app/entity"
	usecases "github.com/coroo/go-starter/app/usecases"
	"github.com/coroo/go-starter/app/middlewares"

	"github.com/gin-gonic/gin"
)

type syOdsEtlPaymentController struct {
	usecases usecases.SyOdsEtlPaymentService
}

func NewSyOdsEtlPaymentController(router *gin.Engine, apiPrefix string, syOdsEtlPaymentService usecases.SyOdsEtlPaymentService) {
	handlerSyOdsEtlPayment := &syOdsEtlPaymentController{
		usecases: syOdsEtlPaymentService,
	}
	
	syOdsETLGroup := router.Group(apiPrefix + "syOdsEtl", middlewares.Auth)
	{
		syOdsETLGroup.GET("payment/index", handlerSyOdsEtlPayment.GetAllSyOdsEtlPayments)
		syOdsETLGroup.POST("payment/create", handlerSyOdsEtlPayment.CreateSyOdsEtlPayment)
		syOdsETLGroup.GET("payment/status/:status", handlerSyOdsEtlPayment.GetSyOdsEtlPaymentByStatus)
		syOdsETLGroup.GET("payment/daily-by-status/:status", handlerSyOdsEtlPayment.GetSyOdsEtlPaymentDailyByStatus)
		syOdsETLGroup.GET("payment/remove-before-map", handlerSyOdsEtlPayment.CancelOutstandingSyOdsEtlPayments)
		syOdsETLGroup.GET("payment/detail/:policyNumber", handlerSyOdsEtlPayment.GetSyOdsEtlPaymentByPolicyNumber)
	}
}

// GetAllSyOdsEtlPayments godoc
// @Param Authorization header string true "Bearer"
// @Security basicAuth
// @Summary List existing syOdsEtlPayments
// @Description Get all the existing syOdsEtlPayments
// @Tags syOdsEtlPayments
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.SyOdsEtlPayment
// @Failure 401 {object} dto.Response
// @Router /syOdsEtl/payment/index [get]
func (deliveries *syOdsEtlPaymentController) GetAllSyOdsEtlPayments(ctx *gin.Context){
	// return syOdsEtlPaymentService.GetAllSyOdsEtlPayments()
	syOdsEtlPayments :=  deliveries.usecases.GetAllSyOdsEtlPayments()
	ctx.JSON(http.StatusOK, gin.H{"data": syOdsEtlPayments})
}

// GetSyOdsEtlPaymentsDetail godoc
// @Param Authorization header string true "Bearer"
// @Security basicAuth
// @Summary Show an existing syOdsEtlPayments
// @Description Get detail the existing syOdsEtlPayments
// @Tags syOdsEtlPayments
// @Accept  json
// @Produce  json
// @Param  policyNumber path int true "Lump Sum Policy Number"
// @Success 200 {array} entity.SyOdsEtlPayment
// @Failure 401 {object} dto.Response
// @Router /syOdsEtl/payment/detail/{policyNumber} [get]
func (deliveries *syOdsEtlPaymentController) GetSyOdsEtlPaymentByPolicyNumber(ctx *gin.Context){
	// return syOdsEtlPaymentService.GetSyOdsEtlPaymentByPolicyNumber(ctx)
	syOdsEtlPayments :=  deliveries.usecases.GetSyOdsEtlPaymentByPolicyNumber(ctx.Param("policyNumber"))
	ctx.JSON(http.StatusOK, gin.H{"data": syOdsEtlPayments})
}

// SyOdsOdsMapEtlLatestPayment godoc
// @Param Authorization header string true "Bearer"
// @Security basicAuth
// @Summary List existing syOdsEtlPayments by Status
// @Description Get all the existing syOdsEtlPayments by Status
// @Tags syOdsEtlPayments
// @Accept  json
// @Produce  json
// @Param  status path string true "Available Status: queue, closed, cancel, success"
// @Success 200 {array} entity.SyOdsEtlPayment
// @Failure 401 {object} dto.Response
// @Router /syOdsEtl/payment/status/{status} [get]
func (deliveries *syOdsEtlPaymentController) GetSyOdsEtlPaymentByStatus(ctx *gin.Context){
	syOdsEtlPayments :=  deliveries.usecases.GetSyOdsEtlPaymentByStatus(ctx.Param("status"))
	ctx.JSON(http.StatusOK, gin.H{"data": syOdsEtlPayments})
	// return syOdsEtlPaymentService.GetSyOdsEtlPaymentByStatus(ctx)
}

// SyOdsEtlPaymentsDailyByStatus godoc
// @Param Authorization header string true "Bearer"
// @Security basicAuth
// @Summary List existing syOdsEtlPayments by Status and Updated by Today
// @Description Get all the existing syOdsEtlPayments by Status and Updated by Today
// @Tags syOdsEtlPayments
// @Accept  json
// @Produce  json
// @Param  status path string true "Available Status: queue, closed, cancel, success"
// @Success 200 {array} entity.SyOdsEtlPayment
// @Failure 401 {object} dto.Response
// @Router /syOdsEtl/payment/daily-by-status/{status} [get]
func (deliveries *syOdsEtlPaymentController) GetSyOdsEtlPaymentDailyByStatus(ctx *gin.Context){
	// return syOdsEtlPaymentService.GetSyOdsEtlPaymentDailyByStatus(ctx)
	syOdsEtlPayment :=  deliveries.usecases.GetSyOdsEtlPaymentDailyByStatus(ctx.Param("status"))
	ctx.JSON(http.StatusOK, gin.H{"data": syOdsEtlPayment})
}

// SyOdsEtlPaymentsCreate godoc
// @Param Authorization header string true "Bearer"
// @Security basicAuth
// @Summary Create new syOdsEtlPayments
// @Description Create a new syOdsEtlPayment
// @Tags syOdsEtlPayments
// @Accept  json
// @Produce  json
// @Param syOdsEtlPayment body entity.SyOdsEtlPayment true "Create syOdsEtlPayment"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /syOdsEtl/payment/create [post]
func (deliveries *syOdsEtlPaymentController) CreateSyOdsEtlPayment(ctx *gin.Context){
	var syOdsEtlPayment entity.SyOdsEtlPayment
	err := ctx.ShouldBindJSON(&syOdsEtlPayment)
	if err != nil {
		ctx.JSON(http.StatusConflict, err)
	}
	// err = validate.Struct(syOdsEtlPayment)
	// if err != nil {
	// 	return err
	// }
	deliveries.usecases.CreateSyOdsEtlPayment(syOdsEtlPayment)
	ctx.JSON(http.StatusOK, syOdsEtlPayment)
}

// RemoveSyOdsEtlBeforeMap godoc
// @Param Authorization header string true "Bearer"
// @Security basicAuth
// @Summary Truncate existing syOdsEtlPayments
// @Description Remove all the existing syOdsEtlPayments
// @Tags syOdsEtlPayments
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.SyOdsEtlPayment
// @Failure 401 {object} dto.Response
// @Router /syOdsEtl/payment/remove-before-map [get]
func (deliveries *syOdsEtlPaymentController) CancelOutstandingSyOdsEtlPayments(ctx *gin.Context){
	cancelSyOdsEtlPayment :=  deliveries.usecases.CancelOutstandingSyOdsEtlPayments()
	ctx.JSON(http.StatusOK, gin.H{"data": cancelSyOdsEtlPayment})
	// return syOdsEtlPaymentService.CancelOutstandingSyOdsEtlPayments()
}
