package form

import "git.zjuqsc.com/rop/ROP-go/model"

type CreateRequest struct {
	Name string `json:"name" binding:"required"`
	Data []*DataItem `json:"data" binding:"required,dive"`
}

type UpdateRequest struct {
	Name string `json:"name" binding:"required"`
	Data []*DataItem `json:"data" binding:"required,dive"`
}

type DataItem struct {
	Tag int `json:"tag" binding:"required"`
	Text string `json:"text" binding:"required"`
	Type string `json:"type" binding:"required"`
	Next int `json:"next" binding:"required"`
	AvailableCnt int `json:"available_cnt" binding:"-"`
	Required bool `json:"required" binding:"-"`
	DefaultJump bool `json:"default_jump" binding:"-"`
	Spec string `json:"spec" binding:"-"`
	Re string `json:"re" binding:"-"`
	Choices []*Choice `json:"choices" binding:"dive"`
}

type Choice struct {
	Tag int `json:"tag" binding:"required"`
	Text string `json:"text" binding:"required"`
	Next int `json:"next" binding:"required"`
}

type ListResponse struct {
	*model.FormModel
	Editable bool `json:"editable"`
}