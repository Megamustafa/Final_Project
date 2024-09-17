package repositories_test

import (
	"aquaculture/models"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllTransactions(t *testing.T) {
	t.Run("Get All Transactions | Valid", func(t *testing.T) {
		transactionRepository.On("GetAll").Return([]models.Transaction{}, nil).Once()

		result, err := transactionService.GetAll()

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Get All Transactions | Invalid", func(t *testing.T) {
		transactionRepository.On("GetAll").Return([]models.Transaction{}, errors.New("whoops")).Once()

		result, err := transactionService.GetAll()

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestGetTransactionByID(t *testing.T) {
	t.Run("Get Transaction by ID | Valid", func(t *testing.T) {
		transactionRepository.On("GetByID", "1").Return(models.Transaction{}, nil).Once()

		result, err := transactionService.GetByID("1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Get Transaction by ID | Invalid", func(t *testing.T) {
		transactionRepository.On("GetByID", "-1").Return(models.Transaction{}, errors.New("whoops")).Once()

		result, err := transactionService.GetByID("-1")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestCreateTransaction(t *testing.T) {
	t.Run("Create Transaction | Valid", func(t *testing.T) {
		transactionRepository.On("Create", models.TransactionRequest{}).Return(models.Transaction{}, nil).Once()

		result, err := transactionService.Create(models.TransactionRequest{})

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Create Transaction | Invalid", func(t *testing.T) {
		transactionRepository.On("Create", models.TransactionRequest{}).Return(models.Transaction{}, errors.New("whoops")).Once()

		result, err := transactionService.Create(models.TransactionRequest{})

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestUpdateTransaction(t *testing.T) {
	t.Run("Update Transaction | Valid", func(t *testing.T) {
		transactionRepository.On("Update", models.TransactionRequest{}, "1").Return(models.Transaction{}, nil).Once()

		result, err := transactionService.Update(models.TransactionRequest{}, "1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Update Transaction | Invalid", func(t *testing.T) {
		transactionRepository.On("Update", models.TransactionRequest{}, "-1").Return(models.Transaction{}, errors.New("whoops")).Once()

		result, err := transactionService.Update(models.TransactionRequest{}, "-1")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}
