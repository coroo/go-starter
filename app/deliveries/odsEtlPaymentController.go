package deliveries

import (
	"net/http"

	entity "github.com/coroo/go-starter/app/entity"
	usecases "github.com/coroo/go-starter/app/usecases"
	"github.com/coroo/go-starter/app/middlewares"

	"github.com/gin-gonic/gin"
)

type odsEtlPaymentController struct {
	usecases usecases.OdsEtlPaymentService
}

func NewOdsEtlPaymentController(router *gin.Engine, apiPrefix string, odsEtlPaymentService usecases.OdsEtlPaymentService) {
	handlerUser := &odsEtlPaymentController{
		usecases: odsEtlPaymentService,
	}
	odsETLGroup := router.Group(apiPrefix + "odsEtl", middlewares.Auth)
	{
		odsETLGroup.GET("payment/index", handlerUser.GetAllOdsEtlPayments)
		odsETLGroup.GET("payment/detail/:id", handlerUser.GetOdsEtlPayment)
		odsETLGroup.POST("payment/create", handlerUser.CreateOdsEtlPayment)
		odsETLGroup.GET("payment/remove-before-map", handlerUser.TruncateTableOdsEtlPayments)
	}
}

// GetAllOdsEtlPayments godoc
// @Param Authorization header string true "Bearer"
// @Security basicAuth
// @Summary List existing odsEtlPayments
// @Description Get all the existing odsEtlPayments
// @Tags odsEtlPayments
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.OdsEtlPayment
// @Failure 401 {object} dto.Response
// @Router /odsEtl/payment/index [get]
func (deliveries *odsEtlPaymentController) GetAllOdsEtlPayments(ctx *gin.Context){
	odsEtlPayments :=  deliveries.usecases.GetAllOdsEtlPayments()
	ctx.JSON(http.StatusOK, gin.H{"data": odsEtlPayments})
}

// GetOdsEtlPaymentsDetail godoc
// @Param Authorization header string true "Bearer"
// @Security basicAuth
// @Summary Show an existing odsEtlPayments
// @Description Get detail the existing odsEtlPayments
// @Tags odsEtlPayments
// @Accept  json
// @Produce  json
// @Param  id path int true "OdsEtlPayment ID"
// @Success 200 {array} entity.OdsEtlPayment
// @Failure 401 {object} dto.Response
// @Router /odsEtl/payment/detail/{id} [get]
func (deliveries *odsEtlPaymentController) GetOdsEtlPayment(ctx *gin.Context){
	odsEtlPayment :=  deliveries.usecases.GetOdsEtlPayment(ctx.Param("id"))
	ctx.JSON(http.StatusOK, gin.H{"data": odsEtlPayment})
}

// OdsEtlPaymentsCreate godoc
// @Param Authorization header string true "Bearer"
// @Security basicAuth
// @Summary Create new odsEtlPayments
// @Description Create a new odsEtlPayment
// @Tags odsEtlPayments
// @Accept  json
// @Produce  json
// @Param odsEtlPayment body entity.OdsEtlPayment true "Create odsEtlPayment"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /odsEtl/payment/create [post]
func (deliveries *odsEtlPaymentController) CreateOdsEtlPayment(ctx *gin.Context){
	var odsEtlPayment entity.OdsEtlPayment
	err := ctx.ShouldBindJSON(&odsEtlPayment)
	if err != nil {
		ctx.JSON(http.StatusConflict, err)
		// ctx.JSON(http.StatusConflict, err)
	}
	
	deliveries.usecases.CreateOdsEtlPayment(odsEtlPayment)
	ctx.JSON(http.StatusOK, gin.H{"data": odsEtlPayment})
}

// RemoveOdsEtlBeforeMap godoc
// @Param Authorization header string true "Bearer"
// @Security basicAuth
// @Summary Truncate existing odsEtlPayments
// @Description Remove all the existing odsEtlPayments
// @Tags odsEtlPayments
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.OdsEtlPayment
// @Failure 401 {object} dto.Response
// @Router /odsEtl/payment/remove-before-map [get]
func (deliveries *odsEtlPaymentController) TruncateTableOdsEtlPayments(ctx *gin.Context) {
	truncateOdsEtlPayments :=  deliveries.usecases.TruncateTableOdsEtlPayments()
	ctx.JSON(http.StatusOK, gin.H{"data": truncateOdsEtlPayments})
}
