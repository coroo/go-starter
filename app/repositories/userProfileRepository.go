package repositories

import (
	"time"
	"fmt"

	entity "github.com/coroo/go-lemonilo/app/entity"
	"github.com/coroo/go-lemonilo/config"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/joho/godotenv/autoload"
)

type UserProfileRepository interface {
	SaveUserProfile(userProfile entity.UserProfile) (int, error)
	UpdateUserProfile(userProfile entity.UserProfile) error
	DeleteUserProfile(userProfile entity.UserProfile) error
	GetAllUserProfiles() []entity.UserProfile
	GetUserProfile(ctx *gin.Context) []entity.UserProfile
}

type userProfileDatabase struct {
	connection *gorm.DB
}

func NewUserProfileRepository() UserProfileRepository {
	db, _ := config.ConnectDB()
	db.AutoMigrate(&entity.UserProfile{})
	return &userProfileDatabase{
		connection: db,
	}
}

func (db *userProfileDatabase) SaveUserProfile(userProfile entity.UserProfile) (int, error) {
	data := &userProfile
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	err := db.connection.Create(data)
	if(err.Error != nil){
		return 0, err.Error
	}
	// result := db.connection.Last(data)
	fmt.Print(data.ID)
	return data.ID, nil
	// return nil
}

func (db *userProfileDatabase) UpdateUserProfile(userProfile entity.UserProfile) error {
	data := &userProfile
	data.UpdatedAt = time.Now()
	db.connection.Save(data)
	return nil
}

func (db *userProfileDatabase) DeleteUserProfile(userProfile entity.UserProfile) error {
	db.connection.Delete(&userProfile)
	return nil
}

func (db *userProfileDatabase) GetAllUserProfiles() []entity.UserProfile {
	var userProfiles []entity.UserProfile
	db.connection.Set("gorm:auto_preload", true).Find(&userProfiles)
	return userProfiles
}

func (db *userProfileDatabase) GetUserProfile(ctx *gin.Context) []entity.UserProfile {
	var userProfile []entity.UserProfile
	db.connection.Set("gorm:auto_preload", true).Where("id = ?", ctx.Param("id")).First(&userProfile)
	return userProfile
}
