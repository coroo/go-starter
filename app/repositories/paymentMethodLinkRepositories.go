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

type PaymentMethodLinkRepository interface {
	SavePaymentMethodLink(paymentMethodLink entity.PaymentMethodLink) (int, error)
	UpdatePaymentMethodLink(paymentMethodLink entity.PaymentMethodLink) error
	DeletePaymentMethodLink(paymentMethodLink entity.PaymentMethodLink) error
	GetAllPaymentMethodLinks() []entity.PaymentMethodLink
	GetPaymentMethodLink(id string) []entity.PaymentMethodLink
	GetPaymentMethodLinkByCode(code string) entity.PaymentMethodLink
}

type paymentMethodLinkDatabase struct {
	connection *gorm.DB
}

func NewPaymentMethodLinkRepository() PaymentMethodLinkRepository {
	db, err := config.ConnectDB()
	if err != nil {
		panic("Failed to connect database")
	}
	// db.AutoMigrate(&entity.PaymentMethodLink{}, &entity.Person{})
	if (os.Getenv("DB_HOST_PAYMENT") != ""){
		db.AutoMigrate(&entity.PaymentMethodLink{})
	}else{
		db.AutoMigrate(&entity.PaymentMethodLinkTesting{})
	}
	return &paymentMethodLinkDatabase{
		connection: db,
	}
}

func (db *paymentMethodLinkDatabase) SavePaymentMethodLink(paymentMethodLink entity.PaymentMethodLink) (int, error) {
	data := &paymentMethodLink
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	err := db.connection.Create(data)
	if(err.Error != nil){
		return 0, err.Error
	}
	return data.ID, nil
}

func (db *paymentMethodLinkDatabase) UpdatePaymentMethodLink(paymentMethodLink entity.PaymentMethodLink) error {
	data := &paymentMethodLink
	data.UpdatedAt = time.Now()
	db.connection.Save(data)
	return nil
}

func (db *paymentMethodLinkDatabase) DeletePaymentMethodLink(paymentMethodLink entity.PaymentMethodLink) error {
	db.connection.Delete(&paymentMethodLink)
	return nil
}

func (db *paymentMethodLinkDatabase) GetAllPaymentMethodLinks() []entity.PaymentMethodLink {
	var paymentMethodLinks []entity.PaymentMethodLink
	db.connection.Preload(clause.Associations).Find(&paymentMethodLinks)
	return paymentMethodLinks
}

func (db *paymentMethodLinkDatabase) GetPaymentMethodLink(id string) []entity.PaymentMethodLink {
	var paymentMethodLink []entity.PaymentMethodLink
	db.connection.Preload(clause.Associations).Where("id = ?", id).First(&paymentMethodLink)
	return paymentMethodLink
}
func (db *paymentMethodLinkDatabase) GetPaymentMethodLinkByCode(code string) entity.PaymentMethodLink {
	var paymentMethodLink entity.PaymentMethodLink
	db.connection.Preload(clause.Associations).Where("code = ?", code).First(&paymentMethodLink)
	return paymentMethodLink
}
