package deliveries

import (
	"net/http"

	usecases "github.com/coroo/go-starter/app/usecases"
	entity "github.com/coroo/go-starter/app/entity"
	// "github.com/coroo/go-starter/app/middlewares"
	
	"github.com/gin-gonic/gin"
)

type GopayLinkingController interface {
	GetGopayLinking(ctx *gin.Context)
	// 	Save(ctx *gin.Context) error
}

type gopayLinkingController struct {
	usecases usecases.GopayLinkingService
}

// var validate *validator.Validate

func NewGopayLinkingController(router *gin.Engine, apiPrefix string, gopayLinkingService usecases.GopayLinkingService) {
	// validate = validator.New()
	// validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	handlerGopayLinking := &gopayLinkingController{
		usecases: gopayLinkingService,
	}
	gopayLinkingsGroup := router.Group(apiPrefix + "gopayLinking")
	{
		gopayLinkingsGroup.GET("index", handlerGopayLinking.GopayLinkingsIndex)
		gopayLinkingsGroup.GET("detail/:id", handlerGopayLinking.GopayLinkingsDetail)
		gopayLinkingsGroup.POST("create", handlerGopayLinking.GopayLinkingCreate)
		gopayLinkingsGroup.PUT("update", handlerGopayLinking.GopayLinkingUpdate)
		gopayLinkingsGroup.DELETE("delete", handlerGopayLinking.GopayLinkingDelete)
	}
}

// GetGopayLinkingsIndex godoc
// @Security basicAuth
// @Summary Show all existing GopayLinkings
// @Description Get all existing GopayLinkings
// @Tags GopayLinkings
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.GopayLinking
// @Failure 401 {object} dto.Response
// @Router /gopayLinking/index [get]
func (deliveries *gopayLinkingController) GopayLinkingsIndex(c *gin.Context) {
	gopayLinkings := deliveries.usecases.GetAllGopayLinkings()
	c.JSON(http.StatusOK, gin.H{"data": gopayLinkings})
}

// GetGopayLinkingsDetail godoc
// @Security basicAuth
// @Summary Show an existing GopayLinkings
// @Description Get detail the existing GopayLinkings
// @Tags GopayLinkings
// @Accept  json
// @Produce  json
// @Param  id path int true "Master Question ID"
// @Success 200 {array} entity.GopayLinking
// @Failure 401 {object} dto.Response
// @Router /gopayLinking/detail/{id} [get]
func (deliveries *gopayLinkingController) GopayLinkingsDetail(c *gin.Context) {
	gopayLinking := deliveries.usecases.GetGopayLinking(c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"data": gopayLinking})
}

// CreateGopayLinkings godoc
// @Security basicAuth
// @Summary Create new GopayLinkings
// @Description Create a new GopayLinkings
// @Tags GopayLinkings
// @Accept  json
// @Produce  json
// @Param gopayLinking body entity.GopayLinking true "Create gopayLinking"
// @Success 200 {object} entity.GopayLinking
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /gopayLinking/create [post]
func (deliveries *gopayLinkingController) GopayLinkingCreate(c *gin.Context) {
	var gopayLinkingEntity entity.GopayLinking
	c.ShouldBindJSON(&gopayLinkingEntity)
	gopayLinkingPK, err := deliveries.usecases.SaveGopayLinking(gopayLinkingEntity)
	if(err!=nil){
		c.JSON(http.StatusConflict, err)
	} else {
		gopayLinkingEntity.ID = gopayLinkingPK
		c.JSON(http.StatusOK, gopayLinkingEntity)
	}
}

// UpdateGopayLinkings godoc
// @Security basicAuth
// @Summary Update GopayLinkings
// @Description Update a GopayLinkings
// @Tags GopayLinkings
// @Accept  json
// @Produce  json
// @Param gopayLinking body entity.GopayLinking true "Update gopayLinking"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /gopayLinking/update [put]
func (deliveries *gopayLinkingController) GopayLinkingUpdate(c *gin.Context) {
	var gopayLinkingEntity entity.GopayLinking
	c.ShouldBindJSON(&gopayLinkingEntity)
	gopayLinking := deliveries.usecases.UpdateGopayLinking(gopayLinkingEntity)
	c.JSON(http.StatusOK, gopayLinking)
}

// DeleteGopayLinkings godoc
// @Security basicAuth
// @Summary Delete GopayLinkings
// @Description Delete a GopayLinkings
// @Tags GopayLinkings
// @Accept  json
// @Produce  json
// @Param gopayLinking body entity.GopayLinking true "Delete gopayLinking"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /gopayLinking/delete [delete]
func (deliveries *gopayLinkingController) GopayLinkingDelete(c *gin.Context) {
	var gopayLinkingEntity entity.GopayLinking
	c.ShouldBindJSON(&gopayLinkingEntity)
	gopayLinking := deliveries.usecases.DeleteGopayLinking(gopayLinkingEntity)
	c.JSON(http.StatusOK, gopayLinking)
}