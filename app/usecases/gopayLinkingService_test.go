package usecases

import (
	"net/http/httptest"
	"testing"
	"time"

	entity "github.com/coroo/go-starter/app/entity"
	repositories "github.com/coroo/go-starter/app/repositories"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// dummy data
var dummyGopayLinking = []entity.GopayLinking{
	entity.GopayLinking{
		ID				        : 1,
		PhoneNumber				: "08123456789",
		AccountId		  		: "ojbqegnlfdam",
		PaymentOptionToken  	: "bnafjlnfasn",
		CreatedAt		: time.Now(),
		UpdatedAt		: time.Now(),
	}, entity.GopayLinking{
		ID				        : 2,
		PhoneNumber				: "08987654321",
		AccountId		  		: "pajlakns",
		PaymentOptionToken  	: "paiisdanfoag",
		CreatedAt		: time.Now(),
		UpdatedAt		: time.Now(),
	},
}

type repoMockGopayLinking struct {
	mock.Mock
}

func (r *repoMockGopayLinking) SaveGopayLinking(gopayLinking entity.GopayLinking) (int, error) {
	return 0, nil
}

func (r *repoMockGopayLinking) UpdateGopayLinking(gopayLinking entity.GopayLinking) error {
	args := r.Called(gopayLinking)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (r *repoMockGopayLinking) DeleteGopayLinking(gopayLinking entity.GopayLinking) error {
	args := r.Called(gopayLinking)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (r *repoMockGopayLinking) GetAllGopayLinkings() []entity.GopayLinking {
	return dummyGopayLinking
}

func (r *repoMockGopayLinking) GetGopayLinking(id string) []entity.GopayLinking {
	return dummyGopayLinking
}

func (r *repoMockGopayLinking) GetGopayLinkingByUuid(uuid string) entity.GopayLinking {
	return dummyGopayLinking[0]
}

func (r *repoMockGopayLinking) GetActiveGopayLinkingByCode(code string) entity.GopayLinking {
	return dummyGopayLinking[0]
}

func (r *repoMockGopayLinking) GetGopayLinkingByCode(code string) entity.GopayLinking {
	return dummyGopayLinking[0]
}

func (r *repoMockGopayLinking) CloseDB() {
}

type GopayLinkingUsecaseTestSuite struct {
	suite.Suite
	repositoryTest repositories.GopayLinkingRepository
}

func (suite *GopayLinkingUsecaseTestSuite) SetupTest() {
	suite.repositoryTest = new(repoMockGopayLinking)
}

func (suite *GopayLinkingUsecaseTestSuite) TestBuildGopayLinkingService() {
	resultTest := NewGopayLinkingService(suite.repositoryTest)
	var dummyImpl *GopayLinkingService
	assert.NotNil(suite.T(), resultTest)
	assert.Implements(suite.T(), dummyImpl, resultTest)
	// assert.NotNil(suite.T(), resultTest.(*GopayLinkingService).repositories)
}

func (suite *GopayLinkingUsecaseTestSuite) TestSaveGopayLinkingUsecase() {
	suite.repositoryTest.(*repoMockGopayLinking).On("SaveGopayLinking", dummyGopayLinking[0]).Return(nil)
	useCaseTest := NewGopayLinkingService(suite.repositoryTest)
	// dummyGopayLinking[0].Password = "Change Password"
	data, _ := useCaseTest.SaveGopayLinking(dummyGopayLinking[0])
	assert.NotNil(suite.T(), data)
}

func (suite *GopayLinkingUsecaseTestSuite) TestUpdateGopayLinkingUsecase() {
	suite.repositoryTest.(*repoMockGopayLinking).On("UpdateGopayLinking", dummyGopayLinking[0]).Return(nil)
	useCaseTest := NewGopayLinkingService(suite.repositoryTest)
	err := useCaseTest.UpdateGopayLinking(dummyGopayLinking[0])
	assert.Nil(suite.T(), err)
}

func (suite *GopayLinkingUsecaseTestSuite) TestDeleteGopayLinkingUsecase() {
	suite.repositoryTest.(*repoMockGopayLinking).On("DeleteGopayLinking", dummyGopayLinking[0]).Return(nil)
	useCaseTest := NewGopayLinkingService(suite.repositoryTest)
	err := useCaseTest.DeleteGopayLinking(dummyGopayLinking[0])
	assert.Nil(suite.T(), err)
}

func (suite *GopayLinkingUsecaseTestSuite) TestGetAllGopayLinkings() {
	suite.repositoryTest.(*repoMockGopayLinking).On("GetAllGopayLinkings", dummyGopayLinking).Return(dummyGopayLinking)
	useCaseTest := NewGopayLinkingService(suite.repositoryTest)
	dummyGopayLinking := useCaseTest.GetAllGopayLinkings()
	assert.Equal(suite.T(), dummyGopayLinking, dummyGopayLinking)
}

func (suite *GopayLinkingUsecaseTestSuite) TestGetGopayLinking() {
	suite.repositoryTest.(*repoMockGopayLinking).On("GetGopayLinking", dummyGopayLinking[0].ID).Return(dummyGopayLinking[0], nil)
	useCaseTest := NewGopayLinkingService(suite.repositoryTest)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}
	dummyGopayLinking := useCaseTest.GetGopayLinking(c.Param("id"))
	assert.NotNil(suite.T(), dummyGopayLinking[0])
	assert.Equal(suite.T(), dummyGopayLinking[0], dummyGopayLinking[0])
}

func TestGopayLinkingUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(GopayLinkingUsecaseTestSuite))
}