package repositories_test

import (
	"aquaculture/models"
	"errors"
	"mime/multipart"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllProducts(t *testing.T) {
	t.Run("Get All Products | Valid", func(t *testing.T) {
		productRepository.On("GetAll").Return([]models.Product{}, nil).Once()

		result, err := productService.GetAll()

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Get All Products | Invalid", func(t *testing.T) {
		productRepository.On("GetAll").Return([]models.Product{}, errors.New("whoops")).Once()

		result, err := productService.GetAll()

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestGetProductByID(t *testing.T) {
	t.Run("Get Product by ID | Valid", func(t *testing.T) {
		productRepository.On("GetByID", "1").Return(models.Product{}, nil).Once()

		result, err := productService.GetByID("1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Get Product by ID | Invalid", func(t *testing.T) {
		productRepository.On("GetByID", "-1").Return(models.Product{}, errors.New("whoops")).Once()

		result, err := productService.GetByID("-1")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestCreateProduct(t *testing.T) {
	t.Run("Create Product | Valid", func(t *testing.T) {
		productRepository.On("Create", models.ProductRequest{}).Return(models.Product{}, nil).Once()

		result, err := productService.Create(models.ProductRequest{})

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Create Product | Invalid", func(t *testing.T) {
		productRepository.On("Create", models.ProductRequest{}).Return(models.Product{}, errors.New("whoops")).Once()

		result, err := productService.Create(models.ProductRequest{})

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestUpdateProduct(t *testing.T) {
	t.Run("Update Product | Valid", func(t *testing.T) {
		productRepository.On("Update", models.ProductRequest{}, "1").Return(models.Product{}, nil).Once()

		result, err := productService.Update(models.ProductRequest{}, "1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Update Product | Invalid", func(t *testing.T) {
		productRepository.On("Update", models.ProductRequest{}, "-1").Return(models.Product{}, errors.New("whoops")).Once()

		result, err := productService.Update(models.ProductRequest{}, "-1")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestDeleteProduct(t *testing.T) {
	t.Run("Delete Product | Valid", func(t *testing.T) {
		articleRepository.On("Delete", "1").Return(nil).Once()

		err := articleService.Delete("1")

		assert.Nil(t, err)
	})

	t.Run("Delete Product | Invalid", func(t *testing.T) {
		productRepository.On("Delete", "-1").Return(errors.New("whoops")).Once()

		err := productService.Delete("-1")

		assert.NotNil(t, err)
	})
}

func TestImportFromCSV(t *testing.T) {
	t.Run("Import Products | Valid", func(t *testing.T) {
		productRepository.On("ImportFromCSV", &multipart.FileHeader{}).Return([]models.Product{}, nil).Once()

		result, err := productService.ImportFromCSV(&multipart.FileHeader{})

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Import Products | Invalid", func(t *testing.T) {
		productRepository.On("ImportFromCSV", &multipart.FileHeader{}).Return([]models.Product{}, errors.New("whoops")).Once()

		result, err := productService.ImportFromCSV(&multipart.FileHeader{})

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}
