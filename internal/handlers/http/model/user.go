package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string  `json:"name"`
	CompanyID int     `json:"company_id"`
	Company   Company `json:"company"`
}
