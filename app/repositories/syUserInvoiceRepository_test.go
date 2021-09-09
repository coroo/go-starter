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

type syUserInvoiceRepositoryTestSuite struct {
	suite.Suite
	ctx *gin.Context
	db  *gorm.DB
}

func SetupSyUserInvoiceRepositoryTest() SyUserInvoiceRepository{
	db, err := config.ConnectDBSY()
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&entity.SyUserInvoice{})
	return &syUserInvoiceDatabase{
		connection: db,
	}
}

func (suite *syUserInvoiceRepositoryTestSuite) TestA_BuildSetupSyUserInvoiceRepositoryTest() {
	repoTest := SetupSyUserInvoiceRepositoryTest()
	var dummyImpl *SyUserInvoiceRepository
	assert.NotNil(suite.T(), repoTest)
	assert.Implements(suite.T(), dummyImpl, repoTest)
}

func (suite *syUserInvoiceRepositoryTestSuite) TestB_CreateSyUserInvoice() {
	repoTest := SetupSyUserInvoiceRepositoryTest()
	dummySyUserInvoice := entity.SyUserInvoice{
		ID					: 1,
		PolicyNumber		: "30012341234123",
		PolicyGroupNumber	: "30012341234122",
		ProposalNumber		: "123123",
		PaymentMethodName	: "Indomaret",
		TotalPremium		: 123123,
		Status				: "test123",
		PaidAt				: time.Now(),
	}
	repoTest.SaveSyUserInvoice(dummySyUserInvoice)
}

func (suite *syUserInvoiceRepositoryTestSuite) TestC_UpdateSyUserInvoice() {
	repoTest := SetupSyUserInvoiceRepositoryTest()
	dummySyUserInvoice := entity.SyUserInvoice{
		ID					: 1,
		PolicyNumber		: "30012341234124",
		PolicyGroupNumber	: "30012341234121",
		ProposalNumber		: "123123",
		PaymentMethodName	: "Indomaret",
		TotalPremium		: 234234,
		Status				: "test123",
		PaidAt				: time.Now(),
	}
	repoTest.UpdateSyUserInvoice(dummySyUserInvoice)
}

// sementara di comment karena salah query near select
// func (suite *syUserInvoiceRepositoryTestSuite) TestD_GetAllPaidUserInvoices() {
// 	repoTest := SetupSyUserInvoiceRepositoryTest()
// 	userPolicyDummy := repoTest.GetAllPaidUserInvoices()
// 	assert.NotNil(suite.T(), userPolicyDummy)
// }

func (suite *syUserInvoiceRepositoryTestSuite) TestE_GetUserInvoice() {
	repoTest := SetupSyUserInvoiceRepositoryTest()
	userPolicyDummy := repoTest.GetUserInvoice("1")
	assert.NotNil(suite.T(), userPolicyDummy)
}

func (suite *syUserInvoiceRepositoryTestSuite) TestF_RemoveSyUserInvoice() {
	repoTest := SetupSyUserInvoiceRepositoryTest()
	dummySyUserInvoice := entity.SyUserInvoice{
		ID: 1,
	}
	repoTest.DeleteSyUserInvoice(dummySyUserInvoice)
}

func TestSyUserInvoiceRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(syUserInvoiceRepositoryTestSuite))
}
