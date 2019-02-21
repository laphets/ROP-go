package form

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"git.zjuqsc.com/rop/ROP-go/pkg/errno"
	. "git.zjuqsc.com/rop/ROP-go/handler"
	"git.zjuqsc.com/rop/ROP-go/model"
	json2 "encoding/json"
)

func Update(c *gin.Context) {
	formId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		SendResponse(c, errno.ErrParam, err)
		return
	}
	if can, err := model.CanFormBeEdited(uint(formId)); !can {
		SendResponse(c, errno.ErrFormCantEdit, err)
		return
	}

	req := UpdateRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		SendResponse(c, errno.ErrBind, err.Error())
		return
	}

	json, err := json2.Marshal(req.Data)
	if err != nil {
		SendResponse(c, errno.ErrParam, nil)
		return
	}

	form := model.FormModel{
		Name:req.Name,
		RootTag: GetRoot(req.Data),
		Data:string(json),
	}


	//form := model.FormModel{}
	//if err := c.ShouldBindJSON(&form); err != nil {
	//	SendResponse(c, errno.ErrBind, err)
	//	return
	//}

	form.ID = uint(formId)

	if err := form.Update(); err != nil {
		SendResponse(c, errno.DBError, err.Error())
		return
	}
	SendResponse(c, nil, nil)
}