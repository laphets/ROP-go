package association

import (
	. "git.zjuqsc.com/rop/ROP-go/handler"
	"git.zjuqsc.com/rop/ROP-go/model"
	"git.zjuqsc.com/rop/ROP-go/pkg/errno"
	"github.com/gin-gonic/gin"
)

func ListAssociaton(c *gin.Context) {
	//userId := c.GetInt("UserId")
	//user, err := model.GetUserById(uint(userId))
	//if err != nil {
	//	SendResponse(c, errno.ErrUserNotFound, err.Error())
	//	return
	//}
	associationList, err := model.ListAssociation()
	if err != nil {
		SendResponse(c, errno.DBError, err.Error())
		return
	}
	SendResponse(c, nil, associationList)
}