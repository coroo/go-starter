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

type PaymentMethodRateRepositoryTestSuite struct {
	suite.Suite
	ctx *gin.Context
	db  *gorm.DB
}

func (suite *PaymentMethodRateRepositoryTestSuite) SetupTest() {
	suite.db, _ = config.ConnectDB()
}

func (suite *PaymentMethodRateRepositoryTestSuite) TestA_BuildNewPaymentMethodRateRepository() {
	repoTest := NewPaymentMethodRateRepository()
	var dummyImpl *PaymentMethodRateRepository
	assert.NotNil(suite.T(), repoTest)
	assert.Implements(suite.T(), dummyImpl, repoTest)
}

func (suite *PaymentMethodRateRepositoryTestSuite) TestB_CreatePaymentMethodRate() {
	repoTest := NewPaymentMethodRateRepository()
	dummyPaymentMethodRate := entity.PaymentMethodRate{
		ID				        : 1,
		PaymentMethodCode		: "indomaret",
		MinTransaction  		: 1,
		MaxTransaction  		: 500000,
		// TransactionFee  		: 3750.00,
		FormulaFee				: "3750.00",
		Premi  					: 0.00,
	}
	_, err := repoTest.SavePaymentMethodRate(dummyPaymentMethodRate)
	assert.Nil(suite.T(), err)
}

func (suite *PaymentMethodRateRepositoryTestSuite) TestC_UpdatePaymentMethodRate() {
	repoTest := NewPaymentMethodRateRepository()
	dummyPaymentMethodRate := entity.PaymentMethodRate{
		ID				: 1,
		PaymentMethodCode		: "indomaret",
		MinTransaction  		: 1,
		MaxTransaction  		: 1000000,
		// TransactionFee  		: 3750.00,
		FormulaFee				: "3750.00",
		Premi  					: 0.00,
	}
	paymentMethodRateDummy := repoTest.UpdatePaymentMethodRate(dummyPaymentMethodRate)
	assert.Nil(suite.T(), paymentMethodRateDummy)
}

func (suite *PaymentMethodRateRepositoryTestSuite) TestE_GetAllPaymentMethodRates() {
	repoTest := NewPaymentMethodRateRepository()
	paymentMethodRateDummy := repoTest.GetAllPaymentMethodRates()
	assert.NotNil(suite.T(), paymentMethodRateDummy)
}

func (suite *PaymentMethodRateRepositoryTestSuite) TestF_GetPaymentMethodRate() {
	repoTest := NewPaymentMethodRateRepository()
	paymentMethodRateDummy := repoTest.GetPaymentMethodRate("1")
	assert.NotNil(suite.T(), paymentMethodRateDummy)
}

func (suite *PaymentMethodRateRepositoryTestSuite) TestH_RemovePaymentMethodRate() {
	repoTest := NewPaymentMethodRateRepository()
	dummyPaymentMethodRate := entity.PaymentMethodRate{
		ID: 1,
	}
	paymentMethodRateDummy := repoTest.DeletePaymentMethodRate(dummyPaymentMethodRate)
	assert.Nil(suite.T(), paymentMethodRateDummy)
}


func TestPaymentMethodRateRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(PaymentMethodRateRepositoryTestSuite))
}