package repositories

import (
	"aquaculture/database"
	"aquaculture/models"
	"time"

	"gorm.io/datatypes"
)

type PromoCodeRepositoryImpl struct{}

func InitPromoCodeRepository() PromoCodeRepository {
	return &PromoCodeRepositoryImpl{}
}

func (pc *PromoCodeRepositoryImpl) GetAll() ([]models.PromoCode, error) {
	var promoCodes []models.PromoCode

	if err := database.DB.Find(&promoCodes).Error; err != nil {
		return []models.PromoCode{}, err
	}

	return promoCodes, nil
}

func (pc *PromoCodeRepositoryImpl) GetByID(id string) (models.PromoCode, error) {
	var promoCode models.PromoCode

	if err := database.DB.First(&promoCode, "id = ?", id).Error; err != nil {
		return models.PromoCode{}, err
	}

	return promoCode, nil
}

func (pc *PromoCodeRepositoryImpl) Create(pcReq models.PromoCodeRequest) (models.PromoCode, error) {
	validFrom, _ := time.Parse("02-01-2006", pcReq.ValidFrom)
	validUntil, _ := time.Parse("02-01-2006", pcReq.ValidUntil)

	var promoCode models.PromoCode = models.PromoCode{
		DiscountPercentage: pcReq.DiscountPercentage,
		ValidFrom:          datatypes.Date(validFrom),
		ValidUntil:         datatypes.Date(validUntil),
		Status:             pcReq.Status,
	}

	result := database.DB.Create(&promoCode)

	if err := result.Error; err != nil {
		return models.PromoCode{}, err
	}

	if err := result.Last(&promoCode).Error; err != nil {
		return models.PromoCode{}, err
	}

	return promoCode, nil
}

func (pc *PromoCodeRepositoryImpl) Update(pcReq models.PromoCodeRequest, id string) (models.PromoCode, error) {
	promoCode, err := pc.GetByID(id)

	if err != nil {
		return models.PromoCode{}, err
	}

	validFrom, _ := time.Parse("02-01-2006", pcReq.ValidFrom)
	validUntil, _ := time.Parse("02-01-2006", pcReq.ValidUntil)

	promoCode.DiscountPercentage = pcReq.DiscountPercentage
	promoCode.ValidFrom = datatypes.Date(validFrom)
	promoCode.ValidUntil = datatypes.Date(validUntil)
	promoCode.Status = pcReq.Status

	if err := database.DB.Save(&promoCode).Error; err != nil {
		return models.PromoCode{}, err
	}

	return promoCode, nil
}
