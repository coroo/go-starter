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

type repoMockSyUserInvoice struct {
	mock.Mock
}

var dummySyUserInvoice = []entity.SyUserInvoice{
	entity.SyUserInvoice{
		ID					: 1,
		PolicyNumber		: "30012341234123",
		PolicyGroupNumber	: "30012341234122",
		ProposalNumber		: "123123",
		PaymentMethodName	: "Indomaret",
		TotalPremium		: 123123,
		Status				: "test123",
		PaidAt				: time.Now(),
	}, entity.SyUserInvoice{
		ID					: 2,
		PolicyNumber		: "30012341234124",
		PolicyGroupNumber	: "30012341234121",
		ProposalNumber		: "123123",
		PaymentMethodName	: "Indomaret",
		TotalPremium		: 234234,
		Status				: "test123",
		PaidAt				: time.Now(),
	},
}

func (r *repoMockSyUserInvoice) SaveSyUserInvoice(syUserInvoice entity.SyUserInvoice) {
}

func (r *repoMockSyUserInvoice) UpdateSyUserInvoice(syUserInvoice entity.SyUserInvoice) {
	
}

func (r *repoMockSyUserInvoice) DeleteSyUserInvoice(syUserInvoice entity.SyUserInvoice) {
	
}

func (r *repoMockSyUserInvoice) GetAllPaidUserInvoices() []entity.SyUserInvoice {
	return dummySyUserInvoice
}

func (r *repoMockSyUserInvoice) GetUserInvoice(id string) []entity.SyUserInvoice {
	return dummySyUserInvoice
}

func (r *repoMockSyUserInvoice) CloseDB() {
}

type SyUserInvoiceUsecaseTestSuite struct {
	suite.Suite
	repositoryTest repositories.SyUserInvoiceRepository
}

func (suite *SyUserInvoiceUsecaseTestSuite) SetupTest() {
	suite.repositoryTest = new(repoMockSyUserInvoice)
}

func (suite *SyUserInvoiceUsecaseTestSuite) TestGetAllSyUserInvoices() {
	suite.repositoryTest.(*repoMockSyUserInvoice).On("GetAllSyUserInvoices").Return(dummySyUserInvoice)
	useCaseTest := NewSyUserInvoiceService(suite.repositoryTest)
	dummyUserInvoices := useCaseTest.GetAllSyUserInvoices()
	assert.Equal(suite.T(), dummyUserInvoices, dummySyUserInvoice)
}

func (suite *SyUserInvoiceUsecaseTestSuite) SyMapEtlLatestPayment() {
	suite.repositoryTest.(*repoMockSyUserInvoice).On("SyMapEtlLatestPayment").Return(dummySyUserInvoice)
	useCaseTest := NewSyUserInvoiceService(suite.repositoryTest)
	dummyUserInvoices := useCaseTest.SyMapEtlLatestPayment()
	assert.Equal(suite.T(), dummyUserInvoices, dummySyUserInvoice)
}

func TestSyUserInvoiceUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(SyUserInvoiceUsecaseTestSuite))
}
