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
var dummyPaymentMethodRate = []entity.PaymentMethodRate{
	entity.PaymentMethodRate{
		ID				: 1,
		PaymentMethodCode		: "indomaret",
		MinTransaction  		: 1,
		MaxTransaction  		: 500000,
		// TransactionFee  		: 3750.00,
		FormulaFee				: "3750.00",
		Premi  					: 0.00,
		CreatedAt		: time.Now(),
		UpdatedAt		: time.Now(),
	}, entity.PaymentMethodRate{
		ID				: 2,
		PaymentMethodCode 		: "indomaret",
		MinTransaction  		: 500001,
		MaxTransaction  		: 1000000,
		// TransactionFee  		: 6500.00,
		FormulaFee				: "6500.00",
		Premi  					: 0.00,
		CreatedAt		: time.Now(),
		UpdatedAt		: time.Now(),
	},
}

type repoMockPaymentMethodRate struct {
	mock.Mock
}

func (r *repoMockPaymentMethodRate) SavePaymentMethodRate(paymentMethodRate entity.PaymentMethodRate) (int, error) {
	return 0, nil
}

func (r *repoMockPaymentMethodRate) UpdatePaymentMethodRate(paymentMethodRate entity.PaymentMethodRate) error {
	args := r.Called(paymentMethodRate)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (r *repoMockPaymentMethodRate) DeletePaymentMethodRate(paymentMethodRate entity.PaymentMethodRate) error {
	args := r.Called(paymentMethodRate)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (r *repoMockPaymentMethodRate) GetAllPaymentMethodRates() []entity.PaymentMethodRate {
	return dummyPaymentMethodRate
}

func (r *repoMockPaymentMethodRate) GetPaymentMethodRate(id string) []entity.PaymentMethodRate {
	return dummyPaymentMethodRate
}

func (r *repoMockPaymentMethodRate) GetPaymentMethodRateByUuid(uuid string) entity.PaymentMethodRate {
	return dummyPaymentMethodRate[0]
}

func (r *repoMockPaymentMethodRate) GetActivePaymentMethodRateByCode(code string) entity.PaymentMethodRate {
	return dummyPaymentMethodRate[0]
}

func (r *repoMockPaymentMethodRate) GetPaymentMethodRateByCode(code string) entity.PaymentMethodRate {
	return dummyPaymentMethodRate[0]
}

func (r *repoMockPaymentMethodRate) CloseDB() {
}

type PaymentMethodRateUsecaseTestSuite struct {
	suite.Suite
	repositoryTest repositories.PaymentMethodRateRepository
}

func (suite *PaymentMethodRateUsecaseTestSuite) SetupTest() {
	suite.repositoryTest = new(repoMockPaymentMethodRate)
}

func (suite *PaymentMethodRateUsecaseTestSuite) TestBuildPaymentMethodRateService() {
	resultTest := NewPaymentMethodRateService(suite.repositoryTest)
	var dummyImpl *PaymentMethodRateService
	assert.NotNil(suite.T(), resultTest)
	assert.Implements(suite.T(), dummyImpl, resultTest)
	// assert.NotNil(suite.T(), resultTest.(*PaymentMethodRateService).repositories)
}

func (suite *PaymentMethodRateUsecaseTestSuite) TestSavePaymentMethodRateUsecase() {
	suite.repositoryTest.(*repoMockPaymentMethodRate).On("SavePaymentMethodRate", dummyPaymentMethodRate[0]).Return(nil)
	useCaseTest := NewPaymentMethodRateService(suite.repositoryTest)
	// dummyPaymentMethodRate[0].Password = "Change Password"
	data, _ := useCaseTest.SavePaymentMethodRate(dummyPaymentMethodRate[0])
	assert.NotNil(suite.T(), data)
}

func (suite *PaymentMethodRateUsecaseTestSuite) TestUpdatePaymentMethodRateUsecase() {
	suite.repositoryTest.(*repoMockPaymentMethodRate).On("UpdatePaymentMethodRate", dummyPaymentMethodRate[0]).Return(nil)
	useCaseTest := NewPaymentMethodRateService(suite.repositoryTest)
	err := useCaseTest.UpdatePaymentMethodRate(dummyPaymentMethodRate[0])
	assert.Nil(suite.T(), err)
}

func (suite *PaymentMethodRateUsecaseTestSuite) TestDeletePaymentMethodRateUsecase() {
	suite.repositoryTest.(*repoMockPaymentMethodRate).On("DeletePaymentMethodRate", dummyPaymentMethodRate[0]).Return(nil)
	useCaseTest := NewPaymentMethodRateService(suite.repositoryTest)
	err := useCaseTest.DeletePaymentMethodRate(dummyPaymentMethodRate[0])
	assert.Nil(suite.T(), err)
}

func (suite *PaymentMethodRateUsecaseTestSuite) TestGetAllPaymentMethodRates() {
	suite.repositoryTest.(*repoMockPaymentMethodRate).On("GetAllPaymentMethodRates", dummyPaymentMethodRate).Return(dummyPaymentMethodRate)
	useCaseTest := NewPaymentMethodRateService(suite.repositoryTest)
	dummyPaymentMethodRate := useCaseTest.GetAllPaymentMethodRates()
	assert.Equal(suite.T(), dummyPaymentMethodRate, dummyPaymentMethodRate)
}

func (suite *PaymentMethodRateUsecaseTestSuite) TestGetPaymentMethodRate() {
	suite.repositoryTest.(*repoMockPaymentMethodRate).On("GetPaymentMethodRate", dummyPaymentMethodRate[0].ID).Return(dummyPaymentMethodRate[0], nil)
	useCaseTest := NewPaymentMethodRateService(suite.repositoryTest)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}
	dummyPaymentMethodRate := useCaseTest.GetPaymentMethodRate(c.Param("id"))
	assert.NotNil(suite.T(), dummyPaymentMethodRate[0])
	assert.Equal(suite.T(), dummyPaymentMethodRate[0], dummyPaymentMethodRate[0])
}

func TestPaymentMethodRateUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(PaymentMethodRateUsecaseTestSuite))
}