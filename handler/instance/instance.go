package instance

import (
	"time"
	"rop/model"
)

type CreateRequest struct {
	Name string `json:"name" binding:"required"`
	Remark string `json:"remark" binding:"-"`
	Association string `json:"association" binding:"required"`
	StartTime time.Time `json:"start_time" binding:"required"`
	EndTime time.Time `json:"end_time" binding:"required"`
	FormId uint `json:"form_id" binding:"required"`
	MaxIntent int `json:"max_intent" binding:"required"`
}

type ListResponse struct {
	*model.InstanceModel
	Status string `json:"status"`
	FreshmanCount int `json:"freshman_count"`
	FormName string `json:"form_name"`
}
