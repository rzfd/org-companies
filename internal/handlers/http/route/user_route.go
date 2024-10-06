package route

import (
	"github.com/rzfd/gorm-ners/internal/handlers/http/controller"

	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

func RegisterUserRoutes(group *echo.Group, db *gorm.DB) {
	uc := &controller.UserController{DB: db}
	group.POST("/users", uc.CreateUser)
	group.GET("/users/:id", uc.GetUser)
	group.PUT("/users/:id", uc.UpdateUser)
	group.DELETE("/users/:id", uc.DeleteUser)
	group.GET("/users", uc.GetAllUsers)
}
