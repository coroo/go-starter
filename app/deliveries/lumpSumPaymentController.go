package deliveries

import (
	"net/http"
	// "github.com/Sequis-Digital-Channel/superyou-ods/models"

	// entity "github.com/coroo/go-starter/app/entity"
	usecases "github.com/coroo/go-starter/app/usecases"
	"github.com/coroo/go-starter/app/middlewares"

	"github.com/gin-gonic/gin"
)

type lumpSumPaymentController struct {
	usecases usecases.LumpSumPaymentService
}

func NewLumpSumPaymentController(router *gin.Engine, apiPrefix string, lumpSumPaymentService usecases.LumpSumPaymentService) {
	handlerUser := &lumpSumPaymentController{
		usecases: lumpSumPaymentService,
	}
	lumpSumPaymentGroup := router.Group(apiPrefix + "lumpSumPayment", middlewares.Auth)
	{
		lumpSumPaymentGroup.GET("index", handlerUser.GetAllLumpSumPayments)
		lumpSumPaymentGroup.GET("detail/:policyNumber", handlerUser.GetLumpSumPayment)
		lumpSumPaymentGroup.GET("map-etl-payment", handlerUser.OdsMapEtlLatestPayment)
	}
}

// GetAllLumpSumPayments godoc
// @Param Authorization header string true "Bearer"
// @Summary List existing ods userInvoices
// @Description Get all the existing ods userInvoices
// @Tags odsLumpSumPayments
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.LumpSumPayment
// @Failure 401 {object} dto.Response
// @Router /lumpSumPayment/index [get]
func (deliveries *lumpSumPaymentController) GetAllLumpSumPayments(ctx *gin.Context) {
	// return c.service.GetAllLumpSumPayments()
	// lumpSumPayments :=  lumpSumPaymentService.GetAllLumpSumPayments()
	// ctx.JSON(http.StatusOK, gin.H{"data": lumpSumPayments})
	lumpSumPayment :=  deliveries.usecases.GetAllLumpSumPayments()
	ctx.JSON(http.StatusOK, gin.H{"data": lumpSumPayment})
}

// OdsMapEtlLatestPayment godoc
// @Security basicAuth
// @Summary Map all existing ods userInvoice to odsEtlPayment
// @Description Map all the existing ods userInvoice to odsEtlPayment
// @Tags odsLumpSumPayments
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.LumpSumPayment
// @Failure 401 {object} dto.Response
// @Router /lumpSumPayment/map-etl-payment [get]
func (deliveries *lumpSumPaymentController) OdsMapEtlLatestPayment(ctx *gin.Context) {
	// return c.service.OdsMapEtlLatestPayment()
	latestLumpSumPayment :=  deliveries.usecases.OdsMapEtlLatestPayment()
	ctx.JSON(http.StatusOK, gin.H{"data": latestLumpSumPayment})
}

// GetLumpSumPaymentsDetail godoc
// @Security basicAuth
// @Summary Show an existing ods userInvoice
// @Description Get detail the existing ods userInvoice
// @Tags odsLumpSumPayments
// @Accept  json
// @Produce  json
// @Param  policyNumber path int true "Lump Sum Policy Number"
// @Success 200 {array} entity.LumpSumPayment
// @Failure 401 {object} dto.Response
// @Router /lumpSumPayment/detail/{policyNumber} [get]
func (deliveries *lumpSumPaymentController) GetLumpSumPayment(ctx *gin.Context) {
	lumpSumPayment :=  deliveries.usecases.GetLumpSumPayment(ctx.Param("policyNumber"))
	ctx.JSON(http.StatusOK, gin.H{"data": lumpSumPayment})
}

// func Save(ctx *gin.Context) error {
// 	var lumpSumPayment entity.LumpSumPayment
// 	err := ctx.ShouldBindJSON(&lumpSumPayment)
// 	if err != nil {
// 		return err
// 	}
// 	err = validate.Struct(lumpSumPayment)
// 	if err != nil {
// 		return err
// 	}
// 	c.service.Save(lumpSumPayment)
// 	return nil
// }
