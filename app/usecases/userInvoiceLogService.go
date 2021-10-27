package usecases

import (
	entity "github.com/coroo/go-starter/app/entity"
	repositories "github.com/coroo/go-starter/app/repositories"

	// "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type UserInvoiceLogService interface {
	SaveUserInvoiceLog(entity.UserInvoiceLog) (int, error)
	UpdateUserInvoiceLog(entity.UserInvoiceLog) error
	DeleteUserInvoiceLog(entity.UserInvoiceLog) error
	GetAllUserInvoiceLogs() []entity.UserInvoiceLog
	GetUserInvoiceLog(id string) []entity.UserInvoiceLog
}

type userInvoiceLogService struct {
	repositories repositories.UserInvoiceLogRepository
}

func NewUserInvoiceLogService(repository repositories.UserInvoiceLogRepository) UserInvoiceLogService {
	return &userInvoiceLogService{
		repositories: repository,
	}
}

func (usecases *userInvoiceLogService) GetAllUserInvoiceLogs() []entity.UserInvoiceLog {
	return usecases.repositories.GetAllUserInvoiceLogs()
}

func (usecases *userInvoiceLogService) GetUserInvoiceLog(id string) []entity.UserInvoiceLog {
	return usecases.repositories.GetUserInvoiceLog(id)
}

func (usecases *userInvoiceLogService) SaveUserInvoiceLog(userInvoiceLog entity.UserInvoiceLog) (int, error) {
	return usecases.repositories.SaveUserInvoiceLog(userInvoiceLog)
}

func (usecases *userInvoiceLogService) UpdateUserInvoiceLog(userInvoiceLog entity.UserInvoiceLog) error {
	return usecases.repositories.UpdateUserInvoiceLog(userInvoiceLog)
}

func (usecases *userInvoiceLogService) DeleteUserInvoiceLog(userInvoiceLog entity.UserInvoiceLog) error {
	return usecases.repositories.DeleteUserInvoiceLog(userInvoiceLog)
}