package seeder

import (
	"time"
	// "log"
	"fmt"
	// "github.com/google/uuid"

	// utils "github.com/coroo/go-starter/app/utils"
	"github.com/coroo/go-starter/config"
	entity "github.com/coroo/go-starter/app/entity"
)
var gopayLinking = []entity.GopayLinking{
	entity.GopayLinking{
		ID				        : 1,
		PhoneNumber				: "08123456789",
		AccountId		  		: "ojbqegnlfdam",
		PaymentOptionToken  	: "bnafjlnfasn",
		CreatedAt				: time.Now(),
		UpdatedAt				: time.Now(),
	},
	entity.GopayLinking{
		ID				        : 2,
		PhoneNumber				: "08987654321",
		AccountId		  		: "pajlakns",
		PaymentOptionToken  	: "paiisdanfoag",
		CreatedAt				: time.Now(),
		UpdatedAt				: time.Now(),
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

func SeedGopayLinkings() {
	db, _ := config.ConnectDB()

	_ = db.AutoMigrate(&entity.GopayLinking{})
	for i, _ := range gopayLinking {
		err := db.Model(&entity.GopayLinking{}).Create(&gopayLinking[i])
		if(err.Error != nil){
			fmt.Println("cannot seed 'Payment Method' table: ", err)
		}
	}
}