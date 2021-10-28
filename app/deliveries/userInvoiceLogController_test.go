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
var dummyUserInvoiceLog = []*entity.UserInvoiceLog{
	&entity.UserInvoiceLog{
		ID               		: 1,
		PaymentMethodCode 		: "indomaret-test",
		SummaryToken  			: "abcd1084ufnouu23af",
		// TransactionFee  		: 6500.00,
		InvoiceNumber			: "12369412964",
		PaymentCycle  			: "YEARLY",
		TransactionFee  		: "20000",
		AgentFee  				: "20000",
		TotalPremium  			: "40000",
		TotalPayment  			: "80000",
		CreatedAt				: time.Now(),
		UpdatedAt				: time.Now(),
	}, &entity.UserInvoiceLog{
		ID               		: 2,
		PaymentMethodCode 		: "visa-master-test",
		SummaryToken  			: "abcd1084ufnouu23af",
		// TransactionFee  		: 6500.00,
		InvoiceNumber			: "12369412964",
		PaymentCycle  			: "YEARLY",
		TransactionFee  		: "20000",
		AgentFee  				: "20000",
		TotalPremium  			: "40000",
		TotalPayment  			: "80000",
		CreatedAt				: time.Now(),
		UpdatedAt				: time.Now(),
	},
}

type userInvoiceLogRouteMock struct {
	mock.Mock
}

func (s *userInvoiceLogRouteMock) SaveUserInvoiceLog(userInvoiceLog entity.UserInvoiceLog) (int, error) {
	return 0, nil
}

func (s *userInvoiceLogRouteMock) UpdateUserInvoiceLog(userInvoiceLog entity.UserInvoiceLog) error {
	return nil
}

func (s *userInvoiceLogRouteMock) DeleteUserInvoiceLog(userInvoiceLog entity.UserInvoiceLog) error {
	return nil
}

func (s *userInvoiceLogRouteMock) GetAllUserInvoiceLogs() []entity.UserInvoiceLog {
	return nil
}

func (s *userInvoiceLogRouteMock) GetUserInvoiceLog(id string) []entity.UserInvoiceLog {
	return nil
}

type UserInvoiceLogRouteTestSuite struct {
	suite.Suite
	serviceTest usecases.UserInvoiceLogService
}

func (suite *UserInvoiceLogRouteTestSuite) SetupTest() {
	suite.serviceTest = new(userInvoiceLogRouteMock)
}

func (suite *UserInvoiceLogRouteTestSuite) TestA_SaveDelivery() {
	// fmt.Println("SAVE USER TEST DELIV")
	// suite.serviceTest.(*userInvoiceLogRouteMock).On("SaveUserInvoiceLog", dummyUserInvoiceLog[0]).Return(nil)
	// c, _ := gin.CreateTestContext(httptest.NewRecorder())
	// c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}
	// UserInvoiceLogCreate(c)

	// fmt.Println("SAVE USER TEST DELIV 2")
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	handlerUserInvoiceLog := &userInvoiceLogController{
		usecases: suite.serviceTest,
	}
	r.POST("userInvoiceLog/create", handlerUserInvoiceLog.UserInvoiceLogCreate)

	jsonValue, _ := json.Marshal(dummyUserInvoiceLog[1])
	req, _ := http.NewRequest(http.MethodPost, "/userInvoiceLog/create", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *UserInvoiceLogRouteTestSuite) TestUpdateDelivery() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	handlerUserInvoiceLog := &userInvoiceLogController{
		usecases: suite.serviceTest,
	}
	r.POST("userInvoiceLog/update", handlerUserInvoiceLog.UserInvoiceLogUpdate)

	jsonValue, _ := json.Marshal(dummyUserInvoiceLog[0])
	req, _ := http.NewRequest(http.MethodPost, "/userInvoiceLog/update", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	fmt.Println("UPDATE USER TEST")
	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *UserInvoiceLogRouteTestSuite) TestDeleteDelivery() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	handlerUserInvoiceLog := &userInvoiceLogController{
		usecases: suite.serviceTest,
	}
	r.POST("userInvoiceLog/delete", handlerUserInvoiceLog.UserInvoiceLogDelete)

	jsonValue, _ := json.Marshal(dummyUserInvoiceLog[0])
	req, _ := http.NewRequest(http.MethodPost, "/userInvoiceLog/delete", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *UserInvoiceLogRouteTestSuite) TestUserInvoiceLogsIndexRoute() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	handlerUserInvoiceLog := &userInvoiceLogController{
		usecases: suite.serviceTest,
	}
	r.GET("userInvoiceLog/index", handlerUserInvoiceLog.UserInvoiceLogsIndex)
	req, _ := http.NewRequest(http.MethodGet, "/userInvoiceLog/index", nil)

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *UserInvoiceLogRouteTestSuite) TestUserInvoiceLogsDetailRoute() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	handlerUserInvoiceLog := &userInvoiceLogController{
		usecases: suite.serviceTest,
	}
	r.GET("userInvoiceLog/detail/:id", handlerUserInvoiceLog.UserInvoiceLogsDetail)
	req, _ := http.NewRequest(http.MethodGet, "/userInvoiceLog/detail/1", nil)

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func TestUserInvoiceLogRouteTestSuite(t *testing.T) {
	suite.Run(t, new(UserInvoiceLogRouteTestSuite))
}