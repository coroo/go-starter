package deliveries

import (
	"fmt"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	usecases "github.com/coroo/go-starter/app/usecases"
	entity "github.com/coroo/go-starter/app/entity"

	"github.com/gin-gonic/gin"
)

// dummy data
var dummyUser = []*entity.User{
	&entity.User{
		ID				: 1,
		Email			: "kuncoro@test.com",
		Password		: "password",
		Name			: "jl lorem ipsum",
		CreatedAt		: time.Now(),
		UpdatedAt		: time.Now(),
	}, &entity.User{
		ID				: 2,
		Email			: "kuncoro2@test.com",
		Password		: "password",
		Name			: "jl lorem ipsum",
		CreatedAt		: time.Now(),
		UpdatedAt		: time.Now(),
	},
}

type userRouteMock struct {
	mock.Mock
}

func (s *userRouteMock) SaveUser(user entity.User) (int, error) {
	fmt.Println("SAVE USER TEST MOCK")
	return 0, nil
}

func (s *userRouteMock) UpdateUser(user entity.User) error {
	return nil
}

func (s *userRouteMock) DeleteUser(user entity.User) error {
	return nil
}

func (s *userRouteMock) GetAllUsers() []entity.User {
	return nil
}

func (s *userRouteMock) GetUser(id string) []entity.User {
	return nil
}

func (s *userRouteMock) AuthUser(user entity.User) (int, entity.User) {
	return 200, user
}

type UserRouteTestSuite struct {
	suite.Suite
	serviceTest usecases.UserService
}

func (suite *UserRouteTestSuite) SetupTest() {
	suite.serviceTest = new(userRouteMock)
}

func (suite *UserRouteTestSuite) TestSaveDelivery() {
	// fmt.Println("SAVE USER TEST DELIV")
	// suite.serviceTest.(*userRouteMock).On("SaveUser", dummyUser[0]).Return(nil)
	// c, _ := gin.CreateTestContext(httptest.NewRecorder())
	// c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}
	// UserCreate(c)

	// fmt.Println("SAVE USER TEST DELIV 2")
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	handlerUser := &userController{
		usecases: suite.serviceTest,
	}
	r.POST("user/create", handlerUser.UserCreate)

	jsonValue, _ := json.Marshal(dummyUser[1])
	req, _ := http.NewRequest(http.MethodPost, "/user/create", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *UserRouteTestSuite) TestUpdateDelivery() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	handlerUser := &userController{
		usecases: suite.serviceTest,
	}
	r.POST("user/update", handlerUser.UserUpdate)

	jsonValue, _ := json.Marshal(dummyUser[0])
	req, _ := http.NewRequest(http.MethodPost, "/user/update", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	fmt.Println("UPDATE USER TEST")
	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *UserRouteTestSuite) TestDeleteDelivery() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	handlerUser := &userController{
		usecases: suite.serviceTest,
	}
	r.POST("user/delete", handlerUser.UserDelete)

	jsonValue, _ := json.Marshal(dummyUser[0])
	req, _ := http.NewRequest(http.MethodPost, "/user/delete", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *UserRouteTestSuite) TestUsersIndexRoute() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	handlerUser := &userController{
		usecases: suite.serviceTest,
	}
	r.GET("user/index", handlerUser.UsersIndex)
	req, _ := http.NewRequest(http.MethodGet, "/user/index", nil)

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *UserRouteTestSuite) TestUsersDetailRoute() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	handlerUser := &userController{
		usecases: suite.serviceTest,
	}
	r.GET("user/detail/:id", handlerUser.UsersDetail)
	req, _ := http.NewRequest(http.MethodGet, "/user/detail/1", nil)

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *UserRouteTestSuite) TestAuthUsersRoute() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	handlerUser := &userController{
		usecases: suite.serviceTest,
	}
	jsonValue, _ := json.Marshal(dummyUser[0])
	// PREPARATION
	r.POST("user/create", handlerUser.UserCreate)

	reqPrep, _ := http.NewRequest(http.MethodPost, "/user/create", bytes.NewBuffer(jsonValue))
	reqPrep.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.ServeHTTP(w, reqPrep)

	// AFTER ANY PREPARATION
	r.POST("user/login", handlerUser.AuthLogin)

	req, _ := http.NewRequest(http.MethodPost, "/user/login", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")	
	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)

	loginValue, _ := json.Marshal(req.Body)
	
	// REFRESH
	r.POST("user/refresh", handlerUser.AuthRefreshToken)

	reqRefresh, _ := http.NewRequest(http.MethodPost, "/user/refresh", bytes.NewBuffer(loginValue))
	reqRefresh.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(w, reqRefresh)
	assert.Equal(suite.T(), w.Code, 200)
		
	// DESTROY
	r.POST("user/logout", handlerUser.AuthDestroyToken)

	reqLogout, _ := http.NewRequest(http.MethodPost, "/user/logout", bytes.NewBuffer(loginValue))
	reqLogout.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(w, reqLogout)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *UserRouteTestSuite) NegativeTestAuthUsersRoute() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	handlerUser := &userController{
		usecases: suite.serviceTest,
	}
	jsonValue, _ := json.Marshal(dummyUser[0])
	r.POST("user/login", handlerUser.AuthLogin)

	req, _ := http.NewRequest(http.MethodPost, "/user/login", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 401)
}

func TestUserRouteTestSuite(t *testing.T) {
	suite.Run(t, new(UserRouteTestSuite))
}
