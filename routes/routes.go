package routes

import (
	"aquaculture/controllers"
	"aquaculture/middlewares"
	"aquaculture/models"
	"aquaculture/utils"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	loggerConfig := middlewares.LoggerConfig{
		Format: "[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}" + "\n",
	}

	loggerMiddleware := loggerConfig.Init()
	e.Use(loggerMiddleware)

	jwtConfig := middlewares.JWTConfig{
		SecretKey: utils.GetConfig("JWT_SECRET_KEY"),
	}

	authMiddlewareConfig := jwtConfig.Init()

	jwtOptions := models.JWTOptions{
		SecretKey:       utils.GetConfig("JWT_SECRET_KEY"),
		ExpiresDuration: 1,
	}

	userController := controllers.InitUserController(jwtOptions)
	userRoutes := e.Group("/api/v1/users")
	adminRoutes := e.Group("/api/v1/admins")

	userRoutes.POST("/register", userController.Register)
	userRoutes.POST("/login", userController.LoginUser)
	userRoutes.GET(
		"/user",
		userController.GetUserInfo,
		echojwt.WithConfig(authMiddlewareConfig),
		middlewares.VerifyToken,
	)

	adminRoutes.POST("/login", userController.LoginAdmin)
	adminRoutes.GET(
		"/admin",
		userController.GetAdminInfo,
		echojwt.WithConfig(authMiddlewareConfig),
	)

	productController := controllers.InitProductController()
	productRoutes := e.Group("/api/v1/products", echojwt.WithConfig(authMiddlewareConfig))
	productRoutes.GET("", productController.GetAll)
	productRoutes.GET("/:id", productController.GetByID)
	productRoutes.POST("", productController.Create)
	productRoutes.PUT("/:id", productController.Update)
	productRoutes.DELETE("/:id", productController.Delete)
	productRoutes.POST("/import", productController.ImportFromCSV)

	productTypeController := controllers.InitProductTypeController()
	productTypeRoutes := e.Group("/api/v1/products/types", echojwt.WithConfig(authMiddlewareConfig))
	productTypeRoutes.GET("", productTypeController.GetAll)
	productTypeRoutes.GET("/:id", productTypeController.GetByID)
	productTypeRoutes.POST("", productTypeController.Create)
	productTypeRoutes.PUT("/:id", productTypeController.Update)
	productTypeRoutes.DELETE("/:id", productTypeController.Delete)

	aquacultureFarmsController := controllers.InitAquacultureFarmsController()
	aquacultureFarmsRoutes := e.Group("/api/v1/aquafarms", echojwt.WithConfig(authMiddlewareConfig))
	aquacultureFarmsRoutes.GET("", aquacultureFarmsController.GetAll)
	aquacultureFarmsRoutes.GET("/:id", aquacultureFarmsController.GetByID)
	aquacultureFarmsRoutes.POST("", aquacultureFarmsController.Create)
	aquacultureFarmsRoutes.PUT("/:id", aquacultureFarmsController.Update)
	aquacultureFarmsRoutes.DELETE("/:id", aquacultureFarmsController.Delete)

	farmController := controllers.InitFarmController()
	farmRoutes := e.Group("/api/v1/farms", echojwt.WithConfig(authMiddlewareConfig))
	farmRoutes.GET("", farmController.GetAll)
	farmRoutes.GET("/:id", farmController.GetByID)
	farmRoutes.POST("", farmController.Create)
	farmRoutes.PUT("/:id", farmController.Update)
	farmRoutes.DELETE("/:id", farmController.Delete)

	farmTypeController := controllers.InitFarmTypeController()
	farmTypeRoutes := e.Group("/api/v1/farms/types", echojwt.WithConfig(authMiddlewareConfig))
	farmTypeRoutes.GET("", farmTypeController.GetAll)
	farmTypeRoutes.GET("/:id", farmTypeController.GetByID)
	farmTypeRoutes.POST("", farmTypeController.Create)
	farmTypeRoutes.PUT("/:id", farmTypeController.Update)
	farmTypeRoutes.DELETE("/:id", farmTypeController.Delete)

	transactionController := controllers.InitTransactionController()
	transactionRoutes := e.Group("/api/v1/transactions", echojwt.WithConfig(authMiddlewareConfig))
	transactionRoutes.GET("", transactionController.GetAll)
	transactionRoutes.GET("/:id", transactionController.GetByID)
	transactionRoutes.POST("", transactionController.Create)
	transactionRoutes.PUT("/:id", transactionController.Update)

	transactionDetailController := controllers.InitTransactionDetailController()
	transactionDetailRoutes := e.Group("/api/v1/transactions/details", echojwt.WithConfig(authMiddlewareConfig))
	transactionDetailRoutes.GET("", transactionDetailController.GetAll)
	transactionDetailRoutes.GET("/:id", transactionDetailController.GetByID)
	transactionDetailRoutes.POST("", transactionDetailController.Create)
	transactionDetailRoutes.PUT("/:id", transactionDetailController.Update)

	articleController := controllers.InitArticleController()
	articleRoutes := e.Group("/api/v1/articles", echojwt.WithConfig(authMiddlewareConfig))
	articleRoutes.GET("", articleController.GetAll)
	articleRoutes.GET("/:id", articleController.GetByID)
	articleRoutes.POST("", articleController.Create)
	articleRoutes.PUT("/:id", articleController.Update)
	articleRoutes.DELETE("/:id", articleController.Delete)

	farmConditionController := controllers.InitFarmConditionController()
	farmConditionRoutes := e.Group("/api/v1/farms/conditions", echojwt.WithConfig(authMiddlewareConfig))
	farmConditionRoutes.GET("", farmConditionController.GetAll)
	farmConditionRoutes.GET("/:id", farmConditionController.GetByID)
	farmConditionRoutes.POST("", farmConditionController.Create)
	farmConditionRoutes.PUT("/:id", farmConditionController.Update)

	promoCodeController := controllers.InitPromoCodeController()
	promoCodeRoutes := e.Group("/api/v1/transactions/promo", echojwt.WithConfig(authMiddlewareConfig))
	promoCodeRoutes.GET("", promoCodeController.GetAll)
	promoCodeRoutes.GET("/:id", promoCodeController.GetByID)
	promoCodeRoutes.POST("", promoCodeController.Create)
	promoCodeRoutes.PUT("/:id", promoCodeController.Update)

}
