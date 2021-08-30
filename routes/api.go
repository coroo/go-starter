package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/coroo/go-starter/app/middlewares"
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

	// for the mean time commented for ods refactoring

	var (
		userRepository repositories.UserRepository = repositories.NewUserRepository()
		userService    usecases.UserService        = usecases.NewUser(userRepository)
	)
	deliveries.NewUserController(router, API_PREFIX, userService)

	userPoliciesGroup := router.Group(API_PREFIX + "userPolicies", middlewares.Auth)
	{
		userPoliciesGroup.GET("index", deliveries.GetAllUserPolicies)
		userPoliciesGroup.GET("detail/:id", deliveries.GetUserPolicy)
	}

	lumpSumPaymentGroup := router.Group(API_PREFIX + "lumpSumPayment", middlewares.Auth)
	{
		lumpSumPaymentGroup.GET("index", deliveries.GetAllLumpSumPayments)
		lumpSumPaymentGroup.GET("detail/:policyNumber", deliveries.GetLumpSumPayment)
		lumpSumPaymentGroup.GET("map-etl-payment", deliveries.OdsMapEtlLatestPayment)
	}

	syUserInvoiceGroup := router.Group(API_PREFIX + "syUserInvoice", middlewares.Auth)
	{
		syUserInvoiceGroup.GET("index", deliveries.GetAllUserPolicies)
		syUserInvoiceGroup.GET("map-etl-payment", deliveries.GetUserPolicy)
	}

	syETLGroup := router.Group(API_PREFIX + "syEtl", middlewares.Auth)
	{
		syETLGroup.GET("payment/index", deliveries.GetAllSyEtlPayments)
		syETLGroup.GET("payment/map-etl-payment", deliveries.SyOdsMapEtlLatestPayment)
		syETLGroup.GET("payment/detail/:policyNumber", deliveries.GetSyEtlPayment)
		syETLGroup.POST("payment/create", deliveries.CreateSyEtlPayment)
		syETLGroup.GET("payment/remove-before-map", deliveries.TruncateTableSyEtlPayments)
	}

	odsETLGroup := router.Group(API_PREFIX + "odsEtl", middlewares.Auth)
	{
		odsETLGroup.GET("payment/index", deliveries.GetAllOdsEtlPayments)
		odsETLGroup.GET("payment/detail/:id", deliveries.GetOdsEtlPayment)
		odsETLGroup.POST("payment/create", deliveries.CreateOdsEtlPayment)
		odsETLGroup.GET("payment/remove-before-map", deliveries.TruncateTableOdsEtlPayments)
	}

	syOdsETLGroup := router.Group(API_PREFIX + "syOdsEtl", middlewares.Auth)
	{
		syOdsETLGroup.GET("payment/index", deliveries.GetAllSyOdsEtlPayments)
		syOdsETLGroup.POST("payment/create", deliveries.CreateSyOdsEtlPayment)
		syOdsETLGroup.GET("payment/status/:status", deliveries.GetSyOdsEtlPaymentByStatus)
		syOdsETLGroup.GET("payment/daily-by-status/:status", deliveries.GetSyOdsEtlPaymentDailyByStatus)
		syOdsETLGroup.GET("payment/remove-before-map", deliveries.CancelOutstandingSyOdsEtlPayments)
		syOdsETLGroup.GET("payment/detail/:policyNumber", deliveries.GetSyOdsEtlPaymentByPolicyNumber)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	console.Schedule()
	router.Run(":"+os.Getenv("MAIN_PORT"))
}