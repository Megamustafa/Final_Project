package controllers

import (
	"aquaculture/middlewares"
	"aquaculture/models"
	"aquaculture/services"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductController struct {
	service     services.ProductService
	userService services.UserService
}

func InitProductController() ProductController {
	return ProductController{
		service:     services.InitProductService(),
		userService: services.InitUserService(models.JWTOptions{}),
	}
}

func verifyAdminP(c echo.Context, pc *ProductController) error {
	claim, err := middlewares.GetUser(c)
	if err != nil {
		return err
	}
	admin, err := pc.userService.GetAdminInfo(strconv.Itoa(claim.ID))
	if err != nil && admin.ID == 0 {
		return err
	}
	return nil
}

func (pc *ProductController) GetAll(c echo.Context) error {
	products, err := pc.service.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "error when fetching products",
		})
	}

	return c.JSON(http.StatusOK, models.Response[[]models.Product]{
		Status:  "success",
		Message: "all products",
		Data:    products,
	})
}

func (pc *ProductController) GetByID(c echo.Context) error {
	productID := c.Param("id")

	product, err := pc.service.GetByID(productID)

	if err != nil {
		return c.JSON(http.StatusNotFound, models.Response[string]{
			Status:  "failed",
			Message: "product not found",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.Product]{
		Status:  "success",
		Message: "product found",
		Data:    product,
	})
}

func (pc *ProductController) Create(c echo.Context) error {
	var productReq models.ProductRequest

	if err := verifyAdminP(c, pc); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	if err := c.Bind(&productReq); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	err := productReq.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "please fill all the required fields",
		})
	}

	product, err := pc.service.Create(productReq)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to create a product",
		})
	}

	return c.JSON(http.StatusCreated, models.Response[models.Product]{
		Status:  "success",
		Message: "product created",
		Data:    product,
	})
}

func (pc *ProductController) Update(c echo.Context) error {
	var productReq models.ProductRequest

	if err := verifyAdminP(c, pc); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}
	if err := c.Bind(&productReq); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	productID := c.Param("id")

	err := productReq.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "please fill all the required fields",
		})
	}

	product, err := pc.service.Update(productReq, productID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to update a product",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.Product]{
		Status:  "success",
		Message: "product updated",
		Data:    product,
	})
}

func (pc *ProductController) Delete(c echo.Context) error {
	productID := c.Param("id")

	if err := verifyAdminP(c, pc); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	err := pc.service.Delete(productID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to delete a product",
		})
	}

	return c.JSON(http.StatusOK, models.Response[string]{
		Status:  "success",
		Message: "product deleted",
	})
}

func (pc *ProductController) ImportFromCSV(c echo.Context) error {

	if err := verifyAdminP(c, pc); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}
	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}
	defer src.Close()

	dst, err := os.Create(file.Filename)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	products, err := pc.service.ImportFromCSV(file)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to import products",
		})
	}

	return c.JSON(http.StatusCreated, models.Response[[]models.Product]{
		Status:  "success",
		Message: "products imported",
		Data:    products,
	})

}
