package model

import (
	"github.com/jinzhu/gorm"
)

type SmsModel struct {
	gorm.Model
	Mobile string `json:"mobile"`
	Text string `json:"text"`
	Result string `json:"result"`
	Error string `json:"error"`
}

func (x *SmsModel) Create() error {
	return DB.Local.Create(&x).Error
}
