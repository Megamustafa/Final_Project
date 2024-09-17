package repositories

import (
	"aquaculture/database"
	"aquaculture/models"
)

type AquacultureFarmsRepositoryImpl struct{}

func InitAquacultureFarmRepository() AquacultureFarmsRepository {
	return &AquacultureFarmsRepositoryImpl{}
}

func (afr *AquacultureFarmsRepositoryImpl) GetAll() ([]models.AquacultureFarms, error) {
	var aquaculturefarms []models.AquacultureFarms

	if err := database.DB.Preload("User").Preload("Farm").Find(&aquaculturefarms).Error; err != nil {
		return []models.AquacultureFarms{}, err
	}

	return aquaculturefarms, nil
}

func (afr *AquacultureFarmsRepositoryImpl) GetByID(id string) (models.AquacultureFarms, error) {
	var aquaculturefarm models.AquacultureFarms

	if err := database.DB.Preload("User").Preload("Farm").First(&aquaculturefarm, "id = ?", id).Error; err != nil {
		return models.AquacultureFarms{}, err
	}

	return aquaculturefarm, nil
}

func (afr *AquacultureFarmsRepositoryImpl) Create(afReqReq models.AquacultureFarmsRequest) (models.AquacultureFarms, error) {
	var aquaculturefarm models.AquacultureFarms = models.AquacultureFarms{
		UserID: afReqReq.UserID,
		FarmID: afReqReq.FarmID,
	}

	result := database.DB.Create(&aquaculturefarm)

	if err := result.Error; err != nil {
		return models.AquacultureFarms{}, err
	}

	if err := result.Preload("User").Preload("Farm").Last(&aquaculturefarm).Error; err != nil {
		return models.AquacultureFarms{}, err
	}

	return aquaculturefarm, nil
}

func (afr *AquacultureFarmsRepositoryImpl) Update(afReq models.AquacultureFarmsRequest, id string) (models.AquacultureFarms, error) {
	aquaculturefarm, err := afr.GetByID(id)

	if err != nil {
		return models.AquacultureFarms{}, err
	}

	aquaculturefarm.UserID = afReq.UserID
	aquaculturefarm.FarmID = afReq.FarmID

	if err := database.DB.Preload("User").Preload("Farm").Save(&aquaculturefarm).Error; err != nil {
		return models.AquacultureFarms{}, err
	}

	return aquaculturefarm, nil
}

func (afr *AquacultureFarmsRepositoryImpl) Delete(id string) error {
	aquaculturefarm, err := afr.GetByID(id)

	if err != nil {
		return err
	}

	if err := database.DB.Delete(&aquaculturefarm).Error; err != nil {
		return err
	}

	return nil
}
