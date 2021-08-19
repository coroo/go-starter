package middlewares

import (
	"fmt"
	"os"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	utils "github.com/coroo/go-starter/app/utils"
)

func BasicAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		os.Getenv("BASIC_AUTH_USERNAME"): os.Getenv("BASIC_AUTH_PASSWORD"),
	})
}

func Auth(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("secret"), nil
	})

	if token != nil && err == nil && !utils.IsInBlacklist(tokenString) {
		fmt.Println("token verified")
	} else if utils.IsInBlacklist(tokenString) {
		result := gin.H{
			"message": "not authorized",
			"error":   "already logout",
		}
		c.JSON(http.StatusUnauthorized, result)
		c.Abort()
	} else {
		result := gin.H{
			"message": "not authorized",
			"error":   err.Error(),
		}
		c.JSON(http.StatusUnauthorized, result)
		c.Abort()
	}
}
