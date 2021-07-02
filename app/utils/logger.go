package utils

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func ErrorLogger(errorDesc error) {
	if flag.Lookup("test.v") == nil {
		f, err := os.OpenFile(`storage/logs/errors/err-`+DateNow("")+`.log`,
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}
		defer f.Close()

		if _, err := f.WriteString(
			fmt.Sprintf("[%s] %s \n",
				time.Now().Format(time.RFC822),
				errorDesc)); err != nil {
			log.Fatal(err)
		}
	}
}

func InfoLogger(infoDesc string) {
	if flag.Lookup("test.v") == nil {
		f, err := os.OpenFile(`storage/logs/informations/info-`+DateNow("")+`.log`,
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}
		defer f.Close()

		if _, err := f.WriteString(
			fmt.Sprintf("[%s] %s \n",
				time.Now().Format(time.RFC822),
				infoDesc)); err != nil {
			log.Fatal(err)
		}
	}
}
