package usecases

import (
	entity "github.com/coroo/go-starter/app/entity"
	repositories "github.com/coroo/go-starter/app/repositories"
	utils "github.com/coroo/go-starter/app/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserService interface {
	SaveUser(entity.User) (int, error)
	UpdateUser(entity.User) error
	DeleteUser(entity.User) error
	GetAllUsers() []entity.User
	GetUser(ctx *gin.Context) []entity.User
	AuthUser(entity.User) (int, entity.User)
}

type userService struct {
	repositories repositories.UserRepository
}

func NewUser(userRepository repositories.UserRepository) UserService {
	return &userService{
		repositories: userRepository,
	}
}

func (usecases *userService) GetAllUsers() []entity.User {
	return usecases.repositories.GetAllUsers()
}

func (usecases *userService) GetUser(ctx *gin.Context) []entity.User {
	return usecases.repositories.GetUser(ctx)
}

func (usecases *userService) SaveUser(user entity.User) (int, error) {
	user.Uuid = uuid.New().String()
	user.Password, _ = utils.HashPassword(user.Password)
	return usecases.repositories.SaveUser(user)
}

func (usecases *userService) UpdateUser(user entity.User) error {
	return usecases.repositories.UpdateUser(user)
}

func (usecases *userService) DeleteUser(user entity.User) error {
	return usecases.repositories.DeleteUser(user)
}

func (usecases *userService) AuthUser(user entity.User) (int, entity.User) {
	res := usecases.repositories.AuthUser(user)
    match := utils.CheckPasswordHash(user.Password, res.Password)
	if(match){
		return 200, res
	}
	return 401, res
}