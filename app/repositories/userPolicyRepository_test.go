package repositories

import (
	"time"
	"net/http/httptest"
	"testing"

	entity "github.com/coroo/go-starter/app/entity"
	"github.com/coroo/go-starter/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type userPolicyRepositoryTestSuite struct {
	suite.Suite
	ctx *gin.Context
	db  *gorm.DB
}

func (suite *userPolicyRepositoryTestSuite) SetupuserPolicyRepositoryTest() {
	suite.db, _ = config.ConnectDB()
}

func (suite *userPolicyRepositoryTestSuite) TestBuildNewUserRepository() {
	repoTest := NewUserPolicyRepository()
	var dummyImpl *UserPolicyRepository
	assert.NotNil(suite.T(), repoTest)
	assert.Implements(suite.T(), dummyImpl, repoTest)
}

func (suite *userPolicyRepositoryTestSuite) TestSaveUserPolicy() {
	repoTest := NewUserPolicyRepository()
	dummyUserPolicy := entity.UserPolicy{
		ID				: 1,
		PolicyNumber	: "30012341234123",
		ProposalNumber	: "30012341234122",
		ProposalDate	: time.Now(),
		CommDate		: time.Now(),
		IssuedDate		: time.Now(),
		ProductCode		: "test123",
		ProductName		: "Test Product 1",
		BenefitLevel	: "A",
		PhName			: "Test Name One",
		PhGender		: "M",
		PhDob			: time.Now(),
		PhProvince		: "Bogor",
		PhCitizenId		: "1234123412341234",
		PhEmail			: "testmail@mail.com",
		PhMobile		: "081234567891",
		PhHomeTel		: "",
		LifeAssured		: "Test Assured one",
		LaDob			: time.Now(),
		LaGender		: "M",
		LaIdNumber		: "1234123412341231",
		LaRelation		: "Diri Sendiri/ Self",
		TotalPremium	: "1000000",
		Beneficiary		: "Test Beneficiary One",
		BeneDob			: time.Now(),
		BeneGender		: "M",
		BeneIdNumber	: "1234123412341232",
		Frequency		: "Monthly",
		PolicyAmount	: 100000,
		LastPaidDate	: time.Now(),
		Ndd				: "2021-12-21",
		OverduePremium	: 10,
		Mop				: "Digital selling Debit Card 4",
		PolicyStatus	: "Waiting for Valodate",
		AdditionalStatus: "Accepted",
		StatusDate		: time.Now(),
		UpdatedAt		: time.Now(),
	}
	repoTest.SaveUserPolicy(dummyUserPolicy)
}

func (suite *UserRepositoryTestSuite) TestUpdateUserPolicy() {
	repoTest := NewUserPolicyRepository()
	dummyUserPolicy := entity.UserPolicy{
		ID				: 1,
		PolicyNumber	: "30012341234124",
		ProposalNumber	: "30012341234121",
		ProposalDate	: time.Now(),
		CommDate		: time.Now(),
		IssuedDate		: time.Now(),
		ProductCode		: "test123",
		ProductName		: "Test Product 2",
		BenefitLevel	: "A",
		PhName			: "Test Name Two",
		PhGender		: "M",
		PhDob			: time.Now(),
		PhProvince		: "Bogor",
		PhCitizenId		: "1234123412341234",
		PhEmail			: "testmail@mail.com",
		PhMobile		: "081234567891",
		PhHomeTel		: "",
		LifeAssured		: "Test Assured Two",
		LaDob			: time.Now(),
		LaGender		: "M",
		LaIdNumber		: "1234123412341231",
		LaRelation		: "Diri Sendiri/ Self",
		TotalPremium	: "1000000",
		Beneficiary		: "Test Beneficiary Two",
		BeneDob			: time.Now(),
		BeneGender		: "M",
		BeneIdNumber	: "1234123412341232",
		Frequency		: "Monthly",
		PolicyAmount	: 100000,
		LastPaidDate	: time.Now(),
		Ndd				: "2021-12-21",
		OverduePremium	: 10,
		Mop				: "Digital selling Debit Card 4",
		PolicyStatus	: "Waiting for Valodate",
		AdditionalStatus: "Accepted",
		StatusDate		: time.Now(),
		UpdatedAt		: time.Now(),
	}
	repoTest.UpdateUserPolicy(dummyUserPolicy)
}

func (suite *UserRepositoryTestSuite) TestGetAllUserPolicies() {
	repoTest := NewUserPolicyRepository()
	userPolicyDummy := repoTest.GetAllUserPolicies("")
	assert.NotNil(suite.T(), userPolicyDummy)
}

func (suite *UserRepositoryTestSuite) TestGetUserPolicy() {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}

	repoTest := NewUserPolicyRepository()
	userPolicyDummy := repoTest.GetUserPolicy(c.Param("id"))
	assert.NotNil(suite.T(), userPolicyDummy)
}

func (suite *UserRepositoryTestSuite) DeleteUserPolicy() {
	repoTest := NewUserPolicyRepository()
	userPolicyDummy := entity.UserPolicy{
		ID: 1,
	}
	repoTest.DeleteUserPolicy(userPolicyDummy)
}

func TestUserPolicyRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(userPolicyRepositoryTestSuite))
}
