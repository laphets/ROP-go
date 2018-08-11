package model

import "github.com/jinzhu/gorm"

type FormModel struct {
	gorm.Model
	Name string
	data string
}

func (x *FormModel) Create() error {
	return DB.Local.Create(&x).Error
}