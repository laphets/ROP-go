package model

import (
	"github.com/jinzhu/gorm"
	"fmt"
)

type FreshmanModel struct {
	gorm.Model
	ZJUid string `gorm:"unique_index"`
	Mobile string
	MainStage string
	SubStage string
	Intent []string
	OtherInfo string
}

func getTableName(instanceId uint) string {
	return fmt.Sprintf("freshman#%d", instanceId)
}
func CreateInsIfNotExist(instanceId uint) {
	if !DB.Local.HasTable(getTableName(instanceId)) {
		DB.Local.CreateTable(getTableName(instanceId))
	}
}

func (x *FreshmanModel) Create(instanceId uint) error {
	CreateInsIfNotExist(instanceId)
	return DB.Local.Create(&x).Error
}
func (x *FreshmanModel) Update(instanceId uint) error {
	CreateInsIfNotExist(instanceId)
	fre := &FreshmanModel{}
	if err := DB.Local.Table(getTableName(instanceId)).Where("ZJUid = ?", x.ZJUid).First(&fre).Error; err != nil {
		return err
	}
	return DB.Local.Table(getTableName(instanceId)).Model(&x).Update(&x).Error
}
func GetFreshmanByZJUid(instanceId uint, ZJUid string) (*FreshmanModel, error) {
	fre := &FreshmanModel{}
	d := DB.Local.Table(getTableName(instanceId)).Where("ZJUid = ?", ZJUid).First(&fre)
	return fre, d.Error
}
func ListFreshman(instanceId uint) ([]*FreshmanModel, error) {
	fre := make([]*FreshmanModel, 0)
	d := DB.Local.Table(getTableName(instanceId)).Find(&fre)
	return fre, d.Error
}