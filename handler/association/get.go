package association

import (
	"github.com/gin-gonic/gin"
	"rop/model"
	. "rop/handler"
	"rop/pkg/errno"
	"strings"
)

func Get(c *gin.Context) {
	associationName := c.Query("associationName")
	association, err := model.GetAssociationByName(associationName)
	if err != nil {
		SendResponse(c, errno.DBError, err)
		return
	}
	res := &GetResponse{
		Department: strings.Split(association.DepartmentList, "&"),
	}
	SendResponse(c, nil, res)
}