package repositories

import (
	"testing"
	"time"

	entity "github.com/coroo/go-starter/app/entity"
	"github.com/coroo/go-starter/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserInvoiceLogRepositoryTestSuite struct {
	suite.Suite
	ctx *gin.Context
	db  *gorm.DB
}

func (suite *UserInvoiceLogRepositoryTestSuite) SetupTest() {
	suite.db, _ = config.ConnectDB()
}

func (suite *UserInvoiceLogRepositoryTestSuite) TestA_BuildNewUserInvoiceLogRepository() {
	repoTest := NewUserInvoiceLogRepository()
	var dummyImpl *UserInvoiceLogRepository
	assert.NotNil(suite.T(), repoTest)
	assert.Implements(suite.T(), dummyImpl, repoTest)
}

func (suite *UserInvoiceLogRepositoryTestSuite) TestB_CreateUserInvoiceLog() {
	repoTest := NewUserInvoiceLogRepository()
	dummyUserInvoiceLog := entity.UserInvoiceLog{
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
	}
	_, err := repoTest.SaveUserInvoiceLog(dummyUserInvoiceLog)
	assert.Nil(suite.T(), err)
}

func (suite *UserInvoiceLogRepositoryTestSuite) TestC_UpdateUserInvoiceLog() {
	repoTest := NewUserInvoiceLogRepository()
	dummyUserInvoiceLog := entity.UserInvoiceLog{
		ID               		: 1,
		PaymentMethodCode 		: "indomaret-test",
		SummaryToken  			: "abcdabcd",
		// TransactionFee  		: 6500.00,
		InvoiceNumber			: "12369412964",
		PaymentCycle  			: "YEARLY",
		TransactionFee  		: "20000",
		AgentFee  				: "20000",
		TotalPremium  			: "40000",
		TotalPayment  			: "80000",
		CreatedAt				: time.Now(),
		UpdatedAt				: time.Now(),
	}
	userInvoiceLogDummy := repoTest.UpdateUserInvoiceLog(dummyUserInvoiceLog)
	assert.Nil(suite.T(), userInvoiceLogDummy)
}

func (suite *UserInvoiceLogRepositoryTestSuite) TestE_GetAllUserInvoiceLogs() {
	repoTest := NewUserInvoiceLogRepository()
	userInvoiceLogDummy := repoTest.GetAllUserInvoiceLogs()
	assert.NotNil(suite.T(), userInvoiceLogDummy)
}

func (suite *UserInvoiceLogRepositoryTestSuite) TestF_GetUserInvoiceLog() {
	repoTest := NewUserInvoiceLogRepository()
	userInvoiceLogDummy := repoTest.GetUserInvoiceLog("1")
	assert.NotNil(suite.T(), userInvoiceLogDummy)
}

func (suite *UserInvoiceLogRepositoryTestSuite) TestH_RemoveUserInvoiceLog() {
	repoTest := NewUserInvoiceLogRepository()
	dummyUserInvoiceLog := entity.UserInvoiceLog{
		ID: 1,
	}
	userInvoiceLogDummy := repoTest.DeleteUserInvoiceLog(dummyUserInvoiceLog)
	assert.Nil(suite.T(), userInvoiceLogDummy)
}


func TestUserInvoiceLogRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserInvoiceLogRepositoryTestSuite))
}