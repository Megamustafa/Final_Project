package repositories

import (
	"aquaculture/database"
	"aquaculture/models"
)

type FarmTypeRepositoryImpl struct{}

func InitFarmTypeRepository() FarmTypeRepository {
	return &FarmTypeRepositoryImpl{}
}

func (ftr *FarmTypeRepositoryImpl) GetAll() ([]models.FarmType, error) {
	var farmTypes []models.FarmType

	if err := database.DB.Find(&farmTypes).Error; err != nil {
		return []models.FarmType{}, err
	}

	return farmTypes, nil
}

func (ftr *FarmTypeRepositoryImpl) GetByID(id string) (models.FarmType, error) {
	var farmType models.FarmType

	if err := database.DB.First(&farmType, "id = ?", id).Error; err != nil {
		return models.FarmType{}, err
	}

	return farmType, nil
}

func (ftr *FarmTypeRepositoryImpl) Create(ftReq models.FarmTypeRequest) (models.FarmType, error) {
	var farmType models.FarmType = models.FarmType{
		Name: ftReq.Name,
	}

	result := database.DB.Create(&farmType)

	if err := result.Error; err != nil {
		return models.FarmType{}, err
	}

	if err := result.Last(&farmType).Error; err != nil {
		return models.FarmType{}, err
	}

	return farmType, nil
}

func (ftr *FarmTypeRepositoryImpl) Update(ftReq models.FarmTypeRequest, id string) (models.FarmType, error) {
	farmType, err := ftr.GetByID(id)

	if err != nil {
		return models.FarmType{}, err
	}

	farmType.Name = ftReq.Name

	if err := database.DB.Save(&farmType).Error; err != nil {
		return models.FarmType{}, err
	}

	return farmType, nil
}

func (ftr *FarmTypeRepositoryImpl) Delete(id string) error {
	farmType, err := ftr.GetByID(id)

	if err != nil {
		return err
	}

	if err := database.DB.Delete(&farmType).Error; err != nil {
		return err
	}

	return nil
}
