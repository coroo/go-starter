package deliveries

import (
	"net/http"

	// "github.com/coroo/go-starter/models"
	usecases "github.com/coroo/go-starter/app/usecases"
	"github.com/coroo/go-starter/app/middlewares"
	
	"github.com/gin-gonic/gin"
)

type UserPolicyController interface {
	GetAllUserPolicies(ctx *gin.Context)
	GetUserPolicy(ctx *gin.Context)
	// 	Save(ctx *gin.Context) error
}

type userPolicyController struct {
	usecases usecases.UserPolicyService
}

// var validate *validator.Validate

func NewUserPolicyController(router *gin.Engine, apiPrefix string, userPolicyService usecases.UserPolicyService) {
	// validate = validator.New()
	// validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	handlerUserPolicy := &userPolicyController{
		usecases: userPolicyService,
	}
	userPoliciesGroup := router.Group(apiPrefix + "userPolicies", middlewares.Auth)
	{
		userPoliciesGroup.GET("index", handlerUserPolicy.GetAllUserPolicies)
		userPoliciesGroup.GET("detail/:id", handlerUserPolicy.GetUserPolicy)
	}
}

func (deliveries *userPolicyController) GetAllUserPolicies(ctx *gin.Context) {
	userPolicies :=  deliveries.usecases.GetAllUserPolicies(ctx.Query("is_overdue"))
	ctx.JSON(http.StatusOK, gin.H{"data": userPolicies})
}

func (deliveries *userPolicyController) GetUserPolicy(ctx *gin.Context) {
	userPolicy :=  deliveries.usecases.GetUserPolicy(ctx.Param("id"))
	ctx.JSON(http.StatusOK, gin.H{"data": userPolicy})
}

// func (c *controller) Save(ctx *gin.Context) error {
// 	var userPolicy entity.UserPolicy
// 	err := ctx.ShouldBindJSON(&userPolicy)
// 	if err != nil {
// 		return err
// 	}
// 	err = validate.Struct(userPolicy)
// 	if err != nil {
// 		return err
// 	}
// 	c.service.Save(userPolicy)
// 	return nil
// }
