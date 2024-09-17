package repositories_test

import (
	"aquaculture/models"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllTransactionDetails(t *testing.T) {
	t.Run("Get All TransactionDetails | Valid", func(t *testing.T) {
		transactionDetailRepository.On("GetAll").Return([]models.TransactionDetail{}, nil).Once()

		result, err := transactionDetailService.GetAll()

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Get All TransactionDetails | Invalid", func(t *testing.T) {
		transactionDetailRepository.On("GetAll").Return([]models.TransactionDetail{}, errors.New("whoops")).Once()

		result, err := transactionDetailService.GetAll()

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestGetTransactionDetailByID(t *testing.T) {
	t.Run("Get TransactionDetail by ID | Valid", func(t *testing.T) {
		transactionDetailRepository.On("GetByID", "1").Return(models.TransactionDetail{}, nil).Once()

		result, err := transactionDetailService.GetByID("1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Get TransactionDetail by ID | Invalid", func(t *testing.T) {
		transactionDetailRepository.On("GetByID", "-1").Return(models.TransactionDetail{}, errors.New("whoops")).Once()

		result, err := transactionDetailService.GetByID("-1")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestCreateTransactionDetail(t *testing.T) {
	t.Run("Create TransactionDetail | Valid", func(t *testing.T) {
		transactionDetailRepository.On("Create", models.TransactionDetailRequest{}).Return(models.TransactionDetail{}, nil).Once()

		result, err := transactionDetailService.Create(models.TransactionDetailRequest{})

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Create TransactionDetail | Invalid", func(t *testing.T) {
		transactionDetailRepository.On("Create", models.TransactionDetailRequest{}).Return(models.TransactionDetail{}, errors.New("whoops")).Once()

		result, err := transactionDetailService.Create(models.TransactionDetailRequest{})

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestUpdateTransactionDetail(t *testing.T) {
	t.Run("Update TransactionDetail | Valid", func(t *testing.T) {
		transactionDetailRepository.On("Update", models.TransactionDetailRequest{}, "1").Return(models.TransactionDetail{}, nil).Once()

		result, err := transactionDetailService.Update(models.TransactionDetailRequest{}, "1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Update TransactionDetail | Invalid", func(t *testing.T) {
		transactionDetailRepository.On("Update", models.TransactionDetailRequest{}, "-1").Return(models.TransactionDetail{}, errors.New("whoops")).Once()

		result, err := transactionDetailService.Update(models.TransactionDetailRequest{}, "-1")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}
