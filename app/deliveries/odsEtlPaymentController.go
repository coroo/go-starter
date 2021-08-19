package deliveries

import (
	"net/http"

	entity "github.com/coroo/go-starter/app/entity"
	usecases "github.com/coroo/go-starter/app/usecases"
	repositories "github.com/coroo/go-starter/app/repositories"

	"github.com/gin-gonic/gin"
)

var (
	odsEtlPaymentRepository repositories.OdsEtlPaymentRepository = repositories.NewOdsEtlPaymentRepository()
	odsEtlPaymentService    usecases.OdsEtlPaymentService = usecases.NewOdsEtlPaymentService(odsEtlPaymentRepository)
	// userController deliveries.UserController   = deliveries.NewUser(userService)
)

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
func GetAllOdsEtlPayments(ctx *gin.Context){
	odsEtlPayments :=  odsEtlPaymentService.GetAllOdsEtlPayments()
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
// @Param  policyNumber path int true "Lump Sum Policy Number"
// @Success 200 {array} entity.OdsEtlPayment
// @Failure 401 {object} dto.Response
// @Router /odsEtl/payment/detail/{policyNumber} [get]
func GetOdsEtlPayment(ctx *gin.Context){
	odsEtlPayment :=  odsEtlPaymentService.GetOdsEtlPayment(ctx.Param("id"))
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
func CreateOdsEtlPayment(ctx *gin.Context){
	var odsEtlPayment entity.OdsEtlPayment
	err := ctx.ShouldBindJSON(&odsEtlPayment)

	if err != nil {
		ctx.JSON(http.StatusConflict, err)
		// ctx.JSON(http.StatusConflict, err)
	}
	
	odsEtlPaymentService.CreateOdsEtlPayment(odsEtlPayment)
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
func TruncateTableOdsEtlPayments(ctx *gin.Context) {
	truncateOdsEtlPayments :=  odsEtlPaymentService.TruncateTableOdsEtlPayments()
	ctx.JSON(http.StatusOK, gin.H{"data": truncateOdsEtlPayments})
}
