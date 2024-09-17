package repositories

import (
	"aquaculture/database"
	"aquaculture/models"
)

type FarmRepositoryImpl struct{}

func InitFarmRepository() FarmRepository {
	return &FarmRepositoryImpl{}
}

func (fr *FarmRepositoryImpl) GetAll() ([]models.Farm, error) {
	var farms []models.Farm

	if err := database.DB.Preload("FarmType").Find(&farms).Error; err != nil {
		return []models.Farm{}, err
	}

	return farms, nil
}

func (fr *FarmRepositoryImpl) GetByID(id string) (models.Farm, error) {
	var farm models.Farm

	if err := database.DB.Preload("FarmType").First(&farm, "id = ?", id).Error; err != nil {
		return models.Farm{}, err
	}

	return farm, nil
}

func (fr *FarmRepositoryImpl) Create(farmReq models.FarmRequest) (models.Farm, error) {
	var farm models.Farm = models.Farm{
		FarmTypeID:  farmReq.FarmTypeID,
		Description: farmReq.Description,
		Price:       farmReq.Price,
	}

	result := database.DB.Create(&farm)

	if err := result.Error; err != nil {
		return models.Farm{}, err
	}

	if err := result.Preload("FarmType").Last(&farm).Error; err != nil {
		return models.Farm{}, err
	}

	return farm, nil
}

func (fr *FarmRepositoryImpl) Update(farmReq models.FarmRequest, id string) (models.Farm, error) {
	farm, err := fr.GetByID(id)

	if err != nil {
		return models.Farm{}, err
	}

	farm.FarmTypeID = farmReq.FarmTypeID
	farm.Description = farmReq.Description
	farm.Price = farmReq.Price

	if err := database.DB.Preload("FarmType").Save(&farm).Error; err != nil {
		return models.Farm{}, err
	}

	return farm, nil
}

func (fr *FarmRepositoryImpl) Delete(id string) error {
	farm, err := fr.GetByID(id)

	if err != nil {
		return err
	}

	if err := database.DB.Delete(&farm).Error; err != nil {
		return err
	}

	return nil
}
