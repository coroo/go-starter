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

type SyEtlPaymentRepository interface {
	CreateSyEtlPayment(syEtlPayment entity.SyEtlPayment)
	UpdateSyEtlPayment(syEtlPayment entity.SyEtlPayment)
	DeleteSyEtlPayment(syEtlPayment entity.SyEtlPayment)
	GetAllSyEtlPayments() []entity.SyEtlPayment
	GetAllLatestGroupSyEtlPayments() []entity.SyEtlPayment
	GetSyEtlPayment(policyNumber string) []entity.SyEtlPayment
	TruncateTableSyEtlPayments()
	CloseDB()
}

type syEtlPaymentDatabase struct {
	connection *gorm.DB
}

func NewSyEtlPaymentRepository() SyEtlPaymentRepository {
	db, err := config.ConnectDB()
	if err != nil {
		panic("Failed to connect syEtlPaymentDatabase")
	}
	// db.AutoMigrate(&entity.SyEtlPayment{}, &entity.Person{})
	db.AutoMigrate(&entity.SyEtlPayment{})
	return &syEtlPaymentDatabase{
		connection: db,
	}
}

func (db *syEtlPaymentDatabase) CloseDB() {
	sqlDB,err := db.connection.DB()
	if err != nil {
		panic("Failed to close database")
	}
	closeDB := sqlDB.Close()
	if closeDB != nil {
		panic("Failed to close database")
	}
}

func (db *syEtlPaymentDatabase) CreateSyEtlPayment(syEtlPayment entity.SyEtlPayment) {
	data := &syEtlPayment
	data.UpdatedAt = time.Now()
	if os.Getenv("DB_TEST") == ""{
		db.connection.Create(data)
	}else if err := db.connection.Where("policy_number = ?", data.OdsPolicyNumber).First(&data).Error; err != nil {
		db.connection.Create(data)
	}
}

func (db *syEtlPaymentDatabase) UpdateSyEtlPayment(syEtlPayment entity.SyEtlPayment) {
	db.connection.Save(&syEtlPayment)
}

func (db *syEtlPaymentDatabase) DeleteSyEtlPayment(syEtlPayment entity.SyEtlPayment) {
	db.connection.Delete(&syEtlPayment)
}

func (db *syEtlPaymentDatabase) GetAllLatestGroupSyEtlPayments() []entity.SyEtlPayment {
	var syEtlPaymentsGroup []entity.SyEtlPayment
	db.connection.Preload(clause.Associations).Find(&syEtlPaymentsGroup)
	return syEtlPaymentsGroup
}

func (db *syEtlPaymentDatabase) GetAllSyEtlPayments() []entity.SyEtlPayment {
	var syEtlPayments []entity.SyEtlPayment
	db.connection.Preload(clause.Associations).Select("*").Joins("left join etl_ods_payments on etl_ods_payments.policy_number = etl_sy_payments.policy_number").Find(&syEtlPayments)
	return syEtlPayments
}

func (db *syEtlPaymentDatabase) GetSyEtlPayment(policyNumber string) []entity.SyEtlPayment {
	var syEtlPayment []entity.SyEtlPayment
	db.connection.Preload(clause.Associations).Where("policy_number = ?", policyNumber).First(&syEtlPayment)
	return syEtlPayment
}

func (db *syEtlPaymentDatabase) TruncateTableSyEtlPayments() {
	if os.Getenv("DB_TEST") != "" {
		db.connection.Exec("DELETE FROM `etl_sy_payments`;")
	} else {
		db.connection.Exec("TRUNCATE `etl_sy_payments`;")
	}
}
