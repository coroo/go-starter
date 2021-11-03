package usecases

import (
	entity "github.com/coroo/go-starter/app/entity"
	repositories "github.com/coroo/go-starter/app/repositories"

	// "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type GopayLinkingService interface {
	SaveGopayLinking(entity.GopayLinking) (int, error)
	UpdateGopayLinking(entity.GopayLinking) error
	DeleteGopayLinking(entity.GopayLinking) error
	GetAllGopayLinkings() []entity.GopayLinking
	GetGopayLinking(id string) []entity.GopayLinking
}

type gopayLinkingService struct {
	repositories repositories.GopayLinkingRepository
}

func NewGopayLinkingService(repository repositories.GopayLinkingRepository) GopayLinkingService {
	return &gopayLinkingService{
		repositories: repository,
	}
}

func (usecases *gopayLinkingService) GetAllGopayLinkings() []entity.GopayLinking {
	return usecases.repositories.GetAllGopayLinkings()
}

func (usecases *gopayLinkingService) GetGopayLinking(id string) []entity.GopayLinking {
	return usecases.repositories.GetGopayLinking(id)
}

func (usecases *gopayLinkingService) SaveGopayLinking(gopayLinking entity.GopayLinking) (int, error) {
	return usecases.repositories.SaveGopayLinking(gopayLinking)
}

func (usecases *gopayLinkingService) UpdateGopayLinking(gopayLinking entity.GopayLinking) error {
	return usecases.repositories.UpdateGopayLinking(gopayLinking)
}

func (usecases *gopayLinkingService) DeleteGopayLinking(gopayLinking entity.GopayLinking) error {
	return usecases.repositories.DeleteGopayLinking(gopayLinking)
}