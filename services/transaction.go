package services

import (
	"aquaculture/models"
	"aquaculture/repositories"
)

type TransactionService struct {
	Repository repositories.TransactionRepository
}

func InitTransactionService() TransactionService {
	return TransactionService{
		Repository: &repositories.TransactionRepositoryImpl{},
	}
}

func (ts *TransactionService) GetAll() ([]models.Transaction, error) {
	return ts.Repository.GetAll()
}

func (ts *TransactionService) GetByID(id string) (models.Transaction, error) {
	return ts.Repository.GetByID(id)
}

func (ts *TransactionService) Create(tReq models.TransactionRequest) (models.Transaction, error) {
	return ts.Repository.Create(tReq)
}

func (ts *TransactionService) Update(tReq models.TransactionStatusRequest, id string) (models.Transaction, error) {
	return ts.Repository.Update(tReq, id)
}
