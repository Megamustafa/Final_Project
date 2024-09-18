package services

import (
	"aquaculture/models"
	"aquaculture/repositories"
)

type PromoCodeService struct {
	Repository repositories.PromoCodeRepository
}

func InitPromoCodeService() PromoCodeService {
	return PromoCodeService{
		Repository: &repositories.PromoCodeRepositoryImpl{},
	}
}

func (pcs *PromoCodeService) GetAll() ([]models.PromoCode, error) {
	return pcs.Repository.GetAll()
}

func (pcs *PromoCodeService) GetByID(id string) (models.PromoCode, error) {
	return pcs.Repository.GetByID(id)
}

func (pcs *PromoCodeService) Create(pcReq models.PromoCodeRequest) (models.PromoCode, error) {
	return pcs.Repository.Create(pcReq)
}

func (pcs *PromoCodeService) Update(pcReq models.PromoCodeRequest, id string) (models.PromoCode, error) {
	return pcs.Repository.Update(pcReq, id)
}
