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
		PaymentMethodCode 		: "visa-master",
		ProcessType	 			: "payment",
		Url 					: "https://superyou.co.id/pembayaran/creditcard/:creditcardtoken",
	},
	entity.PaymentMethodLink{
		PaymentMethodCode 		: "bca-klik-bca",
		ProcessType	 			: "payment",
		Url 					: "https://superyou.co.id/pembayaran-debit-va/:encrypteddata",
	},
	entity.PaymentMethodLink{
		PaymentMethodCode 		: "bca-klikpay",
		ProcessType	 			: "payment",
		Url 					: "https://superyou.co.id/pembayaran-debit-va/:encrypteddata",	
	},
	entity.PaymentMethodLink{
		PaymentMethodCode 		: "cimb-clicks",
		ProcessType	 			: "payment",
		Url 					: "https://superyou.co.id/pembayaran-debit-va/:encrypteddata",	
	},
	entity.PaymentMethodLink{
		PaymentMethodCode 		: "danamon-online-banking",
		ProcessType	 			: "payment",
		Url 					: "https://superyou.co.id/pembayaran-debit-va/:encrypteddata",	
	},
	entity.PaymentMethodLink{
		PaymentMethodCode 		: "mandiri-clickpay",
		ProcessType	 			: "payment",
		Url 					: "https://superyou.co.id/pembayaran-debit-va/:encrypteddata",	
	},
	entity.PaymentMethodLink{
		PaymentMethodCode 		: "mandiri-ecash",
		ProcessType	 			: "payment",
		Url 					: "https://superyou.co.id/pembayaran-debit-va/:encrypteddata",	
	},
	entity.PaymentMethodLink{
		PaymentMethodCode 		: "bri-epay",
		ProcessType	 			: "payment",
		Url 					: "https://superyou.co.id/pembayaran-debit-va/:encrypteddata",	
	},
	entity.PaymentMethodLink{
		PaymentMethodCode 		: "bri-mocash",
		ProcessType	 			: "payment",
		Url 					: "https://superyou.co.id/pembayaran-debit-va/:encrypteddata",	
	},
	entity.PaymentMethodLink{
		PaymentMethodCode 		: "sakuku-bca",
		ProcessType	 			: "payment",
		Url 					: "https://superyou.co.id/pembayaran-debit-va/:encrypteddata",	
	},
	entity.PaymentMethodLink{
		PaymentMethodCode 		: "va-bca",
		ProcessType	 			: "payment",
		Url 					: "https://superyou.co.id/pembayaran-debit-va/:encrypteddata",	
	},
	entity.PaymentMethodLink{
		PaymentMethodCode 		: "atm-mandiri",
		ProcessType	 			: "payment",
		Url 					: "https://superyou.co.id/pembayaran-debit-va/:encrypteddata",	
	},
	entity.PaymentMethodLink{
		PaymentMethodCode 		: "atm-bii",
		ProcessType	 			: "payment",
		Url 					: "https://superyou.co.id/pembayaran-debit-va/:encrypteddata",	
	},
	entity.PaymentMethodLink{
		PaymentMethodCode 		: "atm-permata",
		ProcessType	 			: "payment",
		Url 					: "https://superyou.co.id/pembayaran-debit-va/:encrypteddata",	
	},
	entity.PaymentMethodLink{
		PaymentMethodCode 		: "rekening-ponsel",
		ProcessType	 			: "payment",
		Url 					: "https://superyou.co.id/pembayaran-debit-va/:encrypteddata",	
	},
	entity.PaymentMethodLink{
		PaymentMethodCode 		: "t-cash-telkomsel",
		ProcessType	 			: "payment",
		Url 					: "https://superyou.co.id/pembayaran-debit-va/:encrypteddata",	
	},
	entity.PaymentMethodLink{
		PaymentMethodCode 		: "xl-tunai",
		ProcessType	 			: "payment",
		Url 					: "https://superyou.co.id/pembayaran-debit-va/:encrypteddata",	
	},
	entity.PaymentMethodLink{
		PaymentMethodCode 		: "dompetku-indosat",
		ProcessType	 			: "payment",
		Url 					: "https://superyou.co.id/pembayaran-debit-va/:encrypteddata",	
	},
	entity.PaymentMethodLink{
		PaymentMethodCode 		: "indomaret",
		ProcessType	 			: "payment",
		Url 					: "https://superyou.co.id/pembayaran-debit-va/:encrypteddata",	
	},
	entity.PaymentMethodLink{
		PaymentMethodCode 		: "mynt-artajasa",
		ProcessType	 			: "payment",
		Url 					: "https://superyou.co.id/pembayaran-debit-va/:encrypteddata",	
	},
	entity.PaymentMethodLink{
		PaymentMethodCode 		: "tfp-artajasa",
		ProcessType	 			: "payment",
		Url 					: "https://superyou.co.id/pembayaran-debit-va/:encrypteddata",	
	},
	entity.PaymentMethodLink{
		PaymentMethodCode 		: "gopay",
		ProcessType	 			: "linking",
		Url 					: "https://superyou.co.id/gopay-linking/:encrypteddata",	
	},
	entity.PaymentMethodLink{
		PaymentMethodCode 		: "gopay",
		ProcessType	 			: "payment",
		Url 					: "https://superyou.co.id/pembayaran-debit-va/:encrypteddata",	
	},
	entity.PaymentMethodLink{
		PaymentMethodCode 		: "auto-debet",
		ProcessType	 			: "payment",
		Url 					: "https://superyou.co.id/pembayaran-debit-va/:encrypteddata",	
	},
	entity.PaymentMethodLink{
		PaymentMethodCode 		: "beever-salary-advance",
		ProcessType	 			: "payment",
		Url 					: "https://superyou.co.id/pembayaran-debit-va/:encrypteddata",	
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