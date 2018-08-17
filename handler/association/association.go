package association

type CreateRequest struct {
	Name string `json:"name" binding:"required"`
	Department []string `json:"department" binding:"required"`
}

type GetResponse struct {
	Department []string `json:"department"`
}