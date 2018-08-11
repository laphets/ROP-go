package form

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"rop/pkg/errno"
	. "rop/handler"
	"rop/model"
)

func Create(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		SendResponse(c, errno.ErrParam, err)
		return
	}
	formId := uint(id)
	if can, err := model.CanFormBeEdited(formId); !can {
		SendResponse(c, errno.ErrFormCantEdit, err)
		return
	}
	SendResponse(c, nil, nil)
}