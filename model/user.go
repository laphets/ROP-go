package model

import "github.com/jinzhu/gorm"

type UserModel struct {
	gorm.Model
	ZJUid string	`gorm:"unique_index;column:ZJUid"`
	Name string
	Department string
	Gender string
	InnerId string
	Position string
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