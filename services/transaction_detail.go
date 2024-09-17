package services

import (
	"aquaculture/models"
	"aquaculture/repositories"
)

type TransactionDetailService struct {
	Repository repositories.TransactionDetailRepository
}

func InitTransactionDetailService() TransactionDetailService {
	return TransactionDetailService{
		Repository: repositories.InitTransactionDetailRepository(),
	}
}

func (tds *TransactionDetailService) GetAll() ([]models.TransactionDetail, error) {
	return tds.Repository.GetAll()
}

func (tds *TransactionDetailService) GetByID(id string) (models.TransactionDetail, error) {
	return tds.Repository.GetByID(id)
}

func (tds *TransactionDetailService) Create(tdReq models.TransactionDetailRequest) (models.TransactionDetail, error) {
	return tds.Repository.Create(tdReq)
}

func (tds *TransactionDetailService) Update(tdReq models.TransactionDetailRequest, id string) (models.TransactionDetail, error) {
	return tds.Repository.Update(tdReq, id)
}
