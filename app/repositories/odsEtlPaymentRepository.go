package repositories

import (
	"os"
	"github.com/coroo/go-starter/config"
	entity "github.com/coroo/go-starter/app/entity"
	"time"

	// "github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	_ "gorm.io/driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

type OdsEtlPaymentRepository interface {
	CreateOdsEtlPayment(odsEtlPayment entity.OdsEtlPayment)
	UpdateOdsEtlPayment(odsEtlPayment entity.OdsEtlPayment)
	DeleteOdsEtlPayment(odsEtlPayment entity.OdsEtlPayment)
	GetAllOdsEtlPayments() []entity.OdsEtlPayment
	GetAllLatestGroupOdsEtlPayments() []entity.OdsEtlPayment
	GetOdsEtlPayment(id string) []entity.OdsEtlPayment
	TruncateTableOdsEtlPayments()
	CloseDB()
}

type odsEtlPaymentDatabase struct {
	connection *gorm.DB
}

func NewOdsEtlPaymentRepository() OdsEtlPaymentRepository {
	db, err := config.ConnectDB()
	if err != nil {
		panic("Failed to connect odsEtlPaymentDatabase")
	}
	db.AutoMigrate(&entity.OdsEtlPayment{})
	return &odsEtlPaymentDatabase{
		connection: db,
	}
}

func (db *odsEtlPaymentDatabase) CloseDB() {
	sqlDB,err := db.connection.DB()
	if err != nil {
		panic("Failed to close database")
	}
	closeDB := sqlDB.Close()
	if closeDB != nil {
		panic("Failed to close database")
	}
}

func (db *odsEtlPaymentDatabase) CreateOdsEtlPayment(odsEtlPayment entity.OdsEtlPayment) {
	data := &odsEtlPayment
	data.UpdatedAt = time.Now()
	if os.Getenv("DB_TEST") == ""{
		db.connection.Create(data)
	}else if err := db.connection.Where("policy_number = ?", data.PolicyNumber).First(&data).Error; err != nil {
		// error handling...
		db.connection.Create(data)
	}
}

func (db *odsEtlPaymentDatabase) UpdateOdsEtlPayment(odsEtlPayment entity.OdsEtlPayment) {
	db.connection.Save(&odsEtlPayment)
}

func (db *odsEtlPaymentDatabase) DeleteOdsEtlPayment(odsEtlPayment entity.OdsEtlPayment) {
	db.connection.Delete(&odsEtlPayment)
}

func (db *odsEtlPaymentDatabase) GetAllLatestGroupOdsEtlPayments() []entity.OdsEtlPayment {
	var odsEtlPaymentsGroup []entity.OdsEtlPayment
	db.connection.Preload(clause.Associations).Find(&odsEtlPaymentsGroup)
	return odsEtlPaymentsGroup
}

func (db *odsEtlPaymentDatabase) GetAllOdsEtlPayments() []entity.OdsEtlPayment {
	var odsEtlPayments []entity.OdsEtlPayment
	db.connection.Preload(clause.Associations).Find(&odsEtlPayments)
	return odsEtlPayments
}

func (db *odsEtlPaymentDatabase) GetOdsEtlPayment(id string) []entity.OdsEtlPayment {
	var odsEtlPayment []entity.OdsEtlPayment
	db.connection.Preload(clause.Associations).First(&odsEtlPayment, id)
	return odsEtlPayment
}

func (db *odsEtlPaymentDatabase) TruncateTableOdsEtlPayments() {
	if os.Getenv("DB_TEST") != "" {
		db.connection.Exec("DELETE FROM `etl_ods_payments`;")
	} else {
		db.connection.Exec("TRUNCATE `etl_ods_payments`;")
	}
}
