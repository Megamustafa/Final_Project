package controllers

import (
	"aquaculture/middlewares"
	"aquaculture/models"
	"aquaculture/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type FarmTypeController struct {
	service     services.FarmTypeService
	userService services.UserService
}

func InitFarmTypeController() FarmTypeController {
	return FarmTypeController{
		service:     services.InitFarmTypeService(),
		userService: services.InitUserService(models.JWTOptions{}),
	}
}
func verifyAdminFT(c echo.Context, ftc *FarmTypeController) error {
	claim, err := middlewares.GetUser(c)
	if err != nil {
		return err
	}
	admin, err := ftc.userService.GetAdminInfo(strconv.Itoa(claim.ID))
	if err != nil && admin.ID == 0 {
		return err
	}
	return nil
}

func (ftc *FarmTypeController) GetAll(c echo.Context) error {
	farmTypes, err := ftc.service.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "error when fetching product types",
		})
	}

	return c.JSON(http.StatusOK, models.Response[[]models.FarmType]{
		Status:  "success",
		Message: "all farm types",
		Data:    farmTypes,
	})
}

func (ftc *FarmTypeController) GetByID(c echo.Context) error {
	farmTypeID := c.Param("id")

	farmType, err := ftc.service.GetByID(farmTypeID)

	if err != nil {
		return c.JSON(http.StatusNotFound, models.Response[string]{
			Status:  "failed",
			Message: "farm type not found",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.FarmType]{
		Status:  "success",
		Message: "farm type found",
		Data:    farmType,
	})
}

func (ftc *FarmTypeController) Create(c echo.Context) error {
	var ftReq models.FarmTypeRequest
	if err := verifyAdminFT(c, ftc); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	if err := c.Bind(&ftReq); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	err := ftReq.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "please fill all the required fields",
		})
	}

	farmType, err := ftc.service.Create(ftReq)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to create a farm type",
		})
	}

	return c.JSON(http.StatusCreated, models.Response[models.FarmType]{
		Status:  "success",
		Message: "farm type created",
		Data:    farmType,
	})
}

func (ftc *FarmTypeController) Update(c echo.Context) error {
	var ftReq models.FarmTypeRequest
	if err := verifyAdminFT(c, ftc); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}
	if err := c.Bind(&ftReq); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	farmTypeID := c.Param("id")

	err := ftReq.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "please fill all the required fields",
		})
	}

	farmType, err := ftc.service.Update(ftReq, farmTypeID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to update a farm type",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.FarmType]{
		Status:  "success",
		Message: "farm type updated",
		Data:    farmType,
	})
}

func (ftc *FarmTypeController) Delete(c echo.Context) error {
	farmTypeID := c.Param("id")
	if err := verifyAdminFT(c, ftc); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	err := ftc.service.Delete(farmTypeID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to delete a farm type",
		})
	}

	return c.JSON(http.StatusOK, models.Response[string]{
		Status:  "success",
		Message: "farm type deleted",
	})
}
