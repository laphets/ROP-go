package model

import (
	"github.com/jinzhu/gorm"
)

type InstanceModel struct {
	gorm.Model
	Name string `gorm:"unique_index"`
	Remark string
	Association string
	StartTime string
	EndTime string
	FormId uint
}

func (x *InstanceModel) Create() error {
	return DB.Local.Create(&x).Error
}
func (x *InstanceModel) Update() error {
	ins := &InstanceModel{}
	if err := DB.Local.Where("ID = ?", x.ID).First(&ins).Error; err != nil {
		return err
	}
	return DB.Local.Model(&x).Updates(&x).Error
}
func GetInstanceByName(name string) (*InstanceModel, error) {
	ins := &InstanceModel{}
	d := DB.Local.Where("name = ?", name).First(&ins)
	return ins, d.Error
}
func GetInstanceById(ID uint) (*InstanceModel, error) {
	ins := &InstanceModel{}
	d := DB.Local.Where("ID = ?", ID).First(&ins)
	return ins, d.Error
}
func ListInstance() ([]*InstanceModel, error) {
	ins := make([]*InstanceModel, 0)
	d := DB.Local.Find(&ins)
	return ins, d.Error
}

func CanFormBeEdited(formId uint) (bool, error) {
	ins := &InstanceModel{}
	if err := DB.Local.Where("form_id = ?", formId).First(&ins).Error; err != nil {
		// entity not exist
		return true, err
	}
	return false, nil
}