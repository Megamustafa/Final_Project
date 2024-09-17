package services

import (
	"aquaculture/models"
	"aquaculture/repositories"
)

type AquacultureFarmsService struct {
	Repository repositories.AquacultureFarmsRepository
}

func InitAquacultureFarmsService() AquacultureFarmsService {
	return AquacultureFarmsService{
		Repository: &repositories.AquacultureFarmsRepositoryImpl{},
	}
}

func (afs *AquacultureFarmsService) GetAll() ([]models.AquacultureFarms, error) {
	return afs.Repository.GetAll()
}

func (afs *AquacultureFarmsService) GetByID(id string) (models.AquacultureFarms, error) {
	return afs.Repository.GetByID(id)
}

func (afs *AquacultureFarmsService) Create(afReq models.AquacultureFarmsRequest) (models.AquacultureFarms, error) {
	return afs.Repository.Create(afReq)
}

func (afs *AquacultureFarmsService) Update(afReq models.AquacultureFarmsRequest, id string) (models.AquacultureFarms, error) {
	return afs.Repository.Update(afReq, id)
}

func (afs *AquacultureFarmsService) Delete(id string) error {
	return afs.Repository.Delete(id)
}
