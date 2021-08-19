package deliveries

import (
	"net/http"

	// "github.com/coroo/go-starter/models"
	usecases "github.com/coroo/go-starter/app/usecases"
	repositories "github.com/coroo/go-starter/app/repositories"
	
	
	"github.com/gin-gonic/gin"
)

// type UserPolicyController interface {
// 	GetAllUserPolicies(ctx *gin.Context)
// 	GetUserPolicy(ctx *gin.Context)
// 	// 	Save(ctx *gin.Context) error
// }

var (
	userPolicyRepository repositories.UserPolicyRepository = repositories.NewUserPolicyRepository()
	userPolicyService    usecases.UserPolicyService = usecases.NewUserPolicyService(userPolicyRepository)
	// userController deliveries.UserController   = deliveries.NewUser(userService)
)

// type controller struct {
// 	service service.UserPolicyService
// }

// var validate *validator.Validate

// func NewUserPolicy(service service.UserPolicyService) UserPolicyController {
// 	// validate = validator.New()
// 	// validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
// 	return &controller{
// 		service: service,
// 	}
// }

func GetAllUserPolicies(ctx *gin.Context) {
	userPolicies :=  userPolicyService.GetAllUserPolicies(ctx.Query("is_overdue"))
	ctx.JSON(http.StatusOK, gin.H{"data": userPolicies})
}

func GetUserPolicy(ctx *gin.Context) {
	userPolicy :=  userPolicyService.GetUserPolicy(ctx.Param("id"))
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
