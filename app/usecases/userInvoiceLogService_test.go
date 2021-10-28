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
var dummyUserInvoiceLog = []entity.UserInvoiceLog{
	entity.UserInvoiceLog{
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
	}, entity.UserInvoiceLog{
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

type repoMockUserInvoiceLog struct {
	mock.Mock
}

func (r *repoMockUserInvoiceLog) SaveUserInvoiceLog(userInvoiceLog entity.UserInvoiceLog) (int, error) {
	return 0, nil
}

func (r *repoMockUserInvoiceLog) UpdateUserInvoiceLog(userInvoiceLog entity.UserInvoiceLog) error {
	args := r.Called(userInvoiceLog)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (r *repoMockUserInvoiceLog) DeleteUserInvoiceLog(userInvoiceLog entity.UserInvoiceLog) error {
	args := r.Called(userInvoiceLog)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (r *repoMockUserInvoiceLog) GetAllUserInvoiceLogs() []entity.UserInvoiceLog {
	return dummyUserInvoiceLog
}

func (r *repoMockUserInvoiceLog) GetUserInvoiceLog(id string) []entity.UserInvoiceLog {
	return dummyUserInvoiceLog
}

func (r *repoMockUserInvoiceLog) GetUserInvoiceLogByUuid(uuid string) entity.UserInvoiceLog {
	return dummyUserInvoiceLog[0]
}

func (r *repoMockUserInvoiceLog) GetActiveUserInvoiceLogByCode(code string) entity.UserInvoiceLog {
	return dummyUserInvoiceLog[0]
}

func (r *repoMockUserInvoiceLog) GetUserInvoiceLogByCode(code string) entity.UserInvoiceLog {
	return dummyUserInvoiceLog[0]
}

func (r *repoMockUserInvoiceLog) CloseDB() {
}

type UserInvoiceLogUsecaseTestSuite struct {
	suite.Suite
	repositoryTest repositories.UserInvoiceLogRepository
}

func (suite *UserInvoiceLogUsecaseTestSuite) SetupTest() {
	suite.repositoryTest = new(repoMockUserInvoiceLog)
}

func (suite *UserInvoiceLogUsecaseTestSuite) TestBuildUserInvoiceLogService() {
	resultTest := NewUserInvoiceLogService(suite.repositoryTest)
	var dummyImpl *UserInvoiceLogService
	assert.NotNil(suite.T(), resultTest)
	assert.Implements(suite.T(), dummyImpl, resultTest)
	// assert.NotNil(suite.T(), resultTest.(*UserInvoiceLogService).repositories)
}

func (suite *UserInvoiceLogUsecaseTestSuite) TestSaveUserInvoiceLogUsecase() {
	suite.repositoryTest.(*repoMockUserInvoiceLog).On("SaveUserInvoiceLog", dummyUserInvoiceLog[0]).Return(nil)
	useCaseTest := NewUserInvoiceLogService(suite.repositoryTest)
	// dummyUserInvoiceLog[0].Password = "Change Password"
	data, _ := useCaseTest.SaveUserInvoiceLog(dummyUserInvoiceLog[0])
	assert.NotNil(suite.T(), data)
}

func (suite *UserInvoiceLogUsecaseTestSuite) TestUpdateUserInvoiceLogUsecase() {
	suite.repositoryTest.(*repoMockUserInvoiceLog).On("UpdateUserInvoiceLog", dummyUserInvoiceLog[0]).Return(nil)
	useCaseTest := NewUserInvoiceLogService(suite.repositoryTest)
	err := useCaseTest.UpdateUserInvoiceLog(dummyUserInvoiceLog[0])
	assert.Nil(suite.T(), err)
}

func (suite *UserInvoiceLogUsecaseTestSuite) TestDeleteUserInvoiceLogUsecase() {
	suite.repositoryTest.(*repoMockUserInvoiceLog).On("DeleteUserInvoiceLog", dummyUserInvoiceLog[0]).Return(nil)
	useCaseTest := NewUserInvoiceLogService(suite.repositoryTest)
	err := useCaseTest.DeleteUserInvoiceLog(dummyUserInvoiceLog[0])
	assert.Nil(suite.T(), err)
}

func (suite *UserInvoiceLogUsecaseTestSuite) TestGetAllUserInvoiceLogs() {
	suite.repositoryTest.(*repoMockUserInvoiceLog).On("GetAllUserInvoiceLogs", dummyUserInvoiceLog).Return(dummyUserInvoiceLog)
	useCaseTest := NewUserInvoiceLogService(suite.repositoryTest)
	dummyUserInvoiceLog := useCaseTest.GetAllUserInvoiceLogs()
	assert.Equal(suite.T(), dummyUserInvoiceLog, dummyUserInvoiceLog)
}

func (suite *UserInvoiceLogUsecaseTestSuite) TestGetUserInvoiceLog() {
	suite.repositoryTest.(*repoMockUserInvoiceLog).On("GetUserInvoiceLog", dummyUserInvoiceLog[0].ID).Return(dummyUserInvoiceLog[0], nil)
	useCaseTest := NewUserInvoiceLogService(suite.repositoryTest)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}
	dummyUserInvoiceLog := useCaseTest.GetUserInvoiceLog(c.Param("id"))
	assert.NotNil(suite.T(), dummyUserInvoiceLog[0])
	assert.Equal(suite.T(), dummyUserInvoiceLog[0], dummyUserInvoiceLog[0])
}

func TestUserInvoiceLogUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserInvoiceLogUsecaseTestSuite))
}