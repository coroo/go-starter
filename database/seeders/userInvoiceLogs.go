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
var userInvoiceLog = []entity.UserInvoiceLog{
	entity.UserInvoiceLog{
		ID               		: 1,
		PaymentMethodCode 		: "indomaret-test",
		SummaryToken  			: "abcd1084ufnouu23af",
		// TransactionFee  		: 6500.00,
		InvoiceNumber			: "12369412964",
		PaymentCycle  			: "YEARLY",
		TransactionFee  		: "20000",
		AgentFee  				: "20000",
		TotalPremium  			: "40000",
		TotalPayment  			: "80000",
		CreatedAt				: time.Now(),
		UpdatedAt				: time.Now(),
	},
	entity.UserInvoiceLog{
		ID               		: 1,
		PaymentMethodCode 		: "visa-master-test",
		SummaryToken  			: "abcd1084ufnouu23af",
		// TransactionFee  		: 6500.00,
		InvoiceNumber			: "12369412964",
		PaymentCycle  			: "YEARLY",
		TransactionFee  		: "20000",
		AgentFee  				: "20000",
		TotalPremium  			: "40000",
		TotalPayment  			: "80000",
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

func SeedUserInvoiceLogs() {
	db, _ := config.ConnectDB()

	_ = db.AutoMigrate(&entity.UserInvoiceLog{})
	for i, _ := range userInvoiceLog {
		err := db.Model(&entity.UserInvoiceLog{}).Create(&userInvoiceLog[i])
		if(err.Error != nil){
			fmt.Println("cannot seed 'Payment Method' table: ", err)
		}
	}
}