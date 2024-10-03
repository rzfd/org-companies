package controller

import (
	"net/http"

	"github.com/rzfd/gorm-ners/internal/handlers/http/model"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

type CompanyController struct {
	DB *gorm.DB
}

func (cc *CompanyController) CreateCompany(c echo.Context) error {
	company := new(model.Company)
	if err := c.Bind(company); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}
	if company.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Name is required"})
	}
	if err := cc.DB.Create(company).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create company"})
	}
	return c.JSON(http.StatusCreated, company)
}

func (cc *CompanyController) GetCompany(c echo.Context) error {
	id := c.Param("id")
	var company model.Company
	if err := cc.DB.First(&company, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Company not found"})
	}
	return c.JSON(http.StatusOK, company)
}

func (cc *CompanyController) UpdateCompany(c echo.Context) error {
	id := c.Param("id")
	var company model.Company
	if err := cc.DB.First(&company, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Company not found"})
	}
	if err := c.Bind(&company); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}
	if err := cc.DB.Save(&company).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to update company"})
	}
	return c.JSON(http.StatusOK, company)
}

func (cc *CompanyController) DeleteCompany(c echo.Context) error {
	id := c.Param("id")
	if err := cc.DB.Delete(&model.Company{}, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to delete company"})
	}
	return c.NoContent(http.StatusNoContent)
}

func (cc *CompanyController) GetAllCompanies(c echo.Context) error {
	var companies []model.Company
	if err := cc.DB.Find(&companies).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to retrieve companies"})
	}
	return c.JSON(http.StatusOK, companies)
}
