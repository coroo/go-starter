package repositories

import (
	"github.com/coroo/go-starter/config"
	entity "github.com/coroo/go-starter/app/entity"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/joho/godotenv/autoload"
)

type SyOdsEtlPaymentRepository interface {
	CreateSyOdsEtlPayment(syOdsEtlPayment entity.SyOdsEtlPayment)
	Update(syOdsEtlPayment entity.SyOdsEtlPayment)
	Delete(syOdsEtlPayment entity.SyOdsEtlPayment)
	GetAllSyOdsEtlPayments() []entity.SyOdsEtlPayment
	GetAllLatestGroupSyOdsEtlPayments() []entity.SyOdsEtlPayment
	GetSyOdsEtlPaymentByStatus(status string) []entity.SyOdsEtlPayment
	GetSyOdsEtlPaymentDailyByStatus(status string) []entity.SyOdsEtlPayment
	GetSyOdsEtlPaymentByPolicyNumber(ctx *gin.Context) []entity.SyOdsEtlPayment
	CancelOutstandingSyOdsEtlPayments() []entity.SyOdsEtlPayment
	CloseDB()
}

type syOdsEtlPaymentDatabase struct {
	connection *gorm.DB
}

func NewSyOdsEtlPaymentRepository() SyOdsEtlPaymentRepository {
	db, err := config.ConnectDB()
	if err != nil {
		panic("Failed to connect syOdsEtlPaymentDatabase")
	}
	// db.AutoMigrate(&entity.SyOdsEtlPayment{}, &entity.Person{})
	db.AutoMigrate(&entity.SyOdsEtlPayment{})
	return &syOdsEtlPaymentDatabase{
		connection: db,
	}
}

func (db *syOdsEtlPaymentDatabase) CloseDB() {
	err := db.connection.Close()
	if err != nil {
		panic("Failed to close syOdsEtlPaymentDatabase")
	}
}

func (db *syOdsEtlPaymentDatabase) CreateSyOdsEtlPayment(syOdsEtlPayment entity.SyOdsEtlPayment) {
	data := &syOdsEtlPayment
	data.UpdatedAt = time.Now()
	if err := db.connection.Where("policy_number = ?", data.PolicyNumber).First(&data).Error; err != nil {
		// error handling...
		db.connection.Create(data)
	}
}

func (db *syOdsEtlPaymentDatabase) Update(syOdsEtlPayment entity.SyOdsEtlPayment) {
	db.connection.Save(&syOdsEtlPayment)
}

func (db *syOdsEtlPaymentDatabase) Delete(syOdsEtlPayment entity.SyOdsEtlPayment) {
	db.connection.Delete(&syOdsEtlPayment)
}

func (db *syOdsEtlPaymentDatabase) GetAllLatestGroupSyOdsEtlPayments() []entity.SyOdsEtlPayment {
	var syOdsEtlPaymentsGroup []entity.SyOdsEtlPayment
	db.connection.Set("gorm:auto_preload", true).Find(&syOdsEtlPaymentsGroup)
	return syOdsEtlPaymentsGroup
}

func (db *syOdsEtlPaymentDatabase) GetAllSyOdsEtlPayments() []entity.SyOdsEtlPayment {
	var syOdsEtlPayments []entity.SyOdsEtlPayment
	db.connection.Set("gorm:auto_preload", true).Find(&syOdsEtlPayments)
	return syOdsEtlPayments
}

func (db *syOdsEtlPaymentDatabase) GetSyOdsEtlPaymentByPolicyNumber(ctx *gin.Context) []entity.SyOdsEtlPayment {
	var syOdsEtlPayment []entity.SyOdsEtlPayment
	db.connection.Set("gorm:auto_preload", true).Where("policy_number = ?", ctx.Param("policyNumber")).Order("id desc").Find(&syOdsEtlPayment)
	return syOdsEtlPayment
}

func (db *syOdsEtlPaymentDatabase) GetSyOdsEtlPaymentByStatus(status string) []entity.SyOdsEtlPayment {
	var syOdsEtlPayment []entity.SyOdsEtlPayment
	db.connection.Set("gorm:auto_preload", true).Where("status = ?", status).Find(&syOdsEtlPayment)
	return syOdsEtlPayment
}

func (db *syOdsEtlPaymentDatabase) GetSyOdsEtlPaymentDailyByStatus(status string) []entity.SyOdsEtlPayment {
	var syOdsEtlPayment []entity.SyOdsEtlPayment
	db.connection.Set("gorm:auto_preload", true).Where("status = ? AND DATE(updated_at) = ?", status, time.Now().Format("2006-02-01")).Find(&syOdsEtlPayment)
	return syOdsEtlPayment
}

func (db *syOdsEtlPaymentDatabase) CancelOutstandingSyOdsEtlPayments() []entity.SyOdsEtlPayment {
	var syOdsEtlPayment []entity.SyOdsEtlPayment
	db.connection.Model(&syOdsEtlPayment).Where("status = ?", "queue").UpdateColumn("status", "cancel")
	return syOdsEtlPayment
}
