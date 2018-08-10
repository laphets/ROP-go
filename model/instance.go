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
	FormId string
}

func (x *InstanceModel) Create() error {
	return DB.Local.Create(&x).Error
}
func GetInstanceByName(name string) (*InstanceModel, error) {
	ins := &InstanceModel{}
	d := DB.Local.Where("name = ?", name).First(&ins)
	return ins, d.Error
}