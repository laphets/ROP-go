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
	UserId []uint `json:"user_id" binding:"required"`
}

type AddUserRequest struct {
	ZJUid string `json:"ZJUid" binding:"required"`
	Mobile string `json:"mobile" binding:"required"`
	Name string `json:"name" binding:"required"`
	Department string `json:"department"`
	AssociationId uint `json:"association_id" binding:"required"`
	AdminLevel uint `json:"admin_level" binding:"-"`
}
