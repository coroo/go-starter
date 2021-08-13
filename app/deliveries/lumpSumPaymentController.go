package deliveries

import (
	"net/http"
	// "github.com/Sequis-Digital-Channel/superyou-ods/models"

	// entity "github.com/coroo/go-starter/app/entity"
	usecases "github.com/coroo/go-starter/app/usecases"
	repositories "github.com/coroo/go-starter/app/repositories"

	"github.com/gin-gonic/gin"
)

var (
	lumpSumPaymentRepository 	repositories.LumpSumPaymentRepository = repositories.NewLumpSumPaymentRepository()
	lumpSumPaymentService		usecases.LumpSumPaymentService = usecases.NewLumpSumPaymentService(lumpSumPaymentRepository)
)

// GetAllLumpSumPayments godoc
// @Security basicAuth
// @Summary List existing ods userInvoices
// @Description Get all the existing ods userInvoices
// @Tags odsLumpSumPayments
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.LumpSumPayment
// @Failure 401 {object} dto.Response
// @Router /lumpSumPayment/index [get]
func GetAllLumpSumPayments(ctx *gin.Context) {
	// return c.service.GetAllLumpSumPayments()
	// lumpSumPayments :=  lumpSumPaymentService.GetAllLumpSumPayments()
	// ctx.JSON(http.StatusOK, gin.H{"data": lumpSumPayments})
	lumpSumPayment :=  lumpSumPaymentService.GetAllLumpSumPayments()
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
func OdsMapEtlLatestPayment(ctx *gin.Context) {
	// return c.service.OdsMapEtlLatestPayment()
	latestLumpSumPayment :=  lumpSumPaymentService.OdsMapEtlLatestPayment()
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
func GetLumpSumPayment(ctx *gin.Context) {
	lumpSumPayment :=  lumpSumPaymentService.GetLumpSumPayment(ctx.Param("policyNumber"))
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
