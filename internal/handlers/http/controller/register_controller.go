package controller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rzfd/gorm-ners/internal/handlers/http/entities"
	"github.com/rzfd/gorm-ners/internal/handlers/http/services"
)

func Register(authService *services.AuthService) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := new(entities.Regis)
		if err := c.Bind(u); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Input"})
		}
		if err := authService.RegisterUser(u); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"result": map[string]interface{}{
				"Username": u.Username,
			},
		})
	}
}

func Login(authService *services.AuthService) echo.HandlerFunc {
	return func(e echo.Context) error {
		u := new(entities.Regis)

		if err := e.Bind(u); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Input"})
		}

		user, err := authService.AuthenticateUser(u.Username, u.Password)
		if err != nil {
			return e.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid username or password"})
		}

		token, err := authService.GenerateToken(user.ID)
		if err != nil {
			fmt.Printf("Error generating token: %+v\n", err)
			return e.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not generate token"})
		}

		fmt.Printf("Generated Token: %s\n", token)

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status": "SUCCESS",
			"result": map[string]string{
				"access_token": token,
			},
		})
	}
}
