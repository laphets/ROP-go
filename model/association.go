package model

import "github.com/jinzhu/gorm"

type AssociationModel struct {
	gorm.Model
	Name string `gorm:"unique_index" json:"name"`
	DepartmentList string `gorm:"not null" json:"department_list"`
}

func (x *AssociationModel) Create() error {
	return DB.Local.Create(&x).Error
}

func GetAssociationByName(associationName string) (*AssociationModel, error) {
	association := &AssociationModel{}
	d := DB.Local.Where("name = ?", associationName).First(&association)
	return association, d.Error
}

func ListAssociation() ([]*AssociationModel, error) {
	associations := make([]*AssociationModel, 0)
	d := DB.Local.Find(&associations)
	return associations, d.Error
}

func GetAssociationById(ID uint) (*AssociationModel, error) {
	ins := &AssociationModel{}
	d := DB.Local.Where("ID = ?", ID).First(&ins)
	return ins, d.Error
}