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

