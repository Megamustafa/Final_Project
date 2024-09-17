package services

import (
	"aquaculture/models"
	"aquaculture/repositories"
)

type ArticleService struct {
	Repository repositories.ArticleRepository
}

func InitArticleService() ArticleService {
	return ArticleService{
		Repository: &repositories.ArticleRepositoryImpl{},
	}
}

func (as *ArticleService) GetAll() ([]models.Article, error) {
	return as.Repository.GetAll()
}

func (as *ArticleService) GetByID(id string) (models.Article, error) {
	return as.Repository.GetByID(id)
}

func (as *ArticleService) Create(aReq models.ArticleRequest) (models.Article, error) {
	return as.Repository.Create(aReq)
}

func (as *ArticleService) Update(aReq models.ArticleRequest, id string) (models.Article, error) {
	return as.Repository.Update(aReq, id)
}

func (as *ArticleService) Delete(id string) error {
	return as.Repository.Delete(id)
}
