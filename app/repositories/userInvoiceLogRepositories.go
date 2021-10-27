package repositories

import (
	"os"
	"time"
	"gorm.io/gorm/clause"
	"github.com/coroo/go-starter/config"
	entity "github.com/coroo/go-starter/app/entity"

	// "github.com/gin-gonic/gin"
	"gorm.io/gorm"
	_ "gorm.io/driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

type UserInvoiceLogRepository interface {
	SaveUserInvoiceLog(userInvoiceLog entity.UserInvoiceLog) (int, error)
	UpdateUserInvoiceLog(userInvoiceLog entity.UserInvoiceLog) error
	DeleteUserInvoiceLog(userInvoiceLog entity.UserInvoiceLog) error
	GetAllUserInvoiceLogs() []entity.UserInvoiceLog
	GetUserInvoiceLog(id string) []entity.UserInvoiceLog
}

type userInvoiceLogDatabase struct {
	connection *gorm.DB
}

func NewUserInvoiceLogRepository() UserInvoiceLogRepository {
	db, err := config.ConnectDB()
	if err != nil {
		panic("Failed to connect database")
	}
	// db.AutoMigrate(&entity.UserInvoiceLog{}, &entity.Person{})
	if (os.Getenv("DB_HOST_PAYMENT") != ""){
		db.AutoMigrate(&entity.UserInvoiceLog{})
	} else {
		db.AutoMigrate(&entity.UserInvoiceLogTesting{})
	}
	return &userInvoiceLogDatabase{
		connection: db,
	}
}

func (db *userInvoiceLogDatabase) SaveUserInvoiceLog(userInvoiceLog entity.UserInvoiceLog) (int, error) {
	data := &userInvoiceLog
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	err := db.connection.Create(data)
	if(err.Error != nil){
		return 0, err.Error
	}
	return data.ID, nil
}

func (db *userInvoiceLogDatabase) UpdateUserInvoiceLog(userInvoiceLog entity.UserInvoiceLog) error {
	data := &userInvoiceLog
	data.UpdatedAt = time.Now()
	db.connection.Save(data)
	return nil
}

func (db *userInvoiceLogDatabase) DeleteUserInvoiceLog(userInvoiceLog entity.UserInvoiceLog) error {
	db.connection.Delete(&userInvoiceLog)
	return nil
}

func (db *userInvoiceLogDatabase) GetAllUserInvoiceLogs() []entity.UserInvoiceLog {
	var userInvoiceLogs []entity.UserInvoiceLog
	db.connection.Preload(clause.Associations).Find(&userInvoiceLogs)
	return userInvoiceLogs
}

func (db *userInvoiceLogDatabase) GetUserInvoiceLog(id string) []entity.UserInvoiceLog {
	var userInvoiceLog []entity.UserInvoiceLog
	db.connection.Preload(clause.Associations).Where("id = ?", id).First(&userInvoiceLog)
	return userInvoiceLog
}
