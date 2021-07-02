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

func (r *repoMockUserProfile) SaveUserProfile(userProfile entity.UserProfile) (int, error) {
	return 0, nil
}

func (r *repoMockUserProfile) UpdateUserProfile(userProfile entity.UserProfile) error {
	args := r.Called(userProfile)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (r *repoMockUserProfile) DeleteUserProfile(userProfile entity.UserProfile) error {
	args := r.Called(userProfile)
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

func (r *repoMockUserProfile) AuthUserProfile(userProfile entity.UserProfile) entity.UserProfile {
	return dummyUserProfile[0]
}

func (r *repoMockUserProfile) CloseDB() {
}

type UserProfileUsecaseTestSuite struct {
	suite.Suite
	repositoryTest repositories.UserProfileRepository
}

func (suite *UserProfileUsecaseTestSuite) SetupTest() {
	suite.repositoryTest = new(repoMockUserProfile)
}

func (suite *UserProfileUsecaseTestSuite) TestBuildUserProfileService() {
	resultTest := NewUserProfile(suite.repositoryTest)
	var dummyImpl *UserProfileService
	assert.NotNil(suite.T(), resultTest)
	assert.Implements(suite.T(), dummyImpl, resultTest)
	// assert.NotNil(suite.T(), resultTest.(*UserProfileService).repositories)
}

func (suite *UserProfileUsecaseTestSuite) TestSaveUserProfileUsecase() {
	suite.repositoryTest.(*repoMockUserProfile).On("SaveUserProfile", dummyUserProfile[0]).Return(nil)
	useCaseTest := NewUserProfile(suite.repositoryTest)
	// dummyUserProfile[0].Password = "Change Password"
	data, _ := useCaseTest.SaveUserProfile(dummyUserProfile[0])
	assert.NotNil(suite.T(), data)
}

func (suite *UserProfileUsecaseTestSuite) TestUpdateUserProfileUsecase() {
	suite.repositoryTest.(*repoMockUserProfile).On("UpdateUserProfile", dummyUserProfile[0]).Return(nil)
	useCaseTest := NewUserProfile(suite.repositoryTest)
	err := useCaseTest.UpdateUserProfile(dummyUserProfile[0])
	assert.Nil(suite.T(), err)
}

func (suite *UserProfileUsecaseTestSuite) TestDeleteUserProfileUsecase() {
	suite.repositoryTest.(*repoMockUserProfile).On("DeleteUserProfile", dummyUserProfile[0]).Return(nil)
	useCaseTest := NewUserProfile(suite.repositoryTest)
	err := useCaseTest.DeleteUserProfile(dummyUserProfile[0])
	assert.Nil(suite.T(), err)
}

func (suite *UserProfileUsecaseTestSuite) TestGetAllUserProfiles() {
	suite.repositoryTest.(*repoMockUserProfile).On("GetAllUserProfiles", dummyUserProfile).Return(dummyUserProfile)
	useCaseTest := NewUserProfile(suite.repositoryTest)
	dummyUserProfile := useCaseTest.GetAllUserProfiles()
	assert.Equal(suite.T(), dummyUserProfile, dummyUserProfile)
}

func (suite *UserProfileUsecaseTestSuite) TestGetUserProfile() {
	suite.repositoryTest.(*repoMockUserProfile).On("GetUserProfile", dummyUserProfile[0].ID).Return(dummyUserProfile[0], nil)
	useCaseTest := NewUserProfile(suite.repositoryTest)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}
	dummyUserProfile := useCaseTest.GetUserProfile(c)
	assert.NotNil(suite.T(), dummyUserProfile[0])
	assert.Equal(suite.T(), dummyUserProfile[0], dummyUserProfile[0])
}

func (suite *UserProfileUsecaseTestSuite) TestAuthUserProfileUsecase() {
	suite.repositoryTest.(*repoMockUserProfile).On("AuthUserProfile", dummyUserProfile[0]).Return(nil)
	useCaseTest := NewUserProfile(suite.repositoryTest)
	err := useCaseTest.AuthUserProfile(dummyUserProfile[0])
	assert.NotNil(suite.T(), err)
}

func TestUserProfileUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserProfileUsecaseTestSuite))
}
