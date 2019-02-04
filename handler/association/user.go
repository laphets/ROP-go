package association

import (
	. "git.zjuqsc.com/rop/ROP-go/handler"
	"git.zjuqsc.com/rop/ROP-go/model"
	"git.zjuqsc.com/rop/ROP-go/pkg/errno"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	userList, err := model.ListUser()
	if err != nil {
		SendResponse(c, errno.DBError, err.Error())
		return
	}
	SendResponse(c, nil, userList)
}