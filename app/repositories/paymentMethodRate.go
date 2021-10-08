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

type PaymentMethodRateRepository interface {
	SavePaymentMethodRate(paymentMethod entity.PaymentMethodRate) (int, error)
	UpdatePaymentMethodRate(paymentMethod entity.PaymentMethodRate) error
	DeletePaymentMethodRate(paymentMethod entity.PaymentMethodRate) error
	GetAllPaymentMethodRates() []entity.PaymentMethodRate
	GetPaymentMethodRate(id string) []entity.PaymentMethodRate
}

type paymentMethodRateDatabase struct {
	connection *gorm.DB
}

func NewPaymentMethodRateRepository() PaymentMethodRateRepository {
	db, err := config.ConnectDB()
	if err != nil {
		panic("Failed to connect database")
	}
	// db.AutoMigrate(&entity.PaymentMethodRateRate{}, &entity.Person{})
	db.AutoMigrate(&entity.PaymentMethodRate{})
	return &paymentMethodRateDatabase{
		connection: db,
	}
}

func (db *paymentMethodRateDatabase) SavePaymentMethodRate(paymentMethodRate entity.PaymentMethodRate) (int, error) {
	data := &paymentMethodRate
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	err := db.connection.Create(data)
	if(err.Error != nil){
		return 0, err.Error
	}
	return data.ID, nil
}

func (db *paymentMethodRateDatabase) UpdatePaymentMethodRate(paymentMethodRate entity.PaymentMethodRate) error {
	data := &paymentMethodRate
	data.UpdatedAt = time.Now()
	db.connection.Save(data)
	return nil
}

func (db *paymentMethodRateDatabase) DeletePaymentMethodRate(paymentMethodRate entity.PaymentMethodRate) error {
	db.connection.Delete(&paymentMethodRate)
	return nil
}

func (db *paymentMethodRateDatabase) GetAllPaymentMethodRates() []entity.PaymentMethodRate {
	var paymentMethodRates []entity.PaymentMethodRate
	db.connection.Preload(clause.Associations).Find(&paymentMethodRates)
	return paymentMethodRates
}

func (db *paymentMethodRateDatabase) GetPaymentMethodRate(id string) []entity.PaymentMethodRate {
	var paymentMethodRate []entity.PaymentMethodRate
	db.connection.Preload(clause.Associations).Where("id = ?", id).First(&paymentMethodRate)
	return paymentMethodRate
}
