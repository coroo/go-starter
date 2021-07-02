package usecases

import (
	"net/http/httptest"
	"testing"
	"time"

	entity "github.com/coroo/go-lemonilo/app/entity"
	repositories "github.com/coroo/go-lemonilo/app/repositories"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// dummy data
var dummyUserProfile = []entity.UserProfile{
	entity.UserProfile{
		ID				: 1,
		Email			: "kuncoro@test.com",
		Password		: "password",
		Address			: "jl lorem ipsum",
		CreatedAt		: time.Now(),
		UpdatedAt		: time.Now(),
	}, entity.UserProfile{
		ID				: 2,
		Email			: "kuncoro@test.com",
		Password		: "password",
		Address			: "jl lorem ipsum",
		CreatedAt		: time.Now(),
		UpdatedAt		: time.Now(),
	},
}

type repoMockUserProfile struct {
	mock.Mock
}

func (r *repoMockUserProfile) SaveUserProfile(masterQuestion entity.UserProfile) (int, error) {
	args := r.Called(masterQuestion)
	if args.Get(0) == nil {
		return 0, nil
	}
	return 0, args.Get(0).(error)
}

func (r *repoMockUserProfile) UpdateUserProfile(masterQuestion entity.UserProfile) error {
	args := r.Called(masterQuestion)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (r *repoMockUserProfile) DeleteUserProfile(masterQuestion entity.UserProfile) error {
	args := r.Called(masterQuestion)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (r *repoMockUserProfile) GetAllUserProfiles() []entity.UserProfile {
	return dummyUserProfile
}

func (r *repoMockUserProfile) GetUserProfile(ctx *gin.Context) []entity.UserProfile {
	return dummyUserProfile
}

func (r *repoMockUserProfile) CloseDB() {
}

type UserProfileDeliveryTestSuite struct {
	suite.Suite
	repositoryTest repositories.UserProfileRepository
}

func (suite *UserProfileDeliveryTestSuite) SetupTest() {
	suite.repositoryTest = new(repoMockUserProfile)
}

func (suite *UserProfileDeliveryTestSuite) TestBuildUserProfileService() {
	resultTest := NewUserProfile(suite.repositoryTest)
	var dummyImpl *UserProfileService
	assert.NotNil(suite.T(), resultTest)
	assert.Implements(suite.T(), dummyImpl, resultTest)
	// assert.NotNil(suite.T(), resultTest.(*UserProfileService).repositories)
}

func (suite *UserProfileDeliveryTestSuite) TestSaveUserProfileDelivery() {
	suite.repositoryTest.(*repoMockUserProfile).On("SaveUserProfile", dummyUserProfile[0]).Return(nil)
	useCaseTest := NewUserProfile(suite.repositoryTest)
	_, err := useCaseTest.SaveUserProfile(dummyUserProfile[0])
	assert.Nil(suite.T(), err)
}

func (suite *UserProfileDeliveryTestSuite) TestUpdateUserProfileDelivery() {
	suite.repositoryTest.(*repoMockUserProfile).On("UpdateUserProfile", dummyUserProfile[0]).Return(nil)
	useCaseTest := NewUserProfile(suite.repositoryTest)
	err := useCaseTest.UpdateUserProfile(dummyUserProfile[0])
	assert.Nil(suite.T(), err)
}

func (suite *UserProfileDeliveryTestSuite) TestDeleteUserProfileDelivery() {
	suite.repositoryTest.(*repoMockUserProfile).On("DeleteUserProfile", dummyUserProfile[0]).Return(nil)
	useCaseTest := NewUserProfile(suite.repositoryTest)
	err := useCaseTest.DeleteUserProfile(dummyUserProfile[0])
	assert.Nil(suite.T(), err)
}

func (suite *UserProfileDeliveryTestSuite) TestGetAllUserProfiles() {
	suite.repositoryTest.(*repoMockUserProfile).On("GetAllUserProfiles", dummyUserProfile).Return(dummyUserProfile)
	useCaseTest := NewUserProfile(suite.repositoryTest)
	dummyUserProfile := useCaseTest.GetAllUserProfiles()
	assert.Equal(suite.T(), dummyUserProfile, dummyUserProfile)
}

func (suite *UserProfileDeliveryTestSuite) TestGetUserProfile() {
	suite.repositoryTest.(*repoMockUserProfile).On("GetUserProfile", dummyUserProfile[0].ID).Return(dummyUserProfile[0], nil)
	useCaseTest := NewUserProfile(suite.repositoryTest)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}
	dummyUserProfile := useCaseTest.GetUserProfile(c)
	assert.NotNil(suite.T(), dummyUserProfile[0])
	assert.Equal(suite.T(), dummyUserProfile[0], dummyUserProfile[0])
}

func TestUserProfileDeliveryTestSuite(t *testing.T) {
	suite.Run(t, new(UserProfileDeliveryTestSuite))
}
