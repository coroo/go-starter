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

type PaymentMethodRepositoryTestSuite struct {
	suite.Suite
	ctx *gin.Context
	db  *gorm.DB
}

func (suite *PaymentMethodRepositoryTestSuite) SetupTest() {
	suite.db, _ = config.ConnectDB()
}

func (suite *PaymentMethodRepositoryTestSuite) TestA_BuildNewPaymentMethodRepository() {
	repoTest := NewPaymentMethodRepository()
	var dummyImpl *PaymentMethodRepository
	assert.NotNil(suite.T(), repoTest)
	assert.Implements(suite.T(), dummyImpl, repoTest)
}

func (suite *PaymentMethodRepositoryTestSuite) TestB_CreatePaymentMethod() {
	repoTest := NewPaymentMethodRepository()
	dummyPaymentMethod := entity.PaymentMethod{
		ID				        : 1,
		Code 					: "visa-master",
		InitPaymentCode  		: "PDCC",
		RenewalPaymentCode 		: "CASH",
		FastpayCode 		 	: "cc",
		Name 		 			: "Visa Master",
		PaymentLogo 		 	: "visa-master.png",
		Status 		 			: "active",
	}
	_, err := repoTest.SavePaymentMethod(dummyPaymentMethod)
	assert.Nil(suite.T(), err)
}

func (suite *PaymentMethodRepositoryTestSuite) TestC_UpdatePaymentMethod() {
	repoTest := NewPaymentMethodRepository()
	dummyPaymentMethod := entity.PaymentMethod{
		ID				: 1,
		Code 					: "visa-master-test",
		InitPaymentCode  		: "PDCC",
		RenewalPaymentCode 		: "CASH",
		FastpayCode 		 	: "cc",
		Name 		 			: "Visa Master",
		PaymentLogo 		 	: "visa-master.png",
		Status 		 			: "active",
	}
	paymentMethodDummy := repoTest.UpdatePaymentMethod(dummyPaymentMethod)
	assert.Nil(suite.T(), paymentMethodDummy)
}

func (suite *PaymentMethodRepositoryTestSuite) TestE_GetAllPaymentMethods() {
	repoTest := NewPaymentMethodRepository()
	paymentMethodDummy := repoTest.GetAllPaymentMethods("")
	assert.NotNil(suite.T(), paymentMethodDummy)
}

func (suite *PaymentMethodRepositoryTestSuite) TestF_GetPaymentMethod() {
	repoTest := NewPaymentMethodRepository()
	paymentMethodDummy := repoTest.GetPaymentMethod("1")
	assert.NotNil(suite.T(), paymentMethodDummy)
}

func (suite *PaymentMethodRepositoryTestSuite) TestH_RemovePaymentMethod() {
	repoTest := NewPaymentMethodRepository()
	dummyPaymentMethod := entity.PaymentMethod{
		ID: 1,
	}
	paymentMethodDummy := repoTest.DeletePaymentMethod(dummyPaymentMethod)
	assert.Nil(suite.T(), paymentMethodDummy)
}


func TestPaymentMethodRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(PaymentMethodRepositoryTestSuite))
}