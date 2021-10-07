package seeder

import (
	// "time"
	// "log"
	"fmt"
	// "github.com/google/uuid"

	// utils "github.com/coroo/go-starter/app/utils"
	"github.com/coroo/go-starter/config"
	entity "github.com/coroo/go-starter/app/entity"
)
var paymentMethodRate = []entity.PaymentMethodRate{
	entity.PaymentMethodRate{
		PaymentMethodId 		: 19,
		MinTransaction  		: 1,
		MaxTransaction  		: 500000,
		TransactionFee  		: 3750.00,
		Premi  					: 0.00,
	},
	entity.PaymentMethodRate{
		PaymentMethodId 		: 19,
		MinTransaction  		: 500001,
		MaxTransaction  		: 1000000,
		TransactionFee  		: 6500.00,
		Premi  					: 0.00,
	},
	entity.PaymentMethodRate{
		PaymentMethodId 		: 19,
		MinTransaction  		: 1000001,
		MaxTransaction  		: 5000000,
		TransactionFee  		: 9250.00,
		Premi  					: 0.00,
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

func SeedPaymentMethodRates() {
	db, _ := config.ConnectDB()

	_ = db.AutoMigrate(&entity.PaymentMethodRate{})
	for i, _ := range paymentMethodRate {
		err := db.Model(&entity.PaymentMethodRate{}).Create(&paymentMethodRate[i])
		if(err.Error != nil){
			fmt.Println("cannot seed 'Payment Method' table: ", err)
		}
	}
}