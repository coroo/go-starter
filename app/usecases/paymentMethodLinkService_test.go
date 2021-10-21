package usecases

import (
	"net/http/httptest"
	"testing"
	"time"

	entity "github.com/coroo/go-starter/app/entity"
	repositories "github.com/coroo/go-starter/app/repositories"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// dummy data
var dummyPaymentMethodLink = []entity.PaymentMethodLink{
	entity.PaymentMethodLink{
		ID				: 1,
		PaymentMethodCode 		: "gopay",
		ProcessType	 			: "linking",
		Url 					: "https://www.superyou.com/gopay-linking/:encryptedaccountid",
		CreatedAt		: time.Now(),
		UpdatedAt		: time.Now(),
	}, entity.PaymentMethodLink{
		ID				: 2,
		PaymentMethodCode 		: "gopay",
		ProcessType	 			: "payment",
		Url 					: "https://gopay.payment.link/:accountid",
		CreatedAt		: time.Now(),
		UpdatedAt		: time.Now(),
	},
}

type repoMockPaymentMethodLink struct {
	mock.Mock
}

func (r *repoMockPaymentMethodLink) SavePaymentMethodLink(paymentMethodLink entity.PaymentMethodLink) (int, error) {
	return 0, nil
}

func (r *repoMockPaymentMethodLink) UpdatePaymentMethodLink(paymentMethodLink entity.PaymentMethodLink) error {
	args := r.Called(paymentMethodLink)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (r *repoMockPaymentMethodLink) DeletePaymentMethodLink(paymentMethodLink entity.PaymentMethodLink) error {
	args := r.Called(paymentMethodLink)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (r *repoMockPaymentMethodLink) GetAllPaymentMethodLinks() []entity.PaymentMethodLink {
	return dummyPaymentMethodLink
}

func (r *repoMockPaymentMethodLink) GetPaymentMethodLink(id string) []entity.PaymentMethodLink {
	return dummyPaymentMethodLink
}

func (r *repoMockPaymentMethodLink) GetPaymentMethodLinkByUuid(uuid string) entity.PaymentMethodLink {
	return dummyPaymentMethodLink[0]
}

func (r *repoMockPaymentMethodLink) GetActivePaymentMethodLinkByCode(code string) entity.PaymentMethodLink {
	return dummyPaymentMethodLink[0]
}

func (r *repoMockPaymentMethodLink) GetPaymentMethodLinkByCode(code string) entity.PaymentMethodLink {
	return dummyPaymentMethodLink[0]
}

func (r *repoMockPaymentMethodLink) CloseDB() {
}

type PaymentMethodLinkUsecaseTestSuite struct {
	suite.Suite
	repositoryTest repositories.PaymentMethodLinkRepository
}

func (suite *PaymentMethodLinkUsecaseTestSuite) SetupTest() {
	suite.repositoryTest = new(repoMockPaymentMethodLink)
}

func (suite *PaymentMethodLinkUsecaseTestSuite) TestBuildPaymentMethodLinkService() {
	resultTest := NewPaymentMethodLinkService(suite.repositoryTest)
	var dummyImpl *PaymentMethodLinkService
	assert.NotNil(suite.T(), resultTest)
	assert.Implements(suite.T(), dummyImpl, resultTest)
	// assert.NotNil(suite.T(), resultTest.(*PaymentMethodLinkService).repositories)
}

func (suite *PaymentMethodLinkUsecaseTestSuite) TestSavePaymentMethodLinkUsecase() {
	suite.repositoryTest.(*repoMockPaymentMethodLink).On("SavePaymentMethodLink", dummyPaymentMethodLink[0]).Return(nil)
	useCaseTest := NewPaymentMethodLinkService(suite.repositoryTest)
	// dummyPaymentMethodLink[0].Password = "Change Password"
	data, _ := useCaseTest.SavePaymentMethodLink(dummyPaymentMethodLink[0])
	assert.NotNil(suite.T(), data)
}

func (suite *PaymentMethodLinkUsecaseTestSuite) TestUpdatePaymentMethodLinkUsecase() {
	suite.repositoryTest.(*repoMockPaymentMethodLink).On("UpdatePaymentMethodLink", dummyPaymentMethodLink[0]).Return(nil)
	useCaseTest := NewPaymentMethodLinkService(suite.repositoryTest)
	err := useCaseTest.UpdatePaymentMethodLink(dummyPaymentMethodLink[0])
	assert.Nil(suite.T(), err)
}

func (suite *PaymentMethodLinkUsecaseTestSuite) TestDeletePaymentMethodLinkUsecase() {
	suite.repositoryTest.(*repoMockPaymentMethodLink).On("DeletePaymentMethodLink", dummyPaymentMethodLink[0]).Return(nil)
	useCaseTest := NewPaymentMethodLinkService(suite.repositoryTest)
	err := useCaseTest.DeletePaymentMethodLink(dummyPaymentMethodLink[0])
	assert.Nil(suite.T(), err)
}

func (suite *PaymentMethodLinkUsecaseTestSuite) TestGetAllPaymentMethodLinks() {
	suite.repositoryTest.(*repoMockPaymentMethodLink).On("GetAllPaymentMethodLinks", dummyPaymentMethodLink).Return(dummyPaymentMethodLink)
	useCaseTest := NewPaymentMethodLinkService(suite.repositoryTest)
	dummyPaymentMethodLink := useCaseTest.GetAllPaymentMethodLinks()
	assert.Equal(suite.T(), dummyPaymentMethodLink, dummyPaymentMethodLink)
}

func (suite *PaymentMethodLinkUsecaseTestSuite) TestGetPaymentMethodLink() {
	suite.repositoryTest.(*repoMockPaymentMethodLink).On("GetPaymentMethodLink", dummyPaymentMethodLink[0].ID).Return(dummyPaymentMethodLink[0], nil)
	useCaseTest := NewPaymentMethodLinkService(suite.repositoryTest)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}
	dummyPaymentMethodLink := useCaseTest.GetPaymentMethodLink(c.Param("id"))
	assert.NotNil(suite.T(), dummyPaymentMethodLink[0])
	assert.Equal(suite.T(), dummyPaymentMethodLink[0], dummyPaymentMethodLink[0])
}

func TestPaymentMethodLinkUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(PaymentMethodLinkUsecaseTestSuite))
}