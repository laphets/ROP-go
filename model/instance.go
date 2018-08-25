package model

import (
	"github.com/jinzhu/gorm"
	"time"
	"rop/pkg/timerange"
)

type InstanceModel struct {
	gorm.Model
	Name string `gorm:"not null;unique_index" json:"name"`
	Remark string `json:"remark"`
	Association string `json:"association"`
	StartTime time.Time `json:"start_time"`
	EndTime time.Time `json:"end_time"`
	FormId uint `json:"form_id" gorm:"not null"`
	MaxIntent int `json:"max_intent"`
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
	ins := make([]*InstanceModel, 0)
	if err := DB.Local.Where("form_id = ?", formId).Find(&ins).Error; err != nil {
		// entity not exist
		return false, err
	}
	//log.Debugf("%d", len(ins))
	if len(ins) == 0 {
		return true, nil
	}
	now := time.Now()
	for _, item := range ins {
		if timerange.GetStatusSync(now, item.StartTime, item.EndTime) == "cur" {
			return false, nil
		}
	}
	return true, nil
}