package main

import (
	"log"
	"os"

	"github.com/coroo/go-lemonilo/app/middlewares"
	// "github.com/coroo/go-lemonilo/app/routes"
	"github.com/coroo/go-lemonilo/app/deliveries"
	"github.com/coroo/go-lemonilo/docs"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// Swagger 2.0 Meta Information
	docs.SwaggerInfo.Title = os.Getenv("MAIN_TITLE")
	docs.SwaggerInfo.Description = os.Getenv("MAIN_DESCRIPTION")
	docs.SwaggerInfo.Version = os.Getenv("MAIN_VERSION")
	docs.SwaggerInfo.Host = os.Getenv("MAIN_URL")
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	_, err := os.Stat("storage/logs")

	if os.IsNotExist(err) {
		err_0 := os.Mkdir("storage/logs", 0755)
		if err_0 != nil {
			log.Fatal(err_0)
		}
		err_1 := os.Mkdir("storage/logs/errors", 0755)
		if err_1 != nil {
			log.Fatal(err_1)
		}
		err_2 := os.Mkdir("storage/logs/informations", 0755)
		if err_2 != nil {
			log.Fatal(err_2)
		}
	}

	router := gin.Default()
	// router.Use(middlewares.BasicAuth(), middlewares.Logger())

	API_PREFIX := os.Getenv("API_PREFIX")

	router.GET("/", func(c *gin.Context) {
		c.JSON(404, gin.H{
			"message": os.Getenv("MAIN_DESCRIPTION"),
		})
	})

	userProfilesGroup := router.Group(API_PREFIX + "userProfile")
	{
		userProfilesGroup.POST("login", deliveries.AuthProfilesDetail)
		userProfilesGroup.GET("index", middlewares.Auth, deliveries.UserProfilesIndex)
		userProfilesGroup.GET("detail/:id", middlewares.Auth, deliveries.UserProfilesDetail)
		userProfilesGroup.POST("create", deliveries.UserProfileCreate)
		userProfilesGroup.PUT("update", deliveries.UserProfileUpdate)
		userProfilesGroup.DELETE("delete", deliveries.UserProfileDelete)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":3000")

}
