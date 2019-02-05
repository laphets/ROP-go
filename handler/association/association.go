package association

import (
	"git.zjuqsc.com/rop/ROP-go/model"
)

type CreateRequest struct {
	Name string `json:"name" binding:"required"`
	Department []string `json:"department" binding:"required"`
}

type GetResponse struct {
	*model.AssociationModel
	Department []string `json:"department"`
}

type SendNoticeRequest struct {
	ZJUid []string `json:"ZJUid" binding:"required"`
}