package instance

import "time"

type CreateRequest struct {
	Name string `json:"name" binding:"required"`
	Remark string `json:"remark" binding:"-"`
	Association string `json:"association" binding:"required"`
	StartTime time.Time `json:"start_time" binding:"required"`
	EndTime time.Time `json:"end_time" binding:"required"`
	FormId uint `json:"form_id" binding:"required"`
}