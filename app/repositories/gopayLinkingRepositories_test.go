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

type GopayLinkingRepositoryTestSuite struct {
	suite.Suite
	ctx *gin.Context
	db  *gorm.DB
}

func (suite *GopayLinkingRepositoryTestSuite) SetupTest() {
	suite.db, _ = config.ConnectDB()
}

func (suite *GopayLinkingRepositoryTestSuite) TestA_BuildNewGopayLinkingRepository() {
	repoTest := NewGopayLinkingRepository()
	var dummyImpl *GopayLinkingRepository
	assert.NotNil(suite.T(), repoTest)
	assert.Implements(suite.T(), dummyImpl, repoTest)
}

func (suite *GopayLinkingRepositoryTestSuite) TestB_CreateGopayLinking() {
	repoTest := NewGopayLinkingRepository()
	dummyGopayLinking := entity.GopayLinking{
		ID				        : 1,
		PhoneNumber				: "08123456789",
		AccountId		  		: "ojbqegnlfdam",
		PaymentOptionToken  	: "bnafjlnfasn",
	}
	_, err := repoTest.SaveGopayLinking(dummyGopayLinking)
	assert.Nil(suite.T(), err)
}

func (suite *GopayLinkingRepositoryTestSuite) TestC_UpdateGopayLinking() {
	repoTest := NewGopayLinkingRepository()
	dummyGopayLinking := entity.GopayLinking{
		ID				        : 1,
		PhoneNumber				: "0812345678910",
		AccountId		  		: "ojbqegnlfdam",
		PaymentOptionToken  	: "bnafjlnfasn",
	}
	gopayLinkingDummy := repoTest.UpdateGopayLinking(dummyGopayLinking)
	assert.Nil(suite.T(), gopayLinkingDummy)
}

func (suite *GopayLinkingRepositoryTestSuite) TestE_GetAllGopayLinkings() {
	repoTest := NewGopayLinkingRepository()
	gopayLinkingDummy := repoTest.GetAllGopayLinkings()
	assert.NotNil(suite.T(), gopayLinkingDummy)
}

func (suite *GopayLinkingRepositoryTestSuite) TestF_GetGopayLinking() {
	repoTest := NewGopayLinkingRepository()
	gopayLinkingDummy := repoTest.GetGopayLinking("1")
	assert.NotNil(suite.T(), gopayLinkingDummy)
}

func (suite *GopayLinkingRepositoryTestSuite) TestH_RemoveGopayLinking() {
	repoTest := NewGopayLinkingRepository()
	dummyGopayLinking := entity.GopayLinking{
		ID: 1,
	}
	gopayLinkingDummy := repoTest.DeleteGopayLinking(dummyGopayLinking)
	assert.Nil(suite.T(), gopayLinkingDummy)
}


func TestGopayLinkingRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(GopayLinkingRepositoryTestSuite))
}