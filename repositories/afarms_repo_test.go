package repositories_test

import (
	"aquaculture/models"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllAquacultureFarms(t *testing.T) {
	t.Run("Get All Aquaculture Farms | Valid", func(t *testing.T) {
		aquacultureFarmsRepository.On("GetAll").Return([]models.AquacultureFarms{}, nil).Once()

		result, err := aquacultureFarmsService.GetAll()

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Get All Aquaculture Farms | Invalid", func(t *testing.T) {
		aquacultureFarmsRepository.On("GetAll").Return([]models.AquacultureFarms{}, errors.New("whoops")).Once()

		result, err := aquacultureFarmsService.GetAll()

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestGetAquacultureFarmsByID(t *testing.T) {
	t.Run("Get Aquaculture Farms by ID | Valid", func(t *testing.T) {
		aquacultureFarmsRepository.On("GetByID", "1").Return(models.AquacultureFarms{}, nil).Once()

		result, err := aquacultureFarmsService.GetByID("1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Get Aquaculture Farms by ID | Invalid", func(t *testing.T) {
		aquacultureFarmsRepository.On("GetByID", "-1").Return(models.AquacultureFarms{}, errors.New("whoops")).Once()

		result, err := aquacultureFarmsService.GetByID("-1")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestCreateAquacultureFarms(t *testing.T) {
	t.Run("Create Aquaculture Farms | Valid", func(t *testing.T) {
		aquacultureFarmsRepository.On("Create", models.AquacultureFarmsRequest{}).Return(models.AquacultureFarms{}, nil).Once()

		result, err := aquacultureFarmsService.Create(models.AquacultureFarmsRequest{})

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Create Aquaculture Farms | Invalid", func(t *testing.T) {
		aquacultureFarmsRepository.On("Create", models.AquacultureFarmsRequest{}).Return(models.AquacultureFarms{}, errors.New("whoops")).Once()

		result, err := aquacultureFarmsService.Create(models.AquacultureFarmsRequest{})

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestUpdateAquacultureFarms(t *testing.T) {
	t.Run("Update Aquaculture Farms | Valid", func(t *testing.T) {
		aquacultureFarmsRepository.On("Update", models.AquacultureFarmsRequest{}, "1").Return(models.AquacultureFarms{}, nil).Once()

		result, err := aquacultureFarmsService.Update(models.AquacultureFarmsRequest{}, "1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Update Aquaculture Farms | Invalid", func(t *testing.T) {
		aquacultureFarmsRepository.On("Update", models.AquacultureFarmsRequest{}, "-1").Return(models.AquacultureFarms{}, errors.New("whoops")).Once()

		result, err := aquacultureFarmsService.Update(models.AquacultureFarmsRequest{}, "-1")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestDeleteAquacultureFarms(t *testing.T) {
	t.Run("Delete Aquaculture Farms | Valid", func(t *testing.T) {
		aquacultureFarmsRepository.On("Delete", "1").Return(nil).Once()

		err := aquacultureFarmsService.Delete("1")

		assert.Nil(t, err)
	})

	t.Run("Delete Aquaculture Farms | Invalid", func(t *testing.T) {
		aquacultureFarmsRepository.On("Delete", "-1").Return(errors.New("whoops")).Once()

		err := aquacultureFarmsService.Delete("-1")

		assert.NotNil(t, err)
	})
}
