package usecases

import (
	"time"

	entity "github.com/coroo/go-starter/app/entity"
	repositories "github.com/coroo/go-starter/app/repositories"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type repoMockOdsEtlPayment struct {
	mock.Mock
}

var dummyOdsEtlPayment = []entity.OdsEtlPayment{
	entity.OdsEtlPayment{
		ID					: 1,
		CollectionId		: "1234567",
		ProposalNumber		: "30012341234122",
		PolicyNumber		: "30012341234123",
		FirstPaymentDate	: time.Now(),
		PaymentDate			: time.Now(),
		PaymentMethod		: "Indomaret",
		TotalAmount			: 10000,
		UpdatedAt			: time.Now(),
	}, entity.OdsEtlPayment{
		ID					: 2,
		CollectionId		: "1234568",
		ProposalNumber		: "30012341234124",
		PolicyNumber		: "30012341234121",
		FirstPaymentDate	: time.Now(),
		PaymentDate			: time.Now(),
		PaymentMethod		: "Indomaret",
		TotalAmount			: 10000,
		UpdatedAt			: time.Now(),
	},
}

func (r *repoMockOdsEtlPayment) CreateOdsEtlPayment(odsEtlPayment entity.OdsEtlPayment) {
}

func (r *repoMockOdsEtlPayment) UpdateOdsEtlPayment(odsEtlPayment entity.OdsEtlPayment) {
	
}

func (r *repoMockOdsEtlPayment) DeleteOdsEtlPayment(odsEtlPayment entity.OdsEtlPayment) {
	
}

func (r *repoMockOdsEtlPayment) GetAllOdsEtlPayments() []entity.OdsEtlPayment {
	return dummyOdsEtlPayment
}

func (r *repoMockOdsEtlPayment) GetAllLatestGroupOdsEtlPayments() []entity.OdsEtlPayment {
	return dummyOdsEtlPayment
}

func (r *repoMockOdsEtlPayment) GetOdsEtlPayment(policyNumber string) []entity.OdsEtlPayment {
	return dummyOdsEtlPayment
}

func (r *repoMockOdsEtlPayment) TruncateTableOdsEtlPayments(){
}

func (r *repoMockOdsEtlPayment) CloseDB() {
}

type OdsEtlPaymentUsecaseTestSuite struct {
	suite.Suite
	repositoryTest repositories.OdsEtlPaymentRepository
}

func (suite *OdsEtlPaymentUsecaseTestSuite) SetupOdsEtlPaymentTest() {
	suite.repositoryTest = new(repoMockOdsEtlPayment)
}

func (suite *OdsEtlPaymentUsecaseTestSuite) TestCreateOdsEtlPayment() {
	suite.repositoryTest.(*repoMockOdsEtlPayment).On("CreateOdsEtlPayment", dummyOdsEtlPayment[0]).Return(dummyOdsEtlPayment[0])
	useCaseTest := NewOdsEtlPaymentService(suite.repositoryTest)
	dummyUsecase := useCaseTest.CreateOdsEtlPayment(dummyOdsEtlPayment[0])
	assert.Equal(suite.T(), dummyUsecase, dummyOdsEtlPayment[0])
}

func (suite *OdsEtlPaymentUsecaseTestSuite) TestGetAllOdsEtlPayments() {
	suite.repositoryTest.(*repoMockOdsEtlPayment).On("GetAllOdsEtlPayments").Return(dummyOdsEtlPayment)
	useCaseTest := NewOdsEtlPaymentService(suite.repositoryTest)
	dummyUsecase := useCaseTest.GetAllOdsEtlPayments()
	assert.Equal(suite.T(), dummyUsecase, dummyOdsEtlPayment)
}

func (suite *OdsEtlPaymentUsecaseTestSuite) TestGetOdsEtlPayment() {
	suite.repositoryTest.(*repoMockOdsEtlPayment).On("GetOdsEtlPayment").Return(dummySyEtlPayment[0])
	useCaseTest := NewOdsEtlPaymentService(suite.repositoryTest)
	dummyUsecase := useCaseTest.GetOdsEtlPayment("1")
	assert.Equal(suite.T(), dummyUsecase, dummyOdsEtlPayment[0])
}

func (suite *OdsEtlPaymentUsecaseTestSuite) TestTruncateTableOdsEtlPayments() {
	suite.repositoryTest.(*repoMockOdsEtlPayment).On("TruncateTableOdsEtlPayments").Return(nil)
	useCaseTest := NewOdsEtlPaymentService(suite.repositoryTest)
	dummyUsecase := useCaseTest.TruncateTableOdsEtlPayments()
	assert.Equal(suite.T(), dummyUsecase, nil)
}
