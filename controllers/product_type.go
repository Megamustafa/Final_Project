package controllers

import (
	"aquaculture/middlewares"
	"aquaculture/models"
	"aquaculture/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductTypeController struct {
	service     services.ProductTypeService
	userService services.UserService
}

func InitProductTypeController() ProductTypeController {
	return ProductTypeController{
		service:     services.InitProductTypeService(),
		userService: services.InitUserService(models.JWTOptions{}),
	}
}

func verifyAdminPT(c echo.Context, ptc *ProductTypeController) error {
	claim, err := middlewares.GetUser(c)
	if err != nil {
		return err
	}
	admin, err := ptc.userService.GetAdminInfo(strconv.Itoa(claim.ID))
	if err != nil && admin.ID == 0 {
		return err
	}
	return nil
}

func (ptc *ProductTypeController) GetAll(c echo.Context) error {
	productTypes, err := ptc.service.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "error when fetching product types",
		})
	}

	return c.JSON(http.StatusOK, models.Response[[]models.ProductType]{
		Status:  "success",
		Message: "all product types",
		Data:    productTypes,
	})
}

func (ptc *ProductTypeController) GetByID(c echo.Context) error {
	productTypeID := c.Param("id")

	productType, err := ptc.service.GetByID(productTypeID)

	if err != nil {
		return c.JSON(http.StatusNotFound, models.Response[string]{
			Status:  "failed",
			Message: "product type not found",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.ProductType]{
		Status:  "success",
		Message: "product type found",
		Data:    productType,
	})
}

func (ptc *ProductTypeController) Create(c echo.Context) error {
	var ptReq models.ProductTypeRequest
	if err := verifyAdminPT(c, ptc); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	if err := c.Bind(&ptReq); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	err := ptReq.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "please fill all the required fields",
		})
	}

	productType, err := ptc.service.Create(ptReq)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to create a product type",
		})
	}

	return c.JSON(http.StatusCreated, models.Response[models.ProductType]{
		Status:  "success",
		Message: "product type created",
		Data:    productType,
	})
}

func (ptc *ProductTypeController) Update(c echo.Context) error {
	var ptReq models.ProductTypeRequest
	if err := verifyAdminPT(c, ptc); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}
	if err := c.Bind(&ptReq); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	productTypeID := c.Param("id")

	err := ptReq.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "please fill all the required fields",
		})
	}

	productType, err := ptc.service.Update(ptReq, productTypeID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to update a product type",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.ProductType]{
		Status:  "success",
		Message: "product type updated",
		Data:    productType,
	})
}

func (ptc *ProductTypeController) Delete(c echo.Context) error {
	productTypeID := c.Param("id")
	if err := verifyAdminPT(c, ptc); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}
	err := ptc.service.Delete(productTypeID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to delete product type",
		})
	}

	return c.JSON(http.StatusOK, models.Response[string]{
		Status:  "success",
		Message: "product type deleted",
	})
}
