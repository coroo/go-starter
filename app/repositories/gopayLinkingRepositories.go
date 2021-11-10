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

type GopayLinkingRepository interface {
	SaveGopayLinking(paymentMethod entity.GopayLinking) (int, error)
	UpdateGopayLinking(paymentMethod entity.GopayLinking) error
	DeleteGopayLinking(paymentMethod entity.GopayLinking) error
	GetAllGopayLinkings() []entity.GopayLinking
	GetGopayLinking(id string) []entity.GopayLinking
}

type gopayLinkingDatabase struct {
	connection *gorm.DB
}

func NewGopayLinkingRepository() GopayLinkingRepository {
	db, err := config.ConnectDB()
	if err != nil {
		panic("Failed to connect database")
	}
	// db.AutoMigrate(&entity.GopayLinkingRate{}, &entity.Person{})
	
	if (os.Getenv("DB_HOST_PAYMENT") != ""){
		db.AutoMigrate(&entity.GopayLinking{})
	} else {
		db.AutoMigrate(&entity.GopayLinkingTesting{})
	}
	return &gopayLinkingDatabase{
		connection: db,
	}
}

func (db *gopayLinkingDatabase) SaveGopayLinking(gopayLinking entity.GopayLinking) (int, error) {
	data := &gopayLinking
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	err := db.connection.Create(data)
	if(err.Error != nil){
		return 0, err.Error
	}
	return data.ID, nil
}

func (db *gopayLinkingDatabase) UpdateGopayLinking(gopayLinking entity.GopayLinking) error {
	data := &gopayLinking
	data.UpdatedAt = time.Now()
	db.connection.Save(data)
	return nil
}

func (db *gopayLinkingDatabase) DeleteGopayLinking(gopayLinking entity.GopayLinking) error {
	db.connection.Delete(&gopayLinking)
	return nil
}

func (db *gopayLinkingDatabase) GetAllGopayLinkings() []entity.GopayLinking {
	var gopayLinkings []entity.GopayLinking
	db.connection.Preload(clause.Associations).Find(&gopayLinkings)
	return gopayLinkings
}

func (db *gopayLinkingDatabase) GetGopayLinking(id string) []entity.GopayLinking {
	var gopayLinking []entity.GopayLinking
	db.connection.Preload(clause.Associations).Where("id = ?", id).First(&gopayLinking)
	return gopayLinking
}
