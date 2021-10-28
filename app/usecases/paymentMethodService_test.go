package usecases

import (
	"net/http/httptest"
	"testing"
	"time"

	entity "github.com/coroo/go-starter/app/entity"
	repositories "github.com/coroo/go-starter/app/repositories"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// dummy data
var dummyPaymentMethod = []entity.PaymentMethod{
	entity.PaymentMethod{
		ID				: 1,
		Code 					: "gopay",
		InitPaymentCode  		: "PDDC4",
		RenewalPaymentCode 		: "PDDC4",
		FastpayCode 		 	: "90647",
		BankCode 		 		: "333333",
		Name 		 			: "Go Pay",
		PaymentLogo 		 	: "gopay.png",
		Status 		 			: "active",
		CreatedAt		: time.Now(),
		UpdatedAt		: time.Now(),
	}, entity.PaymentMethod{
		ID				: 2,
		Code 					: "gopay",
		InitPaymentCode  		: "PDDC4",
		RenewalPaymentCode 		: "PDDC4",
		FastpayCode 		 	: "90647",
		BankCode 		 		: "333333",
		Name 		 			: "Go Pay",
		PaymentLogo 		 	: "gopay.png",
		Status 		 			: "active",
		CreatedAt		: time.Now(),
		UpdatedAt		: time.Now(),
	},
}

type repoMockPaymentMethod struct {
	mock.Mock
}

func (r *repoMockPaymentMethod) SavePaymentMethod(paymentMethod entity.PaymentMethod) (int, error) {
	return 0, nil
}

func (r *repoMockPaymentMethod) UpdatePaymentMethod(paymentMethod entity.PaymentMethod) error {
	args := r.Called(paymentMethod)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (r *repoMockPaymentMethod) DeletePaymentMethod(paymentMethod entity.PaymentMethod) error {
	args := r.Called(paymentMethod)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (r *repoMockPaymentMethod) GetAllPaymentMethods(status string) []entity.PaymentMethod {
	return dummyPaymentMethod
}

func (r *repoMockPaymentMethod) GetPaymentMethod(id string) []entity.PaymentMethod {
	return dummyPaymentMethod
}

func (r *repoMockPaymentMethod) GetPaymentMethodByUuid(uuid string) entity.PaymentMethod {
	return dummyPaymentMethod[0]
}

func (r *repoMockPaymentMethod) GetActivePaymentMethodByCode(code string) entity.PaymentMethod {
	return dummyPaymentMethod[0]
}

func (r *repoMockPaymentMethod) GetPaymentMethodByCode(code string) entity.PaymentMethod {
	return dummyPaymentMethod[0]
}

func (r *repoMockPaymentMethod) CloseDB() {
}

type PaymentMethodUsecaseTestSuite struct {
	suite.Suite
	repositoryTest repositories.PaymentMethodRepository
}

func (suite *PaymentMethodUsecaseTestSuite) SetupTest() {
	suite.repositoryTest = new(repoMockPaymentMethod)
}

func (suite *PaymentMethodUsecaseTestSuite) TestBuildPaymentMethodService() {
	resultTest := NewPaymentMethodService(suite.repositoryTest)
	var dummyImpl *PaymentMethodService
	assert.NotNil(suite.T(), resultTest)
	assert.Implements(suite.T(), dummyImpl, resultTest)
	// assert.NotNil(suite.T(), resultTest.(*PaymentMethodService).repositories)
}

func (suite *PaymentMethodUsecaseTestSuite) TestSavePaymentMethodUsecase() {
	suite.repositoryTest.(*repoMockPaymentMethod).On("SavePaymentMethod", dummyPaymentMethod[0]).Return(nil)
	useCaseTest := NewPaymentMethodService(suite.repositoryTest)
	// dummyPaymentMethod[0].Password = "Change Password"
	data, _ := useCaseTest.SavePaymentMethod(dummyPaymentMethod[0])
	assert.NotNil(suite.T(), data)
}

func (suite *PaymentMethodUsecaseTestSuite) TestUpdatePaymentMethodUsecase() {
	suite.repositoryTest.(*repoMockPaymentMethod).On("UpdatePaymentMethod", dummyPaymentMethod[0]).Return(nil)
	useCaseTest := NewPaymentMethodService(suite.repositoryTest)
	err := useCaseTest.UpdatePaymentMethod(dummyPaymentMethod[0])
	assert.Nil(suite.T(), err)
}

func (suite *PaymentMethodUsecaseTestSuite) TestDeletePaymentMethodUsecase() {
	suite.repositoryTest.(*repoMockPaymentMethod).On("DeletePaymentMethod", dummyPaymentMethod[0]).Return(nil)
	useCaseTest := NewPaymentMethodService(suite.repositoryTest)
	err := useCaseTest.DeletePaymentMethod(dummyPaymentMethod[0])
	assert.Nil(suite.T(), err)
}

func (suite *PaymentMethodUsecaseTestSuite) TestGetAllPaymentMethods() {
	suite.repositoryTest.(*repoMockPaymentMethod).On("GetAllPaymentMethods", dummyPaymentMethod).Return(dummyPaymentMethod)
	useCaseTest := NewPaymentMethodService(suite.repositoryTest)
	dummyPaymentMethod := useCaseTest.GetAllPaymentMethods("","")
	assert.Equal(suite.T(), dummyPaymentMethod, dummyPaymentMethod)
}

func (suite *PaymentMethodUsecaseTestSuite) TestGetPaymentMethod() {
	suite.repositoryTest.(*repoMockPaymentMethod).On("GetPaymentMethod", dummyPaymentMethod[0].ID).Return(dummyPaymentMethod[0], nil)
	useCaseTest := NewPaymentMethodService(suite.repositoryTest)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}
	dummyPaymentMethod := useCaseTest.GetPaymentMethod(c.Param("id"))
	assert.NotNil(suite.T(), dummyPaymentMethod[0])
	assert.Equal(suite.T(), dummyPaymentMethod[0], dummyPaymentMethod[0])
}

func (suite *PaymentMethodUsecaseTestSuite) TestGetPaymentMethodByCode() {
	suite.repositoryTest.(*repoMockPaymentMethod).On("GetPaymentMethod", dummyPaymentMethod[0].ID).Return(dummyPaymentMethod[0], nil)
	useCaseTest := NewPaymentMethodService(suite.repositoryTest)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	getByCode := useCaseTest.GetPaymentMethodByCode(c.Param("gopay"))
	assert.NotNil(suite.T(), dummyPaymentMethod)
	assert.Equal(suite.T(), getByCode, dummyPaymentMethod[0])
}

func TestPaymentMethodUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(PaymentMethodUsecaseTestSuite))
}