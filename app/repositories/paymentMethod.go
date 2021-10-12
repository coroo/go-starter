package repositories

import (
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
	GetAllPaymentMethods() []entity.PaymentMethod
	GetPaymentMethod(id string) []entity.PaymentMethod
	GetPaymentMethodByCode(code string) []entity.PaymentMethod
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
	db.AutoMigrate(&entity.PaymentMethod{})
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

func (db *paymentMethodDatabase) GetAllPaymentMethods() []entity.PaymentMethod {
	var paymentMethods []entity.PaymentMethod
	db.connection.Preload(clause.Associations).Find(&paymentMethods)
	return paymentMethods
}

func (db *paymentMethodDatabase) GetPaymentMethod(id string) []entity.PaymentMethod {
	var paymentMethod []entity.PaymentMethod
	db.connection.Preload(clause.Associations).Where("id = ?", id).First(&paymentMethod)
	return paymentMethod
}
func (db *paymentMethodDatabase) GetPaymentMethodByCode(code string) []entity.PaymentMethod {
	var paymentMethod []entity.PaymentMethod
	db.connection.Preload(clause.Associations).Where("code = ?", code).First(&paymentMethod)
	return paymentMethod
}
