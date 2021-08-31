package repositories

import (
	"time"
	"testing"

	entity "github.com/coroo/go-starter/app/entity"
	"github.com/coroo/go-starter/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type syOdsEtlPaymentRepositoryTestSuite struct {
	suite.Suite
	ctx *gin.Context
	db  *gorm.DB
}

func (suite *syOdsEtlPaymentRepositoryTestSuite) SetupSyOdsEtlPaymentRepositoryTest() {
	suite.db, _ = config.ConnectDB()
}

func (suite *syOdsEtlPaymentRepositoryTestSuite) TestBuildNewSyOdsEtlPaymentRepository() {
	repoTest := NewSyOdsEtlPaymentRepository()
	var dummyImpl *SyOdsEtlPaymentRepository
	assert.NotNil(suite.T(), repoTest)
	assert.Implements(suite.T(), dummyImpl, repoTest)
}

func (suite *syOdsEtlPaymentRepositoryTestSuite) CreateSyOdsEtlPayment() {
	repoTest := NewSyOdsEtlPaymentRepository()
	dummySyOdsEtlPayment := entity.SyOdsEtlPayment{
		ID					: 1,
		ProposalNumber		: "30012341234123",
		PolicyNumber		: "30012341234122",
		OdsFirstPaidDate	: time.Now(),
		OdsPaidDate			: time.Now(),
		SyPaidDate			: time.Now(),
		PaymentMethodName	: "Indomaret",
		CollectionId		: "1234567",
		SyTotalAmount		: 10000,
		OdsTotalAmount		: 10000,
		PolicyStatus		: "success",
		Status				: "closed",
		StatusDescription	: "Stated as First Payment, use SY data for First Payment",
		UpdatedAt			: time.Now(),
	}
	repoTest.CreateSyOdsEtlPayment(dummySyOdsEtlPayment)
}

func (suite *syOdsEtlPaymentRepositoryTestSuite) UpdateSyOdsEtlPayment() {
	repoTest := NewSyOdsEtlPaymentRepository()
	dummySyOdsEtlPayment := entity.SyOdsEtlPayment{
		ID					: 1,
		ProposalNumber		: "30012341234124",
		PolicyNumber		: "30012341234121",
		OdsFirstPaidDate	: time.Now(),
		OdsPaidDate			: time.Now(),
		SyPaidDate			: time.Now(),
		PaymentMethodName	: "Indomaret",
		CollectionId		: "1234567",
		SyTotalAmount		: 10000,
		OdsTotalAmount		: 10000,
		PolicyStatus		: "success",
		Status				: "closed",
		StatusDescription	: "Stated as First Payment, use SY data for First Payment",
		UpdatedAt			: time.Now(),
	}
	repoTest.UpdateSyOdsEtlPayment(dummySyOdsEtlPayment)
}

func (suite *syOdsEtlPaymentRepositoryTestSuite) GetAllLatestGroupSyOdsEtlPayments() {
	repoTest := NewSyOdsEtlPaymentRepository()
	userPolicyDummy := repoTest.GetAllLatestGroupSyOdsEtlPayments()
	assert.NotNil(suite.T(), userPolicyDummy)
}

func (suite *syOdsEtlPaymentRepositoryTestSuite) GetAllSyOdsEtlPayments() {
	repoTest := NewSyOdsEtlPaymentRepository()
	userPolicyDummy := repoTest.GetAllSyOdsEtlPayments()
	assert.NotNil(suite.T(), userPolicyDummy)
}

func (suite *syOdsEtlPaymentRepositoryTestSuite) GetSyOdsEtlPaymentByPolicyNumber() {
	repoTest := NewSyOdsEtlPaymentRepository()
	userPolicyDummy := repoTest.GetSyOdsEtlPaymentByPolicyNumber("30012341234122")
	assert.NotNil(suite.T(), userPolicyDummy)
}

func (suite *syOdsEtlPaymentRepositoryTestSuite) GetSyOdsEtlPaymentByStatus() {
	repoTest := NewSyOdsEtlPaymentRepository()
	userPolicyDummy := repoTest.GetSyOdsEtlPaymentByStatus("closed")
	assert.NotNil(suite.T(), userPolicyDummy)
}

func (suite *syOdsEtlPaymentRepositoryTestSuite) GetSyOdsEtlPaymentDailyByStatus() {
	repoTest := NewSyOdsEtlPaymentRepository()
	userPolicyDummy := repoTest.GetSyOdsEtlPaymentDailyByStatus("closed")
	assert.NotNil(suite.T(), userPolicyDummy)
}

func (suite *syOdsEtlPaymentRepositoryTestSuite) CancelOutstandingSyOdsEtlPayments() {
	repoTest := NewSyOdsEtlPaymentRepository()
	userPolicyDummy := repoTest.CancelOutstandingSyOdsEtlPayments()
	assert.NotNil(suite.T(), userPolicyDummy)
}

func (suite *syOdsEtlPaymentRepositoryTestSuite) RemoveSyOdsEtlPayment() {
	repoTest := NewSyOdsEtlPaymentRepository()
	dummySyOdsEtlPayment := entity.SyOdsEtlPayment{
		ID: 1,
	}
	repoTest.DeleteSyOdsEtlPayment(dummySyOdsEtlPayment)
}

func TestSyOdsEtlPaymentRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(syOdsEtlPaymentRepositoryTestSuite))
}
