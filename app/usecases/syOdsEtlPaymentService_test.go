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

type repoMockSyOdsEtlPayment struct {
	mock.Mock
}

var dummySyOdsEtlPayment = []entity.SyOdsEtlPayment{
	entity.SyOdsEtlPayment{
		ID					: 1,
		ProposalNumber		: "30012341234123",
		PolicyNumber		: "30012341234122",
		OdsFirstPaidDate	: time.Now(),
		OdsPaidDate			: time.Now(),
		SyPaidDate			: time.Now(),
		PaymentMethodName	: "Indomaret",
		CollectionId		: "1234567",
		SyTotalAmount		: 10000,
		OdsTotalAmount		: 10000,
		PolicyStatus		: "success",
		Status				: "closed",
		StatusDescription	: "Stated as First Payment, use SY data for First Payment",
		UpdatedAt			: time.Now(),
	}, entity.SyOdsEtlPayment{
		ID					: 2,
		ProposalNumber		: "30012341234124",
		PolicyNumber		: "30012341234121",
		OdsFirstPaidDate	: time.Now(),
		OdsPaidDate			: time.Now(),
		SyPaidDate			: time.Now(),
		PaymentMethodName	: "Indomaret",
		CollectionId		: "1234567",
		SyTotalAmount		: 10000,
		OdsTotalAmount		: 10000,
		PolicyStatus		: "success",
		Status				: "closed",
		StatusDescription	: "Stated as First Payment, use SY data for First Payment",
		UpdatedAt			: time.Now(),
	},
}

func (r *repoMockSyOdsEtlPayment) CreateSyOdsEtlPayment(syOdsEtlPayment entity.SyOdsEtlPayment) {
}

func (r *repoMockSyOdsEtlPayment) UpdateSyOdsEtlPayment(syOdsEtlPayment entity.SyOdsEtlPayment) {
	
}

func (r *repoMockSyOdsEtlPayment) DeleteSyOdsEtlPayment(syOdsEtlPayment entity.SyOdsEtlPayment) {
	
}

func (r *repoMockSyOdsEtlPayment) GetAllSyOdsEtlPayments() []entity.SyOdsEtlPayment {
	return dummySyOdsEtlPayment
}

func (r *repoMockSyOdsEtlPayment) GetAllLatestGroupSyOdsEtlPayments() []entity.SyOdsEtlPayment {
	return dummySyOdsEtlPayment
}

func (r *repoMockSyOdsEtlPayment) GetSyOdsEtlPaymentByStatus(status string) []entity.SyOdsEtlPayment {
	return dummySyOdsEtlPayment
}

func (r *repoMockSyOdsEtlPayment) GetSyOdsEtlPaymentDailyByStatus(status string) []entity.SyOdsEtlPayment {
	return dummySyOdsEtlPayment
}

func (r *repoMockSyOdsEtlPayment) GetSyOdsEtlPaymentByPolicyNumber(policyNumber string) []entity.SyOdsEtlPayment {
	return dummySyOdsEtlPayment
}

func (r *repoMockSyOdsEtlPayment) CancelOutstandingSyOdsEtlPayments() []entity.SyOdsEtlPayment {
	return dummySyOdsEtlPayment
}

func (r *repoMockSyOdsEtlPayment) CloseDB() {
}

type SyOdsEtlPaymentUsecaseTestSuite struct {
	suite.Suite
	repositoryTest repositories.SyOdsEtlPaymentRepository
}

func (suite *SyOdsEtlPaymentUsecaseTestSuite) SetupTest() {
	suite.repositoryTest = new(repoMockSyOdsEtlPayment)
}

func (suite *SyOdsEtlPaymentUsecaseTestSuite) TestCreateSyOdsEtlPayment() {
	suite.repositoryTest.(*repoMockSyOdsEtlPayment).On("CreateSyOdsEtlPayment", dummySyOdsEtlPayment).Return(dummySyOdsEtlPayment)
	useCaseTest := NewSyOdsEtlPaymentService(suite.repositoryTest)
	dummyUsecase := useCaseTest.CreateSyOdsEtlPayment(dummySyOdsEtlPayment[0])
	assert.Equal(suite.T(), dummyUsecase, nil)
}

func (suite *SyOdsEtlPaymentUsecaseTestSuite) TestGetAllSyOdsEtlPayments() {
	suite.repositoryTest.(*repoMockSyOdsEtlPayment).On("GetAllSyOdsEtlPayments").Return(dummySyUserInvoice)
	useCaseTest := NewSyOdsEtlPaymentService(suite.repositoryTest)
	dummyUsecase := useCaseTest.GetAllSyOdsEtlPayments()
	assert.Equal(suite.T(), dummyUsecase, dummySyOdsEtlPayment)
}

func (suite *SyOdsEtlPaymentUsecaseTestSuite) TestGetSyOdsEtlPaymentByPolicyNumber() {
	suite.repositoryTest.(*repoMockSyOdsEtlPayment).On("GetSyOdsEtlPaymentByPolicyNumber").Return(dummySyUserInvoice)
	useCaseTest := NewSyOdsEtlPaymentService(suite.repositoryTest)
	dummyUsecase := useCaseTest.GetSyOdsEtlPaymentByPolicyNumber("021021")
	assert.Equal(suite.T(), dummyUsecase, dummySyOdsEtlPayment)
}

func (suite *SyOdsEtlPaymentUsecaseTestSuite) TestGetSyOdsEtlPaymentByStatus() {
	suite.repositoryTest.(*repoMockSyOdsEtlPayment).On("GetSyOdsEtlPaymentByStatus").Return(dummySyUserInvoice)
	useCaseTest := NewSyOdsEtlPaymentService(suite.repositoryTest)
	dummyUsecase := useCaseTest.GetSyOdsEtlPaymentByStatus("closed")
	assert.Equal(suite.T(), dummyUsecase, dummySyOdsEtlPayment)
}

func (suite *SyOdsEtlPaymentUsecaseTestSuite) TestGetSyOdsEtlPaymentDailyByStatus() {
	suite.repositoryTest.(*repoMockSyOdsEtlPayment).On("GetSyOdsEtlPaymentDailyByStatus").Return(dummySyUserInvoice)
	useCaseTest := NewSyOdsEtlPaymentService(suite.repositoryTest)
	dummyUsecase := useCaseTest.GetSyOdsEtlPaymentDailyByStatus("closed")
	assert.Equal(suite.T(), dummyUsecase, dummySyOdsEtlPayment)
}

func (suite *SyOdsEtlPaymentUsecaseTestSuite) TestCancelOutstandingSyOdsEtlPayments() {
	suite.repositoryTest.(*repoMockSyOdsEtlPayment).On("CancelOutstandingSyOdsEtlPayments").Return(dummySyUserInvoice)
	useCaseTest := NewSyOdsEtlPaymentService(suite.repositoryTest)
	dummyUsecase := useCaseTest.CancelOutstandingSyOdsEtlPayments()
	assert.Equal(suite.T(), dummyUsecase, dummySyOdsEtlPayment)
}

func TestSyOdsEtlPaymentUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(SyOdsEtlPaymentUsecaseTestSuite))
}