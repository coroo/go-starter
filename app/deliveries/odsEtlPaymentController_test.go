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

var dummyOdsEtlPayment = []*entity.OdsEtlPayment{
	&entity.OdsEtlPayment{
		ID					: 1,
		CollectionId		: "1234567",
		ProposalNumber		: "12345678",
		PolicyNumber		: "6001234567",
		FirstPaymentDate	: time.Now(),
		PaymentDate			: time.Now(),
		PaymentMethod		: "indomaret",
		TotalAmount			: 100000,
		UpdatedAt			: time.Now(),
	},
}

type odsEtlPaymentRouteMock struct {
	mock.Mock
}

func (s *odsEtlPaymentRouteMock) CreateOdsEtlPayment(odsEtlPayment entity.OdsEtlPayment) error {
	return nil
}

func (s *odsEtlPaymentRouteMock) GetAllOdsEtlPayments() []entity.OdsEtlPayment {
	return nil
}
func (s *odsEtlPaymentRouteMock) GetOdsEtlPayment(id string) []entity.OdsEtlPayment {
	return nil
}
func (s *odsEtlPaymentRouteMock) TruncateTableOdsEtlPayments() error {
	return nil
}

type odsEtlPaymentRouteTestSuite struct {
	suite.Suite
	serviceTest usecases.OdsEtlPaymentService
}

func (suite *odsEtlPaymentRouteTestSuite) SetupOdsEtlPaymentTest() {
	suite.serviceTest = new(odsEtlPaymentRouteMock)
}

func (suite *odsEtlPaymentRouteTestSuite) TestGetAllOdsEtlPaymentsRoute() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	handlerUser := &odsEtlPaymentController{
		usecases: suite.serviceTest,
	}
	r.GET("odsEtl/payment/index", handlerUser.GetAllOdsEtlPayments)
	req, _ := http.NewRequest(http.MethodGet, "/odsEtl/payment/index", nil)

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *odsEtlPaymentRouteTestSuite) TestGetOdsEtlPaymentRoute() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	handlerUser := &odsEtlPaymentController{
		usecases: suite.serviceTest,
	}
	r.GET("odsEtl/payment/detail/:id", handlerUser.GetOdsEtlPayment)
	req, _ := http.NewRequest(http.MethodGet, "/odsEtl/payment/detail/1", nil)

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *odsEtlPaymentRouteTestSuite) TestCreateOdsEtlPaymentRoute() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	handlerUser := &odsEtlPaymentController{
		usecases: suite.serviceTest,
	}
	r.POST("odsEtl/payment/create", handlerUser.CreateOdsEtlPayment)

	jsonValue, _ := json.Marshal(dummyOdsEtlPayment[0])
	req, _ := http.NewRequest(http.MethodPost, "odsEtl/payment/create", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *odsEtlPaymentRouteTestSuite) TestTruncateTableOdsEtlPaymentsRoute() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	handlerUser := &odsEtlPaymentController{
		usecases: suite.serviceTest,
	}
	r.GET("odsEtl/payment/remove-before-map", handlerUser.TruncateTableOdsEtlPayments)
	req, _ := http.NewRequest(http.MethodGet, "/odsEtl/payment/remove-before-map", nil)

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}
