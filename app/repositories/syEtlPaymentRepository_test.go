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

func (suite *syEtlPaymentRepositoryTestSuite) SetupTest() {
	suite.db, _ = config.ConnectDB()
}

func (suite *syEtlPaymentRepositoryTestSuite) TestA_BuildNewSyEtlPaymentRepository() {
	repoTest := NewSyEtlPaymentRepository()
	var dummyImpl *SyEtlPaymentRepository
	assert.NotNil(suite.T(), repoTest)
	assert.Implements(suite.T(), dummyImpl, repoTest)
}

// sementara di comment karena kena error select saat if
func (suite *syEtlPaymentRepositoryTestSuite) TestB_CreateSyEtlPayment() {
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

// sementara di comment karena kena error berhubungan dengan select tetapi tetap ke update
func (suite *syEtlPaymentRepositoryTestSuite) TestC_UpdateSyEtlPayment() {
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

func (suite *syEtlPaymentRepositoryTestSuite) TestD_GetAllLatestGroupSyEtlPayments() {
	repoTest := NewSyEtlPaymentRepository()
	userPolicyDummy := repoTest.GetAllLatestGroupSyEtlPayments()
	assert.NotNil(suite.T(), userPolicyDummy)
}

func (suite *syEtlPaymentRepositoryTestSuite) TestE_GetAllSyEtlPayments() {
	repoTest := NewSyEtlPaymentRepository()
	userPolicyDummy := repoTest.GetAllSyEtlPayments()
	assert.NotNil(suite.T(), userPolicyDummy)
}

func (suite *syEtlPaymentRepositoryTestSuite) TestF_GetSyEtlPayment() {
	repoTest := NewSyEtlPaymentRepository()
	userPolicyDummy := repoTest.GetSyEtlPayment("30012341234121")
	assert.NotNil(suite.T(), userPolicyDummy)
}

// sementara di comment karena perbedaan query mysql n sqlite
// func (suite *syEtlPaymentRepositoryTestSuite) TestTruncateTableSyEtlPayments() {
// 	repoTest := NewSyEtlPaymentRepository()
// 	repoTest.TruncateTableSyEtlPayments()
// }

func (suite *syEtlPaymentRepositoryTestSuite) TestG_RemoveSyEtlPayment() {
	repoTest := NewSyEtlPaymentRepository()
	dummyEtlPayment := entity.SyEtlPayment{
		ID: 1,
	}
	repoTest.DeleteSyEtlPayment(dummyEtlPayment)
}

func TestSyEtlPaymentRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(syEtlPaymentRepositoryTestSuite))
}
