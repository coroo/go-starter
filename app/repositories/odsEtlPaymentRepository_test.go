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

type odsEtlPaymentRepositoryTestSuite struct {
	suite.Suite
	ctx *gin.Context
	db  *gorm.DB
}

func (suite *odsEtlPaymentRepositoryTestSuite) SetupOdsEtlPaymentRepositoryTest() {
	suite.db, _ = config.ConnectDB()
}

func (suite *odsEtlPaymentRepositoryTestSuite) TestBuildNewOdsEtlPaymentRepository() {
	repoTest := NewOdsEtlPaymentRepository()
	var dummyImpl *OdsEtlPaymentRepository
	assert.NotNil(suite.T(), repoTest)
	assert.Implements(suite.T(), dummyImpl, repoTest)
}

func (suite *odsEtlPaymentRepositoryTestSuite) CreateOdsEtlPayment() {
	repoTest := NewOdsEtlPaymentRepository()
	dummySyOdsEtlPayment := entity.OdsEtlPayment{
		ID					: 1,
		CollectionId		: "1234567",
		ProposalNumber		: "30012341234122",
		PolicyNumber		: "30012341234123",
		FirstPaymentDate	: time.Now(),
		PaymentDate			: time.Now(),
		PaymentMethod		: "Indomaret",
		TotalAmount			: 10000,
		UpdatedAt			: time.Now(),
	}
	repoTest.CreateOdsEtlPayment(dummySyOdsEtlPayment)
}

func (suite *odsEtlPaymentRepositoryTestSuite) UpdateOdsEtlPayment() {
	repoTest := NewOdsEtlPaymentRepository()
	dummyOdsEtlPayment := entity.OdsEtlPayment{
		ID					: 1,
		CollectionId		: "1234568",
		ProposalNumber		: "30012341234124",
		PolicyNumber		: "30012341234121",
		FirstPaymentDate	: time.Now(),
		PaymentDate			: time.Now(),
		PaymentMethod		: "Indomaret",
		TotalAmount			: 10000,
		UpdatedAt			: time.Now(),
	}
	repoTest.UpdateOdsEtlPayment(dummyOdsEtlPayment)
}

func (suite *odsEtlPaymentRepositoryTestSuite) GetAllLatestGroupOdsEtlPayments() {
	repoTest := NewOdsEtlPaymentRepository()
	odsEtlPaymentDummy := repoTest.GetAllLatestGroupOdsEtlPayments()
	assert.NotNil(suite.T(), odsEtlPaymentDummy)
}

func (suite *odsEtlPaymentRepositoryTestSuite) GetAllOdsEtlPayments() {
	repoTest := NewOdsEtlPaymentRepository()
	userPolicyDummy := repoTest.GetAllOdsEtlPayments()
	assert.NotNil(suite.T(), userPolicyDummy)
}

func (suite *odsEtlPaymentRepositoryTestSuite) GetOdsEtlPayment() {
	repoTest := NewOdsEtlPaymentRepository()
	userPolicyDummy := repoTest.GetOdsEtlPayment("1")
	assert.NotNil(suite.T(), userPolicyDummy)
}

func (suite *odsEtlPaymentRepositoryTestSuite) TruncateTableOdsEtlPayments() {
	repoTest := NewOdsEtlPaymentRepository()
	repoTest.TruncateTableOdsEtlPayments()
}

func (suite *odsEtlPaymentRepositoryTestSuite) TestDeleteOdsEtlPayment() {
	repoTest := NewOdsEtlPaymentRepository()
	dummyOdsEtlPayment := entity.OdsEtlPayment{
		ID: 1,
	}
	repoTest.DeleteOdsEtlPayment(dummyOdsEtlPayment)
}

func TestOdsEtlPaymentRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(odsEtlPaymentRepositoryTestSuite))
}
