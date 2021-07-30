package seeder

import (
	"log"
	"github.com/google/uuid"

	utils "github.com/coroo/go-starter/app/utils"
	"github.com/coroo/go-starter/config"
	entity "github.com/coroo/go-starter/app/entity"
)

var defaultPass, _ = utils.HashPassword("secretsekali")

var users = []entity.User{
	entity.User{
		Name	: "Kuncoro Wicaksono",
		Uuid	: uuid.New().String(),
		Email	: "coroo.wicaksono@gmail.com",
		Password: defaultPass,
	},
	entity.User{
		Name	: "Peter James",
		Uuid	: uuid.New().String(),
		Email	: "pjw.peterjames@gmail.com",
		Password: defaultPass,
	},
}

// var posts = []entity.Post{
// 	entity.Post{
// 		Title:   "Title 1",
// 		Content: "Hello world 1",
// 	},
// 	entity.Post{
// 		Title:   "Title 2",
// 		Content: "Hello world 2",
// 	},
// }

// func Load(db *gorm.DB) {
func SeedUsers() {
	db, _ := config.ConnectDB()

	_ = db.AutoMigrate(&entity.User{})
	for i, _ := range users {
		err := db.Model(&entity.User{}).Create(&users[i])
		if(err.Error != nil){
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}