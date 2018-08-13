package model

import "github.com/jinzhu/gorm"

type FormModel struct {
	gorm.Model
	Name string `gorm:"unique_index" json:"name"`
	Data string `json:"data"`
}

func (x *FormModel) Create() error {
	return DB.Local.Create(&x).Error
}

func GetFormByID(formId uint) (*FormModel, error) {
	form := &FormModel{}
	d := DB.Local.Where("ID = ?", formId).First(&form)
	
	return form, d.Error
}