package repositories

import (
	"aquaculture/database"
	"aquaculture/models"
)

type TransactionRepositoryImpl struct{}

func InitTransactionRepository() TransactionRepository {
	return &TransactionRepositoryImpl{}
}

func (tr *TransactionRepositoryImpl) GetAll() ([]models.Transaction, error) {
	var transactions []models.Transaction

	if err := database.DB.Preload("TransactionDetails").Find(&transactions).Error; err != nil {
		return []models.Transaction{}, err
	}

	return transactions, nil
}

func (tr *TransactionRepositoryImpl) GetByID(id string) (models.Transaction, error) {
	var transaction models.Transaction

	if err := database.DB.Preload("TransactionDetails").First(&transaction, "id = ?", id).Error; err != nil {
		return models.Transaction{}, err
	}

	return transaction, nil
}

func (tr *TransactionRepositoryImpl) Create(tReq models.TransactionRequest) (models.Transaction, error) {
	var transaction models.Transaction = models.Transaction{
		UserID:        tReq.UserID,
		TotalAmount:   0,
		Status:        tReq.Status,
		PaymentMethod: tReq.PaymentMethod,
	}

	if tReq.PromoCodeID != 0 {
		transaction.PromoCodeID = tReq.PromoCodeID
	}
	result := database.DB.Create(&transaction)

	if err := result.Error; err != nil {
		return models.Transaction{}, err
	}

	if err := result.Preload("TransactionDetails").Last(&transaction).Error; err != nil {
		return models.Transaction{}, err
	}

	return transaction, nil
}

func (tr *TransactionRepositoryImpl) Update(tReq models.TransactionRequest, id string) (models.Transaction, error) {
	transaction, err := tr.GetByID(id)

	if err != nil {
		return models.Transaction{}, err
	}

	transaction.UserID = tReq.UserID
	transaction.TotalAmount = float64(tReq.TotalAmount)
	transaction.Status = tReq.Status
	transaction.PaymentMethod = tReq.PaymentMethod

	if err := database.DB.Preload("TransactionDetails").Save(&transaction).Error; err != nil {
		return models.Transaction{}, err
	}

	return transaction, nil
}
