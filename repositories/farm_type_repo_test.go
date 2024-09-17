package repositories_test

import (
	"aquaculture/models"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllFarmTypes(t *testing.T) {
	t.Run("Get All Farm Types | Valid", func(t *testing.T) {
		farmTypeRepository.On("GetAll").Return([]models.FarmType{}, nil).Once()

		result, err := farmTypeService.GetAll()

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Get All Farm Types | Invalid", func(t *testing.T) {
		farmTypeRepository.On("GetAll").Return([]models.FarmType{}, errors.New("whoops")).Once()

		result, err := farmTypeService.GetAll()

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestGetFarmTypeByID(t *testing.T) {
	t.Run("Get Farm Type by ID | Valid", func(t *testing.T) {
		farmTypeRepository.On("GetByID", "1").Return(models.FarmType{}, nil).Once()

		result, err := farmTypeService.GetByID("1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Get Farm Type by ID | Invalid", func(t *testing.T) {
		farmTypeRepository.On("GetByID", "-1").Return(models.FarmType{}, errors.New("whoops")).Once()

		result, err := farmTypeService.GetByID("-1")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestCreateFarmType(t *testing.T) {
	t.Run("Create Farm Type | Valid", func(t *testing.T) {
		farmTypeRepository.On("Create", models.FarmTypeRequest{}).Return(models.FarmType{}, nil).Once()

		result, err := farmTypeService.Create(models.FarmTypeRequest{})

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Create Farm Type | Invalid", func(t *testing.T) {
		farmTypeRepository.On("Create", models.FarmTypeRequest{}).Return(models.FarmType{}, errors.New("whoops")).Once()

		result, err := farmTypeService.Create(models.FarmTypeRequest{})

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestUpdateFarmType(t *testing.T) {
	t.Run("Update Farm Type | Valid", func(t *testing.T) {
		farmTypeRepository.On("Update", models.FarmTypeRequest{}, "1").Return(models.FarmType{}, nil).Once()

		result, err := farmTypeService.Update(models.FarmTypeRequest{}, "1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Update Farm Type | Invalid", func(t *testing.T) {
		farmTypeRepository.On("Update", models.FarmTypeRequest{}, "-1").Return(models.FarmType{}, errors.New("whoops")).Once()

		result, err := farmTypeService.Update(models.FarmTypeRequest{}, "-1")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestDeleteFarmType(t *testing.T) {
	t.Run("Delete Farm Type | Valid", func(t *testing.T) {
		farmTypeRepository.On("Delete", "1").Return(nil).Once()

		err := farmTypeService.Delete("1")

		assert.Nil(t, err)
	})

	t.Run("Delete Farm Type | Invalid", func(t *testing.T) {
		farmTypeRepository.On("Delete", "-1").Return(errors.New("whoops")).Once()

		err := farmTypeService.Delete("-1")

		assert.NotNil(t, err)
	})
}
