package route

import (
	"github.com/rzfd/gorm-ners/internal/handlers/http/controller"
	_ "github.com/swaggo/echo-swagger/example/docs"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

func RegisterCompanyRoutes(group *echo.Group, db *gorm.DB) {
	cc := &controller.CompanyController{DB: db}
	group.POST("/companies", cc.CreateCompany)
	group.GET("/companies/:id", cc.GetCompany)
	group.PUT("/companies/:id", cc.UpdateCompany)
	group.DELETE("/companies/:id", cc.DeleteCompany)
	group.GET("/companies", cc.GetAllCompanies)
}
