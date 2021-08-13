package repositories

import (
	"github.com/coroo/go-starter/config"
	entity "github.com/coroo/go-starter/app/entity"

	// "github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/joho/godotenv/autoload"
)

type LumpSumPaymentRepository interface {
	Save(lumpSumPayment entity.LumpSumPayment)
	Update(lumpSumPayment entity.LumpSumPayment)
	Delete(lumpSumPayment entity.LumpSumPayment)
	GetAllLumpSumPayments() []entity.LumpSumPayment
	GetAllLatestGroupLumpSumPayments() []entity.LumpSumPayment
	GetLumpSumPayment(policyNumber string) []entity.LumpSumPayment
	CloseDB()
}

type database struct {
	connection *gorm.DB
}

func NewLumpSumPaymentRepository() LumpSumPaymentRepository {
	db, err := config.ConnectDB()
	if err != nil {
		panic("Failed to connect database")
	}
	// db.AutoMigrate(&entity.LumpSumPayment{}, &entity.Person{})
	db.AutoMigrate(&entity.LumpSumPayment{}).DropColumn("first_effective_date")
	return &database{
		connection: db,
	}
}

func (db *database) CloseDB() {
	err := db.connection.Close()
	if err != nil {
		panic("Failed to close database")
	}
}

func (db *database) Save(lumpSumPayment entity.LumpSumPayment) {
	if err := db.connection.Where("policy_number = ?", &lumpSumPayment.PolicyNumber).First(&lumpSumPayment).Error; err != nil {
		// error handling...
		db.connection.Create(&lumpSumPayment)
	}
}

func (db *database) Update(lumpSumPayment entity.LumpSumPayment) {
	db.connection.Save(&lumpSumPayment)
}

func (db *database) Delete(lumpSumPayment entity.LumpSumPayment) {
	db.connection.Delete(&lumpSumPayment)
}

func (db *database) GetAllLatestGroupLumpSumPayments() []entity.LumpSumPayment {
	var lumpSumPaymentsGroup []entity.LumpSumPayment
	subQuery := db.connection.
		Select("policy_number, min(effective_date) as first_effective_date").
		Table("lump_sum_payments").
		Group("policy_number").
		SubQuery()
	db.connection.Set("gorm:auto_preload", true).Select("lump_sum_payments.*, t1.first_effective_date").Joins("LEFT JOIN ? AS t1 ON t1.policy_number = lump_sum_payments.policy_number", subQuery).Where("(lump_sum_payments.policy_number, effective_date) IN ?", db.connection.Table("lump_sum_payments").Select("policy_number, max(effective_date) as effective_date").Group("policy_number").SubQuery()).Find(&lumpSumPaymentsGroup)
	return lumpSumPaymentsGroup
}

func (db *database) GetAllLumpSumPayments() []entity.LumpSumPayment {
	var lumpSumPayments []entity.LumpSumPayment
	db.connection.Set("gorm:auto_preload", true).Find(&lumpSumPayments)
	return lumpSumPayments
}

func (db *database) GetLumpSumPayment(policyNumber string) []entity.LumpSumPayment {
	var lumpSumPayment []entity.LumpSumPayment
	db.connection.Set("gorm:auto_preload", true).Where("(policy_number, effective_date) IN ?", db.connection.Table("lump_sum_payments").Select("policy_number, max(effective_date) as effective_date").Where("policy_number = ?", policyNumber).Group("policy_number").SubQuery()).First(&lumpSumPayment)
	return lumpSumPayment
}
