package repositories

import (
	"github.com/coroo/go-starter/config"
	entity "github.com/coroo/go-starter/app/entity"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/joho/godotenv/autoload"
)

type SyUserInvoiceRepository interface {
	Save(syUserInvoice entity.SyUserInvoice)
	Update(syUserInvoice entity.SyUserInvoice)
	Delete(syUserInvoice entity.SyUserInvoice)
	GetAllPaidUserInvoices() []entity.SyUserInvoice
	GetUserInvoice(ctx *gin.Context) []entity.SyUserInvoice
	CloseDB()
}

type syUserInvoiceDatabase struct {
	connection *gorm.DB
}

func NewSyUserInvoiceRepository() SyUserInvoiceRepository {
	db, err := config.ConnectDBSY()
	if err != nil {
		panic("Failed to connect database")
	}
	// db.AutoMigrate(&entity.SyUserInvoice{}, &entity.Person{})
	// db.AutoMigrate(&entity.SyUserInvoice{})
	return &syUserInvoiceDatabase{
		connection: db,
	}
}

func (db *syUserInvoiceDatabase) CloseDB() {
	err := db.connection.Close()
	if err != nil {
		panic("Failed to close database")
	}
}

func (db *syUserInvoiceDatabase) Save(syUserInvoice entity.SyUserInvoice) {
	db.connection.Create(&syUserInvoice)
}

func (db *syUserInvoiceDatabase) Update(syUserInvoice entity.SyUserInvoice) {
	db.connection.Save(&syUserInvoice)
}

func (db *syUserInvoiceDatabase) Delete(syUserInvoice entity.SyUserInvoice) {
	db.connection.Delete(&syUserInvoice)
}

func (db *syUserInvoiceDatabase) GetAllPaidUserInvoices() []entity.SyUserInvoice {
	var syUserInvoices []entity.SyUserInvoice
	db.connection.Set("gorm:auto_preload", true).Select("*, payment_methods.name as payment_method_name").Where("(user_invoices.policy_group_number, paid_at) IN ?", db.connection.Table("user_invoices").Select("user_invoices.policy_group_number, max(paid_at) as paid_at").Group("user_invoices.policy_group_number").SubQuery()).Joins("left join user_policies on user_policies.policy_group_number = user_invoices.policy_group_number").Joins("left join payment_methods on payment_methods.id = user_invoices.payment_method_id").Where("policy_number IS NOT NULL").Find(&syUserInvoices)
	return syUserInvoices
}

func (db *syUserInvoiceDatabase) GetUserInvoice(ctx *gin.Context) []entity.SyUserInvoice {
	var syUserInvoice []entity.SyUserInvoice
	db.connection.Set("gorm:auto_preload", true).First(&syUserInvoice, ctx.Param("id"))
	return syUserInvoice
}
