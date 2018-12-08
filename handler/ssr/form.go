package ssr

import (
	"github.com/gin-gonic/gin"
	"git.zjuqsc.com/rop/ROP-go/pkg/errno"
	"strconv"
	. "git.zjuqsc.com/rop/ROP-go/handler"
	"git.zjuqsc.com/rop/ROP-go/model"
)

func GetFormByIns(c *gin.Context) {
	//res, err := service.Encrypt("1")
	//log.Debug(res)
	//uid := c.DefaultQuery("uid", "")
	//if uid == "" {
	//	SendResponse(c, errno.ErrParam, "param is null")
	//	return
	//}
	//instanceIdString, err := service.Decrypt(uid)
	//if err != nil {
	//	SendResponse(c, errno.ErrParam, err)
	//	return
	//}

	//instanceId, err := strconv.ParseUint(instanceIdString, 10, 64)
	instanceId, err := strconv.ParseUint(c.Query("instanceId"), 10, 64)

	if err != nil {
		SendResponse(c, errno.ErrParam, err)
		return
	}

	instance, err := model.GetInstanceById(uint(instanceId))
	if err != nil {
		SendResponse(c, errno.DBError, err)
		return
	}

	form, err := model.GetFormByID(instance.FormId)
	if err != nil {
		SendResponse(c, errno.DBError, err)
		return
	}
	SendResponse(c, nil ,form)
}