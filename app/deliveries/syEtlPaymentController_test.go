package deliveries

import (
	"time"
	"bytes"
	"encoding/json"

	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	usecases "github.com/coroo/go-starter/app/usecases"
	entity "github.com/coroo/go-starter/app/entity"

	"github.com/gin-gonic/gin"
)

var dummySyEtlPayment = []*entity.SyEtlPayment{
	&entity.SyEtlPayment{
		ID					: 1,
		ProposalNumber		: "600012345678",
		PolicyNumber		: "3001234567",
		PaidDate			: time.Now(),
		PaymentMethodName	: "Indomaret",
		PolicyStatus		: "success",
		TotalPremium		: 100000,
		UpdatedAt			: time.Now(),
	},
}

type syEtlPaymentRouteMock struct {
	mock.Mock
}

func (s *syEtlPaymentRouteMock) CreateSyEtlPayment(SyEtlPayment entity.SyEtlPayment) error {
	return nil
}

func (s *syEtlPaymentRouteMock) SyOdsMapEtlLatestPayment() []entity.SyEtlPayment {
	return nil
}
func (s *syEtlPaymentRouteMock) GetAllSyEtlPayments() []entity.SyEtlPayment {
	return nil
}
func (s *syEtlPaymentRouteMock) GetSyEtlPayment(policyNumber string) []entity.SyEtlPayment {
	return nil
}
func (s *syEtlPaymentRouteMock) TruncateTableSyEtlPayments() error {
	return nil
}

type syEtlPaymentRouteTestSuite struct {
	suite.Suite
	serviceTest usecases.SyEtlPaymentService
}

func (suite *syEtlPaymentRouteTestSuite) SetupsyOdsEtlPaymentTest() {
	suite.serviceTest = new(syEtlPaymentRouteMock)
}

func (suite *syEtlPaymentRouteTestSuite) TestGetAllSyEtlPaymentsRoute() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("syEtl/payment/index", GetAllSyEtlPayments)
	req, _ := http.NewRequest(http.MethodGet, "/syEtl/payment/index", nil)

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *syEtlPaymentRouteTestSuite) TestSyOdsMapEtlLatestPaymentRoute() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("syEtl/payment/map-etl-payment", SyOdsMapEtlLatestPayment)
	req, _ := http.NewRequest(http.MethodGet, "/syEtl/payment/map-etl-payment", nil)

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *syEtlPaymentRouteTestSuite) TestGetSyEtlPaymentRoute() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("syEtl/payment/detail/:policyNumber", GetSyEtlPayment)
	req, _ := http.NewRequest(http.MethodGet, "/syEtl/payment/detail/300123456", nil)

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *syEtlPaymentRouteTestSuite) TestCreateSyEtlPaymentRoute() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	r.POST("syOdsEtl/payment/create", CreateSyEtlPayment)

	jsonValue, _ := json.Marshal(dummySyEtlPayment[0])
	req, _ := http.NewRequest(http.MethodPost, "syOdsEtl/payment/create", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *syEtlPaymentRouteTestSuite) TestTruncateTableSyEtlPaymentsRoute() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("syEtl/payment/remove-before-map", TruncateTableSyEtlPayments)
	req, _ := http.NewRequest(http.MethodGet, "/syEtl/payment/remove-before-map", nil)

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}
