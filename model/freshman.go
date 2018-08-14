package model

import (
	"github.com/jinzhu/gorm"
)

type FreshmanModel struct {
	gorm.Model
	InstanceId uint `gorm:"not null;unique_index:idx_instance_ZJUid" json:"instance_id"`
	ZJUid string `gorm:"not null;column:ZJUid;unique_index:idx_instance_ZJUid" json:"ZJUid"`
	Mobile string `json:"mobile"`
	//MainStage string `json:"main_stage"`
	//SubStage string `json:"sub_stage"`
	//Intent []string `json:"intent"`
	OtherInfo string `json:"other_info"`
}


func (x *FreshmanModel) Create() error {
	return DB.Local.Create(&x).Error
}
func (x *FreshmanModel) Update(instanceId uint) error {
	fre := &FreshmanModel{}
	if err := DB.Local.Where("instance_id = ?", instanceId).Where("ZJUid = ?", x.ZJUid).First(&fre).Error; err != nil {
		return err
	}
	return DB.Local.Where("instance_id = ?", instanceId).Model(&x).Update(&x).Error
}
func GetFreshmanCountByID(instanceId uint) (int, error) {
	//log.Debugf("%d",instanceId)
	count := 0
	d := DB.Local.Model(&FreshmanModel{}).Where("instance_id = ?", instanceId).Count(&count)
	return count, d.Error
}
func GetFreshmanByZJUid(instanceId uint, ZJUid string) (*FreshmanModel, error) {
	fre := &FreshmanModel{}
	d := DB.Local.Where("instance_id = ?", instanceId).Where("ZJUid = ?", ZJUid).First(&fre)
	return fre, d.Error
}
func ListFreshman(instanceId uint) ([]*FreshmanModel, error) {
	fre := make([]*FreshmanModel, 0)
	d := DB.Local.Where("instance_id = ?", instanceId).Find(&fre)
	return fre, d.Error
}