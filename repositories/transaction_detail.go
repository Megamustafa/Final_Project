package repositories

import (
	"aquaculture/database"
	"aquaculture/models"
	"strconv"
)

type TransactionDetailRepositoryImpl struct {
	productRepo     ProductRepository
	transactionRepo TransactionRepository
}

func InitTransactionDetailRepository() TransactionDetailRepository {
	return &TransactionDetailRepositoryImpl{
		productRepo:     InitProductRepository(),
		transactionRepo: InitTransactionRepository(),
	}
}

func (tdr *TransactionDetailRepositoryImpl) GetAll() ([]models.TransactionDetail, error) {
	var transactionDetails []models.TransactionDetail

	if err := database.DB.Find(&transactionDetails).Error; err != nil {
		return []models.TransactionDetail{}, err
	}

	return transactionDetails, nil
}

func (tdr *TransactionDetailRepositoryImpl) GetByID(id string) (models.TransactionDetail, error) {
	var transactionDetail models.TransactionDetail

	if err := database.DB.First(&transactionDetail, "id = ?", id).Error; err != nil {
		return models.TransactionDetail{}, err
	}

	return transactionDetail, nil
}

func (tdr *TransactionDetailRepositoryImpl) Create(tdReq models.TransactionDetailRequest) (models.TransactionDetail, error) {
	//find product by id
	//amount=product price* quantity
	//product repo
	product, err := tdr.productRepo.GetByID(strconv.Itoa(int(tdReq.ProductID)))

	if err != nil {
		return models.TransactionDetail{}, err
	}

	amount := float64(tdReq.Quantity) * float64(product.Price)
	var transactionDetail models.TransactionDetail = models.TransactionDetail{
		TransactionID: tdReq.TransactionID,
		ProductID:     tdReq.ProductID,
		Quantity:      tdReq.Quantity,
		Amount:        amount,
	}

	result := database.DB.Create(&transactionDetail)

	if err := result.Error; err != nil {
		return models.TransactionDetail{}, err
	}

	if err := result.Last(&transactionDetail).Error; err != nil {
		return models.TransactionDetail{}, err
	}

	transactionData, err := tdr.transactionRepo.GetByID(strconv.Itoa(int(tdReq.TransactionID)))

	if err != nil {
		return models.TransactionDetail{}, err
	}

	currentDetails := transactionData.TransactionDetails

	var totalAmounttmp float64

	for _, detail := range currentDetails {
		totalAmounttmp += detail.Amount
	}

	transactionData.TotalAmount = totalAmounttmp
	if transactionData.PromoCodeID != 0 {
		discountedTotalAmount := totalAmounttmp - (totalAmounttmp * transactionData.PromoCode.DiscountPercentage)

		transactionData.TotalAmount = discountedTotalAmount
	}
	if err := database.DB.Save(&transactionData).Error; err != nil {
		return models.TransactionDetail{}, err
	}

	return transactionDetail, nil
}

func (tdr *TransactionDetailRepositoryImpl) Update(tdReq models.TransactionDetailRequest, id string) (models.TransactionDetail, error) {
	transactionDetail, err := tdr.GetByID(id)

	if err != nil {
		return models.TransactionDetail{}, err
	}

	product, err := tdr.productRepo.GetByID(strconv.Itoa(int(tdReq.ProductID)))

	if err != nil {
		return models.TransactionDetail{}, err
	}

	amount := float64(tdReq.Quantity) * float64(product.Price)

	transactionDetail.ProductID = tdReq.ProductID
	transactionDetail.Quantity = tdReq.Quantity
	transactionDetail.Amount = amount

	if err := database.DB.Save(&transactionDetail).Error; err != nil {
		return models.TransactionDetail{}, err
	}

	transactionData, err := tdr.transactionRepo.GetByID(strconv.Itoa(int(tdReq.TransactionID)))

	if err != nil {
		return models.TransactionDetail{}, err
	}

	currentDetails := transactionData.TransactionDetails
	var totalAmounttmp float64

	for _, detail := range currentDetails {
		totalAmounttmp += detail.Amount
	}

	transactionData.TotalAmount = totalAmounttmp

	if transactionData.PromoCodeID != 0 {
		discountedTotalAmount := totalAmounttmp - (totalAmounttmp * transactionData.PromoCode.DiscountPercentage)

		transactionData.TotalAmount = discountedTotalAmount
	}

	if err := database.DB.Save(&transactionData).Error; err != nil {
		return models.TransactionDetail{}, err
	}

	return transactionDetail, nil
}
