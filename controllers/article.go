package controllers

import (
	"aquaculture/middlewares"
	"aquaculture/models"
	"aquaculture/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ArticleController struct {
	service     services.ArticleService
	userService services.UserService
}

func InitArticleController() ArticleController {
	return ArticleController{
		service:     services.InitArticleService(),
		userService: services.InitUserService(models.JWTOptions{}),
	}
}

func verifyAdminA(c echo.Context, ac *ArticleController) error {
	claim, err := middlewares.GetUser(c)
	if err != nil {
		return err
	}
	admin, err := ac.userService.GetAdminInfo(strconv.Itoa(claim.ID))
	if err != nil && admin.ID == 0 {
		return err
	}
	return nil
}

func (ac *ArticleController) GetAll(c echo.Context) error {
	articles, err := ac.service.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "error when fetching articles",
		})
	}

	return c.JSON(http.StatusOK, models.Response[[]models.Article]{
		Status:  "success",
		Message: "all articles",
		Data:    articles,
	})
}

func (ac *ArticleController) GetByID(c echo.Context) error {
	articleID := c.Param("id")

	article, err := ac.service.GetByID(articleID)

	if err != nil {
		return c.JSON(http.StatusNotFound, models.Response[string]{
			Status:  "failed",
			Message: "article not found",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.Article]{
		Status:  "success",
		Message: "article found",
		Data:    article,
	})
}

func (ac *ArticleController) Create(c echo.Context) error {
	var aReq models.ArticleRequest
	if err := verifyAdminA(c, ac); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	if err := c.Bind(&aReq); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	err := aReq.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "please fill all the required fields",
		})
	}

	article, err := ac.service.Create(aReq)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to create a article",
		})
	}

	return c.JSON(http.StatusCreated, models.Response[models.Article]{
		Status:  "success",
		Message: "article created",
		Data:    article,
	})
}

func (ac *ArticleController) Update(c echo.Context) error {
	var aReq models.ArticleRequest
	if err := verifyAdminA(c, ac); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	if err := c.Bind(&aReq); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	articleID := c.Param("id")

	err := aReq.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "please fill all the required fields",
		})
	}

	article, err := ac.service.Update(aReq, articleID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to update a article",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.Article]{
		Status:  "success",
		Message: "article updated",
		Data:    article,
	})
}

func (ac *ArticleController) Delete(c echo.Context) error {
	articleID := c.Param("id")
	if err := verifyAdminA(c, ac); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	err := ac.service.Delete(articleID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to delete an article",
		})
	}

	return c.JSON(http.StatusOK, models.Response[string]{
		Status:  "success",
		Message: "article deleted",
	})
}
