package user

import (
	"git.zjuqsc.com/rop/ROP-go/model"
)

type GetInfoResponse struct {

}

type AvatarRequest struct {
	URL string `json:"url" binding:"required"`
}

type LoginByPasswordRequest struct {
	AssociationId uint `json:"association_id" binding:"required"`
	ZJUid string `json:"ZJUid" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type InfoResponse struct {
	*model.UserModel
	Association *model.AssociationModel `json:"association"`
}

type RegisterRequest struct {
	Uid string `json:"uid" binding:"required"`
	Password string `json:"password" binding:"required"`
	Department string `json:"department" binding:"required"`
}
