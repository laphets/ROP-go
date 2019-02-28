package model

import (
	"git.zjuqsc.com/rop/ROP-go/pkg/timerange"
	"github.com/jinzhu/gorm"
	"time"
)

type InterviewModel struct {
	gorm.Model
	InstanceId uint `gorm:"not null;unique_index:idx_instance_interview" json:"instance_id"`
	Name string `gorm:"not null;unique_index:idx_instance_interview" json:"name"`
	InterviewType uint `gorm:"not null;unique_index:idx_instance_interview" json:"interview_type"`
	Department string `gorm:"not null;unique_index:idx_instance_interview" json:"department"`
	Director string `gorm:"not null" json:"director"`
	AutoJoinable int `json:"auto_joinable"`
	Interviewers string `json:"interviewer"`
	StartTime time.Time `json:"start_time"`
	EndTime time.Time `json:"end_time"`
	Capacity int `json:"capacity"`
	Remark string `json:"remark"`
	Location string `json:"location"`
	Available int `json:"available"`
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
func ListInterview(instanceId uint, conditions *InterviewModel) ([]*InterviewModel, error) {
	interviews := make([]*InterviewModel, 0)

	d := DB.Local.Where("instance_id = ?", instanceId).Where(conditions).Find(&interviews)
	return interviews, d.Error

}
func GetInterviewByID(interviewId uint) (*InterviewModel, error) {
	interview := &InterviewModel{}
	d := DB.Local.Where("ID = ?", interviewId).First(&interview)
	return interview, d.Error
}
func EnableInterview(interviewId uint) (error) {
	interview, err := GetInterviewByID(interviewId)
	if err != nil {
		return err
	}
	interview.Available = 1
	return DB.Local.Save(&interview).Error
}
func DisableInterview(interviewId uint) (error) {
	interview, err := GetInterviewByID(interviewId)
	if err != nil {
		return err
	}
	interview.Available = 0
	return DB.Local.Save(&interview).Error
}

type FullInterview struct {
	*InterviewModel
	Status string `json:"status"`
	Participants []*FullIntent `json:"participants"`
}
func ListFulInterview(instanceId uint, conditions *InterviewModel) ([]*FullInterview, error) {
	interviews, err := ListInterview(instanceId, conditions)
	if err != nil {
		return nil, err
	}
	fulInterviews := make([]*FullInterview, 0)
	for _, item := range interviews {
		tmp, err := GetFulInterviewByID(item.ID)
		if err != nil {
			return nil, err
		}
		fulInterviews = append(fulInterviews, tmp)
	}
	return fulInterviews, nil
}
func GetFulInterviewByID(interviewId uint) (*FullInterview, error) {
	interview, err := GetInterviewByID(interviewId)
	if err != nil {
		return nil, err
	}
	fulInterview := &FullInterview{
		InterviewModel: interview,
		Status: timerange.GetStatus(interview.StartTime, interview.EndTime),
	}
	fulIntents, err := ListIntentByInterview(interviewId)
	if err != nil {
		return nil, err
	}
	fulInterview.Participants = fulIntents
	return fulInterview, nil
}