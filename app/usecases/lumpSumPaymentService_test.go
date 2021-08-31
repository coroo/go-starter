package usecases

import (
	"time"

	entity "github.com/coroo/go-starter/app/entity"
	repositories "github.com/coroo/go-starter/app/repositories"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type repoMockLumpSumPayment struct {
	mock.Mock
}

var dummyLumpSumPayment = []entity.LumpSumPayment{
	entity.LumpSumPayment{
		ID					: 1,
		FeeId				: "1234567",
		CollectionId		: "1234567",
		ProposalNumber		: "30012341234122",
		PolicyNumber		: "30012341234123",
		FirstEffectiveDate	: time.Now(),
		EffectiveDate		: time.Now(),
		SettledDate			: time.Now(),
		PaymentMethod		: "Indomaret",
		TotalAmount			: 10000,
		BankName			: "BCA",
	}, entity.LumpSumPayment{
		ID					: 2,
		FeeId				: "1234568",
		CollectionId		: "1234568",
		ProposalNumber		: "30012341234124",
		PolicyNumber		: "30012341234121",
		FirstEffectiveDate	: time.Now(),
		EffectiveDate		: time.Now(),
		SettledDate			: time.Now(),
		PaymentMethod		: "Indomaret",
		TotalAmount			: 10000,
		BankName			: "BCA",
	},
}

func (r *repoMockLumpSumPayment) CreateLumSumPayment(lumpSumPayment entity.LumpSumPayment) {
}

func (r *repoMockLumpSumPayment) UpdateLumSumPayment(lumpSumPayment entity.LumpSumPayment) {
	
}

func (r *repoMockLumpSumPayment) DeleteLumSumPayment(lumpSumPayment entity.LumpSumPayment) {
	
}

func (r *repoMockLumpSumPayment) GetAllLumpSumPayments() []entity.LumpSumPayment {
	return dummyLumpSumPayment
}

func (r *repoMockLumpSumPayment) GetAllLatestGroupLumpSumPayments() []entity.LumpSumPayment {
	return dummyLumpSumPayment
}

func (r *repoMockLumpSumPayment) GetLumpSumPayment(policyNumber string) []entity.LumpSumPayment {
	return dummyLumpSumPayment
}

func (r *repoMockLumpSumPayment) CloseDB() {
}

type LumpSumPaymentUsecaseTestSuite struct {
	suite.Suite
	repositoryTest repositories.LumpSumPaymentRepository
}

func (suite *LumpSumPaymentUsecaseTestSuite) SetupLumpSumPaymentTest() {
	suite.repositoryTest = new(repoMockLumpSumPayment)
}

func (suite *LumpSumPaymentUsecaseTestSuite) TestGetAllLumpSumPayments() {
	suite.repositoryTest.(*repoMockLumpSumPayment).On("GetAllLumpSumPayments", dummyOdsEtlPayment).Return(dummyOdsEtlPayment)
	useCaseTest := NewLumpSumPaymentService(suite.repositoryTest)
	dummyUsecase := useCaseTest.GetAllLumpSumPayments()
	assert.Equal(suite.T(), dummyUsecase, dummyOdsEtlPayment[0])
}

func (suite *LumpSumPaymentUsecaseTestSuite) TestOdsMapEtlLatestPayment() {
	suite.repositoryTest.(*repoMockLumpSumPayment).On("GetAllLatestGroupLumpSumPayments").Return(dummyOdsEtlPayment)
	useCaseTest := NewLumpSumPaymentService(suite.repositoryTest)
	dummyUsecase := useCaseTest.OdsMapEtlLatestPayment()
	assert.Equal(suite.T(), dummyUsecase, dummyOdsEtlPayment)
}

func (suite *LumpSumPaymentUsecaseTestSuite) TestGetLumpSumPayment() {
	suite.repositoryTest.(*repoMockLumpSumPayment).On("GetLumpSumPayment").Return(dummySyEtlPayment[0])
	useCaseTest := NewLumpSumPaymentService(suite.repositoryTest)
	dummyUsecase := useCaseTest.GetLumpSumPayment("300123123123")
	assert.Equal(suite.T(), dummyUsecase, dummyOdsEtlPayment[0])
}

