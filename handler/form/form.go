package form

import "rop/model"

type CreateRequest struct {
	Name string `json:"name" binding:"required"`
	Data string `json:"data" binding:"required"`
}

type ListResponse struct {
	*model.FormModel
	Editable bool `json:"editable"`
}