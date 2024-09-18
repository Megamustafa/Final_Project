package controllers

import (
	"aquaculture/middlewares"
	"aquaculture/models"
	"aquaculture/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PromoCodeController struct {
	service     services.PromoCodeService
	userService services.UserService
}

func InitPromoCodeController() PromoCodeController {
	return PromoCodeController{
		service:     services.InitPromoCodeService(),
		userService: services.InitUserService(models.JWTOptions{}),
	}
}
func verifyUserPC(c echo.Context, pcc *PromoCodeController) error {
	claim, err := middlewares.GetUser(c)
	if err != nil {
		return err
	}
	user, err := pcc.userService.GetUserInfo(strconv.Itoa(claim.ID))
	if err != nil && user.ID == 0 {
		return err
	}
	return nil
}

func verifyAdminPC(c echo.Context, pcc *PromoCodeController) error {
	claim, err := middlewares.GetUser(c)
	if err != nil {
		return err
	}
	admin, err := pcc.userService.GetAdminInfo(strconv.Itoa(claim.ID))
	if err != nil && admin.ID == 0 {
		return err
	}
	return nil
}

func (pcc *PromoCodeController) GetAll(c echo.Context) error {
	if err := verifyUserPC(c, pcc); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}
	promoCodes, err := pcc.service.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "error when fetching promo codes",
		})
	}

	return c.JSON(http.StatusOK, models.Response[[]models.PromoCode]{
		Status:  "success",
		Message: "all promo codes",
		Data:    promoCodes,
	})
}

func (pcc *PromoCodeController) GetByID(c echo.Context) error {
	promoCodeID := c.Param("id")
	if err := verifyUserPC(c, pcc); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}
	promoCode, err := pcc.service.GetByID(promoCodeID)

	if err != nil {
		return c.JSON(http.StatusNotFound, models.Response[string]{
			Status:  "failed",
			Message: "promo code not found",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.PromoCode]{
		Status:  "success",
		Message: "promo code found",
		Data:    promoCode,
	})
}

func (pcc *PromoCodeController) Create(c echo.Context) error {
	var pcReq models.PromoCodeRequest
	if err := verifyAdminPC(c, pcc); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}
	if err := c.Bind(&pcReq); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	err := pcReq.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "please fill all the required fields",
		})
	}

	promoCode, err := pcc.service.Create(pcReq)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to create a promo code",
		})
	}

	return c.JSON(http.StatusCreated, models.Response[models.PromoCode]{
		Status:  "success",
		Message: "promo code created",
		Data:    promoCode,
	})
}

func (pcc *PromoCodeController) Update(c echo.Context) error {
	var pcReq models.PromoCodeRequest
	if err := verifyAdminPC(c, pcc); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	if err := c.Bind(&pcReq); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	promoCodeID := c.Param("id")

	err := pcReq.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "please fill all the required fields",
		})
	}

	promoCode, err := pcc.service.Update(pcReq, promoCodeID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to update a promo code",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.PromoCode]{
		Status:  "success",
		Message: "promo code updated",
		Data:    promoCode,
	})
}
