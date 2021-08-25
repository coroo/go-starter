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

type syEtlPaymentRepositoryTestSuite struct {
	suite.Suite
	ctx *gin.Context
	db  *gorm.DB
}

func (suite *syEtlPaymentRepositoryTestSuite) SetupSyEtlPaymentRepositoryTest() {
	suite.db, _ = config.ConnectDB()
}

func (suite *syEtlPaymentRepositoryTestSuite) TestBuildNewSyEtlPaymentRepository() {
	repoTest := NewSyEtlPaymentRepository()
	var dummyImpl *SyOdsEtlPaymentRepository
	assert.NotNil(suite.T(), repoTest)
	assert.Implements(suite.T(), dummyImpl, repoTest)
}

func (suite *syEtlPaymentRepositoryTestSuite) TestCreateSyEtlPayment() {
	repoTest := NewSyEtlPaymentRepository()
	dummyEtlPayment := entity.SyEtlPayment{
		ID					: 1,
		ProposalNumber		: "30012341234123",
		PolicyNumber		: "30012341234122",
		PaidDate			: time.Now(),
		PaymentMethodName	: "Indomaret",
		PolicyStatus		: "closed",
		TotalPremium		: 10000,
		UpdatedAt			: time.Now(),
	}
	repoTest.CreateSyEtlPayment(dummyEtlPayment)
}

func (suite *syEtlPaymentRepositoryTestSuite) TestUpdateSyEtlPayment() {
	repoTest := NewSyEtlPaymentRepository()
	dummyEtlPayment := entity.SyEtlPayment{
		ID					: 1,
		ProposalNumber		: "30012341234124",
		PolicyNumber		: "30012341234121",
		PaidDate			: time.Now(),
		PaymentMethodName	: "Indomaret",
		PolicyStatus		: "closed",
		TotalPremium		: 11000,
		UpdatedAt			: time.Now(),
	}
	repoTest.UpdateSyEtlPayment(dummyEtlPayment)
}

func (suite *syEtlPaymentRepositoryTestSuite) GetAllLatestGroupSyEtlPayments() {
	repoTest := NewSyEtlPaymentRepository()
	userPolicyDummy := repoTest.GetAllLatestGroupSyEtlPayments()
	assert.NotNil(suite.T(), userPolicyDummy)
}

func (suite *syEtlPaymentRepositoryTestSuite) GetAllSyEtlPayments() {
	repoTest := NewSyEtlPaymentRepository()
	userPolicyDummy := repoTest.GetAllSyEtlPayments()
	assert.NotNil(suite.T(), userPolicyDummy)
}

func (suite *syEtlPaymentRepositoryTestSuite) GetSyEtlPayment() {
	repoTest := NewSyEtlPaymentRepository()
	userPolicyDummy := repoTest.GetSyEtlPayment("30012341234121")
	assert.NotNil(suite.T(), userPolicyDummy)
}

func (suite *syEtlPaymentRepositoryTestSuite) TruncateTableSyEtlPayments() {
	repoTest := NewSyEtlPaymentRepository()
	repoTest.TruncateTableSyEtlPayments()
}

func (suite *syEtlPaymentRepositoryTestSuite) DeleteSyEtlPayment() {
	repoTest := NewSyEtlPaymentRepository()
	dummyEtlPayment := entity.SyEtlPayment{
		ID: 1,
	}
	repoTest.DeleteSyEtlPayment(dummyEtlPayment)
}

func TestSyEtlPaymentRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(syEtlPaymentRepositoryTestSuite))
}
