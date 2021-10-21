package repositories

import (
	"testing"

	entity "github.com/coroo/go-starter/app/entity"
	"github.com/coroo/go-starter/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PaymentMethodLinkRepositoryTestSuite struct {
	suite.Suite
	ctx *gin.Context
	db  *gorm.DB
}

func (suite *PaymentMethodLinkRepositoryTestSuite) SetupTest() {
	suite.db, _ = config.ConnectDB()
}

func (suite *PaymentMethodLinkRepositoryTestSuite) TestA_BuildNewPaymentMethodLinkRepository() {
	repoTest := NewPaymentMethodLinkRepository()
	var dummyImpl *PaymentMethodLinkRepository
	assert.NotNil(suite.T(), repoTest)
	assert.Implements(suite.T(), dummyImpl, repoTest)
}

func (suite *PaymentMethodLinkRepositoryTestSuite) TestB_CreatePaymentMethodLink() {
	repoTest := NewPaymentMethodLinkRepository()
	dummyPaymentMethodLink := entity.PaymentMethodLink{
		ID				        : 1,
		PaymentMethodCode 		: "gopay",
		ProcessType	 			: "linking",
		Url 					: "https://www.superyou.com/gopay-linking/:encryptedaccountid",
	}
	_, err := repoTest.SavePaymentMethodLink(dummyPaymentMethodLink)
	assert.Nil(suite.T(), err)
}

func (suite *PaymentMethodLinkRepositoryTestSuite) TestC_UpdatePaymentMethodLink() {
	repoTest := NewPaymentMethodLinkRepository()
	dummyPaymentMethodLink := entity.PaymentMethodLink{
		ID				: 1,
		PaymentMethodCode 		: "gopay",
		ProcessType	 			: "linking",
		Url 					: "https://www.superyou.com/gopay-linking/:encryptedaccountid",
	}
	paymentMethodLinkDummy := repoTest.UpdatePaymentMethodLink(dummyPaymentMethodLink)
	assert.Nil(suite.T(), paymentMethodLinkDummy)
}

func (suite *PaymentMethodLinkRepositoryTestSuite) TestE_GetAllPaymentMethodLinks() {
	repoTest := NewPaymentMethodLinkRepository()
	paymentMethodLinkDummy := repoTest.GetAllPaymentMethodLinks()
	assert.NotNil(suite.T(), paymentMethodLinkDummy)
}

func (suite *PaymentMethodLinkRepositoryTestSuite) TestF_GetPaymentMethodLink() {
	repoTest := NewPaymentMethodLinkRepository()
	paymentMethodLinkDummy := repoTest.GetPaymentMethodLink("1")
	assert.NotNil(suite.T(), paymentMethodLinkDummy)
}

func (suite *PaymentMethodLinkRepositoryTestSuite) TestH_RemovePaymentMethodLink() {
	repoTest := NewPaymentMethodLinkRepository()
	dummyPaymentMethodLink := entity.PaymentMethodLink{
		ID: 1,
	}
	paymentMethodLinkDummy := repoTest.DeletePaymentMethodLink(dummyPaymentMethodLink)
	assert.Nil(suite.T(), paymentMethodLinkDummy)
}


func TestPaymentMethodLinkRepositoryTestSuite(t *testing.T) {
	
	suite.Run(t, new(PaymentMethodLinkRepositoryTestSuite))
}