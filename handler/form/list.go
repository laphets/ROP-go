package form

import (
	"github.com/gin-gonic/gin"
	"git.zjuqsc.com/rop/ROP-go/model"
	. "git.zjuqsc.com/rop/ROP-go/handler"
	"git.zjuqsc.com/rop/ROP-go/pkg/errno"
)

func List(c *gin.Context)  {
	forms, err := model.ListForm()
	if err != nil {
		SendResponse(c, errno.DBError, err.Error())
		return
	}

	listRes := make([]*ListResponse, 0)

	for _, item := range forms {
		tmp := &ListResponse{FormModel: item}
		editable, err := model.CanFormBeEdited(item.ID)
		if err != nil {
			SendResponse(c, errno.DBError, err.Error())
			return
		}
		tmp.Editable = editable
		listRes = append(listRes, tmp)
	}
	SendResponse(c, nil, listRes)
}