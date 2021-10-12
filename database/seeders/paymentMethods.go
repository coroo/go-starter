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
var paymentMethod = []entity.PaymentMethod{
	entity.PaymentMethod{
		Code 					: "visa-master",
		InitPaymentCode  		: "PDCC",
		RenewalPaymentCode 		: "CASH",
		FastpayCode 		 	: "cc",
		// BankCode 		 		: nil,
		Name 		 			: "Visa Master",
		PaymentLogo 		 	: "visa-master.png",
		// RateApplied 		 	: false,
		// TransactionFee 		 	: 0.00,
		// PercentageFee 		 	: 0,
		Status 		 			: "active",
		// CreatedAt        		: time.Now(),
		// UpdatedAt        		: time.Now(),
	},
	entity.PaymentMethod{
		Code 					: "bca-klik-bca",
		// InitPaymentCode  		: "PDCC",
		// RenewalPaymentCode 		: "CASH",
		FastpayCode 		 	: "",
		// BankCode 		 		: nil,
		Name 		 			: "BCA Klik-BCA",
		PaymentLogo 		 	: "bca-klik-bca.png",
		// RateApplied 		 	: false,
		// TransactionFee 		 	: 2000.00,
		// PercentageFee 		 	: 0,
		Status 		 			: "inactive",
		// CreatedAt        		: time.Now(),
		// UpdatedAt        		: time.Now(),
	},
	entity.PaymentMethod{
		Code 					: "bca-klikpay",
		// InitPaymentCode  		: "PDCC",
		// RenewalPaymentCode 		: "CASH",
		FastpayCode 		 	: "",
		// BankCode 		 		: nil,
		Name 		 			: "BCA KlikPay",
		PaymentLogo 		 	: "bca-klikpay.png",
		// RateApplied 		 	: false,
		// TransactionFee 		 	: 1500.00,
		// PercentageFee 		 	: 0,
		Status 		 			: "inactive",
		// CreatedAt        		: time.Now(),
		// UpdatedAt        		: time.Now(),
	},
	entity.PaymentMethod{
		Code 					: "cimb-clicks",
		InitPaymentCode  		: "PDDC1",
		RenewalPaymentCode 		: "PDDC1",
		FastpayCode 		 	: "700",
		// BankCode 		 		: nil,
		Name 		 			: "CIMB Clicks",
		PaymentLogo 		 	: "cimb-clicks.png",
		// RateApplied 		 	: false,
		// TransactionFee 		 	: 4700.00,
		// PercentageFee 		 	: 0,
		Status 		 			: "inactive",
		// CreatedAt        		: time.Now(),
		// UpdatedAt        		: time.Now(),
	},
	entity.PaymentMethod{
		Code 					: "danamon-online-banking",
		// InitPaymentCode  		: "PDDC1",
		// RenewalPaymentCode 		: "PDDC1",
		FastpayCode 		 	: "",
		// BankCode 		 		: nil,
		Name 		 			: "Danamon Online Banking",
		PaymentLogo 		 	: "danamon-online-banking.png",
		// RateApplied 		 	: false,
		// TransactionFee 		 	: 5000.00,
		// PercentageFee 		 	: 0,
		Status 		 			: "inactive",
		// CreatedAt        		: time.Now(),
		// UpdatedAt        		: time.Now(),
	},
	entity.PaymentMethod{
		Code 					: "mandiri-clickpay",
		InitPaymentCode  		: "PDDC5",
		RenewalPaymentCode 		: "PDDC5",
		FastpayCode 		 	: "406",
		// BankCode 		 		: nil,
		Name 		 			: "Mandiri Clickpay",
		PaymentLogo 		 	: "mandiri-clickpay.png",
		// RateApplied 		 	: false,
		// TransactionFee 		 	: 4700.00,
		// PercentageFee 		 	: 0,
		Status 		 			: "inactive",
		// CreatedAt        		: time.Now(),
		// UpdatedAt        		: time.Now(),
	},
	entity.PaymentMethod{
		Code 					: "mandiri-ecash",
		// InitPaymentCode  		: "PDDC5",
		// RenewalPaymentCode 		: "PDDC5",
		FastpayCode 		 	: "",
		// BankCode 		 		: nil,
		Name 		 			: "Mandiri ECash",
		PaymentLogo 		 	: "mandiri-ecash.png",
		// RateApplied 		 	: false,
		// TransactionFee 		 	: 0.00,
		// PercentageFee 		 	: 1,
		Status 		 			: "inactive",
		// CreatedAt        		: time.Now(),
		// UpdatedAt        		: time.Now(),
	},
	entity.PaymentMethod{
		Code 					: "bri-epay",
		InitPaymentCode  		: "PDDC2",
		RenewalPaymentCode 		: "PDDC2",
		FastpayCode 		 	: "401",
		// BankCode 		 		: nil,
		Name 		 			: "BRI Epay",
		PaymentLogo 		 	: "bri-epay.png",
		// RateApplied 		 	: false,
		// TransactionFee 		 	: 5000.00,
		// PercentageFee 		 	: 0,
		Status 		 			: "inactive",
		// CreatedAt        		: time.Now(),
		// UpdatedAt        		: time.Now(),
	},
	entity.PaymentMethod{
		Code 					: "bri-mocash",
		// InitPaymentCode  		: "PDDC2",
		// RenewalPaymentCode 		: "PDDC2",
		FastpayCode 		 	: "",
		// BankCode 		 		: nil,
		Name 		 			: "BRI Mocash",
		PaymentLogo 		 	: "bri-mocash.png",
		// RateApplied 		 	: false,
		// TransactionFee 		 	: 2500.00,
		// PercentageFee 		 	: 0,
		Status 		 			: "inactive",
		// CreatedAt        		: time.Now(),
		// UpdatedAt        		: time.Now(),
	},
	entity.PaymentMethod{
		Code 					: "sakuku-bca",
		// InitPaymentCode  		: "PDDC2",
		// RenewalPaymentCode 		: "PDDC2",
		FastpayCode 		 	: "",
		// BankCode 		 		: nil,
		Name 		 			: "Sakuku BCA",
		PaymentLogo 		 	: "sakuku-bca.png",
		// RateApplied 		 	: false,
		// TransactionFee 		 	: 500.00,
		// PercentageFee 		 	: 0,
		Status 		 			: "inactive",
		// CreatedAt        		: time.Now(),
		// UpdatedAt        		: time.Now(),
	},
	entity.PaymentMethod{
		Code 					: "va-bca",
		// InitPaymentCode  		: "PDDC2",
		// RenewalPaymentCode 		: "PDDC2",
		FastpayCode 		 	: "",
		// BankCode 		 		: nil,
		Name 		 			: "VA BCA",
		PaymentLogo 		 	: "va-bca.png",
		// RateApplied 		 	: false,
		// TransactionFee 		 	: 2000.00,
		// PercentageFee 		 	: 0,
		Status 		 			: "inactive",
		// CreatedAt        		: time.Now(),
		// UpdatedAt        		: time.Now(),
	},
	entity.PaymentMethod{
		Code 					: "atm-mandiri",
		InitPaymentCode  		: "PDDC5",
		RenewalPaymentCode 		: "PDDC5",
		FastpayCode 		 	: "703",
		BankCode 		 		: "842802",
		Name 		 			: "ATM Mandiri",
		PaymentLogo 		 	: "atm-mandiri.png",
		// RateApplied 		 	: false,
		// TransactionFee 		 	: 5000.00,
		// PercentageFee 		 	: 0,
		Status 		 			: "active",
		// CreatedAt        		: time.Now(),
		// UpdatedAt        		: time.Now(),
	},
	entity.PaymentMethod{
		Code 					: "atm-bii",
		// InitPaymentCode  		: "PDDC5",
		// RenewalPaymentCode 		: "PDDC5",
		FastpayCode 		 	: "",
		// BankCode 		 		: "842802",
		Name 		 			: "ATM BII",
		PaymentLogo 		 	: "atm-bii.png",
		// RateApplied 		 	: false,
		// TransactionFee 		 	: 4000.00,
		// PercentageFee 		 	: 0,
		Status 		 			: "inactive",
		// CreatedAt        		: time.Now(),
		// UpdatedAt        		: time.Now(),
	},
	entity.PaymentMethod{
		Code 					: "atm-permata",
		InitPaymentCode  		: "PDDC4",
		RenewalPaymentCode 		: "PDDC4",
		FastpayCode 		 	: "402",
		BankCode 		 		: "898508",
		Name 		 			: "ATM Permata",
		PaymentLogo 		 	: "atm-permata.png",
		// RateApplied 		 	: false,
		// TransactionFee 		 	: 3000.00,
		// PercentageFee 		 	: 0,
		Status 		 			: "active",
		// CreatedAt        		: time.Now(),
		// UpdatedAt        		: time.Now(),
	},
	entity.PaymentMethod{
		Code 					: "rekening-ponsel",
		// InitPaymentCode  		: "PDDC4",
		// RenewalPaymentCode 		: "PDDC4",
		FastpayCode 		 	: "",
		// BankCode 		 		: "898508",
		Name 		 			: "Rekening Ponsel",
		PaymentLogo 		 	: "rekening-ponsel.png",
		// RateApplied 		 	: false,
		// TransactionFee 		 	: 5000.00,
		// PercentageFee 		 	: 0,
		Status 		 			: "inactive",
		// CreatedAt        		: time.Now(),
		// UpdatedAt        		: time.Now(),
	},
	entity.PaymentMethod{
		Code 					: "t-cash-telkomsel",
		// InitPaymentCode  		: "PDDC4",
		// RenewalPaymentCode 		: "PDDC4",
		FastpayCode 		 	: "",
		// BankCode 		 		: "898508",
		Name 		 			: "T-CASH Telkomsel",
		PaymentLogo 		 	: "t-cash-telkomsel.png",
		// RateApplied 		 	: false,
		// TransactionFee 		 	: 3000.00,
		// PercentageFee 		 	: 0,
		Status 		 			: "inactive",
		// CreatedAt        		: time.Now(),
		// UpdatedAt        		: time.Now(),
	},
	entity.PaymentMethod{
		Code 					: "xl-tunai",
		// InitPaymentCode  		: "PDDC4",
		// RenewalPaymentCode 		: "PDDC4",
		FastpayCode 		 	: "",
		// BankCode 		 		: "898508",
		Name 		 			: "XL Tunai",
		PaymentLogo 		 	: "xl-tunai.png",
		// RateApplied 		 	: false,
		// TransactionFee 		 	: 3000.00,
		// PercentageFee 		 	: 0,
		Status 		 			: "inactive",
		// CreatedAt        		: time.Now(),
		// UpdatedAt        		: time.Now(),
	},
	entity.PaymentMethod{
		Code 					: "dompetku-indosat",
		// InitPaymentCode  		: "PDDC4",
		// RenewalPaymentCode 		: "PDDC4",
		FastpayCode 		 	: "",
		// BankCode 		 		: "898508",
		Name 		 			: "Dompetku Indosat",
		PaymentLogo 		 	: "dompetku-indosat.png",
		// RateApplied 		 	: false,
		// TransactionFee 		 	: 3000.00,
		// PercentageFee 		 	: 0,
		Status 		 			: "inactive",
		// CreatedAt        		: time.Now(),
		// UpdatedAt        		: time.Now(),
	},
	entity.PaymentMethod{
		Code 					: "indomaret",
		InitPaymentCode  		: "PDDC1",
		RenewalPaymentCode 		: "PDDC1",
		FastpayCode 		 	: "706",
		BankCode 		 		: "312530",
		Name 		 			: "Indomaret",
		PaymentLogo 		 	: "indomaret.png",
		// RateApplied 		 	: true,
		// TransactionFee 		 	: 3750.00,
		// PercentageFee 		 	: 0,
		Status 		 			: "active",
		// CreatedAt        		: time.Now(),
		// UpdatedAt        		: time.Now(),
	},
	entity.PaymentMethod{
		Code 					: "mynt-artajasa",
		// InitPaymentCode  		: "PDDC1",
		// RenewalPaymentCode 		: "PDDC1",
		FastpayCode 		 	: "",
		// BankCode 		 		: "312530",
		Name 		 			: "Mynt Artajasa",
		PaymentLogo 		 	: "mynt-artajasa.png",
		// RateApplied 		 	: false,
		// TransactionFee 		 	: 3000.00,
		// PercentageFee 		 	: 0,
		Status 		 			: "inactive",
		// CreatedAt        		: time.Now(),
		// UpdatedAt        		: time.Now(),
	},
	entity.PaymentMethod{
		Code 					: "tfp-artajasa",
		// InitPaymentCode  		: "PDDC1",
		// RenewalPaymentCode 		: "PDDC1",
		FastpayCode 		 	: "",
		// BankCode 		 		: "312530",
		Name 		 			: "TFP Artajasa",
		PaymentLogo 		 	: "tfp-artajasa.png",
		// RateApplied 		 	: false,
		// TransactionFee 		 	: 5000.00,
		// PercentageFee 		 	: 0,
		Status 		 			: "inactive",
		// CreatedAt        		: time.Now(),
		// UpdatedAt        		: time.Now(),
	},
	entity.PaymentMethod{
		Code 					: "gopay",
		InitPaymentCode  		: "PDDC4",
		RenewalPaymentCode 		: "PDDC4",
		FastpayCode 		 	: "90647",
		BankCode 		 		: "333333",
		Name 		 			: "Go Pay",
		PaymentLogo 		 	: "gopay.png",
		// RateApplied 		 	: true,
		// TransactionFee 		 	: 0.00,
		// PercentageFee 		 	: 2,
		Status 		 			: "active",
		// CreatedAt        		: time.Now(),
		// UpdatedAt        		: time.Now(),
	},
	entity.PaymentMethod{
		Code 					: "auto-debet",
		InitPaymentCode  		: "PDAD",
		RenewalPaymentCode 		: "PDAD",
		FastpayCode 		 	: "",
		// BankCode 		 		: "333333",
		Name 		 			: "Auto Debet",
		PaymentLogo 		 	: "auto-debet.png",
		// RateApplied 		 	: false,
		// TransactionFee 		 	: 0.00,
		// PercentageFee 		 	: 0,
		Status 		 			: "inactive",
		// CreatedAt        		: time.Now(),
		// UpdatedAt        		: time.Now(),
	},
	entity.PaymentMethod{
		Code 					: "beever-salary-advance",
		InitPaymentCode  		: "PDDC4",
		RenewalPaymentCode 		: "PDDC4",
		FastpayCode 		 	: "402",
		BankCode 		 		: "898508",
		Name 		 			: "Beever Salary Advance",
		PaymentLogo 		 	: "beever-salary-advance.png",
		// RateApplied 		 	: false,
		// TransactionFee 		 	: 3000.00,
		// PercentageFee 		 	: 0,
		Status 		 			: "inactive",
		// CreatedAt        		: time.Now(),
		// UpdatedAt        		: time.Now(),
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

func SeedPaymentMethods() {
	db, _ := config.ConnectDB()

	_ = db.AutoMigrate(&entity.PaymentMethod{})
	for i, _ := range paymentMethod {
		err := db.Model(&entity.PaymentMethod{}).Create(&paymentMethod[i])
		if(err.Error != nil){
			fmt.Println("cannot seed 'Payment Method' table: ", err)
		}
	}
}