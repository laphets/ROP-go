package instance

type CreateRequest struct {
	Name string `json:"name" binding:"required"`
	Remark string `json:"remark" binding:"-"`
	Association string `json:"association" binding:"required"`
	StartTime string `json:"start_time" binding:"required"`
	EndTime string `json:"end_time" binding:"required"`
	FormId string `json:"form_id" binding:"required"`
}