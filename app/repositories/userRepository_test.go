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

func (suite *UserRepositoryTestSuite) TestA_BuildNewUserRepository() {
	repoTest := NewUserRepository()
	var dummyImpl *UserRepository
	assert.NotNil(suite.T(), repoTest)
	assert.Implements(suite.T(), dummyImpl, repoTest)
}

func (suite *UserRepositoryTestSuite) TestB_CreateUser() {
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

func (suite *UserRepositoryTestSuite) TestC_UpdateUser() {
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

func (suite *UserRepositoryTestSuite) TestD_AuthUser() {
	repoTest := NewUserRepository()
	dummyUser2 := entity.User{
		Email			: "kuncoro3@test.com",
		Password		: "password",
	}
	userDummy := repoTest.AuthUser(dummyUser2)
	assert.NotNil(suite.T(), userDummy)
}

func (suite *UserRepositoryTestSuite) TestE_GetAllUsers() {
	repoTest := NewUserRepository()
	userDummy := repoTest.GetAllUsers()
	assert.NotNil(suite.T(), userDummy)
}

func (suite *UserRepositoryTestSuite) TestF_GetUser() {
	repoTest := NewUserRepository()
	userDummy := repoTest.GetUser("1")
	assert.NotNil(suite.T(), userDummy)
}

func (suite *UserRepositoryTestSuite) TestG_GetUserByUuid() {
	repoTest := NewUserRepository()
	userDummy := repoTest.GetUser("1")
	uuid := userDummy[0].Uuid
	userDummy2 := repoTest.GetUserByUuid(uuid)
	assert.NotNil(suite.T(), userDummy2)
}

func (suite *UserRepositoryTestSuite) TestH_RemoveUser() {
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
