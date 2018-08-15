package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type InterviewModel struct {
	gorm.Model
	InstanceId uint `json:"instance_id"`
	InterviewType uint `json:"interview_type"`
	Department string `json:"department"`
	Director string `json:"director"`
	Interviewer string `json:"interviewer"`
	StartTime time.Time `json:"start_time"`
	EndTime time.Time `json:"end_time"`
}

func (x *InterviewModel) Create() error {
	return DB.Local.Create(&x).Error
}
func (x *InterviewModel) Update() error {
	interview := &InterviewModel{}
	if err := DB.Local.Where("ID = ?", x.ID).First(&interview).Error; err != nil {
		return err
	}
	return DB.Local.Model(&x).Updates(&x).Error
}
func DeleteInterview(interviewId uint) error {
	interview := &InterviewModel{}
	interview.ID = interviewId
	return DB.Local.Delete(&interview).Error
}
func ListInterview(instanceId uint) ([]*InterviewModel, error) {
	interviews := make([]*InterviewModel, 0)
	d := DB.Local.Where("instance_id = ?", instanceId).Find(&interviews)
	return interviews, d.Error
}
func GetInterviewByID(interviewId uint) (*InterviewModel, error) {
	interview := &InterviewModel{}
	d := DB.Local.Where("ID = ?", interviewId).First(&interview)
	return interview, d.Error
}