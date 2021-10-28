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
var dummyPaymentMethodLink = []*entity.PaymentMethodLink{
	&entity.PaymentMethodLink{
		ID               		: 1,
		PaymentMethodCode 		: "visa-master-test",
		ProcessType	 			: "payment",
		Url 					: "https://visa.payment.link/:accountid",
		CreatedAt				: time.Now(),
		UpdatedAt				: time.Now(),
	}, &entity.PaymentMethodLink{
		ID               		: 2,
		PaymentMethodCode 		: "gopay-test",
		ProcessType	 			: "payment",
		Url 					: "https://gopay.payment.link/:accountid",
		CreatedAt				: time.Now(),
		UpdatedAt				: time.Now(),
	},
}

type paymentMethodLinkRouteMock struct {
	mock.Mock
}

func (s *paymentMethodLinkRouteMock) SavePaymentMethodLink(paymentMethodLink entity.PaymentMethodLink) (int, error) {
	return 0, nil
}

func (s *paymentMethodLinkRouteMock) UpdatePaymentMethodLink(paymentMethodLink entity.PaymentMethodLink) error {
	return nil
}

func (s *paymentMethodLinkRouteMock) DeletePaymentMethodLink(paymentMethodLink entity.PaymentMethodLink) error {
	return nil
}

func (s *paymentMethodLinkRouteMock) GetAllPaymentMethodLinks() []entity.PaymentMethodLink {
	return nil
}

func (s *paymentMethodLinkRouteMock) GetPaymentMethodLink(id string) []entity.PaymentMethodLink {
	return nil
}

func (s *paymentMethodLinkRouteMock) GetPaymentMethodLinkByCode(code string) entity.PaymentMethodLink {
	result := entity.PaymentMethodLink{}
	result = *dummyPaymentMethodLink[0]
	return result
}

func (s *paymentMethodLinkRouteMock) GetPaymentMethodLinkByCodeAndProcessType(code string, processType string) entity.PaymentMethodLink {
	result := entity.PaymentMethodLink{}
	result = *dummyPaymentMethodLink[0]
	return result
}

func (s *paymentMethodLinkRouteMock) ProcessPaymentForSettlement(proposalNumber string, paymentMethodLink entity.PaymentMethodLink) {
	
}

type PaymentMethodLinkRouteTestSuite struct {
	suite.Suite
	serviceTest usecases.PaymentMethodLinkService
}

func (suite *PaymentMethodLinkRouteTestSuite) SetupTest() {
	suite.serviceTest = new(paymentMethodLinkRouteMock)
}

func (suite *PaymentMethodLinkRouteTestSuite) TestA_SaveDelivery() {
	// fmt.Println("SAVE USER TEST DELIV")
	// suite.serviceTest.(*paymentMethodLinkRouteMock).On("SavePaymentMethodLink", dummyPaymentMethodLink[0]).Return(nil)
	// c, _ := gin.CreateTestContext(httptest.NewRecorder())
	// c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}
	// PaymentMethodLinkCreate(c)

	// fmt.Println("SAVE USER TEST DELIV 2")
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	handlerPaymentMethodLink := &paymentMethodLinkController{
		usecases: suite.serviceTest,
	}
	r.POST("paymentMethodLink/create", handlerPaymentMethodLink.PaymentMethodLinkCreate)

	jsonValue, _ := json.Marshal(dummyPaymentMethodLink[1])
	req, _ := http.NewRequest(http.MethodPost, "/paymentMethodLink/create", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *PaymentMethodLinkRouteTestSuite) TestUpdateDelivery() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	handlerPaymentMethodLink := &paymentMethodLinkController{
		usecases: suite.serviceTest,
	}
	r.POST("paymentMethodLink/update", handlerPaymentMethodLink.PaymentMethodLinkUpdate)

	jsonValue, _ := json.Marshal(dummyPaymentMethodLink[0])
	req, _ := http.NewRequest(http.MethodPost, "/paymentMethodLink/update", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	fmt.Println("UPDATE USER TEST")
	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *PaymentMethodLinkRouteTestSuite) TestDeleteDelivery() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	handlerPaymentMethodLink := &paymentMethodLinkController{
		usecases: suite.serviceTest,
	}
	r.POST("paymentMethodLink/delete", handlerPaymentMethodLink.PaymentMethodLinkDelete)

	jsonValue, _ := json.Marshal(dummyPaymentMethodLink[0])
	req, _ := http.NewRequest(http.MethodPost, "/paymentMethodLink/delete", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *PaymentMethodLinkRouteTestSuite) TestPaymentMethodLinksIndexRoute() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	handlerPaymentMethodLink := &paymentMethodLinkController{
		usecases: suite.serviceTest,
	}
	r.GET("paymentMethodLink/index", handlerPaymentMethodLink.PaymentMethodLinksIndex)
	req, _ := http.NewRequest(http.MethodGet, "/paymentMethodLink/index", nil)

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *PaymentMethodLinkRouteTestSuite) TestPaymentMethodLinksDetailRoute() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	handlerPaymentMethodLink := &paymentMethodLinkController{
		usecases: suite.serviceTest,
	}
	r.GET("paymentMethodLink/detail/:id", handlerPaymentMethodLink.PaymentMethodLinksDetail)
	req, _ := http.NewRequest(http.MethodGet, "/paymentMethodLink/detail/1", nil)

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func TestPaymentMethodLinkRouteTestSuite(t *testing.T) {
	suite.Run(t, new(PaymentMethodLinkRouteTestSuite))
}