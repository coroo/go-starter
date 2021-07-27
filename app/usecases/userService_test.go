package usecases

import (
	"net/http/httptest"
	"testing"
	"time"

	entity "github.com/coroo/go-pawoon-user/app/entity"
	repositories "github.com/coroo/go-pawoon-user/app/repositories"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// dummy data
var dummyUser = []entity.User{
	entity.User{
		ID				: 1,
		Email			: "kuncoro@test.com",
		Password		: "password",
		Name			: "jl lorem ipsum",
		CreatedAt		: time.Now(),
		UpdatedAt		: time.Now(),
	}, entity.User{
		ID				: 2,
		Email			: "kuncoro@test.com",
		Password		: "password",
		Name			: "jl lorem ipsum",
		CreatedAt		: time.Now(),
		UpdatedAt		: time.Now(),
	},
}

type repoMockUser struct {
	mock.Mock
}

func (r *repoMockUser) SaveUser(user entity.User) (int, error) {
	return 0, nil
}

func (r *repoMockUser) UpdateUser(user entity.User) error {
	args := r.Called(user)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (r *repoMockUser) DeleteUser(user entity.User) error {
	args := r.Called(user)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (r *repoMockUser) GetAllUsers() []entity.User {
	return dummyUser
}

func (r *repoMockUser) GetUser(ctx *gin.Context) []entity.User {
	return dummyUser
}

func (r *repoMockUser) AuthUser(user entity.User) entity.User {
	return dummyUser[0]
}

func (r *repoMockUser) CloseDB() {
}

type UserUsecaseTestSuite struct {
	suite.Suite
	repositoryTest repositories.UserRepository
}

func (suite *UserUsecaseTestSuite) SetupTest() {
	suite.repositoryTest = new(repoMockUser)
}

func (suite *UserUsecaseTestSuite) TestBuildUserService() {
	resultTest := NewUser(suite.repositoryTest)
	var dummyImpl *UserService
	assert.NotNil(suite.T(), resultTest)
	assert.Implements(suite.T(), dummyImpl, resultTest)
	// assert.NotNil(suite.T(), resultTest.(*UserService).repositories)
}

func (suite *UserUsecaseTestSuite) TestSaveUserUsecase() {
	suite.repositoryTest.(*repoMockUser).On("SaveUser", dummyUser[0]).Return(nil)
	useCaseTest := NewUser(suite.repositoryTest)
	// dummyUser[0].Password = "Change Password"
	data, _ := useCaseTest.SaveUser(dummyUser[0])
	assert.NotNil(suite.T(), data)
}

func (suite *UserUsecaseTestSuite) TestUpdateUserUsecase() {
	suite.repositoryTest.(*repoMockUser).On("UpdateUser", dummyUser[0]).Return(nil)
	useCaseTest := NewUser(suite.repositoryTest)
	err := useCaseTest.UpdateUser(dummyUser[0])
	assert.Nil(suite.T(), err)
}

func (suite *UserUsecaseTestSuite) TestDeleteUserUsecase() {
	suite.repositoryTest.(*repoMockUser).On("DeleteUser", dummyUser[0]).Return(nil)
	useCaseTest := NewUser(suite.repositoryTest)
	err := useCaseTest.DeleteUser(dummyUser[0])
	assert.Nil(suite.T(), err)
}

func (suite *UserUsecaseTestSuite) TestGetAllUsers() {
	suite.repositoryTest.(*repoMockUser).On("GetAllUsers", dummyUser).Return(dummyUser)
	useCaseTest := NewUser(suite.repositoryTest)
	dummyUser := useCaseTest.GetAllUsers()
	assert.Equal(suite.T(), dummyUser, dummyUser)
}

func (suite *UserUsecaseTestSuite) TestGetUser() {
	suite.repositoryTest.(*repoMockUser).On("GetUser", dummyUser[0].ID).Return(dummyUser[0], nil)
	useCaseTest := NewUser(suite.repositoryTest)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}
	dummyUser := useCaseTest.GetUser(c)
	assert.NotNil(suite.T(), dummyUser[0])
	assert.Equal(suite.T(), dummyUser[0], dummyUser[0])
}

func (suite *UserUsecaseTestSuite) TestAuthUserUsecase() {
	suite.repositoryTest.(*repoMockUser).On("AuthUser", dummyUser[0]).Return(nil)
	useCaseTest := NewUser(suite.repositoryTest)
	err := useCaseTest.AuthUser(dummyUser[0])
	assert.NotNil(suite.T(), err)
}

func TestUserUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserUsecaseTestSuite))
}
