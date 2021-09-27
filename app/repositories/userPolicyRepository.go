package repositories

import (
	"github.com/coroo/go-starter/config"
	entity "github.com/coroo/go-starter/app/entity"
	"strings"

	// "github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	_ "gorm.io/driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

type UserPolicyRepository interface {
	SaveUserPolicy(userPolicy entity.UserPolicy)
	UpdateUserPolicy(userPolicy entity.UserPolicy)
	DeleteUserPolicy(userPolicy entity.UserPolicy)
	GetAllUserPolicies(is_overdue string) []entity.UserPolicy
	GetUserPolicy(id string) []entity.UserPolicy
	CloseDB()
}

type userPolicyDatabase struct {
	connection *gorm.DB
}

func NewUserPolicyRepository() UserPolicyRepository {
	db, err := config.ConnectDB()
	if err != nil {
		panic("Failed to connect database")
	}
	// db.AutoMigrate(&entity.UserPolicy{}, &entity.Person{})
	db.AutoMigrate(&entity.UserPolicy{})
	return &userPolicyDatabase{
		connection: db,
	}
}

func (db *userPolicyDatabase) CloseDB() {
	sqlDB,err := db.connection.DB()
	if err != nil {
		panic("Failed to close database")
	}
	closeDB := sqlDB.Close()
	if closeDB != nil {
		panic("Failed to close database")
	}
}

func (db *userPolicyDatabase) SaveUserPolicy(userPolicy entity.UserPolicy) {
	db.connection.Create(&userPolicy)
}

func (db *userPolicyDatabase) UpdateUserPolicy(userPolicy entity.UserPolicy) {
	db.connection.Save(&userPolicy)
}

func (db *userPolicyDatabase) DeleteUserPolicy(userPolicy entity.UserPolicy) {
	db.connection.Delete(&userPolicy)
}

func (db *userPolicyDatabase) GetAllUserPolicies(is_overdue string) []entity.UserPolicy {
	var userPolicies []entity.UserPolicy
	query := db.connection.Preload(clause.Associations)

	// if param overdue exist, do query 
	if is_overdue == "1" || strings.ToLower(is_overdue) == "true"{
		query = query.Where("overdue_premium > ?", 0)
	} else if is_overdue == "0" || strings.ToLower(is_overdue) == "false"{
		query = query.Where("overdue_premium = ?", 0)
	}
	query.Find(&userPolicies)
	return userPolicies
}

func (db *userPolicyDatabase) GetUserPolicy(id string) []entity.UserPolicy {
	var userPolicy []entity.UserPolicy
	db.connection.Preload(clause.Associations).First(&userPolicy, id)
	return userPolicy
}
