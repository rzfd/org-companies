package route

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB) {
	RegisterUserRoutes(e, db)
	RegisterCompanyRoutes(e, db)
}
