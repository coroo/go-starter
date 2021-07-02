package routes

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	deliveries "github.com/coroo/go-lemonilo/app/deliveries"
	entity "github.com/coroo/go-lemonilo/app/entity"

	"github.com/gin-gonic/gin"
)

// dummy data
var dummyUserProfile = []*entity.UserProfile{
	&entity.UserProfile{
		ID				: 1,
		Email			: "kuncoro@test.com",
		Password		: "password",
		Address			: "jl lorem ipsum",
		CreatedAt		: time.Now(),
		UpdatedAt		: time.Now(),
	}, &entity.UserProfile{
		ID				: 2,
		Email			: "kuncoro@test.com",
		Password		: "password",
		Address			: "jl lorem ipsum",
		CreatedAt		: time.Now(),
		UpdatedAt		: time.Now(),
	},
}

type userProfileRouteMock struct {
	mock.Mock
}

func (s *userProfileRouteMock) Save(userProfile entity.UserProfile) (int, error) {
	return 0, nil
}

func (s *userProfileRouteMock) Update(userProfile entity.UserProfile) error {
	return nil
}

func (s *userProfileRouteMock) Delete(userProfile entity.UserProfile) error {
	return nil
}

func (s *userProfileRouteMock) GetAllUserProfiles() []entity.UserProfile {
	return nil
}

func (s *userProfileRouteMock) GetUserProfile(ctx *gin.Context) []entity.UserProfile {
	return nil
}

func (s *userProfileRouteMock) GetUserProfileByQuestion(ctx *gin.Context) []entity.UserProfile {
	return nil
}

type UserProfileRouteTestSuite struct {
	suite.Suite
	serviceTest deliveries.UserProfileController
}

func (suite *UserProfileRouteTestSuite) SetupTest() {
	suite.serviceTest = new(userProfileRouteMock)
}

func (suite *UserProfileRouteTestSuite) TestSaveDelivery() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	r.POST("userProfile/create", UserProfileCreate)

	jsonValue, _ := json.Marshal(dummyUserProfile[0])
	req, _ := http.NewRequest(http.MethodPost, "/userProfile/create", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *UserProfileRouteTestSuite) TestUpdateDelivery() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	r.POST("userProfile/update", UserProfileUpdate)

	jsonValue, _ := json.Marshal(dummyUserProfile[0])
	req, _ := http.NewRequest(http.MethodPost, "/userProfile/update", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *UserProfileRouteTestSuite) TestDeleteDelivery() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	r.POST("userProfile/delete", UserProfileDelete)

	jsonValue, _ := json.Marshal(dummyUserProfile[0])
	req, _ := http.NewRequest(http.MethodPost, "/userProfile/delete", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *UserProfileRouteTestSuite) TestUserProfilesIndexRoute() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("userProfile/index", UserProfilesIndex)
	req, _ := http.NewRequest(http.MethodGet, "/userProfile/index", nil)

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *UserProfileRouteTestSuite) TestUserProfilesDetailRoute() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("userProfile/detail/1", UserProfilesDetail)
	req, _ := http.NewRequest(http.MethodGet, "/userProfile/detail/1", nil)

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func TestUserProfileRouteTestSuite(t *testing.T) {
	suite.Run(t, new(UserProfileRouteTestSuite))
}
