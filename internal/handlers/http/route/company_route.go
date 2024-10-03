package route

import (
	"github.com/rzfd/gorm-ners/internal/handlers/http/controller"
	_ "github.com/swaggo/echo-swagger/example/docs"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

func RegisterCompanyRoutes(e *echo.Echo, db *gorm.DB) {
	cc := &controller.CompanyController{DB: db}
	e.POST("/companies", cc.CreateCompany)
	e.GET("/companies/:id", cc.GetCompany)
	e.PUT("/companies/:id", cc.UpdateCompany)
	e.DELETE("/companies/:id", cc.DeleteCompany)
	e.GET("/companies", cc.GetAllCompanies)
}
