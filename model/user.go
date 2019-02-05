package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type UserModel struct {
	gorm.Model
	ZJUid string	`gorm:"not null;unique_index;column:ZJUid"`
	Name string		`json:"name"`
	Department string	`json:"department"`
	Gender string	`json:"gender"`
	InnerId string	`json:"inner_id"`
	Position string	`json:"position"`
	Mobile string `json:"mobile"`
	LastSeen time.Time `json:"last_seen"`
	AssociationId string `json:"association_id"`
	Avatar string `json:"avatar"`
}

func (x *UserModel) Create() error {
	return DB.Local.Create(&x).Error
}
func (x *UserModel) Update() error {
	return DB.Local.Save(x).Error
}
func Delete(ZJUid string) error {
	user := UserModel{}
	user.ZJUid = ZJUid
	return DB.Local.Delete(&user).Error
}
func GetUserByZJUid(ZJUid string) (*UserModel, error) {
	user := &UserModel{}
	d := DB.Local.Where("ZJUid = ?", ZJUid).First(&user)
	return user, d.Error
}
func ListUser() ([]*UserModel, error) {
	users := make([]*UserModel, 0)
	d := DB.Local.Find(&users)
	return users, d.Error
}
func UpdateLastSeen(ZJUid string) (error) {
	user, err := GetUserByZJUid(ZJUid)
	if err != nil {
		return err
	}
	user.LastSeen = time.Now()
	return DB.Local.Save(&user).Error
}
func UpdateAvatar(ZJUid, URL string) (error) {
	user, err := GetUserByZJUid(ZJUid)
	if err != nil {
		return err
	}
	user.Avatar = URL
	return DB.Local.Save(&user).Error
}
