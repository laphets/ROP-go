package model

import "github.com/jinzhu/gorm"

type FormModel struct {
	gorm.Model
	Name string `json:"name"`
	Data string `json:"data"`
}

func (x *FormModel) Create() error {
	return DB.Local.Create(&x).Error
}