package model

import "github.com/jinzhu/gorm"

type IntentModel struct {
	gorm.Model
	FreshmanId uint `gorm:"not null;unique_index:idx_freshman_department" json:"freshman_id"`
	Department string `gorm:"not null;unique_index:idx_freshman_department" json:"department"`
	InterviewId uint `json:"interview_id"`
	MainStage string `json:"main_stage"`
	SubStage string `json:"sub_stage"`
}

func (x *IntentModel) Create() error {
	// Specify FreshmanId
	return DB.Local.Create(&x).Error
}

//func (x *IntentModel) Update(freshmanId uint) error {
//	intent := &IntentModel{}
//	if err := DB.Local.Where("freshman_id = ?", freshmanId).
//}