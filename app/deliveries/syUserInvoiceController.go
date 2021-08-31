package deliveries

import (
	"net/http"

	// entity "github.com/coroo/go-starter/app/entity"
	usecases "github.com/coroo/go-starter/app/usecases"
	"github.com/coroo/go-starter/app/middlewares"

	"github.com/gin-gonic/gin"
)

type syUserInvoiceController struct {
	usecases usecases.SyUserInvoiceService
}

func NewSyUserInvoiceController(router *gin.Engine, apiPrefix string, syUserInvoiceService usecases.SyUserInvoiceService) {
	handlerSyUserInvoice := &syUserInvoiceController{
		usecases: syUserInvoiceService,
	}

	syUserInvoiceGroup := router.Group(apiPrefix + "syUserInvoice", middlewares.Auth)
	{
		syUserInvoiceGroup.GET("index", handlerSyUserInvoice.GetAllSyUserInvoices)
		syUserInvoiceGroup.GET("map-etl-payment", handlerSyUserInvoice.SyMapEtlLatestPayment)
	}
}

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
func (deliveries *syUserInvoiceController) GetAllSyUserInvoices(ctx *gin.Context) {
	// return c.service.GetAllSyUserInvoices()
	syUserInvoices :=  deliveries.usecases.GetAllSyUserInvoices()
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
func (deliveries *syUserInvoiceController) SyMapEtlLatestPayment(ctx *gin.Context) {
	// return c.service.SyMapEtlLatestPayment()
	syUserInvoices :=  deliveries.usecases.SyMapEtlLatestPayment()
	ctx.JSON(http.StatusOK, gin.H{"data": syUserInvoices})
}
