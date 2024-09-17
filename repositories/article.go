package repositories

import (
	"aquaculture/database"
	"aquaculture/models"
)

type ArticleRepositoryImpl struct{}

func InitArticleRepository() ArticleRepository {
	return &ArticleRepositoryImpl{}
}

func (ar *ArticleRepositoryImpl) GetAll() ([]models.Article, error) {
	var articles []models.Article

	if err := database.DB.Find(&articles).Error; err != nil {
		return []models.Article{}, err
	}

	return articles, nil
}

func (ar *ArticleRepositoryImpl) GetByID(id string) (models.Article, error) {
	var article models.Article

	if err := database.DB.First(&article, "id = ?", id).Error; err != nil {
		return models.Article{}, err
	}

	return article, nil
}

func (ar *ArticleRepositoryImpl) Create(aReq models.ArticleRequest) (models.Article, error) {
	var article models.Article = models.Article{
		Title: aReq.Title,
		Body:  aReq.Body,
	}

	result := database.DB.Create(&article)

	if err := result.Error; err != nil {
		return models.Article{}, err
	}

	if err := result.Last(&article).Error; err != nil {
		return models.Article{}, err
	}

	return article, nil
}

func (ar *ArticleRepositoryImpl) Update(aReq models.ArticleRequest, id string) (models.Article, error) {
	article, err := ar.GetByID(id)

	if err != nil {
		return models.Article{}, err
	}

	article.Title = aReq.Title
	article.Body = aReq.Body

	if err := database.DB.Save(&article).Error; err != nil {
		return models.Article{}, err
	}

	return article, nil
}

func (ar *ArticleRepositoryImpl) Delete(id string) error {
	article, err := ar.GetByID(id)

	if err != nil {
		return err
	}

	if err := database.DB.Delete(&article).Error; err != nil {
		return err
	}

	return nil
}
