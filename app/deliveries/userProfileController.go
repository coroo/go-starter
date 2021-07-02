package deliveries

import (
	entity "github.com/coroo/go-lemonilo/app/entity"
	usecases "github.com/coroo/go-lemonilo/app/usecases"

	"github.com/gin-gonic/gin"
)

type UserProfileController interface {
	Save(userProfile entity.UserProfile) (int, error)
	Update(userProfile entity.UserProfile) error
	Delete(userProfile entity.UserProfile) error
	GetAllUserProfiles() []entity.UserProfile
	GetUserProfile(ctx *gin.Context) []entity.UserProfile
}

type userProfileDeliveries struct {
	usecases usecases.UserProfileService
}

func NewUserProfile(usecases usecases.UserProfileService) UserProfileController {
	return &userProfileDeliveries{
		usecases: usecases,
	}
}

func (c *userProfileDeliveries) GetAllUserProfiles() []entity.UserProfile {
	return c.usecases.GetAllUserProfiles()
}

func (c *userProfileDeliveries) GetUserProfile(ctx *gin.Context) []entity.UserProfile {
	return c.usecases.GetUserProfile(ctx)
}

func (c *userProfileDeliveries) Save(userProfile entity.UserProfile) (int, error) {
	return c.usecases.SaveUserProfile(userProfile)
}

func (c *userProfileDeliveries) Update(userProfile entity.UserProfile) error {
	c.usecases.UpdateUserProfile(userProfile)
	return nil
}

func (c *userProfileDeliveries) Delete(userProfile entity.UserProfile) error {
	c.usecases.DeleteUserProfile(userProfile)
	return nil
}
