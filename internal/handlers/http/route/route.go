package route

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterRoutes(group *echo.Group, db *gorm.DB) {
	RegisterUserRoutes(group, db)
	RegisterCompanyRoutes(group, db)
}
