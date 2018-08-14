package form

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"rop/pkg/errno"
	. "rop/handler"
	"rop/model"
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


	form := model.FormModel{}
	if err := c.ShouldBindJSON(&form); err != nil {
		SendResponse(c, errno.ErrBind, err)
		return
	}

	form.ID = uint(formId)

	if err := form.Update(); err != nil {
		SendResponse(c, errno.DBError, err.Error())
		return
	}
	SendResponse(c, nil, nil)
}