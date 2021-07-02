package routes

import (
	"net/http"

	deliveries "github.com/coroo/go-lemonilo/app/deliveries"
	entity "github.com/coroo/go-lemonilo/app/entity"
	repositories "github.com/coroo/go-lemonilo/app/repositories"
	usecases "github.com/coroo/go-lemonilo/app/usecases"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var (
	userProfileRepository repositories.UserProfileRepository = repositories.NewUserProfileRepository()
	userProfileService    usecases.UserProfileService        = usecases.NewUserProfile(userProfileRepository)
	userProfileController deliveries.UserProfileController   = deliveries.NewUserProfile(userProfileService)
)

// GetUserProfilesIndex godoc
// @Security basicAuth
// @Summary Show all existing User_Profiles
// @Description Get all existing User_Profiles
// @Tags User_Profiles
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.UserProfile
// @Failure 401 {object} dto.Response
// @Router /userProfile/index [get]
func UserProfilesIndex(c *gin.Context) {
	userProfiles := userProfileController.GetAllUserProfiles()
	c.JSON(http.StatusOK, gin.H{"data": userProfiles})
}

// GetUserProfilesDetail godoc
// @Security basicAuth
// @Summary Show an existing User_Profiles
// @Description Get detail the existing User_Profiles
// @Tags User_Profiles
// @Accept  json
// @Produce  json
// @Param  id path int true "Master Question ID"
// @Success 200 {array} entity.UserProfile
// @Failure 401 {object} dto.Response
// @Router /userProfile/detail/{id} [get]
func UserProfilesDetail(c *gin.Context) {
	userProfile := userProfileController.GetUserProfile(c)
	c.JSON(http.StatusOK, gin.H{"data": userProfile})
}

// CreateUserProfiles godoc
// @Security basicAuth
// @Summary Create new User_Profiles
// @Description Create a new User_Profiles
// @Tags User_Profiles
// @Accept  json
// @Produce  json
// @Param userProfile body entity.UserProfileCreate true "Create userProfile"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /userProfile/create [post]
func UserProfileCreate(c *gin.Context) {
	var userProfileEntity entity.UserProfile
	c.ShouldBindJSON(&userProfileEntity)
	userProfilePK, err := userProfileController.Save(userProfileEntity)
	if(err!=nil){
		c.JSON(http.StatusConflict, err)
	} else {
		userProfileEntity.ID = userProfilePK
		c.JSON(http.StatusOK, userProfileEntity)
	}
}

// UpdateUserProfiles godoc
// @Security basicAuth
// @Summary Update User_Profiles
// @Description Update a User_Profiles
// @Tags User_Profiles
// @Accept  json
// @Produce  json
// @Param userProfile body entity.UserProfile true "Update userProfile"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /userProfile/update [put]
func UserProfileUpdate(c *gin.Context) {
	var userProfileEntity entity.UserProfile
	c.ShouldBindJSON(&userProfileEntity)
	userProfile := userProfileController.Update(userProfileEntity)
	c.JSON(http.StatusOK, userProfile)
}

// DeleteUserProfiles godoc
// @Security basicAuth
// @Summary Delete User_Profiles
// @Description Delete a User_Profiles
// @Tags User_Profiles
// @Accept  json
// @Produce  json
// @Param userProfile body entity.UserProfileDelete true "Delete userProfile"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /userProfile/delete [delete]
func UserProfileDelete(c *gin.Context) {
	var userProfileEntity entity.UserProfile
	c.ShouldBindJSON(&userProfileEntity)
	userProfile := userProfileController.Delete(userProfileEntity)
	c.JSON(http.StatusOK, userProfile)
}

// Login godoc
// @Security basicAuth
// @Summary Login new Users
// @Description Login a new Users
// @Tags Users
// @Accept  json
// @Produce  json
// @Param userProfile body entity.UserProfileLogin true "Login userProfile"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /userProfile/login [post]
func Login(c *gin.Context) {
	var userProfileEntity entity.UserProfile
	c.ShouldBindJSON(&userProfileEntity)
	userProfilePK, err := userProfileController.Save(userProfileEntity)
	if(err!=nil){
		c.JSON(http.StatusConflict, err)
	} else {
		userProfileEntity.ID = userProfilePK
		c.JSON(http.StatusOK, userProfileEntity)
	}
}