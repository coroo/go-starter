package repositories

import (
	"os"
	"github.com/coroo/go-starter/config"
	entity "github.com/coroo/go-starter/app/entity"

	// "github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	_ "gorm.io/driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

type LumpSumPaymentRepository interface {
	CreateLumSumPayment(lumpSumPayment entity.LumpSumPayment)
	UpdateLumSumPayment(lumpSumPayment entity.LumpSumPayment)
	DeleteLumSumPayment(lumpSumPayment entity.LumpSumPayment)
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
	db.AutoMigrate(&entity.LumpSumPayment{})
	db.Migrator().DropColumn(&entity.LumpSumPayment{}, "first_effective_date")
	return &database{
		connection: db,
	}
}

func (db *database) CloseDB() {
	sqlDB,err := db.connection.DB()
	if err != nil {
		panic("Failed to close database")
	}
	closeDB := sqlDB.Close()
	if closeDB != nil {
		panic("Failed to close database")
	}
}

func (db *database) CreateLumSumPayment(lumpSumPayment entity.LumpSumPayment) {
	if os.Getenv("DB_TEST") == ""{
		db.connection.Create(&lumpSumPayment)
	}else if err := db.connection.Where("policy_number = ?", &lumpSumPayment.PolicyNumber).First(&lumpSumPayment).Error; err != nil {
		// error handling...
		db.connection.Create(&lumpSumPayment)
	}
}

func (db *database) UpdateLumSumPayment(lumpSumPayment entity.LumpSumPayment) {
	db.connection.Save(&lumpSumPayment)
}

func (db *database) DeleteLumSumPayment(lumpSumPayment entity.LumpSumPayment) {
	db.connection.Delete(&lumpSumPayment)
}

func (db *database) GetAllLatestGroupLumpSumPayments() []entity.LumpSumPayment {
	var lumpSumPaymentsGroup []entity.LumpSumPayment
	subQuery := db.connection.
		Select("policy_number, min(effective_date) as first_effective_date").
		Table("lump_sum_payments").
		Group("policy_number")
	db.connection.Preload(clause.Associations).Select("lump_sum_payments.*, t1.first_effective_date").Joins("LEFT JOIN (?) AS t1 ON t1.policy_number = lump_sum_payments.policy_number", subQuery).Where("(lump_sum_payments.policy_number, effective_date) IN (?)", db.connection.Table("lump_sum_payments").Select("policy_number, max(effective_date) as effective_date").Group("policy_number")).Find(&lumpSumPaymentsGroup)
	return lumpSumPaymentsGroup
}

func (db *database) GetAllLumpSumPayments() []entity.LumpSumPayment {
	var lumpSumPayments []entity.LumpSumPayment
	db.connection.Preload(clause.Associations).Find(&lumpSumPayments)
	return lumpSumPayments
}

func (db *database) GetLumpSumPayment(policyNumber string) []entity.LumpSumPayment {
	var lumpSumPayment []entity.LumpSumPayment
	db.connection.Preload(clause.Associations).Where("(`policy_number`, `effective_date`) IN (?)", db.connection.Table("lump_sum_payments").Select("`policy_number`, max(`effective_date`) as `effective_date`").Where("`policy_number` = ?", policyNumber).Group("policy_number")).First(&lumpSumPayment)
	return lumpSumPayment
}
