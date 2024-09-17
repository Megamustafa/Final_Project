package services

import (
	"aquaculture/models"
	"aquaculture/repositories"
)

type FarmService struct {
	Repository repositories.FarmRepository
}

func InitFarmService() FarmService {
	return FarmService{
		Repository: &repositories.FarmRepositoryImpl{},
	}
}

func (fs *FarmService) GetAll() ([]models.Farm, error) {
	return fs.Repository.GetAll()
}

func (fs *FarmService) GetByID(id string) (models.Farm, error) {
	return fs.Repository.GetByID(id)
}

func (fs *FarmService) Create(farmReq models.FarmRequest) (models.Farm, error) {
	return fs.Repository.Create(farmReq)
}

func (fs *FarmService) Update(farmReq models.FarmRequest, id string) (models.Farm, error) {
	return fs.Repository.Update(farmReq, id)
}

func (fs *FarmService) Delete(id string) error {
	return fs.Repository.Delete(id)
}
