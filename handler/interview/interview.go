package interview

import "time"

type CreateRequest struct {
	//InstanceId uint `json:"instance_id" binding:"required"`
	Name string `json:"name" binding:"required"`
	InterviewType uint `json:"interview_type" binding:"required"`
	Department string `json:"department" binding:"required"`
	Director string `json:"director" binding:"required"`
	AutoJoinable int `json:"auto_joinable" binding:"required"`
	Interviewers string `json:"interviewers" binding:"required"`
	StartTime time.Time `json:"start_time" binding:"required"`
	EndTime time.Time `json:"end_time" binding:"required"`
	Capacity int `json:"capacity" binding:"required"`
	Remark string `json:"remark" binding:"-"`
	Location string `json:"location" binding:"required"`
}

type JoinRequest struct {
	Intents []uint `json:"intents"`
}