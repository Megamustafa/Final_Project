package services

import (
	"aquaculture/models"
	"aquaculture/repositories"
)

type FarmConditionService struct {
	Repository repositories.FarmConditionRepository
}

func InitFarmConditionService() FarmConditionService {
	return FarmConditionService{
		Repository: &repositories.FarmConditionRepositoryImpl{},
	}
}

func (fcs *FarmConditionService) GetAll() ([]models.FarmCondition, error) {
	return fcs.Repository.GetAll()
}

func (fcs *FarmConditionService) GetByID(id string) (models.FarmCondition, error) {
	return fcs.Repository.GetByID(id)
}

func (fcs *FarmConditionService) Create(fcReq models.FarmConditionRequest) (models.FarmCondition, error) {
	return fcs.Repository.Create(fcReq)
}

func (fcs *FarmConditionService) Update(fcReq models.FarmConditionRequest, id string) (models.FarmCondition, error) {
	return fcs.Repository.Update(fcReq, id)
}
