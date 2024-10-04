package controller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rzfd/gorm-ners/internal/handlers/http/middleware"
	"github.com/rzfd/gorm-ners/internal/handlers/http/model"
	"gorm.io/gorm"
)

func Regis(db *gorm.DB, jwtSecret string) echo.HandlerFunc {
	return func(e echo.Context) error {
		u := new(model.Regis)
		if err := e.Bind(u); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Input"})
		}
		if err := db.Create(u).Error; err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]string{"error": "Registration failed"})
		}

		return e.JSON(http.StatusCreated, map[string]interface{}{
			"result": map[string]interface{}{
				"Username": u.Uname,
			},
		})
	}
}

func Login(db *gorm.DB, jwtSecret string) echo.HandlerFunc {
	return func(e echo.Context) error {
		u := new(model.Regis)
		if err := e.Bind(u); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Input"})
		}
		token, err := middleware.GetToken(u.ID, jwtSecret)
		if err != nil {
			fmt.Printf("JWT error: %+v\n", err)
			return e.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not generate token"})
		}
		return e.JSON(http.StatusOK, map[string]interface{}{
			"status": "SUCCESS",
			"result": map[string]string{
				"access_token":  token,
				"refresh_token": "",
			},
		})
	}
}
