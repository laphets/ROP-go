package model

import "github.com/jinzhu/gorm"

type FormModel struct {
	gorm.Model
	Name string `gorm:"unique_index" json:"name"`
	RootTag int `json:"root_tag"`
	Data string `json:"data" gorm:"type:varchar(20000)"`
}

func (x *FormModel) Create() error {
	return DB.Local.Create(&x).Error
}
func (x *FormModel) Update() error {
	form := &FormModel{}
	if err := DB.Local.Where("ID = ?", x.ID).First(&form).Error; err != nil {
		return err
	}
	return DB.Local.Model(&x).Update(&x).Error
}
func GetFormByID(formId uint) (*FormModel, error) {
	form := &FormModel{}
	d := DB.Local.Where("ID = ?", formId).First(&form)
	return form, d.Error
}
func ListForm() ([]*FormModel, error) {
	forms := make([]*FormModel, 0)
	d := DB.Local.Find(&forms)
	return forms, d.Error
}