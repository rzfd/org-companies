package controller

import (
	"net/http"

	"github.com/rzfd/gorm-ners/internal/handlers/http/entities"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	DB *gorm.DB
}

func (uc *UserController) CreateUser(c echo.Context) error {
	user := new(entities.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}
	if user.Name == "" || user.CompanyID == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Name and CompanyID are required"})
	}
	if err := uc.DB.Create(user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create user"})
	}
	return c.JSON(http.StatusCreated, user)
}

func (uc *UserController) GetUser(c echo.Context) error {
	id := c.Param("id")
	var user entities.User
	if err := uc.DB.Preload("Company").First(&user, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "User not found"})
	}
	return c.JSON(http.StatusOK, user)
}

func (uc *UserController) UpdateUser(c echo.Context) error {
	id := c.Param("id")
	var user entities.User
	if err := uc.DB.First(&user, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "User not found"})
	}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}
	if err := uc.DB.Save(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to update user"})
	}
	return c.JSON(http.StatusOK, user)
}

func (uc *UserController) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	if err := uc.DB.Delete(&entities.User{}, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to delete user"})
	}
	return c.NoContent(http.StatusNoContent)
}

func (uc *UserController) GetAllUsers(c echo.Context) error {
	var users []entities.User
	if err := uc.DB.Preload("Company").Find(&users).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, users)
}
