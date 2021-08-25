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

type syUserInvoiceRouteMock struct {
	mock.Mock
}

func (s *syUserInvoiceRouteMock) GetAllSyUserInvoices() []entity.SyUserInvoice {
	return nil
}

func (s *syUserInvoiceRouteMock) SyMapEtlLatestPayment() []entity.SyUserInvoice {
	return nil
}

type syUserInvoiceRouteTestSuite struct {
	suite.Suite
	serviceTest usecases.SyUserInvoiceService
}

func (suite *syUserInvoiceRouteTestSuite) SetupSyUserInvoiceTest() {
	suite.serviceTest = new(syUserInvoiceRouteMock)
}

func (suite *syUserInvoiceRouteTestSuite) TestsyUserInvoiceIndexRoute() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("syUserInvoice/index", GetAllUserPolicies)
	req, _ := http.NewRequest(http.MethodGet, "/syUserInvoice/index", nil)

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *syUserInvoiceRouteTestSuite) TestSyMapEtlLatestPaymentRoute() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("syUserInvoice/map-etl-payment", SyMapEtlLatestPayment)
	req, _ := http.NewRequest(http.MethodGet, "/syUserInvoice/map-etl-payment", nil)

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}
