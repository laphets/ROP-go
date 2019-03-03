package model

import (
	"time"
)

type UserModel struct {
	ID        uint `gorm:"primary_key" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"unique_index:idx_user_association"`
	AssociationId uint `gorm:"not null;unique_index:idx_user_association" json:"association_id"`
	ZJUid string	`gorm:"not null;unique_index:idx_user_association;column:ZJUid" json:"ZJUid"`
	Password string `json:"password"`
	Name string		`json:"name"`
	Department string	`json:"department"`
	Gender string	`json:"gender"`
	InnerId string	`json:"inner_id"`
	Position string	`json:"position"`
	Mobile string `json:"mobile"`
	LastSeen time.Time `json:"last_seen"`
	Avatar string `json:"avatar"`
	AdminLevel uint `json:"admin_level"`	// 0 for none, 1 for association level, 2 for system level
}

func (x *UserModel) Create() error {
	return DB.Local.Create(&x).Error
}
func (x *UserModel) Update() error {
	return DB.Local.Save(x).Error
}
func Delete(UserId uint) error {
	user := UserModel{}
	user.ID = UserId
	return DB.Local.Delete(&user).Error
}

func GetUserById(UserId uint) (*UserModel, error) {
	user := &UserModel{}
	d := DB.Local.Where("id = ?", UserId).First(&user)
	return user, d.Error
}
func GetUserByAssociationAndZJUid(associationId uint, ZJUid string) (*UserModel, error) {
	user := &UserModel{}
	d := DB.Local.Where(&UserModel{AssociationId: associationId, ZJUid: ZJUid}).First(&user)
	return user, d.Error
}

type FullUserModel struct{
	*UserModel
	AssociationName string `json:"association_name"`
}

func getFullUser(userList []*UserModel) ([]*FullUserModel) {
	final := make([]*FullUserModel, 0)
	for _, user := range userList {
		fulUser := &FullUserModel{
			UserModel: user,
		}
		association, err := GetAssociationById(user.AssociationId)
		if err != nil {
			fulUser.AssociationName = ""
		} else {
			fulUser.AssociationName = association.Name
		}
		final = append(final, fulUser)
	}
	return final
}

func ListUser() ([]*FullUserModel, error) {
	users := make([]*UserModel, 0)
	d := DB.Local.Find(&users)
	if d.Error != nil {
		return nil, d.Error
	} else {
		return getFullUser(users), nil
	}
}
func ListUserByAssociation(associationId uint) ([]*FullUserModel, error) {
	users := make([]*UserModel, 0)
	d := DB.Local.Where(&UserModel{AssociationId: associationId}).Find(&users)
	if d.Error != nil {
		return nil, d.Error
	} else {
		return getFullUser(users), nil
	}
}
func UpdateLastSeen(UserId uint) (error) {
	user, err := GetUserById(UserId)
	if err != nil {
		return err
	}
	user.LastSeen = time.Now()
	return DB.Local.Save(&user).Error
}
func UpdateAvatar(UserId uint, URL string) (error) {
	user, err := GetUserById(UserId)
	if err != nil {
		return err
	}
	user.Avatar = URL
	return DB.Local.Save(&user).Error
}
