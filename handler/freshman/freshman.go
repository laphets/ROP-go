package freshman

type SubmitRequest struct {
	Data []*submitData	`json:"data" binding:"required,dive"`
}
type submitData struct {
	Key uint		`json:"key" binding:"required"`
	Value string	`json:"value" binding:"required"`
}