package controllers

import (
	"aquaculture/middlewares"
	"aquaculture/models"
	"aquaculture/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TransactionController struct {
	service     services.TransactionService
	userService services.UserService
}

func InitTransactionController() TransactionController {
	return TransactionController{
		service:     services.InitTransactionService(),
		userService: services.InitUserService(models.JWTOptions{}),
	}
}
func verifyUserT(c echo.Context, tc *TransactionController) error {
	claim, err := middlewares.GetUser(c)
	if err != nil {
		return err
	}
	user, err := tc.userService.GetUserInfo(strconv.Itoa(claim.ID))
	if err != nil && user.ID == 0 {
		return err
	}
	return nil
}

func (tc *TransactionController) GetAll(c echo.Context) error {
	if err := verifyUserT(c, tc); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}
	transactions, err := tc.service.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "error when fetching transactions",
		})
	}

	return c.JSON(http.StatusOK, models.Response[[]models.Transaction]{
		Status:  "success",
		Message: "all transactions",
		Data:    transactions,
	})
}

func (tc *TransactionController) GetByID(c echo.Context) error {
	transactionID := c.Param("id")
	if err := verifyUserT(c, tc); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}
	transaction, err := tc.service.GetByID(transactionID)

	if err != nil {
		return c.JSON(http.StatusNotFound, models.Response[string]{
			Status:  "failed",
			Message: "transaction not found",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.Transaction]{
		Status:  "success",
		Message: "transaction found",
		Data:    transaction,
	})
}

func (tc *TransactionController) Create(c echo.Context) error {
	claim, err := middlewares.GetUser(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "user not found",
		})
	}

	var tReq models.TransactionRequest

	tReq.UserID = uint(claim.ID)

	if err := verifyUserT(c, tc); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	if err := c.Bind(&tReq); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	err = tReq.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "please fill all the required fields",
		})
	}

	transaction, err := tc.service.Create(tReq)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to create a transaction",
		})
	}

	return c.JSON(http.StatusCreated, models.Response[models.Transaction]{
		Status:  "success",
		Message: "transaction created",
		Data:    transaction,
	})
}

func (tc *TransactionController) Update(c echo.Context) error {
	claim, err := middlewares.GetUser(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "user not found",
		})
	}

	var tReq models.TransactionRequest

	tReq.UserID = uint(claim.ID)

	if err := verifyUserT(c, tc); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}
	if err := c.Bind(&tReq); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	transactionID := c.Param("id")

	err = tReq.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "please fill all the required fields",
		})
	}

	transaction, err := tc.service.Update(tReq, transactionID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to update a transaction",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.Transaction]{
		Status:  "success",
		Message: "transaction updated",
		Data:    transaction,
	})
}
