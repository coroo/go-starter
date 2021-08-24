package deliveries

import (
	"net/http"

	// entity "github.com/coroo/go-starter/app/entity"
	usecases "github.com/coroo/go-starter/app/usecases"
	repositories "github.com/coroo/go-starter/app/repositories"

	"github.com/gin-gonic/gin"
)

var (
	syUserInvoiceRepository 	repositories.SyUserInvoiceRepository = repositories.NewSyUserInvoiceRepository()
	syUserInvoiceService		usecases.SyUserInvoiceService = usecases.NewSyUserInvoiceService(syUserInvoiceRepository)
)

// type SyUserInvoiceController interface {
// 	GetAllSyUserInvoices() []entity.SyUserInvoice
// 	SyMapEtlLatestPayment() []entity.SyUserInvoice
// }

// type syUserInvoiceController struct {
// 	service service.SyUserInvoiceService
// }

// func NewSyUserInvoice(service service.SyUserInvoiceService) SyUserInvoiceController {
// 	return &syUserInvoiceController{
// 		service: service,
// 	}
// }

// GetUserInvoices godoc
// @Security basicAuth
// @Param Authorization header string true "Bearer"
// @Summary List existing superyou userInvoices
// @Description Get all the existing superyou userInvoices
// @Tags syUserInvoices
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.SyUserInvoice
// @Failure 401 {object} dto.Response
// @Router /syUserInvoice/index [get]
func GetAllSyUserInvoices(ctx *gin.Context) {
	// return c.service.GetAllSyUserInvoices()
	syUserInvoices :=  syUserInvoiceService.GetAllSyUserInvoices()
	ctx.JSON(http.StatusOK, gin.H{"data": syUserInvoices})
}

// SyMapEtlLatestPayment godoc
// @Param Authorization header string true "Bearer"
// @Security basicAuth
// @Summary Map existing superyou userInvoices to syEtlPayments
// @Description Map all the existing superyou userInvoices to ODS syEtlPayments db
// @Tags syUserInvoices
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.SyUserInvoice
// @Failure 401 {object} dto.Response
// @Router /syUserInvoice/map-etl-payment [get]
func SyMapEtlLatestPayment(ctx *gin.Context) {
	// return c.service.SyMapEtlLatestPayment()
	syUserInvoices :=  syUserInvoiceService.SyMapEtlLatestPayment()
	ctx.JSON(http.StatusOK, gin.H{"data": syUserInvoices})
}
