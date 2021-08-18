package console

import (
	"os"
	"github.com/coroo/go-starter/app/utils"

	"gopkg.in/robfig/cron.v2"
)

var schedulerInfo = "Running Job: "

func Schedule() {
	c := cron.New()

	// Truncate Sy ETL Payment & ODS ETL Payment : Runs At : 7:30
	c.AddFunc("15 16 * * *", func() {
		_, errSy := utils.CreateHttpRequest("GET", os.Getenv("MAIN_SCHEMES")+"://"+os.Getenv("MAIN_URL")+"/"+os.Getenv("API_PREFIX")+"syEtl/payment/remove-before-map", nil)
		if errSy != nil {
			utils.ErrorLogger(errSy)
		} else {
			utils.InfoLogger(schedulerInfo + "Truncate Sy ETL Payment")
		}

		_, errOds := utils.CreateHttpRequest("GET", os.Getenv("MAIN_SCHEMES")+"://"+os.Getenv("MAIN_URL")+"/"+os.Getenv("API_PREFIX")+"odsEtl/payment/remove-before-map", nil)
		if errOds != nil {
			utils.ErrorLogger(errOds)
		} else {
			utils.InfoLogger(schedulerInfo + "Truncate Ods ETL Payment")
		}
	})

	// Mapping Ods ETL Payment : Runs At : 7:31
	c.AddFunc("16 16 * * *", func() {
		_, err := utils.CreateHttpRequest("GET", os.Getenv("MAIN_SCHEMES")+"://"+os.Getenv("MAIN_URL")+"/"+os.Getenv("API_PREFIX")+"lumpSumPayment/map-etl-payment", nil)
		if err != nil {
			utils.ErrorLogger(err)
		} else {
			utils.InfoLogger(schedulerInfo + "Mapping Ods ETL Payment")
		}
	})

	// Mapping Sy ETL Payment & Truncate Sy Ods ETL Payment: Runs At : 7:33
	c.AddFunc("18 16 * * *", func() {
		_, err := utils.CreateHttpRequest("GET", os.Getenv("MAIN_SCHEMES")+"://"+os.Getenv("MAIN_URL")+"/"+os.Getenv("API_PREFIX")+"syUserInvoice/map-etl-payment", nil)
		if err != nil {
			utils.ErrorLogger(err)
		} else {
			utils.InfoLogger(schedulerInfo + "Mapping Sy ETL Payment")
		}

		_, errOds := utils.CreateHttpRequest("GET", os.Getenv("MAIN_SCHEMES")+"://"+os.Getenv("MAIN_URL")+"/"+os.Getenv("API_PREFIX")+"syOdsEtl/payment/remove-before-map", nil)
		if errOds != nil {
			utils.ErrorLogger(errOds)
		} else {
			utils.InfoLogger(schedulerInfo + "Truncate Sy Ods ETL Payment")
		}
	})

	// Mapping Sy Ods ETL Payment : Runs At : 7:35
	c.AddFunc("20 16 * * *", func() {
		_, err := utils.CreateHttpRequest("GET", os.Getenv("MAIN_SCHEMES")+"://"+os.Getenv("MAIN_URL")+"/"+os.Getenv("API_PREFIX")+"syEtl/payment/map-etl-payment", nil)
		if err != nil {
			utils.ErrorLogger(err)
		} else {
			utils.InfoLogger(schedulerInfo + "Mapping Sy Ods ETL Payment")
		}
	})

	c.Start()
}
