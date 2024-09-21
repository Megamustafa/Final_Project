package repositories

import (
	"aquaculture/database"
	"aquaculture/models"
	"encoding/csv"
	"os"
	"strconv"
)

type ProductRepositoryImpl struct{}

func InitProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (pr *ProductRepositoryImpl) GetAll() ([]models.Product, error) {
	var products []models.Product

	if err := database.DB.Preload("ProductType").Find(&products).Error; err != nil {
		return []models.Product{}, err
	}

	return products, nil
}

func (pr *ProductRepositoryImpl) GetByID(id string) (models.Product, error) {
	var product models.Product

	if err := database.DB.Preload("ProductType").First(&product, "id = ?", id).Error; err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (pr *ProductRepositoryImpl) Create(productReq models.ProductRequest) (models.Product, error) {
	var product models.Product = models.Product{
		ProductTypeID: productReq.ProductTypeID,
		Description:   productReq.Description,
		Price:         productReq.Price,
	}

	result := database.DB.Create(&product)

	if err := result.Error; err != nil {
		return models.Product{}, err
	}

	if err := result.Preload("ProductType").Last(&product).Error; err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (pr *ProductRepositoryImpl) Update(productReq models.ProductRequest, id string) (models.Product, error) {
	product, err := pr.GetByID(id)

	if err != nil {
		return models.Product{}, err
	}

	product.ProductTypeID = productReq.ProductTypeID
	product.Description = productReq.Description
	product.Price = productReq.Price

	if err := database.DB.Preload("ProductType").Save(&product).Error; err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (pr *ProductRepositoryImpl) Delete(id string) error {
	product, err := pr.GetByID(id)

	if err != nil {
		return err
	}

	if err := database.DB.Delete(&product).Error; err != nil {
		return err
	}

	return nil
}

func (pr *ProductRepositoryImpl) ImportFromCSV(filename string) ([]models.Product, error) {

	csvFile, err := os.Open(filename)

	if err != nil {
		return []models.Product{}, err
	}

	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	records, err := reader.ReadAll()

	if err != nil {
		return []models.Product{}, err
	}

	var products []models.Product = []models.Product{}

	for idx, eachrecord := range records {

		if idx == 0 {
			continue
		}
		price, _ := strconv.Atoi(eachrecord[1])

		products = append(products, models.Product{Description: eachrecord[0], Price: price})
	}

	if err := database.DB.Create(&products).Error; err != nil {
		return []models.Product{}, err
	}

	return products, nil

}
