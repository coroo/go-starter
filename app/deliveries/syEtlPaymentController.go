package deliveries

import (
	"net/http"
	// "github.com/Sequis-Digital-Channel/superyou-ods/models"

	entity "github.com/coroo/go-starter/app/entity"
	usecases "github.com/coroo/go-starter/app/usecases"
	repositories "github.com/coroo/go-starter/app/repositories"

	"github.com/gin-gonic/gin"
)

var (
	syEtlPaymentRepository repositories.SyEtlPaymentRepository = repositories.NewSyEtlPaymentRepository()
	syEtlPaymentService    usecases.SyEtlPaymentService = usecases.NewSyEtlPaymentService(syEtlPaymentRepository)
)

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
// @Security basicAuth
// @Summary List existing syEtlPayments
// @Description Get all the existing syEtlPayments
// @Tags syEtlPayments
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.SyEtlPayment
// @Failure 401 {object} dto.Response
// @Router /syEtl/payment/index [get]
func GetAllSyEtlPayments(ctx *gin.Context){
	syEtlPayments :=  syEtlPaymentService.GetAllSyEtlPayments()
	ctx.JSON(http.StatusOK, gin.H{"data": syEtlPayments})
}

// SyOdsMapEtlLatestPayment godoc
// @Security basicAuth
// @Summary Map existing syEtlPayments to syOdsEtlPayments
// @Description Map all the existing syEtlPayments to syOdsEtlPayments
// @Tags syEtlPayments
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.SyEtlPayment
// @Failure 401 {object} dto.Response
// @Router /syEtl/payment/map-etl-payment [get]
func SyOdsMapEtlLatestPayment(ctx *gin.Context){
	latestSyEtlPayments :=  syEtlPaymentService.SyOdsMapEtlLatestPayment()
	ctx.JSON(http.StatusOK, gin.H{"data": latestSyEtlPayments})
}

// GetSyEtlPaymentsDetail godoc
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
func GetSyEtlPayment(ctx *gin.Context){
	syEtlPayment :=  syEtlPaymentService.GetSyEtlPayment(ctx.Param("policyNumber"))
	ctx.JSON(http.StatusOK, gin.H{"data": syEtlPayment})
}

// SyEtlPaymentsCreate godoc
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
func Save(ctx *gin.Context){
	var syEtlPayment entity.SyEtlPayment
	err := ctx.ShouldBindJSON(&syEtlPayment)
	if err != nil {
		ctx.JSON(http.StatusConflict, err)
	}
	// err = validate.Struct(syEtlPayment)
	// if err != nil {
	// 	return err
	// }
	syEtlPaymentService.Save(syEtlPayment)
	ctx.JSON(http.StatusOK, syEtlPayment)
}

// RemoveSyEtlBeforeMap godoc
// @Security basicAuth
// @Summary Truncate existing syEtlPayments
// @Description Remove all the existing syEtlPayments
// @Tags syEtlPayments
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.SyEtlPayment
// @Failure 401 {object} dto.Response
// @Router /syEtl/payment/remove-before-map [get]
func TruncateTableSyEtlPayments(ctx *gin.Context){
	truncateSyEtlPayment :=  syEtlPaymentService.TruncateTableSyEtlPayments()
	ctx.JSON(http.StatusOK, gin.H{"data": truncateSyEtlPayment})
}
