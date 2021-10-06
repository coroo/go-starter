package repositories

import (
	"github.com/coroo/go-starter/config"
	entity "github.com/coroo/go-starter/app/entity"

	// "github.com/gin-gonic/gin"
	"gorm.io/gorm"
	_ "gorm.io/driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

type PaymentMethodRateRepository interface {
	SavePaymentMethodRate()
}

type paymentMethodRateDatabase struct {
	connection *gorm.DB
}

func NewPaymentMethodRateRepository() PaymentMethodRateRepository {
	db, err := config.ConnectDB()
	if err != nil {
		panic("Failed to connect database")
	}
	// db.AutoMigrate(&entity.PaymentMethodRate{}, &entity.Person{})
	db.AutoMigrate(&entity.PaymentMethodRate{})
	return &paymentMethodRateDatabase{
		connection: db,
	}
}

func (db *paymentMethodRateDatabase) SavePaymentMethodRate() {

}
