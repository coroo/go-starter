package middlewares

// import (
// 	"fmt"
// 	"os"
// 	"net/http"
// 	"strings"

// 	jwt "github.com/dgrijalva/jwt-go"
// 	"github.com/gin-gonic/gin"
// 	_ "github.com/go-sql-driver/mysql"
// 	utils "github.com/coroo/go-starter/app/utils"
// 	"github.com/coroo/go-starter/app/repositories"
// )

// func BasicAuth() gin.HandlerFunc {
// 	return gin.BasicAuth(gin.Accounts{
// 		os.Getenv("BASIC_AUTH_USERNAME"): os.Getenv("BASIC_AUTH_PASSWORD"),
// 	})
// }

// func Auth(c *gin.Context) {
// 	tokenString := c.Request.Header.Get("Authorization")
// 	tokenArray := strings.Split(tokenString, " ")

// 	// if authentication using uuid
// 	if tokenArray[0] == "uuid" {
// 		// get user by uuid
// 		userRepo := repositories.NewUserRepository()
// 		getUser := userRepo.GetUserByUuid(tokenArray[1])
// 		// check if user founded with the uuid
// 		if getUser.ID != 0 {
// 			fmt.Println("uuid verified")
// 		}else{
// 			result := gin.H{
// 				"message": "not authorized",
// 				"error":   "user not found",
// 			}
// 			c.JSON(http.StatusUnauthorized, result)
// 			c.Abort()
// 		}
// 	}else{
// 		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 			if jwt.GetSigningMethod("HS256") != token.Method {
// 				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 			}
	
// 			return []byte("secret"), nil
// 		})

// 		if token != nil && err == nil && !utils.IsInBlacklist(tokenString) {
// 			fmt.Println("token verified")
// 		} else if utils.IsInBlacklist(tokenString) {
// 			result := gin.H{
// 				"message": "not authorized",
// 				"error":   "already logout",
// 			}
// 			c.JSON(http.StatusUnauthorized, result)
// 			c.Abort()
// 		} else {
// 			result := gin.H{
// 				"message": "not authorized",
// 				"error":   err.Error(),
// 			}
// 			c.JSON(http.StatusUnauthorized, result)
// 			c.Abort()
// 		}
// 	}
// }
