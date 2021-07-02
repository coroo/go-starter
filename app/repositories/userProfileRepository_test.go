package repositories

import (
	"net/http/httptest"
	"testing"

	entity "github.com/coroo/go-lemonilo/app/entity"
	"github.com/coroo/go-lemonilo/config"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserProfileRepositoryTestSuite struct {
	suite.Suite
	ctx *gin.Context
	db  *gorm.DB
}

func (suite *UserProfileRepositoryTestSuite) SetupTest() {
	suite.db, _ = config.ConnectDB()
}

func (suite *UserProfileRepositoryTestSuite) TestBuildNewUserProfileRepository() {
	repoTest := NewUserProfileRepository()
	var dummyImpl *UserProfileRepository
	assert.NotNil(suite.T(), repoTest)
	assert.Implements(suite.T(), dummyImpl, repoTest)
}

func (suite *UserProfileRepositoryTestSuite) TestUserProfileCreate() {
	repoTest := NewUserProfileRepository()
	dummyUserProfile := entity.UserProfile{
		Email			: "kuncoro@test.com",
		Password		: "password",
		Address			: "jl lorem ipsum",
	}
	_, err := repoTest.SaveUserProfile(dummyUserProfile)
	assert.Nil(suite.T(), err)
}

func (suite *UserProfileRepositoryTestSuite) TestUserProfileUpdate() {
	repoTest := NewUserProfileRepository()
	dummyUserProfile := entity.UserProfile{
		ID				: 1,
		Email			: "kuncoro@test.com",
		Password		: "password",
		Address			: "jl lorem ipsum",
	}
	userProfileDummy := repoTest.UpdateUserProfile(dummyUserProfile)
	assert.Nil(suite.T(), userProfileDummy)
}

func (suite *UserProfileRepositoryTestSuite) TestUserProfileDelete() {
	repoTest := NewUserProfileRepository()
	dummyUserProfile := entity.UserProfile{
		ID: 1,
	}
	userProfileDummy := repoTest.DeleteUserProfile(dummyUserProfile)
	assert.Nil(suite.T(), userProfileDummy)
}

func (suite *UserProfileRepositoryTestSuite) TestGetAllUserProfiles() {
	repoTest := NewUserProfileRepository()
	userProfileDummy := repoTest.GetAllUserProfiles()
	assert.NotNil(suite.T(), userProfileDummy)
}

func (suite *UserProfileRepositoryTestSuite) TestGetUserProfile() {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}

	repoTest := NewUserProfileRepository()
	userProfileDummy := repoTest.GetUserProfile(c)
	assert.NotNil(suite.T(), userProfileDummy)
}

func TestUserProfileRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserProfileRepositoryTestSuite))
}
