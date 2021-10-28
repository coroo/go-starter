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

type PaymentMethodRepository interface {
	SavePaymentMethod(paymentMethod entity.PaymentMethod) (int, error)
	UpdatePaymentMethod(paymentMethod entity.PaymentMethod) error
	DeletePaymentMethod(paymentMethod entity.PaymentMethod) error
	GetAllPaymentMethods(status string) []entity.PaymentMethod
	GetPaymentMethod(id string) []entity.PaymentMethod
	GetPaymentMethodByCode(code string) entity.PaymentMethod
	GetActivePaymentMethodByCode(code string) entity.PaymentMethod
}

type paymentMethodDatabase struct {
	connection *gorm.DB
}

func NewPaymentMethodRepository() PaymentMethodRepository {
	db, err := config.ConnectDB()
	if err != nil {
		panic("Failed to connect database")
	}
	// db.AutoMigrate(&entity.PaymentMethod{}, &entity.Person{})
	if (os.Getenv("DB_HOST_PAYMENT") != ""){
		db.AutoMigrate(&entity.PaymentMethod{})
	} else {
		db.AutoMigrate(&entity.PaymentMethodTesting{})
	}
	return &paymentMethodDatabase{
		connection: db,
	}
}

func (db *paymentMethodDatabase) SavePaymentMethod(paymentMethod entity.PaymentMethod) (int, error) {
	data := &paymentMethod
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	err := db.connection.Create(data)
	if(err.Error != nil){
		return 0, err.Error
	}
	return data.ID, nil
}

func (db *paymentMethodDatabase) UpdatePaymentMethod(paymentMethod entity.PaymentMethod) error {
	data := &paymentMethod
	data.UpdatedAt = time.Now()
	db.connection.Save(data)
	return nil
}

func (db *paymentMethodDatabase) DeletePaymentMethod(paymentMethod entity.PaymentMethod) error {
	db.connection.Delete(&paymentMethod)
	return nil
}

func (db *paymentMethodDatabase) GetAllPaymentMethods(status string) []entity.PaymentMethod {
	var paymentMethods []entity.PaymentMethod
	query := db.connection.Preload(clause.Associations)
	if status == "active" || status == "inactive"{
		query = query.Where("status = ?", status)
	}
	query = query.Find(&paymentMethods)
	return paymentMethods
}

func (db *paymentMethodDatabase) GetPaymentMethod(id string) []entity.PaymentMethod {
	var paymentMethod []entity.PaymentMethod
	db.connection.Preload(clause.Associations).Where("id = ?", id).First(&paymentMethod)
	return paymentMethod
}

func (db *paymentMethodDatabase) GetPaymentMethodByCode(code string) entity.PaymentMethod {
	var paymentMethod entity.PaymentMethod
	db.connection.Preload(clause.Associations).Where("code = ?", code).First(&paymentMethod)
	return paymentMethod
}

func (db *paymentMethodDatabase) GetActivePaymentMethodByCode(code string) entity.PaymentMethod {
	var paymentMethod entity.PaymentMethod
	db.connection.Preload(clause.Associations).Where("code = ?", code).Where("status = ?", "active").First(&paymentMethod)
	return paymentMethod
}
