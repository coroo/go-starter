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

type repoMockSyEtlPayment struct {
	mock.Mock
}

var dummySyEtlPayment = []entity.SyEtlPayment{
	entity.SyEtlPayment{
		ID					: 1,
		ProposalNumber		: "30012341234123",
		PolicyNumber		: "30012341234122",
		PaidDate			: time.Now(),
		PaymentMethodName	: "Indomaret",
		PolicyStatus		: "closed",
		TotalPremium		: 10000,
		UpdatedAt			: time.Now(),
	}, entity.SyEtlPayment{
		ID					: 2,
		ProposalNumber		: "30012341234124",
		PolicyNumber		: "30012341234121",
		PaidDate			: time.Now(),
		PaymentMethodName	: "Indomaret",
		PolicyStatus		: "closed",
		TotalPremium		: 11000,
		UpdatedAt			: time.Now(),
	},
}

func (r *repoMockSyEtlPayment) CreateSyEtlPayment(syEtlPayment entity.SyEtlPayment) {
}

func (r *repoMockSyEtlPayment) UpdateSyEtlPayment(syEtlPayment entity.SyEtlPayment) {
	
}

func (r *repoMockSyEtlPayment) DeleteSyEtlPayment(syEtlPayment entity.SyEtlPayment) {
	
}

func (r *repoMockSyEtlPayment) GetAllSyEtlPayments() []entity.SyEtlPayment {
	return dummySyEtlPayment
}

func (r *repoMockSyEtlPayment) GetAllLatestGroupSyEtlPayments() []entity.SyEtlPayment {
	return dummySyEtlPayment
}

func (r *repoMockSyEtlPayment) GetSyEtlPayment(policyNumber string) []entity.SyEtlPayment {
	return dummySyEtlPayment
}

func (r *repoMockSyEtlPayment) TruncateTableSyEtlPayments(){
}

func (r *repoMockSyEtlPayment) CloseDB() {
}

type SyEtlPaymentUsecaseTestSuite struct {
	suite.Suite
	repositoryTest repositories.SyEtlPaymentRepository
}

func (suite *SyEtlPaymentUsecaseTestSuite) SetupTest() {
	suite.repositoryTest = new(repoMockSyEtlPayment)
}

func (suite *SyEtlPaymentUsecaseTestSuite) TestCreateSyEtlPayment() {
	suite.repositoryTest.(*repoMockSyEtlPayment).On("CreateSyEtlPayment", dummySyEtlPayment[0]).Return(dummySyOdsEtlPayment[0])
	useCaseTest := NewSyEtlPaymentService(suite.repositoryTest)
	dummyUsecase := useCaseTest.CreateSyEtlPayment(dummySyEtlPayment[0])
	assert.Equal(suite.T(), dummyUsecase, nil)
}

// tanya soal httprequest create
func (suite *SyEtlPaymentUsecaseTestSuite) TestSyOdsMapEtlLatestPayment() {
	suite.repositoryTest.(*repoMockSyEtlPayment).On("GetAllSyEtlPayments").Return(dummySyEtlPayment)
	useCaseTest := NewSyEtlPaymentService(suite.repositoryTest)
	dummyUsecase := useCaseTest.GetAllSyEtlPayments()
	assert.Equal(suite.T(), dummyUsecase, dummySyEtlPayment)
}

func (suite *SyEtlPaymentUsecaseTestSuite) TestGetAllSyEtlPayments() {
	suite.repositoryTest.(*repoMockSyEtlPayment).On("GetAllSyEtlPayments").Return(dummySyEtlPayment)
	useCaseTest := NewSyEtlPaymentService(suite.repositoryTest)
	dummyUsecase := useCaseTest.GetAllSyEtlPayments()
	assert.Equal(suite.T(), dummyUsecase, dummySyEtlPayment)
}

func (suite *SyEtlPaymentUsecaseTestSuite) TestGetSyEtlPayment() {
	suite.repositoryTest.(*repoMockSyEtlPayment).On("GetSyEtlPayment").Return(dummySyEtlPayment)
	useCaseTest := NewSyEtlPaymentService(suite.repositoryTest)
	dummyUsecase := useCaseTest.GetSyEtlPayment("300123123123")
	assert.Equal(suite.T(), dummyUsecase, dummySyEtlPayment)
}

func (suite *SyEtlPaymentUsecaseTestSuite) TestTruncateTableSyEtlPayments() {
	suite.repositoryTest.(*repoMockSyEtlPayment).On("TruncateTableSyEtlPayments")
	useCaseTest := NewSyEtlPaymentService(suite.repositoryTest)
	dummyUsecase := useCaseTest.TruncateTableSyEtlPayments()
	assert.Equal(suite.T(), dummyUsecase, nil)
}

func TestSyEtlPaymentUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(SyEtlPaymentUsecaseTestSuite))
}