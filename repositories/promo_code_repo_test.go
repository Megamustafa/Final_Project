package repositories_test

import (
	"aquaculture/models"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllPromoCodes(t *testing.T) {
	t.Run("Get All PromoCodes | Valid", func(t *testing.T) {
		promoCodeRepository.On("GetAll").Return([]models.PromoCode{}, nil).Once()

		result, err := promoCodeService.GetAll()

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Get All PromoCodes | Invalid", func(t *testing.T) {
		promoCodeRepository.On("GetAll").Return([]models.PromoCode{}, errors.New("whoops")).Once()

		result, err := promoCodeService.GetAll()

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestGetPromoCodeByID(t *testing.T) {
	t.Run("Get PromoCode by ID | Valid", func(t *testing.T) {
		promoCodeRepository.On("GetByID", "1").Return(models.PromoCode{}, nil).Once()

		result, err := promoCodeService.GetByID("1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Get PromoCode by ID | Invalid", func(t *testing.T) {
		promoCodeRepository.On("GetByID", "-1").Return(models.PromoCode{}, errors.New("whoops")).Once()

		result, err := promoCodeService.GetByID("-1")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestCreatePromoCode(t *testing.T) {
	t.Run("Create PromoCode | Valid", func(t *testing.T) {
		promoCodeRepository.On("Create", models.PromoCodeRequest{}).Return(models.PromoCode{}, nil).Once()

		result, err := promoCodeService.Create(models.PromoCodeRequest{})

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Create PromoCode | Invalid", func(t *testing.T) {
		promoCodeRepository.On("Create", models.PromoCodeRequest{}).Return(models.PromoCode{}, errors.New("whoops")).Once()

		result, err := promoCodeService.Create(models.PromoCodeRequest{})

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestUpdatePromoCode(t *testing.T) {
	t.Run("Update PromoCode | Valid", func(t *testing.T) {
		promoCodeRepository.On("Update", models.PromoCodeRequest{}, "1").Return(models.PromoCode{}, nil).Once()

		result, err := promoCodeService.Update(models.PromoCodeRequest{}, "1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Update PromoCode | Invalid", func(t *testing.T) {
		promoCodeRepository.On("Update", models.PromoCodeRequest{}, "-1").Return(models.PromoCode{}, errors.New("whoops")).Once()

		result, err := promoCodeService.Update(models.PromoCodeRequest{}, "-1")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}
