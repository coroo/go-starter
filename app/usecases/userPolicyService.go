package usecases

import (
	entity "github.com/coroo/go-starter/app/entity"
	repositories "github.com/coroo/go-starter/app/repositories"

	// "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type UserPolicyService interface {
	// Save(entity.UserPolicy) entity.UserPolicy
	GetAllUserPolicies(is_overdue string) []entity.UserPolicy
	GetUserPolicy(id string) []entity.UserPolicy
}

type userPolicyService struct {
	repositories repositories.UserPolicyRepository
}

func NewUserPolicyService(repository repositories.UserPolicyRepository) UserPolicyService {
	return &userPolicyService{
		repositories: repository,
	}
}

func (usecases *userPolicyService) GetAllUserPolicies(is_overdue string) []entity.UserPolicy {
	return usecases.repositories.GetAllUserPolicies(is_overdue)
}

func (usecases *userPolicyService) GetUserPolicy(id string) []entity.UserPolicy {
	return usecases.repositories.GetUserPolicy(id)
}
