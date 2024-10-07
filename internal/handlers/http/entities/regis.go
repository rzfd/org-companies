package entities

type Regis struct {
	ID       int    `gorm:"type:char(36);primary_key" json:"user_id"`
	Uname    string `json:"uname"`
	Password string `json:"pwd"`
}
