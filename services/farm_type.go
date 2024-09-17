package services

import (
	"aquaculture/models"
	"aquaculture/repositories"
)

type FarmTypeService struct {
	Repository repositories.FarmTypeRepository
}

func InitFarmTypeService() FarmTypeService {
	return FarmTypeService{
		Repository: &repositories.FarmTypeRepositoryImpl{},
	}
}

func (fts *FarmTypeService) GetAll() ([]models.FarmType, error) {
	return fts.Repository.GetAll()
}

func (fts *FarmTypeService) GetByID(id string) (models.FarmType, error) {
	return fts.Repository.GetByID(id)
}

func (fts *FarmTypeService) Create(ftReq models.FarmTypeRequest) (models.FarmType, error) {
	return fts.Repository.Create(ftReq)
}

func (fts *FarmTypeService) Update(ftReq models.FarmTypeRequest, id string) (models.FarmType, error) {
	return fts.Repository.Update(ftReq, id)
}

func (fts *FarmTypeService) Delete(id string) error {
	return fts.Repository.Delete(id)
}
