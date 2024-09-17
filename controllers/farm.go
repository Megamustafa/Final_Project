package controllers

import (
	"aquaculture/middlewares"
	"aquaculture/models"
	"aquaculture/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type FarmController struct {
	service     services.FarmService
	userService services.UserService
}

func InitFarmController() FarmController {
	return FarmController{
		service:     services.InitFarmService(),
		userService: services.InitUserService(models.JWTOptions{}),
	}
}
func verifyAdminF(c echo.Context, fc *FarmController) error {
	claim, err := middlewares.GetUser(c)
	if err != nil {
		return err
	}
	admin, err := fc.userService.GetAdminInfo(strconv.Itoa(claim.ID))
	if err != nil && admin.ID == 0 {
		return err
	}
	return nil
}

func (fc *FarmController) GetAll(c echo.Context) error {
	farms, err := fc.service.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "error when fetching farms",
		})
	}

	return c.JSON(http.StatusOK, models.Response[[]models.Farm]{
		Status:  "success",
		Message: "all farms",
		Data:    farms,
	})
}

func (fc *FarmController) GetByID(c echo.Context) error {
	farmID := c.Param("id")

	farm, err := fc.service.GetByID(farmID)

	if err != nil {
		return c.JSON(http.StatusNotFound, models.Response[string]{
			Status:  "failed",
			Message: "farm not found",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.Farm]{
		Status:  "success",
		Message: "farm found",
		Data:    farm,
	})
}

func (fc *FarmController) Create(c echo.Context) error {
	var farmReq models.FarmRequest
	if err := verifyAdminF(c, fc); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}
	if err := c.Bind(&farmReq); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	err := farmReq.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "please fill all the required fields",
		})
	}

	farm, err := fc.service.Create(farmReq)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to create a farm",
		})
	}

	return c.JSON(http.StatusCreated, models.Response[models.Farm]{
		Status:  "success",
		Message: "farm created",
		Data:    farm,
	})
}

func (fc *FarmController) Update(c echo.Context) error {
	var farmReq models.FarmRequest
	if err := verifyAdminF(c, fc); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	if err := c.Bind(&farmReq); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	farmID := c.Param("id")

	err := farmReq.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "please fill all the required fields",
		})
	}

	farm, err := fc.service.Update(farmReq, farmID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to update a farm",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.Farm]{
		Status:  "success",
		Message: "farm updated",
		Data:    farm,
	})
}

func (fc *FarmController) Delete(c echo.Context) error {
	farmID := c.Param("id")
	if err := verifyAdminF(c, fc); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}
	err := fc.service.Delete(farmID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to delete a farm",
		})
	}

	return c.JSON(http.StatusOK, models.Response[string]{
		Status:  "success",
		Message: "farm deleted",
	})
}
