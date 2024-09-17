package repositories_test

import (
	"aquaculture/models"
	"aquaculture/repositories/mocks"
	"aquaculture/services"
	"os"
	"testing"
)

var (
	userRepository mocks.UserRepository
	userService    services.UserService

	aquacultureFarmsRepository mocks.AquacultureFarmsRepository
	aquacultureFarmsService    services.AquacultureFarmsService

	articleRepository mocks.ArticleRepository
	articleService    services.ArticleService

	farmTypeRepository mocks.FarmTypeRepository
	farmTypeService    services.FarmTypeService

	farmRepository mocks.FarmRepository
	farmService    services.FarmService

	productTypeRepository mocks.ProductTypeRepository
	productTypeService    services.ProductTypeService

	productRepository mocks.ProductRepository
	productService    services.ProductService

	transactionDetailRepository mocks.TransactionDetailRepository
	transactionDetailService    services.TransactionDetailService

	transactionRepository mocks.TransactionRepository
	transactionService    services.TransactionService
)

func TestMain(m *testing.M) {

	userService = services.InitUserService(models.JWTOptions{})
	userService.Repository = &userRepository

	aquacultureFarmsService = services.InitAquacultureFarmsService()
	aquacultureFarmsService.Repository = &aquacultureFarmsRepository

	articleService = services.InitArticleService()
	articleService.Repository = &articleRepository

	farmTypeService = services.InitFarmTypeService()
	farmTypeService.Repository = &farmTypeRepository

	farmService = services.InitFarmService()
	farmService.Repository = &farmRepository

	productTypeService = services.InitProductTypeService()
	productTypeService.Repository = &productTypeRepository

	productService = services.InitProductService()
	productService.Repository = &productRepository

	transactionDetailService = services.InitTransactionDetailService()
	transactionDetailService.Repository = &transactionDetailRepository

	transactionService = services.InitTransactionService()
	transactionService.Repository = &transactionRepository

	os.Exit(m.Run())
}
