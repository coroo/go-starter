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

type UserRepositoryTestSuite struct {
	suite.Suite
	ctx *gin.Context
	db  *gorm.DB
}

func (suite *UserRepositoryTestSuite) SetupTest() {
	suite.db, _ = config.ConnectDB()
}

func (suite *UserRepositoryTestSuite) BuildNewUserRepository() {
	repoTest := NewUserRepository()
	var dummyImpl *UserRepository
	assert.NotNil(suite.T(), repoTest)
	assert.Implements(suite.T(), dummyImpl, repoTest)
}

func (suite *UserRepositoryTestSuite) CreateUser() {
	repoTest := NewUserRepository()
	dummyUser := entity.User{
		ID				: 1,
		Email			: "kuncoro@test.com",
		Password		: "password",
		Name			: "jl lorem ipsum",
	}
	_, err := repoTest.SaveUser(dummyUser)
	assert.Nil(suite.T(), err)
}

func (suite *UserRepositoryTestSuite) UpdateUser() {
	repoTest := NewUserRepository()
	dummyUser := entity.User{
		ID				: 1,
		Email			: "kuncoro3@test.com",
		Password		: "password",
		Name			: "jl lorem ipsum",
	}
	userDummy := repoTest.UpdateUser(dummyUser)
	assert.Nil(suite.T(), userDummy)
}

func (suite *UserRepositoryTestSuite) AuthUser() {
	repoTest := NewUserRepository()
	dummyUser := entity.User{
		Email			: "kuncoro@test.com",
		Password		: "password",
	}
	userDummy := repoTest.AuthUser(dummyUser)
	assert.NotNil(suite.T(), userDummy)
}

func (suite *UserRepositoryTestSuite) GetAllUsers() {
	repoTest := NewUserRepository()
	userDummy := repoTest.GetAllUsers()
	assert.NotNil(suite.T(), userDummy)
}

func (suite *UserRepositoryTestSuite) GetUser() {
	repoTest := NewUserRepository()
	userDummy := repoTest.GetUser("1")
	assert.NotNil(suite.T(), userDummy)
}

func (suite *UserRepositoryTestSuite) RemoveUser() {
	repoTest := NewUserRepository()
	dummyUser := entity.User{
		ID: 1,
	}
	userDummy := repoTest.DeleteUser(dummyUser)
	assert.Nil(suite.T(), userDummy)
}

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
