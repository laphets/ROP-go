package model

import (
	"github.com/jinzhu/gorm"
)

type LogModel struct {
	gorm.Model
	UserId uint `json:"user_id"`
	IP string `json:"ip"`
	URL string `json:"url"`
	UA string `json:"ua"`
	Info string `json:"info"`
}

func (x *LogModel) Create() error {
	return DB.Local.Create(&x).Error
}
