package route

import (
	"github.com/rzfd/gorm-ners/internal/handlers/http/controller"

	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

func RegisterUserRoutes(e *echo.Echo, db *gorm.DB) {
	uc := &controller.UserController{DB: db}
	e.POST("/users", uc.CreateUser)
	e.GET("/users/:id", uc.GetUser)
	e.PUT("/users/:id", uc.UpdateUser)
	e.DELETE("/users/:id", uc.DeleteUser)
	e.GET("/users", uc.GetAllUsers)
}
