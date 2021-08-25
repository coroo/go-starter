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

var dummyPayment = []*entity.SyOdsEtlPayment{
	&entity.SyOdsEtlPayment{
		ID					: 1,
		ProposalNumber		: "600012345678",
		PolicyNumber		: "3001234567",
		OdsFirstPaidDate	: time.Now(),
		OdsPaidDate			: time.Now(),
		SyPaidDate			: time.Now(),
		PaymentMethodName	: "Indomaret",
		CollectionId		: "1234567",
		SyTotalAmount		: 100000,
		OdsTotalAmount		: 100000,
		PolicyStatus		: "success",
		Status				: "closed",
		StatusDescription	: "Stated as first payment (TEST)",
		UpdatedAt			: time.Now(),
	},
}

type syOdsEtlPaymentRouteMock struct {
	mock.Mock
}

func (s *syOdsEtlPaymentRouteMock) CreateSyOdsEtlPayment(SyOdsEtlPayment entity.SyOdsEtlPayment) error {
	return nil
}

func (s *syOdsEtlPaymentRouteMock) GetAllSyOdsEtlPayments() []entity.SyOdsEtlPayment {
	return nil
}
func (s *syOdsEtlPaymentRouteMock) GetSyOdsEtlPaymentByPolicyNumber(policyNumber string) []entity.SyOdsEtlPayment {
	return nil
}
func (s *syOdsEtlPaymentRouteMock) GetSyOdsEtlPaymentByStatus(status string) []entity.SyOdsEtlPayment {
	return nil
}
func (s *syOdsEtlPaymentRouteMock) GetSyOdsEtlPaymentDailyByStatus(status string) []entity.SyOdsEtlPayment {
	return nil
}
func (s *syOdsEtlPaymentRouteMock) CancelOutstandingSyOdsEtlPayments() []entity.SyOdsEtlPayment {
	return nil
}

type syOdsEtlPaymentRouteTestSuite struct {
	suite.Suite
	serviceTest usecases.SyOdsEtlPaymentService
}

func (suite *syOdsEtlPaymentRouteTestSuite) SetupsyOdsEtlPaymentTest() {
	suite.serviceTest = new(syOdsEtlPaymentRouteMock)
}

func (suite *syOdsEtlPaymentRouteTestSuite) TestSyOdsEtlPaymentIndexRoute() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("syOdsEtl/payment/index", GetAllSyOdsEtlPayments)
	req, _ := http.NewRequest(http.MethodGet, "/syOdsEtl/payment/index", nil)

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *syOdsEtlPaymentRouteTestSuite) TestSyOdsEtlPaymentByPolicyNumberRoute() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("syOdsEtl/payment/detail/:policyNumber", GetSyOdsEtlPaymentByPolicyNumber)
	req, _ := http.NewRequest(http.MethodGet, "/syOdsEtl/payment/detail/3006088126", nil)

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *syOdsEtlPaymentRouteTestSuite) TestSyOdsEtlPaymentIndexByStatusRoute() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("syOdsEtl/payment/status/:status", GetUserPolicy)
	req, _ := http.NewRequest(http.MethodGet, "/syOdsEtl/payment/status/closed", nil)

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *syOdsEtlPaymentRouteTestSuite) TestSyOdsEtlPaymentDailyByStatusRoute() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("syOdsEtl/payment/daily-by-status/:status", GetSyOdsEtlPaymentByStatus)
	req, _ := http.NewRequest(http.MethodGet, "/syOdsEtl/payment/daily-by-status/closed", nil)

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *syOdsEtlPaymentRouteTestSuite) TestCreateSyOdsEtlPaymentRoute() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	r.POST("syOdsEtl/payment/create", UserCreate)

	jsonValue, _ := json.Marshal(dummyPayment[0])
	req, _ := http.NewRequest(http.MethodPost, "syOdsEtl/payment/create", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *syOdsEtlPaymentRouteTestSuite) TestCancelOutstandingSyOdsEtlPaymentsRoute() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("syOdsEtl/payment/remove-before-map", CancelOutstandingSyOdsEtlPayments)
	req, _ := http.NewRequest(http.MethodGet, "/syOdsEtl/payment/remove-before-map", nil)

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}
