package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/coroo/go-starter/app/deliveries"
	"github.com/coroo/go-starter/app/usecases"
	"github.com/coroo/go-starter/app/repositories"
	"github.com/coroo/go-starter/app/console"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Api() {
	router := gin.Default()
	// router.Use(middlewares.BasicAuth())

	API_PREFIX := os.Getenv("API_PREFIX")

	router.GET("/", func(c *gin.Context) {
		c.JSON(404, gin.H{
			"message": os.Getenv("MAIN_DESCRIPTION"),
		})
	})
	// PAYMENT METHOD
	var (
		paymentMethodRepository repositories.PaymentMethodRepository = repositories.NewPaymentMethodRepository()
		paymentMethodService    usecases.PaymentMethodService = usecases.NewPaymentMethodService(paymentMethodRepository)
		// userController deliveries.UserController   = deliveries.NewUser(userService)
	)
	deliveries.NewPaymentMethodController(router, API_PREFIX, paymentMethodService)

	// PAYMENT METHOD RATE
	var (
		paymentMethodRateRepository repositories.PaymentMethodRateRepository = repositories.NewPaymentMethodRateRepository()
		paymentMethodRateService    usecases.PaymentMethodRateService = usecases.NewPaymentMethodRateService(paymentMethodRateRepository)
		// userController deliveries.UserController   = deliveries.NewUser(userService)
	)
	deliveries.NewPaymentMethodRateController(router, API_PREFIX, paymentMethodRateService)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	console.Schedule()
	router.Run(":"+os.Getenv("MAIN_PORT"))
}