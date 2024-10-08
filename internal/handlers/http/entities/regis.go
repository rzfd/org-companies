package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Regis struct {
	ID       uuid.UUID `gorm:"type:char(36);primary_key" json:"user_id"`
	Username string    `json:"username"`
	Password string    `json:"pwd"`
}

func (u *Regis) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
