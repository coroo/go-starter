package repositories

import (
	"github.com/coroo/go-starter/config"
	entity "github.com/coroo/go-starter/app/entity"

	// "github.com/gin-gonic/gin"
	"gorm.io/gorm"
	_ "gorm.io/driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

type PaymentMethodRepository interface {
	SavePaymentMethod()
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

func (db *paymentMethodDatabase) SavePaymentMethod() {

}
