package usecases

import (
	entity "github.com/coroo/go-lemonilo/app/entity"
	repositories "github.com/coroo/go-lemonilo/app/repositories"
	// utils "github.com/coroo/go-lemonilo/app/utils"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type UserProfileService interface {
	SaveUserProfile(entity.UserProfile) (int, error)
	UpdateUserProfile(entity.UserProfile) error
	DeleteUserProfile(entity.UserProfile) error
	GetAllUserProfiles() []entity.UserProfile
	GetUserProfile(ctx *gin.Context) []entity.UserProfile
}

type userProfileService struct {
	// userProfile []entity.UserProfile
	// models models.UserProfileModel
	repositories repositories.UserProfileRepository
}

func NewUserProfile(userProfileRepository repositories.UserProfileRepository) UserProfileService {
	// return &userProfileService{
	// 	models: models,
	// }
	return &userProfileService{
		repositories: userProfileRepository,
	}
}

func (usecases *userProfileService) GetAllUserProfiles() []entity.UserProfile {
	return usecases.repositories.GetAllUserProfiles()
}

func (usecases *userProfileService) GetUserProfile(ctx *gin.Context) []entity.UserProfile {
	return usecases.repositories.GetUserProfile(ctx)
}

func (usecases *userProfileService) SaveUserProfile(userProfile entity.UserProfile) (int, error) {
	// userProfile.Password, _ = utils.HashPassword(userProfile.Password)
	return usecases.repositories.SaveUserProfile(userProfile)
}

func (usecases *userProfileService) UpdateUserProfile(userProfile entity.UserProfile) error {
	return usecases.repositories.UpdateUserProfile(userProfile)
}

func (usecases *userProfileService) DeleteUserProfile(userProfile entity.UserProfile) error {
	return usecases.repositories.DeleteUserProfile(userProfile)
}
