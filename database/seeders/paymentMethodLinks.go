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
var paymentMethodLink = []entity.PaymentMethodLink{
	entity.PaymentMethodLink{
		PaymentMethodCode 		: "gopay",
		ProcessType	 			: "linking",
		Url 					: "https://www.superyou.com/gopay-linking/:encryptedaccountid",
	},
	entity.PaymentMethodLink{
		PaymentMethodCode 		: "gopay",
		ProcessType	 			: "payment",
		Url 					: "https://gopay.payment.link/:accountid",
	},
	entity.PaymentMethodLink{
		PaymentMethodCode 		: "visa-master",
		ProcessType	 			: "payment",
		Url 					: "https://visa.payment.link/:accountid",	
	},
	entity.PaymentMethodLink{
		PaymentMethodCode 		: "indomaret",
		ProcessType	 			: "payment",
		Url 					: "https://midtrans.payment.link/:accountid",	
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

func SeedPaymentMethodLinks() {
	db, _ := config.ConnectDB()

	_ = db.AutoMigrate(&entity.PaymentMethodLink{})
	for i, _ := range paymentMethodLink {
		err := db.Model(&entity.PaymentMethodLink{}).Create(&paymentMethodLink[i])
		if(err.Error != nil){
			fmt.Println("cannot seed 'Payment Method Link' table: ", err)
		}
	}
}