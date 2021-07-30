package middlewares

import (
	"fmt"
	"log"
	"os"
	"github.com/coroo/go-starter/app/utils"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger returns a middleware with the specified log format function.
func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		f, err := os.OpenFile(`storage/logs/`+utils.DateNow("")+`.log`,
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}
		defer f.Close()

		if _, err := f.WriteString(
			fmt.Sprintf("%s - [%s] %s %s %d %s \n",
				param.ClientIP,
				param.TimeStamp.Format(time.RFC822),
				param.Method,
				param.Path,
				param.StatusCode,
				param.Latency,
			)); err != nil {
			log.Fatal(err)
		}

		return ``
	})
}
