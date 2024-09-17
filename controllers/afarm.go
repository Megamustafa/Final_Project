package controllers

import (
	"aquaculture/middlewares"
	"aquaculture/models"
	"aquaculture/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AquacultureFarmsController struct {
	service     services.AquacultureFarmsService
	userService services.UserService
}

func InitAquacultureFarmsController() AquacultureFarmsController {
	return AquacultureFarmsController{
		service:     services.InitAquacultureFarmsService(),
		userService: services.InitUserService(models.JWTOptions{}),
	}
}

func verifyAdminAF(c echo.Context, afc *AquacultureFarmsController) error {
	claim, err := middlewares.GetUser(c)
	if err != nil {
		return err
	}
	admin, err := afc.userService.GetAdminInfo(strconv.Itoa(claim.ID))
	if err != nil && admin.ID == 0 {
		return err
	}
	return nil
}

func (afc *AquacultureFarmsController) GetAll(c echo.Context) error {
	aquaculturefarms, err := afc.service.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "error when fetching aquaculture farms",
		})
	}

	return c.JSON(http.StatusOK, models.Response[[]models.AquacultureFarms]{
		Status:  "success",
		Message: "all aquaculture farms",
		Data:    aquaculturefarms,
	})
}

func (afc *AquacultureFarmsController) GetByID(c echo.Context) error {
	aquaculturefarmID := c.Param("id")

	aquaculturefarm, err := afc.service.GetByID(aquaculturefarmID)

	if err != nil {
		return c.JSON(http.StatusNotFound, models.Response[string]{
			Status:  "failed",
			Message: "aquaculture farm not found",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.AquacultureFarms]{
		Status:  "success",
		Message: "aquaculture farm found",
		Data:    aquaculturefarm,
	})
}

func (afc *AquacultureFarmsController) Create(c echo.Context) error {
	var afReq models.AquacultureFarmsRequest

	if err := verifyAdminAF(c, afc); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	if err := c.Bind(&afReq); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	err := afReq.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "please fill all the required fields",
		})
	}

	aquaculturefarm, err := afc.service.Create(afReq)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to create an aquaculture farm",
		})
	}

	return c.JSON(http.StatusCreated, models.Response[models.AquacultureFarms]{
		Status:  "success",
		Message: "aquaculture farm created",
		Data:    aquaculturefarm,
	})
}

func (afc *AquacultureFarmsController) Update(c echo.Context) error {
	var afReq models.AquacultureFarmsRequest

	if err := verifyAdminAF(c, afc); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}
	if err := c.Bind(&afReq); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	aquaculturefarmID := c.Param("id")

	err := afReq.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "please fill all the required fields",
		})
	}

	aquaculturefarm, err := afc.service.Update(afReq, aquaculturefarmID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to update an aquaculture farm",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.AquacultureFarms]{
		Status:  "success",
		Message: "aquaculture farm updated",
		Data:    aquaculturefarm,
	})
}

func (afc *AquacultureFarmsController) Delete(c echo.Context) error {
	aquaculturefarmID := c.Param("id")

	if err := verifyAdminAF(c, afc); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	err := afc.service.Delete(aquaculturefarmID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to delete aquaculture farm",
		})
	}

	return c.JSON(http.StatusOK, models.Response[string]{
		Status:  "success",
		Message: "aquaculture farm deleted",
	})
}
