package services

import (
	"aquaculture/models"
	"aquaculture/repositories"
)

type ProductTypeService struct {
	Repository repositories.ProductTypeRepository
}

func InitProductTypeService() ProductTypeService {
	return ProductTypeService{
		Repository: &repositories.ProductTypeRepositoryImpl{},
	}
}

func (pts *ProductTypeService) GetAll() ([]models.ProductType, error) {
	return pts.Repository.GetAll()
}

func (pts *ProductTypeService) GetByID(id string) (models.ProductType, error) {
	return pts.Repository.GetByID(id)
}

func (pts *ProductTypeService) Create(ptReq models.ProductTypeRequest) (models.ProductType, error) {
	return pts.Repository.Create(ptReq)
}

func (pts *ProductTypeService) Update(ptReq models.ProductTypeRequest, id string) (models.ProductType, error) {
	return pts.Repository.Update(ptReq, id)
}

func (pts *ProductTypeService) Delete(id string) error {
	return pts.Repository.Delete(id)
}
