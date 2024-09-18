package repositories

import (
	"aquaculture/database"
	"aquaculture/models"
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
	var promoCode models.PromoCode = models.PromoCode{
		DiscountPercentage: pcReq.DiscountPercentage,
		ValidFrom:          pcReq.ValidFrom,
		ValidUntil:         pcReq.ValidUntil,
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

	promoCode.DiscountPercentage = pcReq.DiscountPercentage
	promoCode.ValidFrom = pcReq.ValidFrom
	promoCode.ValidUntil = pcReq.ValidUntil
	promoCode.Status = pcReq.Status

	if err := database.DB.Save(&promoCode).Error; err != nil {
		return models.PromoCode{}, err
	}

	return promoCode, nil
}
