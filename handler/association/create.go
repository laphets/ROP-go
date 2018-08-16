package association

import (
	"github.com/gin-gonic/gin"
	. "rop/handler"
	"rop/pkg/errno"
	"rop/model"
)

func Create(c *gin.Context) {
	req := CreateRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		SendResponse(c, errno.ErrBind, err.Error())
		return
	}

	if len(req.Department) == 0 {
		SendResponse(c, errno.ErrBind, "department not fit")
		return
	}

	departList := ""
	for _, dep := range req.Department {
		departList += dep+"&"
	}

	association := model.AssociationModel{
		Name: req.Name,
		DepartmentList:departList,
	}

	if err := association.Create(); err != nil {
		SendResponse(c, errno.DBError, err.Error())
		return
	}
	SendResponse(c, nil, nil)
}