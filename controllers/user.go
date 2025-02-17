package controllers

import (
	"net/http"
	"siki/configs"
	"siki/models"

	"github.com/labstack/echo/v4"
	"github.com/microcosm-cc/bluemonday"
	"gorm.io/gorm"
)

func Create(c echo.Context) error {
	ctx := c.Request().Context()

	policy := bluemonday.UGCPolicy()
	name := c.FormValue("name")
	email := c.FormValue("email")

	cleanName := policy.Sanitize(name)
	cleanEmail := policy.Sanitize(email)

	if name == "" || email == "" {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "Bad Request",
		})
	}

	data := models.User{
		Name:  cleanName,
		Email: cleanEmail,
	}

	if err := models.Create(ctx, data); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "Internal Server Error",
			"error":   err,
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "OK",
		"data":    data,
	})
}

func Read(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")

	var result interface{}
	var err error

	if id != "" {
		var user models.User
		err = configs.Connection.WithContext(ctx).
			Where("id = ?", id).
			First(&user).Error
		result = user

		// Handle not found case
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]any{
				"message": "User not found",
				"id":      id,
			})
		}
	} else {
		var users []models.User
		err = configs.Connection.WithContext(ctx).Find(&users).Error
		result = users
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "OK",
		"data":    result,
	})
}

func Update(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")

	// Sanitize input
	policy := bluemonday.UGCPolicy()
	name := c.FormValue("name")
	email := c.FormValue("email")

	cleanName := policy.Sanitize(name)
	cleanEmail := policy.Sanitize(email)

	// Validate input
	if name == "" || email == "" {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "Bad Request",
		})
	}

	data := models.User{
		Name:  cleanName,
		Email: cleanEmail,
	}

	if err := models.Update(ctx, id, data); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "Internal Server Error",
			"error":   err,
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "OK",
		"data":    data,
	})
}

func Delete(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")

	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "ID is required",
		})
	}

	if err := models.Delete(ctx, id); err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]any{
				"message": "User not found",
				"id":      id,
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "User deleted successfully",
	})
}
