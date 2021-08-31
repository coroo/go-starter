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

	// USER
	var (
		userRepository repositories.UserRepository = repositories.NewUserRepository()
		userService    usecases.UserService        = usecases.NewUser(userRepository)
	)
	deliveries.NewUserController(router, API_PREFIX, userService)

	// USER POLICY
	var (
		userPolicyRepository repositories.UserPolicyRepository = repositories.NewUserPolicyRepository()
		userPolicyService    usecases.UserPolicyService = usecases.NewUserPolicyService(userPolicyRepository)
		// userController deliveries.UserController   = deliveries.NewUser(userService)
	)
	deliveries.NewUserPolicyController(router, API_PREFIX, userPolicyService)

	// SY USER INVOICE
	var (
		syUserInvoiceRepository 	repositories.SyUserInvoiceRepository = repositories.NewSyUserInvoiceRepository()
		syUserInvoiceService		usecases.SyUserInvoiceService = usecases.NewSyUserInvoiceService(syUserInvoiceRepository)
	)
	deliveries.NewSyUserInvoiceController(router, API_PREFIX, syUserInvoiceService)

	// SY ODS ETL PAYMENT
	var (
		syOdsEtlPaymentRepository repositories.SyOdsEtlPaymentRepository = repositories.NewSyOdsEtlPaymentRepository()
		syOdsEtlPaymentService    usecases.SyOdsEtlPaymentService = usecases.NewSyOdsEtlPaymentService(syOdsEtlPaymentRepository)
	)
	deliveries.NewSyOdsEtlPaymentController(router, API_PREFIX, syOdsEtlPaymentService)

	// SY ETL PAYMENT
	var (
		syEtlPaymentRepository repositories.SyEtlPaymentRepository = repositories.NewSyEtlPaymentRepository()
		syEtlPaymentService    usecases.SyEtlPaymentService = usecases.NewSyEtlPaymentService(syEtlPaymentRepository)
	)
	deliveries.NewSyEtlPaymentController(router, API_PREFIX, syEtlPaymentService)

	// ODS ETL PAYMENT
	var (
		odsEtlPaymentRepository repositories.OdsEtlPaymentRepository = repositories.NewOdsEtlPaymentRepository()
		odsEtlPaymentService    usecases.OdsEtlPaymentService = usecases.NewOdsEtlPaymentService(odsEtlPaymentRepository)
		// userController deliveries.UserController   = deliveries.NewUser(userService)
	)
	deliveries.NewOdsEtlPaymentController(router, API_PREFIX, odsEtlPaymentService)

	// LUMP SUM PAYMENT
	var (
		lumpSumPaymentRepository 	repositories.LumpSumPaymentRepository = repositories.NewLumpSumPaymentRepository()
		lumpSumPaymentService		usecases.LumpSumPaymentService = usecases.NewLumpSumPaymentService(lumpSumPaymentRepository)
	)
	deliveries.NewLumpSumPaymentController(router, API_PREFIX, lumpSumPaymentService)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	console.Schedule()
	router.Run(":"+os.Getenv("MAIN_PORT"))
}