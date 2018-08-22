package freshman

type SubmitRequest struct {
	Data []*submitData	`json:"data" binding:"required,dive"`
}
type submitData struct {
	Key int		`json:"key" binding:"required"`
	Value []string	`json:"value" binding:"required"`
}

type OtherInfo struct {
	Key string `json:"key"`
	Value string `json:"value"`
}