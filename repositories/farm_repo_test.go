package repositories_test

import (
	"aquaculture/models"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllFarms(t *testing.T) {
	t.Run("Get All Farms | Valid", func(t *testing.T) {
		farmRepository.On("GetAll").Return([]models.Farm{}, nil).Once()

		result, err := farmService.GetAll()

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Get All Farms | Invalid", func(t *testing.T) {
		farmRepository.On("GetAll").Return([]models.Farm{}, errors.New("whoops")).Once()

		result, err := farmService.GetAll()

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestGetFarmByID(t *testing.T) {
	t.Run("Get Farm by ID | Valid", func(t *testing.T) {
		farmRepository.On("GetByID", "1").Return(models.Farm{}, nil).Once()

		result, err := farmService.GetByID("1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Get Farm by ID | Invalid", func(t *testing.T) {
		farmRepository.On("GetByID", "-1").Return(models.Farm{}, errors.New("whoops")).Once()

		result, err := farmService.GetByID("-1")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestCreateFarm(t *testing.T) {
	t.Run("Create Farm | Valid", func(t *testing.T) {
		farmRepository.On("Create", models.FarmRequest{}).Return(models.Farm{}, nil).Once()

		result, err := farmService.Create(models.FarmRequest{})

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Create Farm | Invalid", func(t *testing.T) {
		farmRepository.On("Create", models.FarmRequest{}).Return(models.Farm{}, errors.New("whoops")).Once()

		result, err := farmService.Create(models.FarmRequest{})

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestUpdateFarm(t *testing.T) {
	t.Run("Update Farm | Valid", func(t *testing.T) {
		farmRepository.On("Update", models.FarmRequest{}, "1").Return(models.Farm{}, nil).Once()

		result, err := farmService.Update(models.FarmRequest{}, "1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Update Farm | Invalid", func(t *testing.T) {
		farmRepository.On("Update", models.FarmRequest{}, "-1").Return(models.Farm{}, errors.New("whoops")).Once()

		result, err := farmService.Update(models.FarmRequest{}, "-1")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestDeleteFarm(t *testing.T) {
	t.Run("Delete Farm | Valid", func(t *testing.T) {
		farmRepository.On("Delete", "1").Return(nil).Once()

		err := farmService.Delete("1")

		assert.Nil(t, err)
	})

	t.Run("Delete Farm | Invalid", func(t *testing.T) {
		farmRepository.On("Delete", "-1").Return(errors.New("whoops")).Once()

		err := farmService.Delete("-1")

		assert.NotNil(t, err)
	})
}
