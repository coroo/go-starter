package deliveries

import (
	"net/http/httptest"
	"testing"
	"time"

	entity "github.com/coroo/go-lemonilo/app/entity"
	usecases "github.com/coroo/go-lemonilo/app/usecases"
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

type serviceMockUserProfile struct {
	mock.Mock
}

func (r *serviceMockUserProfile) SaveUserProfile(userProfile entity.UserProfile) (int, error) {
	args := r.Called(userProfile)
	if args.Get(0) == nil {
		return 0, nil
	}
	return 0, args.Get(0).(error)
}

func (r *serviceMockUserProfile) UpdateUserProfile(userProfile entity.UserProfile) error {
	args := r.Called(userProfile)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (r *serviceMockUserProfile) DeleteUserProfile(userProfile entity.UserProfile) error {
	args := r.Called(userProfile)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (r *serviceMockUserProfile) GetAllUserProfiles() []entity.UserProfile {
	return dummyUserProfile
}

func (r *serviceMockUserProfile) GetUserProfile(ctx *gin.Context) []entity.UserProfile {
	return dummyUserProfile
}

type UserProfileDeliveryTestSuite struct {
	suite.Suite
	serviceTest usecases.UserProfileService
}

func (suite *UserProfileDeliveryTestSuite) SetupTest() {
	suite.serviceTest = new(serviceMockUserProfile)
}

func (suite *UserProfileDeliveryTestSuite) TestBuildUserProfileController() {
	resultTest := NewUserProfile(suite.serviceTest)
	var dummyImpl *UserProfileController
	assert.NotNil(suite.T(), resultTest)
	assert.Implements(suite.T(), dummyImpl, resultTest)
}

func (suite *UserProfileDeliveryTestSuite) TestSaveUserProfileDelivery() {
	suite.serviceTest.(*serviceMockUserProfile).On("SaveUserProfile", dummyUserProfile[0]).Return(nil)
	deliveryTest := NewUserProfile(suite.serviceTest)
	_, err := deliveryTest.Save(dummyUserProfile[0])
	assert.Nil(suite.T(), err)
}

func (suite *UserProfileDeliveryTestSuite) TestUpdateUserProfileDelivery() {
	suite.serviceTest.(*serviceMockUserProfile).On("UpdateUserProfile", dummyUserProfile[0]).Return(nil)
	deliveryTest := NewUserProfile(suite.serviceTest)
	err := deliveryTest.Update(dummyUserProfile[0])
	assert.Nil(suite.T(), err)
}

func (suite *UserProfileDeliveryTestSuite) TestDeleteUserProfileDelivery() {
	suite.serviceTest.(*serviceMockUserProfile).On("DeleteUserProfile", dummyUserProfile[0]).Return(nil)
	deliveryTest := NewUserProfile(suite.serviceTest)
	err := deliveryTest.Delete(dummyUserProfile[0])
	assert.Nil(suite.T(), err)
}

func (suite *UserProfileDeliveryTestSuite) TestGetAllUserProfiles() {
	suite.serviceTest.(*serviceMockUserProfile).On("GetAllUserProfiles", dummyUserProfile).Return(dummyUserProfile)
	deliveryTest := NewUserProfile(suite.serviceTest)
	dummyUserProfile := deliveryTest.GetAllUserProfiles()
	assert.Equal(suite.T(), dummyUserProfile, dummyUserProfile)
}

func (suite *UserProfileDeliveryTestSuite) TestGetUserProfile() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	suite.serviceTest.(*serviceMockUserProfile).On("GetUserProfile", dummyUserProfile[0].ID).Return(dummyUserProfile[0], nil)
	deliveryTest := NewUserProfile(suite.serviceTest)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}
	dummyUserProfile := deliveryTest.GetUserProfile(c)
	assert.NotNil(suite.T(), dummyUserProfile[0])
	assert.Equal(suite.T(), dummyUserProfile[0], dummyUserProfile[0])
}

func TestUserProfileDeliveryTestSuite(t *testing.T) {
	suite.Run(t, new(UserProfileDeliveryTestSuite))
}
