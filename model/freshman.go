package model

import (
	"time"
)

type FreshmanModel struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"unique_index:idx_instance_ZJUid"`
	InstanceId uint `gorm:"not null;unique_index:idx_instance_ZJUid" json:"instance_id"`
	ZJUid string `gorm:"not null;column:ZJUid;unique_index:idx_instance_ZJUid" json:"ZJUid"`
	Mobile string `json:"mobile"`
	//MainStage string `json:"main_stage"`
	//SubStage string `json:"sub_stage"`
	//Intent []string `json:"intent"`
	OtherInfo string `json:"other_info"`
}

// When calling this method, you need make sure instanceId and ZJUid exists in your freshmanModel
func (x *FreshmanModel) Create() (error) {
	curFreshman := &FreshmanModel{}
	//log.Debugf("%d %s", x.InstanceId, x.ZJUid)
	if !DB.Local.Where("instance_id = ?", x.InstanceId).Where("ZJUid = ?", x.ZJUid).First(&curFreshman).RecordNotFound() {
		// if record exist, then replace
		if err := DeleteFreshman(curFreshman.ID); err != nil {
			return err
		}
	}
	//log.Debugf("%+v", curFreshman)
	return DB.Local.Create(&x).Error
}

// Rewrite. Abandoned.
func (x *FreshmanModel) Update(instanceId uint) error {
	fre := &FreshmanModel{}
	if err := DB.Local.Where("instance_id = ?", instanceId).Where("ZJUid = ?", x.ZJUid).First(&fre).Error; err != nil {
		return err
	}
	return DB.Local.Where("instance_id = ?", instanceId).Model(&x).Update(&x).Error
}

func DeleteFreshman(freshmanId uint) error {
	freshman := &FreshmanModel{}
	freshman.ID = freshmanId
	//if err := DeleteIntent()
	if err := DB.Local.Where("freshman_id = ?", freshmanId).Delete(&IntentModel{}).Error; err != nil {
		return err
	}
	return DB.Local.Delete(&freshman).Error
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

func GetFreshmanById(freshmanId uint) (*FreshmanModel, error) {
	freshman := &FreshmanModel{}
	d := DB.Local.Where("ID = ?", freshmanId).First(&freshman)
	return freshman, d.Error
}