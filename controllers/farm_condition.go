package controllers

import (
	"aquaculture/middlewares"
	"aquaculture/models"
	"aquaculture/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type FarmConditionController struct {
	service     services.FarmConditionService
	userService services.UserService
}

func InitFarmConditionController() FarmConditionController {
	return FarmConditionController{
		service:     services.InitFarmConditionService(),
		userService: services.InitUserService(models.JWTOptions{}),
	}
}
func verifyUserFC(c echo.Context, fcc *FarmConditionController) error {
	claim, err := middlewares.GetUser(c)
	if err != nil {
		return err
	}
	user, err := fcc.userService.GetUserInfo(strconv.Itoa(claim.ID))
	if err != nil && user.ID == 0 {
		return err
	}
	return nil
}

func verifyAdminFC(c echo.Context, fcc *FarmConditionController) error {
	claim, err := middlewares.GetUser(c)
	if err != nil {
		return err
	}
	admin, err := fcc.userService.GetAdminInfo(strconv.Itoa(claim.ID))
	if err != nil && admin.ID == 0 {
		return err
	}
	return nil
}

func (fcc *FarmConditionController) GetAll(c echo.Context) error {
	if err := verifyUserFC(c, fcc); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}
	farmConditions, err := fcc.service.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "error when fetching farm conditions",
		})
	}

	return c.JSON(http.StatusOK, models.Response[[]models.FarmCondition]{
		Status:  "success",
		Message: "all farm conditions",
		Data:    farmConditions,
	})
}

func (fcc *FarmConditionController) GetByID(c echo.Context) error {
	farmConditionID := c.Param("id")
	if err := verifyUserFC(c, fcc); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}
	farmCondition, err := fcc.service.GetByID(farmConditionID)

	if err != nil {
		return c.JSON(http.StatusNotFound, models.Response[string]{
			Status:  "failed",
			Message: "farm condition not found",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.FarmCondition]{
		Status:  "success",
		Message: "farm condition found",
		Data:    farmCondition,
	})
}

func (fcc *FarmConditionController) Create(c echo.Context) error {
	var fcReq models.FarmConditionRequest
	if err := verifyAdminFC(c, fcc); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}
	if err := c.Bind(&fcReq); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	err := fcReq.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "please fill all the required fields",
		})
	}

	farmCondition, err := fcc.service.Create(fcReq)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to create a farm condition",
		})
	}

	return c.JSON(http.StatusCreated, models.Response[models.FarmCondition]{
		Status:  "success",
		Message: "farm condition created",
		Data:    farmCondition,
	})
}

func (fcc *FarmConditionController) Update(c echo.Context) error {
	var fcReq models.FarmConditionRequest
	if err := verifyAdminFC(c, fcc); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	if err := c.Bind(&fcReq); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	farmConditionID := c.Param("id")

	err := fcReq.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "please fill all the required fields",
		})
	}

	farmCondition, err := fcc.service.Update(fcReq, farmConditionID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to update a farm condition",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.FarmCondition]{
		Status:  "success",
		Message: "farm condition updated",
		Data:    farmCondition,
	})
}
