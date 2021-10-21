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
var dummyPaymentMethodRate = []*entity.PaymentMethodRate{
	&entity.PaymentMethodRate{
		ID               		: 1,
		PaymentMethodCode 		: "indomaret",
		MinTransaction  		: 500001,
		MaxTransaction  		: 1000000,
		// TransactionFee  		: 6500.00,
		FormulaFee				: "6500.00",
		Premi  					: 0.00,
		CreatedAt				: time.Now(),
		UpdatedAt				: time.Now(),
	}, &entity.PaymentMethodRate{
		ID               		: 2,
		PaymentMethodCode 		: "visa-master-test",
		MinTransaction  		: 1,
		MaxTransaction  		: 1000000,
		// TransactionFee  		: 6500.00,
		FormulaFee				: "6500.00",
		Premi  					: 0.00,
		CreatedAt				: time.Now(),
		UpdatedAt				: time.Now(),
	},
}

type paymentMethodRateRouteMock struct {
	mock.Mock
}

func (s *paymentMethodRateRouteMock) SavePaymentMethodRate(paymentMethodRate entity.PaymentMethodRate) (int, error) {
	return 0, nil
}

func (s *paymentMethodRateRouteMock) UpdatePaymentMethodRate(paymentMethodRate entity.PaymentMethodRate) error {
	return nil
}

func (s *paymentMethodRateRouteMock) DeletePaymentMethodRate(paymentMethodRate entity.PaymentMethodRate) error {
	return nil
}

func (s *paymentMethodRateRouteMock) GetAllPaymentMethodRates() []entity.PaymentMethodRate {
	return nil
}

func (s *paymentMethodRateRouteMock) GetPaymentMethodRate(id string) []entity.PaymentMethodRate {
	return nil
}

type PaymentMethodRateRouteTestSuite struct {
	suite.Suite
	serviceTest usecases.PaymentMethodRateService
}

func (suite *PaymentMethodRateRouteTestSuite) SetupTest() {
	suite.serviceTest = new(paymentMethodRateRouteMock)
}

func (suite *PaymentMethodRateRouteTestSuite) TestA_SaveDelivery() {
	// fmt.Println("SAVE USER TEST DELIV")
	// suite.serviceTest.(*paymentMethodRateRouteMock).On("SavePaymentMethodRate", dummyPaymentMethodRate[0]).Return(nil)
	// c, _ := gin.CreateTestContext(httptest.NewRecorder())
	// c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}
	// PaymentMethodRateCreate(c)

	// fmt.Println("SAVE USER TEST DELIV 2")
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	handlerPaymentMethodRate := &paymentMethodRateController{
		usecases: suite.serviceTest,
	}
	r.POST("paymentMethodRate/create", handlerPaymentMethodRate.PaymentMethodRateCreate)

	jsonValue, _ := json.Marshal(dummyPaymentMethodRate[1])
	req, _ := http.NewRequest(http.MethodPost, "/paymentMethodRate/create", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *PaymentMethodRateRouteTestSuite) TestUpdateDelivery() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	handlerPaymentMethodRate := &paymentMethodRateController{
		usecases: suite.serviceTest,
	}
	r.POST("paymentMethodRate/update", handlerPaymentMethodRate.PaymentMethodRateUpdate)

	jsonValue, _ := json.Marshal(dummyPaymentMethodRate[0])
	req, _ := http.NewRequest(http.MethodPost, "/paymentMethodRate/update", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	fmt.Println("UPDATE USER TEST")
	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *PaymentMethodRateRouteTestSuite) TestDeleteDelivery() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	handlerPaymentMethodRate := &paymentMethodRateController{
		usecases: suite.serviceTest,
	}
	r.POST("paymentMethodRate/delete", handlerPaymentMethodRate.PaymentMethodRateDelete)

	jsonValue, _ := json.Marshal(dummyPaymentMethodRate[0])
	req, _ := http.NewRequest(http.MethodPost, "/paymentMethodRate/delete", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *PaymentMethodRateRouteTestSuite) TestPaymentMethodRatesIndexRoute() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	handlerPaymentMethodRate := &paymentMethodRateController{
		usecases: suite.serviceTest,
	}
	r.GET("paymentMethodRate/index", handlerPaymentMethodRate.PaymentMethodRatesIndex)
	req, _ := http.NewRequest(http.MethodGet, "/paymentMethodRate/index", nil)

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *PaymentMethodRateRouteTestSuite) TestPaymentMethodRatesDetailRoute() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	handlerPaymentMethodRate := &paymentMethodRateController{
		usecases: suite.serviceTest,
	}
	r.GET("paymentMethodRate/detail/:id", handlerPaymentMethodRate.PaymentMethodRatesDetail)
	req, _ := http.NewRequest(http.MethodGet, "/paymentMethodRate/detail/1", nil)

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func TestPaymentMethodRateRouteTestSuite(t *testing.T) {
	suite.Run(t, new(PaymentMethodRateRouteTestSuite))
}