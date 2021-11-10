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
var dummyGopayLinking = []*entity.GopayLinking{
	&entity.GopayLinking{
		ID				        : 1,
		PhoneNumber				: "08123456789",
		AccountId		  		: "ojbqegnlfdam",
		PaymentOptionToken  	: "bnafjlnfasn",
		CreatedAt				: time.Now(),
		UpdatedAt				: time.Now(),
	}, &entity.GopayLinking{
		ID				        : 2,
		PhoneNumber				: "0898654321",
		AccountId		  		: "jasjarjlasr",
		PaymentOptionToken  	: "afnlanrouwqnf",
		CreatedAt				: time.Now(),
		UpdatedAt				: time.Now(),
	},
}

type gopayLinkingRouteMock struct {
	mock.Mock
}

func (s *gopayLinkingRouteMock) SaveGopayLinking(gopayLinking entity.GopayLinking) (int, error) {
	return 0, nil
}

func (s *gopayLinkingRouteMock) UpdateGopayLinking(gopayLinking entity.GopayLinking) error {
	return nil
}

func (s *gopayLinkingRouteMock) DeleteGopayLinking(gopayLinking entity.GopayLinking) error {
	return nil
}

func (s *gopayLinkingRouteMock) GetAllGopayLinkings() []entity.GopayLinking {
	return nil
}

func (s *gopayLinkingRouteMock) GetGopayLinking(id string) []entity.GopayLinking {
	return nil
}

type GopayLinkingRouteTestSuite struct {
	suite.Suite
	serviceTest usecases.GopayLinkingService
}

func (suite *GopayLinkingRouteTestSuite) SetupTest() {
	suite.serviceTest = new(gopayLinkingRouteMock)
}

func (suite *GopayLinkingRouteTestSuite) TestA_SaveDelivery() {
	// fmt.Println("SAVE USER TEST DELIV")
	// suite.serviceTest.(*gopayLinkingRouteMock).On("SaveGopayLinking", dummyGopayLinking[0]).Return(nil)
	// c, _ := gin.CreateTestContext(httptest.NewRecorder())
	// c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}
	// GopayLinkingCreate(c)

	// fmt.Println("SAVE USER TEST DELIV 2")
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	handlerGopayLinking := &gopayLinkingController{
		usecases: suite.serviceTest,
	}
	r.POST("gopayLinking/create", handlerGopayLinking.GopayLinkingCreate)

	jsonValue, _ := json.Marshal(dummyGopayLinking[1])
	req, _ := http.NewRequest(http.MethodPost, "/gopayLinking/create", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *GopayLinkingRouteTestSuite) TestUpdateDelivery() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	handlerGopayLinking := &gopayLinkingController{
		usecases: suite.serviceTest,
	}
	r.POST("gopayLinking/update", handlerGopayLinking.GopayLinkingUpdate)

	jsonValue, _ := json.Marshal(dummyGopayLinking[0])
	req, _ := http.NewRequest(http.MethodPost, "/gopayLinking/update", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	fmt.Println("UPDATE USER TEST")
	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *GopayLinkingRouteTestSuite) TestDeleteDelivery() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	handlerGopayLinking := &gopayLinkingController{
		usecases: suite.serviceTest,
	}
	r.POST("gopayLinking/delete", handlerGopayLinking.GopayLinkingDelete)

	jsonValue, _ := json.Marshal(dummyGopayLinking[0])
	req, _ := http.NewRequest(http.MethodPost, "/gopayLinking/delete", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *GopayLinkingRouteTestSuite) TestGopayLinkingsIndexRoute() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	handlerGopayLinking := &gopayLinkingController{
		usecases: suite.serviceTest,
	}
	r.GET("gopayLinking/index", handlerGopayLinking.GopayLinkingsIndex)
	req, _ := http.NewRequest(http.MethodGet, "/gopayLinking/index", nil)

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *GopayLinkingRouteTestSuite) TestGopayLinkingsDetailRoute() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	handlerGopayLinking := &gopayLinkingController{
		usecases: suite.serviceTest,
	}
	r.GET("gopayLinking/detail/:id", handlerGopayLinking.GopayLinkingsDetail)
	req, _ := http.NewRequest(http.MethodGet, "/gopayLinking/detail/1", nil)

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func TestGopayLinkingRouteTestSuite(t *testing.T) {
	suite.Run(t, new(GopayLinkingRouteTestSuite))
}