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

type lumpSumPaymentRouteMock struct {
	mock.Mock
}

func (s *lumpSumPaymentRouteMock) GetAllLumpSumPayments() []entity.LumpSumPayment {
	return nil
}

func (s *lumpSumPaymentRouteMock) OdsMapEtlLatestPayment() []entity.LumpSumPayment {
	return nil
}
func (s *lumpSumPaymentRouteMock) GetLumpSumPayment(policyNumber string) []entity.LumpSumPayment {
	return nil
}

type lumpSumPaymentRouteTestSuite struct {
	suite.Suite
	serviceTest usecases.LumpSumPaymentService
}

func (suite *lumpSumPaymentRouteTestSuite) SetupOdsEtlPaymentTest() {
	suite.serviceTest = new(lumpSumPaymentRouteMock)
}

func (suite *lumpSumPaymentRouteTestSuite) TestGetAllLumpSumPaymentsRoute() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("lumpSumPayment/index", GetAllLumpSumPayments)
	req, _ := http.NewRequest(http.MethodGet, "/lumpSumPayment/index", nil)

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *lumpSumPaymentRouteTestSuite) TestOdsMapEtlLatestPaymentRoute() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("lumpSumPayment/map-etl-payment", OdsMapEtlLatestPayment)
	req, _ := http.NewRequest(http.MethodGet, "/lumpSumPayment/map-etl-payment", nil)

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *lumpSumPaymentRouteTestSuite) TestGetLumpSumPaymentRoute() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("lumpSumPayment/detail/:policyNumber", TruncateTableOdsEtlPayments)
	req, _ := http.NewRequest(http.MethodGet, "/lumpSumPayment/detail/300123456", nil)

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}
