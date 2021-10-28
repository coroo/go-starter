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
var dummyPaymentMethod = []*entity.PaymentMethod{
	&entity.PaymentMethod{
		ID               		: 1,
		Code 					: "visa-master-test",
		InitPaymentCode  		: "PDCC",
		RenewalPaymentCode 		: "CASH",
		FastpayCode 		 	: "cc",
		Name 		 			: "Visa Master Test",
		PaymentLogo 		 	: "visa-master.png",
		Status 		 			: "active",
		CreatedAt				: time.Now(),
		UpdatedAt				: time.Now(),
	}, &entity.PaymentMethod{
		ID               		: 2,
		Code 					: "gopay",
		InitPaymentCode  		: "PDDC4",
		RenewalPaymentCode 		: "PDDC4",
		FastpayCode 		 	: "90647",
		BankCode 		 		: "333333",
		Name 		 			: "Go Pay",
		PaymentLogo 		 	: "gopay.png",
		Status 		 			: "active",
		CreatedAt				: time.Now(),
		UpdatedAt				: time.Now(),
	},
}

type paymentMethodRouteMock struct {
	mock.Mock
}

func (s *paymentMethodRouteMock) SavePaymentMethod(paymentMethod entity.PaymentMethod) (int, error) {
	return 0, nil
}

func (s *paymentMethodRouteMock) UpdatePaymentMethod(paymentMethod entity.PaymentMethod) error {
	return nil
}

func (s *paymentMethodRouteMock) DeletePaymentMethod(paymentMethod entity.PaymentMethod) error {
	return nil
}

func (s *paymentMethodRouteMock) GetAllPaymentMethods(total_premium string, status string) []entity.PaymentMethodWithPremium {
	return nil
}

func (s *paymentMethodRouteMock) GetPaymentMethod(id string) []entity.PaymentMethod {
	return nil
}

func (s *paymentMethodRouteMock) ConnectionTest() *http.Response{
	return nil
}

func (s *paymentMethodRouteMock) GenerateVaSignature(code string, proposal_group_number string) string{
	return ""
}

func (s *paymentMethodRouteMock) GetPaymentMethodByCode(code string) entity.PaymentMethod{
	result := entity.PaymentMethod{}
	result = *dummyPaymentMethod[0]
	return result
}

type PaymentMethodRouteTestSuite struct {
	suite.Suite
	serviceTest usecases.PaymentMethodService
}

func (suite *PaymentMethodRouteTestSuite) SetupTest() {
	suite.serviceTest = new(paymentMethodRouteMock)
}

func (suite *PaymentMethodRouteTestSuite) TestA_SaveDelivery() {
	// fmt.Println("SAVE USER TEST DELIV")
	// suite.serviceTest.(*paymentMethodRouteMock).On("SavePaymentMethod", dummyPaymentMethod[0]).Return(nil)
	// c, _ := gin.CreateTestContext(httptest.NewRecorder())
	// c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}
	// PaymentMethodCreate(c)

	// fmt.Println("SAVE USER TEST DELIV 2")
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	handlerPaymentMethod := &paymentMethodController{
		usecases: suite.serviceTest,
	}
	r.POST("paymentMethod/create", handlerPaymentMethod.PaymentMethodCreate)

	jsonValue, _ := json.Marshal(dummyPaymentMethod[1])
	req, _ := http.NewRequest(http.MethodPost, "/paymentMethod/create", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *PaymentMethodRouteTestSuite) TestUpdateDelivery() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	handlerPaymentMethod := &paymentMethodController{
		usecases: suite.serviceTest,
	}
	r.POST("paymentMethod/update", handlerPaymentMethod.PaymentMethodUpdate)

	jsonValue, _ := json.Marshal(dummyPaymentMethod[0])
	req, _ := http.NewRequest(http.MethodPost, "/paymentMethod/update", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	fmt.Println("UPDATE USER TEST")
	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *PaymentMethodRouteTestSuite) TestDeleteDelivery() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	handlerPaymentMethod := &paymentMethodController{
		usecases: suite.serviceTest,
	}
	r.POST("paymentMethod/delete", handlerPaymentMethod.PaymentMethodDelete)

	jsonValue, _ := json.Marshal(dummyPaymentMethod[0])
	req, _ := http.NewRequest(http.MethodPost, "/paymentMethod/delete", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *PaymentMethodRouteTestSuite) TestPaymentMethodsIndexRoute() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	handlerPaymentMethod := &paymentMethodController{
		usecases: suite.serviceTest,
	}
	r.GET("paymentMethod/index", handlerPaymentMethod.PaymentMethodsIndex)
	req, _ := http.NewRequest(http.MethodGet, "/paymentMethod/index", nil)

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *PaymentMethodRouteTestSuite) TestPaymentMethodsDetailRoute() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	handlerPaymentMethod := &paymentMethodController{
		usecases: suite.serviceTest,
	}
	r.GET("paymentMethod/detail/:id", handlerPaymentMethod.PaymentMethodsDetail)
	req, _ := http.NewRequest(http.MethodGet, "/paymentMethod/detail/1", nil)

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *PaymentMethodRouteTestSuite) TestPaymentMethodsDetailByCodeRoute() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	handlerPaymentMethod := &paymentMethodController{
		usecases: suite.serviceTest,
	}
	r.GET("paymentMethod/detail-by-code/visa-master-test", handlerPaymentMethod.PaymentMethodsDetailByCode)
	req, _ := http.NewRequest(http.MethodGet, "/paymentMethod/detail-by-code/visa-master-test", nil)

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *PaymentMethodRouteTestSuite) TestGenerateVaSignatureRoute() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	handlerPaymentMethod := &paymentMethodController{
		usecases: suite.serviceTest,
	}
	r.GET("paymentMethod/fastpay/generate-va-signature", handlerPaymentMethod.GenerateVaSignature)
	req, _ := http.NewRequest(http.MethodGet, "/paymentMethod/fastpay/generate-va-signature?payment_channel=visa-master-test;proposal_group_number=12391235471243", nil)

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *PaymentMethodRouteTestSuite) TestConnectionTestRoute() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	handlerPaymentMethod := &paymentMethodController{
		usecases: suite.serviceTest,
	}
	r.GET("paymentMethod/fastpay/connection-test", handlerPaymentMethod.GenerateVaSignature)
	req, _ := http.NewRequest(http.MethodGet, "/paymentMethod/fastpay/connection-test", nil)

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func TestPaymentMethodRouteTestSuite(t *testing.T) {
	suite.Run(t, new(PaymentMethodRouteTestSuite))
}