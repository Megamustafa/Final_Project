package repositories

import (
	"aquaculture/database"
	"aquaculture/models"
)

type FarmConditionRepositoryImpl struct{}

func InitFarmConditionRepository() FarmConditionRepository {
	return &FarmConditionRepositoryImpl{}
}

func (fc *FarmConditionRepositoryImpl) GetAll() ([]models.FarmCondition, error) {
	var farmConditions []models.FarmCondition

	if err := database.DB.Preload("Farm").Find(&farmConditions).Error; err != nil {
		return []models.FarmCondition{}, err
	}
	return []models.FarmCondition{}, nil
}

func (fc *FarmConditionRepositoryImpl) GetByID(id string) (models.FarmCondition, error) {
	var farmCondition models.FarmCondition

	if err := database.DB.Preload("Farm").Find(&farmCondition).Error; err != nil {
		return models.FarmCondition{}, err
	}

	return farmCondition, nil
}

func (fc *FarmConditionRepositoryImpl) Create(fcReq models.FarmConditionRequest) (models.FarmCondition, error) {
	var farmCondition models.FarmCondition = models.FarmCondition{
		FarmID:      fcReq.FarmID,
		Temperature: fcReq.Temperature,
		PH:          fcReq.PH,
		OxygenLevel: fcReq.OxygenLevel,
	}

	result := database.DB.Create(&farmCondition)

	if err := result.Error; err != nil {
		return models.FarmCondition{}, err
	}

	if err := result.Preload("Farm").Last(&farmCondition).Error; err != nil {
		return models.FarmCondition{}, err
	}

	return farmCondition, nil
}

func (fc *FarmConditionRepositoryImpl) Update(fcReq models.FarmConditionRequest, id string) (models.FarmCondition, error) {
	farmCondition, err := fc.GetByID(id)

	if err != nil {
		return models.FarmCondition{}, err
	}

	farmCondition.FarmID = fcReq.FarmID
	farmCondition.Temperature = fcReq.Temperature
	farmCondition.PH = fcReq.PH
	farmCondition.OxygenLevel = fcReq.OxygenLevel

	if err := database.DB.Preload("Farm").Save(&farmCondition).Error; err != nil {
		return models.FarmCondition{}, err
	}

	return farmCondition, nil
}
