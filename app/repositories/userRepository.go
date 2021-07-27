package repositories

import (
	"time"

	entity "github.com/coroo/go-pawoon-user/app/entity"
	"github.com/coroo/go-pawoon-user/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	_ "gorm.io/driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

type UserRepository interface {
	SaveUser(user entity.User) (int, error)
	UpdateUser(user entity.User) error
	DeleteUser(user entity.User) error
	GetAllUsers() []entity.User
	GetUser(ctx *gin.Context) []entity.User
	AuthUser(user entity.User) entity.User
}

type userDatabase struct {
	connection *gorm.DB
}

func NewUserRepository() UserRepository {
	db, _ := config.ConnectDB()
	db.AutoMigrate(&entity.User{})
	return &userDatabase{
		connection: db,
	}
}

func (db *userDatabase) SaveUser(user entity.User) (int, error) {
	data := &user
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	err := db.connection.Create(data)
	if(err.Error != nil){
		return 0, err.Error
	}
	return data.ID, nil
}

func (db *userDatabase) UpdateUser(user entity.User) error {
	data := &user
	data.UpdatedAt = time.Now()
	db.connection.Save(data)
	return nil
}

func (db *userDatabase) DeleteUser(user entity.User) error {
	db.connection.Delete(&user)
	return nil
}

func (db *userDatabase) GetAllUsers() []entity.User {
	var users []entity.User
	db.connection.Set("gorm:auto_preload", true).Find(&users)
	return users
}

func (db *userDatabase) GetUser(ctx *gin.Context) []entity.User {
	var user []entity.User
	db.connection.Set("gorm:auto_preload", true).Where("id = ?", ctx.Param("id")).First(&user)
	return user
}

func (db *userDatabase) AuthUser(user entity.User) entity.User {
	data := &user
	db.connection.Set("gorm:auto_preload", true).Where("email = ?", data.Email).First(&user)
	return user
}
