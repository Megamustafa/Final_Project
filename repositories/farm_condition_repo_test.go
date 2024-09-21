package repositories_test

import (
	"aquaculture/models"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllFarmConditions(t *testing.T) {
	t.Run("Get All Farm Conditions | Valid", func(t *testing.T) {
		farmConditionRepository.On("GetAll").Return([]models.FarmCondition{}, nil).Once()

		result, err := farmConditionService.GetAll()

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Get All Farm Conditions | Invalid", func(t *testing.T) {
		farmConditionRepository.On("GetAll").Return([]models.FarmCondition{}, errors.New("whoops")).Once()

		result, err := farmConditionService.GetAll()

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestGetFarmConditionByID(t *testing.T) {
	t.Run("Get Farm Condition by ID | Valid", func(t *testing.T) {
		farmConditionRepository.On("GetByID", "1").Return(models.FarmCondition{}, nil).Once()

		result, err := farmConditionService.GetByID("1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Get Farm Condition by ID | Invalid", func(t *testing.T) {
		farmConditionRepository.On("GetByID", "-1").Return(models.FarmCondition{}, errors.New("whoops")).Once()

		result, err := farmConditionService.GetByID("-1")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestCreateFarmCondition(t *testing.T) {
	t.Run("Create Farm Condition | Valid", func(t *testing.T) {
		farmConditionRepository.On("Create", models.FarmConditionRequest{}).Return(models.FarmCondition{}, nil).Once()

		result, err := farmConditionService.Create(models.FarmConditionRequest{})

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Create Farm Condition | Invalid", func(t *testing.T) {
		farmConditionRepository.On("Create", models.FarmConditionRequest{}).Return(models.FarmCondition{}, errors.New("whoops")).Once()

		result, err := farmConditionService.Create(models.FarmConditionRequest{})

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestUpdateFarmCondition(t *testing.T) {
	t.Run("Update Farm Condition | Valid", func(t *testing.T) {
		farmConditionRepository.On("Update", models.FarmConditionRequest{}, "1").Return(models.FarmCondition{}, nil).Once()

		result, err := farmConditionService.Update(models.FarmConditionRequest{}, "1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Update Farm Condition | Invalid", func(t *testing.T) {
		farmConditionRepository.On("Update", models.FarmConditionRequest{}, "-1").Return(models.FarmCondition{}, errors.New("whoops")).Once()

		result, err := farmConditionService.Update(models.FarmConditionRequest{}, "-1")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}
