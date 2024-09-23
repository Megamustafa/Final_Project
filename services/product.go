package services

import (
	"aquaculture/models"
	"aquaculture/repositories"
	"mime/multipart"
)

type ProductService struct {
	Repository repositories.ProductRepository
}

func InitProductService() ProductService {
	return ProductService{
		Repository: &repositories.ProductRepositoryImpl{},
	}
}

func (ps *ProductService) GetAll() ([]models.Product, error) {
	return ps.Repository.GetAll()
}

func (ps *ProductService) GetByID(id string) (models.Product, error) {
	return ps.Repository.GetByID(id)
}

func (ps *ProductService) Create(productReq models.ProductRequest) (models.Product, error) {
	return ps.Repository.Create(productReq)
}

func (ps *ProductService) Update(productReq models.ProductRequest, id string) (models.Product, error) {
	return ps.Repository.Update(productReq, id)
}

func (ps *ProductService) Delete(id string) error {
	return ps.Repository.Delete(id)
}

func (ps *ProductService) ImportFromCSV(file *multipart.FileHeader) ([]models.Product, error) {
	return ps.Repository.ImportFromCSV(file)
}
