package deliveries

import (
	"net/http"
	"fmt"
    // "time"

	utils "github.com/coroo/go-starter/app/utils"
	entity "github.com/coroo/go-starter/app/entity"
	usecases "github.com/coroo/go-starter/app/usecases"
	"github.com/coroo/go-starter/app/middlewares"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type userController struct {
	usecases usecases.UserService
}

func NewUserController(router *gin.Engine, apiPrefix string, userService usecases.UserService) {
	handlerUser := &userController{
		usecases: userService,
	}
	usersGroup := router.Group(apiPrefix + "user")
	{
		usersGroup.POST("login", handlerUser.AuthLogin)
		usersGroup.POST("refresh", handlerUser.AuthRefreshToken)
		usersGroup.POST("logout", handlerUser.AuthDestroyToken)
		usersGroup.GET("index", middlewares.Auth, handlerUser.UsersIndex)
		usersGroup.GET("detail/:id", middlewares.Auth, handlerUser.UsersDetail)
		usersGroup.POST("create", handlerUser.UserCreate)
		usersGroup.PUT("update", handlerUser.UserUpdate)
		usersGroup.DELETE("delete", handlerUser.UserDelete)
	}
}

// GetUsersIndex godoc
// @Param Authorization header string true "Bearer"
// @Security basicAuth
// @Summary Show all existing Users
// @Description Get all existing Users
// @Tags Users
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.User
// @Failure 401 {object} dto.Response
// @Router /user/index [get]
func (deliveries *userController) UsersIndex(c *gin.Context) {
	users := deliveries.usecases.GetAllUsers()
	c.JSON(http.StatusOK, gin.H{"data": users})
}

// GetUsersDetail godoc
// @Param Authorization header string true "Bearer"
// @Security basicAuth
// @Summary Show an existing Users
// @Description Get detail the existing Users
// @Tags Users
// @Accept  json
// @Produce  json
// @Param  id path int true "Master Question ID"
// @Success 200 {array} entity.User
// @Failure 401 {object} dto.Response
// @Router /user/detail/{id} [get]
func (deliveries *userController) UsersDetail(c *gin.Context) {
	user := deliveries.usecases.GetUser(c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// CreateUsers godoc
// @Security basicAuth
// @Summary Create new Users
// @Description Create a new Users
// @Tags Users
// @Accept  json
// @Produce  json
// @Param user body entity.UserCreate true "Create user"
// @Success 200 {object} entity.User
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /user/create [post]
func (deliveries *userController) UserCreate(c *gin.Context) {
	var userEntity entity.User
	c.ShouldBindJSON(&userEntity)
	userPK, err := deliveries.usecases.SaveUser(userEntity)
	if(err!=nil){
		c.JSON(http.StatusConflict, err)
	} else {
		userEntity.ID = userPK
		c.JSON(http.StatusOK, userEntity)
	}
}

// UpdateUsers godoc
// @Security basicAuth
// @Summary Update Users
// @Description Update a Users
// @Tags Users
// @Accept  json
// @Produce  json
// @Param user body entity.User true "Update user"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /user/update [put]
func (deliveries *userController) UserUpdate(c *gin.Context) {
	var userEntity entity.User
	c.ShouldBindJSON(&userEntity)
	user := deliveries.usecases.UpdateUser(userEntity)
	c.JSON(http.StatusOK, user)
}

// DeleteUsers godoc
// @Security basicAuth
// @Summary Delete Users
// @Description Delete a Users
// @Tags Users
// @Accept  json
// @Produce  json
// @Param user body entity.UserDelete true "Delete user"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /user/delete [delete]
func (deliveries *userController) UserDelete(c *gin.Context) {
	var userEntity entity.User
	c.ShouldBindJSON(&userEntity)
	user := deliveries.usecases.DeleteUser(userEntity)
	c.JSON(http.StatusOK, user)
}

// Login godoc
// @Security basicAuth
// @Summary Login new Users
// @Description Login a new Users
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param user body entity.UserLogin true "Login user"
// @Success 200 {object} dto.Token
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /user/login [post]
func (deliveries *userController) AuthLogin(c *gin.Context) {
	var userEntity entity.User
	c.ShouldBindJSON(&userEntity)
	stats, res := deliveries.usecases.AuthUser(userEntity)
	if(stats==200){
		authRes, err := utils.CreateToken(res)
		// sign := jwt.New(jwt.GetSigningMethod("HS256"))
		// // Set claims
        // // This is the information which frontend can use
        // // The backend can also decode the token and get admin etc.
        // claims := sign.Claims.(jwt.MapClaims)
        // claims["user_id"] = res.ID
        // claims["email"] = res.Email
        // claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
		// token, err := sign.SignedString([]byte("secret"))
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{
		// 		"message": err.Error(),
		// 	})
		// 	c.Abort()
		// }

		// refreshToken := jwt.New(jwt.SigningMethodHS256)
		// rtClaims := refreshToken.Claims.(jwt.MapClaims)
		// rtClaims["user_id"] = res.ID
		// rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()
		// rt, err := refreshToken.SignedString([]byte("secret"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			c.Abort()
		}

		c.JSON(http.StatusOK, gin.H{
			"access_token": authRes.AccessToken,
			"refresh_token": authRes.RefreshToken,
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Terjadi kesalahan pada penulisan email atau password kamu, harap periksa kembali"})
	}
}

// Refresh godoc
// @Security basicAuth
// @Summary Refresh Users Token
// @Description Refresh a Users Token
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param user body dto.Token true "Login user"
// @Success 200 {object} dto.Token
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /user/refresh [post]
func (deliveries *userController) AuthRefreshToken(c *gin.Context) {
	var tokenRequest entity.TokenReqBody
	c.ShouldBindJSON(&tokenRequest)
	token, _ := jwt.Parse(tokenRequest.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("secret"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var userEntity entity.User
		fmt.Println(claims)
		userEntity.Email = claims["email"].(string)
		// Get the user record from database or
		// run through your business logic to verify if the user can log in
		authRes, err := utils.CreateToken(userEntity)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			c.Abort()
		}

		c.JSON(http.StatusOK, gin.H{
			"access_token": authRes.AccessToken,
			"refresh_token": authRes.RefreshToken,
		})
	}
}

// Destroy godoc
// @Security basicAuth
// @Summary Destroy Users Token
// @Description Destroy a Users Token
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param user body dto.Token true "Login user"
// @Success 200 {object} dto.Token
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /user/logout [post]
func (deliveries *userController) AuthDestroyToken(c *gin.Context) {
	var tokenRequest entity.Token
	c.ShouldBindJSON(&tokenRequest)
	token, _ := jwt.Parse(tokenRequest.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		
		err := utils.Logout(tokenRequest.AccessToken, token)
		return []byte("secret"), err
	})
	

	// if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
	// 	var userEntity entity.User
	// 	fmt.Println(claims)
	// 	userEntity.Email = claims["email"].(string)
	// 	// Get the user record from database or
	// 	// run through your business logic to verify if the user can log in
	// 	authRes, err := utils.CreateToken(userEntity)
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{
	// 			"message": err.Error(),
	// 		})
	// 		c.Abort()
	// 	}

		c.JSON(http.StatusOK, gin.H{
			"access_token": token,
			// "refresh_token": authRes.RefreshToken,
		})
	// }
}