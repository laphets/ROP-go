package user

import (
	"git.zjuqsc.com/rop/ROP-go/pkg/errno"
	"git.zjuqsc.com/rop/ROP-go/service"
	"github.com/gin-gonic/gin"
	. "git.zjuqsc.com/rop/ROP-go/handler"
)

func GetSMSAccount(c *gin.Context) {
	account, err := service.GetAccountInfo()
	if err != nil {
		SendResponse(c, errno.ErrSMS, err.Error())
		return
	}
	SendResponse(c, nil, account)
}