package deliveries

import (
	"net/http"

	usecases "github.com/coroo/go-starter/app/usecases"
	entity "github.com/coroo/go-starter/app/entity"
	// "github.com/coroo/go-starter/app/middlewares"
	
	"github.com/gin-gonic/gin"
)

type UserInvoiceLogController interface {
	GetUserInvoiceLog(ctx *gin.Context)
	// 	Save(ctx *gin.Context) error
}

type userInvoiceLogController struct {
	usecases usecases.UserInvoiceLogService
}

// var validate *validator.Validate

func NewUserInvoiceLogController(router *gin.Engine, apiPrefix string, userInvoiceLogService usecases.UserInvoiceLogService) {
	// validate = validator.New()
	// validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	handlerUserInvoiceLog := &userInvoiceLogController{
		usecases: userInvoiceLogService,
	}
	userInvoiceLogsGroup := router.Group(apiPrefix + "userInvoiceLog")
	{
		userInvoiceLogsGroup.GET("index", handlerUserInvoiceLog.UserInvoiceLogsIndex)
		userInvoiceLogsGroup.GET("detail/:id", handlerUserInvoiceLog.UserInvoiceLogsDetail)
		userInvoiceLogsGroup.POST("create", handlerUserInvoiceLog.UserInvoiceLogCreate)
		userInvoiceLogsGroup.PUT("update", handlerUserInvoiceLog.UserInvoiceLogUpdate)
		userInvoiceLogsGroup.DELETE("delete", handlerUserInvoiceLog.UserInvoiceLogDelete)
	}
}

// GetUserInvoiceLogsIndex godoc
// @Security basicAuth
// @Summary Show all existing UserInvoiceLogs
// @Description Get all existing UserInvoiceLogs
// @Tags UserInvoiceLogs
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.UserInvoiceLog
// @Failure 401 {object} dto.Response
// @Router /userInvoiceLog/index [get]
func (deliveries *userInvoiceLogController) UserInvoiceLogsIndex(c *gin.Context) {
	userInvoiceLogs := deliveries.usecases.GetAllUserInvoiceLogs()
	c.JSON(http.StatusOK, gin.H{"data": userInvoiceLogs})
}

// GetUserInvoiceLogsDetail godoc
// @Security basicAuth
// @Summary Show an existing UserInvoiceLogs
// @Description Get detail the existing UserInvoiceLogs
// @Tags UserInvoiceLogs
// @Accept  json
// @Produce  json
// @Param  id path int true "Master Question ID"
// @Success 200 {array} entity.UserInvoiceLog
// @Failure 401 {object} dto.Response
// @Router /userInvoiceLog/detail/{id} [get]
func (deliveries *userInvoiceLogController) UserInvoiceLogsDetail(c *gin.Context) {
	userInvoiceLog := deliveries.usecases.GetUserInvoiceLog(c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"data": userInvoiceLog})
}

// CreateUserInvoiceLogs godoc
// @Security basicAuth
// @Summary Create new UserInvoiceLogs
// @Description Create a new UserInvoiceLogs
// @Tags UserInvoiceLogs
// @Accept  json
// @Produce  json
// @Param userInvoiceLog body entity.UserInvoiceLog true "Create userInvoiceLog"
// @Success 200 {object} entity.UserInvoiceLog
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /userInvoiceLog/create [post]
func (deliveries *userInvoiceLogController) UserInvoiceLogCreate(c *gin.Context) {
	var userInvoiceLogEntity entity.UserInvoiceLog
	c.ShouldBindJSON(&userInvoiceLogEntity)
	userInvoiceLogPK, err := deliveries.usecases.SaveUserInvoiceLog(userInvoiceLogEntity)
	if(err!=nil){
		c.JSON(http.StatusConflict, err)
	} else {
		userInvoiceLogEntity.ID = userInvoiceLogPK
		c.JSON(http.StatusOK, userInvoiceLogEntity)
	}
}

// UpdateUserInvoiceLogs godoc
// @Security basicAuth
// @Summary Update UserInvoiceLogs
// @Description Update a UserInvoiceLogs
// @Tags UserInvoiceLogs
// @Accept  json
// @Produce  json
// @Param userInvoiceLog body entity.UserInvoiceLog true "Update userInvoiceLog"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /userInvoiceLog/update [put]
func (deliveries *userInvoiceLogController) UserInvoiceLogUpdate(c *gin.Context) {
	var userInvoiceLogEntity entity.UserInvoiceLog
	c.ShouldBindJSON(&userInvoiceLogEntity)
	userInvoiceLog := deliveries.usecases.UpdateUserInvoiceLog(userInvoiceLogEntity)
	c.JSON(http.StatusOK, userInvoiceLog)
}

// DeleteUserInvoiceLogs godoc
// @Security basicAuth
// @Summary Delete UserInvoiceLogs
// @Description Delete a UserInvoiceLogs
// @Tags UserInvoiceLogs
// @Accept  json
// @Produce  json
// @Param userInvoiceLog body entity.UserInvoiceLog true "Delete userInvoiceLog"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /userInvoiceLog/delete [delete]
func (deliveries *userInvoiceLogController) UserInvoiceLogDelete(c *gin.Context) {
	var userInvoiceLogEntity entity.UserInvoiceLog
	c.ShouldBindJSON(&userInvoiceLogEntity)
	userInvoiceLog := deliveries.usecases.DeleteUserInvoiceLog(userInvoiceLogEntity)
	c.JSON(http.StatusOK, userInvoiceLog)
}