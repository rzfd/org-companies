package controller

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/rzfd/gorm-ners/internal/handlers/http/entities"
	"github.com/rzfd/gorm-ners/internal/handlers/http/security"
	"gorm.io/gorm"
)

func Regis(db *gorm.DB, jwtSecret string) echo.HandlerFunc {
	return func(e echo.Context) error {
		u := new(entities.Regis)
		if err := e.Bind(u); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Input"})
		}
		var existingUser entities.Regis
		if err := db.Where("uname = ?", u.Uname).First(&existingUser).Error; err == nil {
			return e.JSON(http.StatusBadRequest, map[string]string{"error": "User already exists"})
		}
		hash, err := security.HashPassword(u.Password)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]string{"error": "Error When Hashing"})
		}

		u.Password = strings.ToLower(strings.ReplaceAll(u.Uname, " ", ""))
		u.Password = strings.ReplaceAll(u.Password, " ", "")
		u.Password = hash
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
		u := new(entities.Regis)
		if err := e.Bind(u); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Input"})
		}
		var existingUser entities.Regis
		if err := db.Where("uname = ?", u.Uname).First(&existingUser).Error; err != nil {
			return e.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid username or password"})
		}
		if !security.CheckPasswordHash(u.Password, existingUser.Password) {
			return e.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid username or password"})
		}
		token, err := security.GetToken(u.ID, jwtSecret)
		if err != nil {
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
