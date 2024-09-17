package repositories_test

import (
	"aquaculture/models"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllProductTypes(t *testing.T) {
	t.Run("Get All Product Types | Valid", func(t *testing.T) {
		productTypeRepository.On("GetAll").Return([]models.ProductType{}, nil).Once()

		result, err := productTypeService.GetAll()

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Get All Product Types | Invalid", func(t *testing.T) {
		productTypeRepository.On("GetAll").Return([]models.ProductType{}, errors.New("whoops")).Once()

		result, err := productTypeService.GetAll()

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestGetProductTypeByID(t *testing.T) {
	t.Run("Get Product Type by ID | Valid", func(t *testing.T) {
		productTypeRepository.On("GetByID", "1").Return(models.ProductType{}, nil).Once()

		result, err := productTypeService.GetByID("1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Get Product Type by ID | Invalid", func(t *testing.T) {
		productTypeRepository.On("GetByID", "-1").Return(models.ProductType{}, errors.New("whoops")).Once()

		result, err := productTypeService.GetByID("-1")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestCreateProductType(t *testing.T) {
	t.Run("Create Product Type | Valid", func(t *testing.T) {
		productTypeRepository.On("Create", models.ProductTypeRequest{}).Return(models.ProductType{}, nil).Once()

		result, err := productTypeService.Create(models.ProductTypeRequest{})

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Create Product Type | Invalid", func(t *testing.T) {
		productTypeRepository.On("Create", models.ProductTypeRequest{}).Return(models.ProductType{}, errors.New("whoops")).Once()

		result, err := productTypeService.Create(models.ProductTypeRequest{})

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestUpdateProductType(t *testing.T) {
	t.Run("Update Product Type | Valid", func(t *testing.T) {
		productTypeRepository.On("Update", models.ProductTypeRequest{}, "1").Return(models.ProductType{}, nil).Once()

		result, err := productTypeService.Update(models.ProductTypeRequest{}, "1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Update Product Type | Invalid", func(t *testing.T) {
		productTypeRepository.On("Update", models.ProductTypeRequest{}, "-1").Return(models.ProductType{}, errors.New("whoops")).Once()

		result, err := productTypeService.Update(models.ProductTypeRequest{}, "-1")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestDeleteProductType(t *testing.T) {
	t.Run("Delete Product Type | Valid", func(t *testing.T) {
		productTypeRepository.On("Delete", "1").Return(nil).Once()

		err := productTypeService.Delete("1")

		assert.Nil(t, err)
	})

	t.Run("Delete Product Type | Invalid", func(t *testing.T) {
		productTypeRepository.On("Delete", "-1").Return(errors.New("whoops")).Once()

		err := productTypeService.Delete("-1")

		assert.NotNil(t, err)
	})
}
