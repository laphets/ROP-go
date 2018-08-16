package intent

type AssignRequest struct {
	AssignMode string `json:"assign_mode" binding:"required"` // assign mode +> auto or manual
	Intents []uint `json:"intents" binding:"required"`
	TargetInterviewId uint `json:"target_interview_id" binding:"-"`
}
