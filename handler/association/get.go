package association

import (
	"github.com/gin-gonic/gin"
	"git.zjuqsc.com/rop/ROP-go/model"
	. "git.zjuqsc.com/rop/ROP-go/handler"
	"git.zjuqsc.com/rop/ROP-go/pkg/errno"
	"strings"
)

func Get(c *gin.Context) {
	//associationName := c.Query("associationName")
	associationName := c.Param("name")
	association, err := model.GetAssociationByName(associationName)

	if err != nil {
		SendResponse(c, errno.DBError, err)
		return
	}
	res := &GetResponse{
		AssociationModel: association,
		Department: strings.Split(association.DepartmentList, "&"),
	}
	SendResponse(c, nil, res)
}