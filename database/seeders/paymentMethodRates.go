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
		PaymentMethodCode		: "indomaret",
		MinTransaction  		: 1,
		MaxTransaction  		: 500000,
		// TransactionFee  		: 3750.00,
		FormulaFee				: "3750.00",
		Premi  					: 0.00,
	},
	entity.PaymentMethodRate{
		PaymentMethodCode 		: "indomaret",
		MinTransaction  		: 500001,
		MaxTransaction  		: 1000000,
		// TransactionFee  		: 6500.00,
		FormulaFee				: "6500.00",
		Premi  					: 0.00,
	},
	entity.PaymentMethodRate{
		PaymentMethodCode 		: "indomaret",
		MinTransaction  		: 1000001,
		MaxTransaction  		: 5000000,
		// TransactionFee  		: 9250.00,
		FormulaFee				: "9250.00",
		Premi  					: 0.00,
	},
	entity.PaymentMethodRate{
		PaymentMethodCode 		: "visa-master",
		MinTransaction  		: 1,
		MaxTransaction  		: 10000000,
		// TransactionFee  		: 0.00,
		FormulaFee				: "0.00",
		Premi  					: 0.00,
	},
	entity.PaymentMethodRate{
		PaymentMethodCode 		: "bca-klik-bca",
		MinTransaction  		: 1,
		MaxTransaction  		: 10000000,
		// TransactionFee  		: 2000.00,
		FormulaFee				: "2000.00",
		Premi  					: 0.00,
	},
	entity.PaymentMethodRate{
		PaymentMethodCode 		: "bca-klikpay",
		MinTransaction  		: 1,
		MaxTransaction  		: 10000000,
		// TransactionFee  		: 1500.00,
		FormulaFee				: "1500.00",
		Premi  					: 0.00,
	},
	entity.PaymentMethodRate{
		PaymentMethodCode 		: "cimb-clicks",
		MinTransaction  		: 1,
		MaxTransaction  		: 10000000,
		// TransactionFee  		: 4700.00,
		FormulaFee				: "4700.00",
		Premi  					: 0.00,
	},
	entity.PaymentMethodRate{
		PaymentMethodCode 		: "danamon-online-banking",
		MinTransaction  		: 1,
		MaxTransaction  		: 10000000,
		// TransactionFee  		: 5000.00,
		FormulaFee				: "5000.00",
		Premi  					: 0.00,
	},
	entity.PaymentMethodRate{
		PaymentMethodCode 		: "mandiri-clickpay",
		MinTransaction  		: 1,
		MaxTransaction  		: 10000000,
		// TransactionFee  		: 4700.00,
		FormulaFee				: "4700.00",
		Premi  					: 0.00,
	},
	entity.PaymentMethodRate{
		PaymentMethodCode 		: "mandiri-ecash",
		MinTransaction  		: 1,
		MaxTransaction  		: 10000000,
		// TransactionFee  		: 0.00,
		FormulaFee				: "0.00",
		Premi  					: 0.00,
	},
	entity.PaymentMethodRate{
		PaymentMethodCode 		: "bri-epay",
		MinTransaction  		: 1,
		MaxTransaction  		: 10000000,
		// TransactionFee  		: 5000.00,
		FormulaFee				: "5000.00",
		Premi  					: 0.00,
	},
	entity.PaymentMethodRate{
		PaymentMethodCode 		: "bri-mocash",
		MinTransaction  		: 1,
		MaxTransaction  		: 10000000,
		// TransactionFee  		: 2500.00,
		FormulaFee				: "2500.00",
		Premi  					: 0.00,
	},
	entity.PaymentMethodRate{
		PaymentMethodCode 		: "sakuku-bca",
		MinTransaction  		: 1,
		MaxTransaction  		: 10000000,
		// TransactionFee  		: 500.00,
		FormulaFee				: "500.00",
		Premi  					: 0.00,
	},
	entity.PaymentMethodRate{
		PaymentMethodCode 		: "va-bca",
		MinTransaction  		: 1,
		MaxTransaction  		: 10000000,
		// TransactionFee  		: 2000.00,
		FormulaFee				: "2000.00",
		Premi  					: 0.00,
	},
	entity.PaymentMethodRate{
		PaymentMethodCode 		: "atm-mandiri",
		MinTransaction  		: 1,
		MaxTransaction  		: 10000000,
		// TransactionFee  		: 5000.00,
		FormulaFee				: "5000.00",
		Premi  					: 0.00,
	},
	entity.PaymentMethodRate{
		PaymentMethodCode 		: "atm-bii",
		MinTransaction  		: 1,
		MaxTransaction  		: 10000000,
		// TransactionFee  		: 4000.00,
		FormulaFee				: "4000.00",
		Premi  					: 0.00,
	},
	entity.PaymentMethodRate{
		PaymentMethodCode 		: "atm-permata",
		MinTransaction  		: 1,
		MaxTransaction  		: 10000000,
		// TransactionFee  		: 3000.00,
		FormulaFee				: "3000.00",
		Premi  					: 0.00,
	},
	entity.PaymentMethodRate{
		PaymentMethodCode 		: "rekening-ponsel",
		MinTransaction  		: 1,
		MaxTransaction  		: 10000000,
		// TransactionFee  		: 5000.00,
		FormulaFee				: "5000.00",
		Premi  					: 0.00,
	},
	entity.PaymentMethodRate{
		PaymentMethodCode 		: "t-cash-telkomsel",
		MinTransaction  		: 1,
		MaxTransaction  		: 10000000,
		// TransactionFee  		: 3000.00,
		FormulaFee				: "3000.00",
		Premi  					: 0.00,
	},
	entity.PaymentMethodRate{
		PaymentMethodCode 		: "xl-tunai",
		MinTransaction  		: 1,
		MaxTransaction  		: 10000000,
		// TransactionFee  		: 3000.00,
		FormulaFee				: "3000.00",
		Premi  					: 0.00,
	},
	entity.PaymentMethodRate{
		PaymentMethodCode 		: "dompetku-indosat",
		MinTransaction  		: 1,
		MaxTransaction  		: 10000000,
		// TransactionFee  		: 3000.00,
		FormulaFee				: "3000.00",
		Premi  					: 0.00,
	},
	entity.PaymentMethodRate{
		PaymentMethodCode 		: "mynt-artajasa",
		MinTransaction  		: 1,
		MaxTransaction  		: 10000000,
		// TransactionFee  		: 3000.00,
		FormulaFee				: "3000.00",
		Premi  					: 0.00,
	},
	entity.PaymentMethodRate{
		PaymentMethodCode 		: "tfp-artajasa",
		MinTransaction  		: 1,
		MaxTransaction  		: 10000000,
		// TransactionFee  		: 5000.00,
		FormulaFee				: "5000.00",
		Premi  					: 0.00,
	},
	entity.PaymentMethodRate{
		PaymentMethodCode 		: "gopay",
		MinTransaction  		: 1,
		MaxTransaction  		: 10000000,
		// TransactionFee  		: 0.00,
		FormulaFee				: "(premi * 100/98) - premi",
		Premi  					: 0.00,
	},
	entity.PaymentMethodRate{
		PaymentMethodCode 		: "auto-debet",
		MinTransaction  		: 1,
		MaxTransaction  		: 10000000,
		// TransactionFee  		: 0.00,
		FormulaFee				: "1",
		Premi  					: 0.00,
	},
	entity.PaymentMethodRate{
		PaymentMethodCode 		: "beever-salary-advance",
		MinTransaction  		: 1,
		MaxTransaction  		: 10000000,
		// TransactionFee  		: 3000.00,
		FormulaFee				: "3000.00",
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