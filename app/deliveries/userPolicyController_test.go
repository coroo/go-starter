package deliveries

import (
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	usecases "github.com/coroo/go-starter/app/usecases"
	entity "github.com/coroo/go-starter/app/entity"

	"github.com/gin-gonic/gin"
)

type userPolicyRouteMock struct {
	mock.Mock
}

func (s *userPolicyRouteMock) GetAllUserPolicies(is_overdue string) []entity.UserPolicy {
	return nil
}

func (s *userPolicyRouteMock) GetUserPolicy(id string) []entity.UserPolicy {
	return nil
}

type UserPolicyRouteTestSuite struct {
	suite.Suite
	serviceTest usecases.UserPolicyService
}

func (suite *UserPolicyRouteTestSuite) SetupUserPolicyTest() {
	suite.serviceTest = new(userPolicyRouteMock)
}

func (suite *UserPolicyRouteTestSuite) TestUserPolicyIndexRoute() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	handlerUserPolicy := &userPolicyController{
		usecases: suite.serviceTest,
	}
	r.GET("userPolicies/index", handlerUserPolicy.GetAllUserPolicies)
	req, _ := http.NewRequest(http.MethodGet, "/userPolicies/index", nil)

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *UserPolicyRouteTestSuite) TestUserPolicyDetailRoute() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	handlerUserPolicy := &userPolicyController{
		usecases: suite.serviceTest,
	}
	r.GET("userPolicies/detail/1", handlerUserPolicy.GetUserPolicy)
	req, _ := http.NewRequest(http.MethodGet, "/userPolicies/detail/1", nil)

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}
