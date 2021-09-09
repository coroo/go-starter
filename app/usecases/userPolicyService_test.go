package usecases

import (
	"time"
	"testing"

	entity "github.com/coroo/go-starter/app/entity"
	repositories "github.com/coroo/go-starter/app/repositories"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type repoMockUserPolicy struct {
	mock.Mock
}

var dummyUserPolicy = []entity.UserPolicy{
	entity.UserPolicy{
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
	}, entity.UserPolicy{
		ID				: 2,
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
	},
}

func (r *repoMockUserPolicy) SaveUserPolicy(userPolicy entity.UserPolicy) {
}

func (r *repoMockUserPolicy) UpdateUserPolicy(userPolicy entity.UserPolicy) {
	
}

func (r *repoMockUserPolicy) DeleteUserPolicy(userPolicy entity.UserPolicy) {
	
}

func (r *repoMockUserPolicy) GetAllUserPolicies(is_overdue string) []entity.UserPolicy {
	return dummyUserPolicy
}

func (r *repoMockUserPolicy) GetUserPolicy(id string) []entity.UserPolicy {
	return dummyUserPolicy
}

func (r *repoMockUserPolicy) CloseDB() {
}

type UserPolicyUsecaseTestSuite struct {
	suite.Suite
	repositoryTest repositories.UserPolicyRepository
}

func (suite *UserPolicyUsecaseTestSuite) SetupTest() {
	suite.repositoryTest = new(repoMockUserPolicy)
}

func (suite *UserPolicyUsecaseTestSuite) TestGetAllUserPolicies() {
	suite.repositoryTest.(*repoMockUserPolicy).On("GetAllUserPolicies", dummyUserPolicy).Return(dummyUserPolicy)
	useCaseTest := NewUserPolicyService(suite.repositoryTest)
	dummyUser := useCaseTest.GetAllUserPolicies("")
	assert.Equal(suite.T(), dummyUser, dummyUserPolicy)
}

func (suite *UserPolicyUsecaseTestSuite) TestGetUserPolicy() {
	suite.repositoryTest.(*repoMockUserPolicy).On("GetUserPolicy", dummyUserPolicy[0]).Return(dummyUserPolicy[0])
	useCaseTest := NewUserPolicyService(suite.repositoryTest)
	dummyUser := useCaseTest.GetUserPolicy("1")
	assert.Equal(suite.T(), dummyUser, dummyUserPolicy)
}

func TestUserPolicyUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserPolicyUsecaseTestSuite))
}
