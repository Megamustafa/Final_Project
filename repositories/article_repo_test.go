package repositories_test

import (
	"aquaculture/models"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllArticles(t *testing.T) {
	t.Run("Get All Articles | Valid", func(t *testing.T) {
		articleRepository.On("GetAll").Return([]models.Article{}, nil).Once()

		result, err := articleService.GetAll()

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Get All Articles | Invalid", func(t *testing.T) {
		articleRepository.On("GetAll").Return([]models.Article{}, errors.New("whoops")).Once()

		result, err := articleService.GetAll()

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestGetArticleByID(t *testing.T) {
	t.Run("Get Article by ID | Valid", func(t *testing.T) {
		articleRepository.On("GetByID", "1").Return(models.Article{}, nil).Once()

		result, err := articleService.GetByID("1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Get Article by ID | Invalid", func(t *testing.T) {
		articleRepository.On("GetByID", "-1").Return(models.Article{}, errors.New("whoops")).Once()

		result, err := articleService.GetByID("-1")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestCreateArticle(t *testing.T) {
	t.Run("Create Article | Valid", func(t *testing.T) {
		articleRepository.On("Create", models.ArticleRequest{}).Return(models.Article{}, nil).Once()

		result, err := articleService.Create(models.ArticleRequest{})

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Create Article | Invalid", func(t *testing.T) {
		articleRepository.On("Create", models.ArticleRequest{}).Return(models.Article{}, errors.New("whoops")).Once()

		result, err := articleService.Create(models.ArticleRequest{})

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestUpdateArticle(t *testing.T) {
	t.Run("Update Article | Valid", func(t *testing.T) {
		articleRepository.On("Update", models.ArticleRequest{}, "1").Return(models.Article{}, nil).Once()

		result, err := articleService.Update(models.ArticleRequest{}, "1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Update Article | Invalid", func(t *testing.T) {
		articleRepository.On("Update", models.ArticleRequest{}, "-1").Return(models.Article{}, errors.New("whoops")).Once()

		result, err := articleService.Update(models.ArticleRequest{}, "-1")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestDeleteArticle(t *testing.T) {
	t.Run("Delete Article | Valid", func(t *testing.T) {
		articleRepository.On("Delete", "1").Return(nil).Once()

		err := articleService.Delete("1")

		assert.Nil(t, err)
	})

	t.Run("Delete Article | Invalid", func(t *testing.T) {
		articleRepository.On("Delete", "-1").Return(errors.New("whoops")).Once()

		err := articleService.Delete("-1")

		assert.NotNil(t, err)
	})
}
