package intent

import (
	"github.com/gin-gonic/gin"
	"strconv"
	. "git.zjuqsc.com/rop/ROP-go/handler"
	"git.zjuqsc.com/rop/ROP-go/pkg/errno"
	"git.zjuqsc.com/rop/ROP-go/model"
)

func List(c *gin.Context) {
	instanceId, err := strconv.ParseUint(c.Query("instanceId"), 10, 64)
	if err != nil {
		SendResponse(c, errno.ErrInstanceNotSelected, err)
		return
	}


	if _, err := model.GetInstanceById(uint(instanceId)); err != nil {
		SendResponse(c, errno.ErrInstanceNotFound, err)
		return
	}

	mainStage := c.DefaultQuery("mainStage", "")
	department := c.DefaultQuery("department", "")

	fullIntents, err := model.ListFullIntentByInstance(uint(instanceId), mainStage, department)

	if err != nil {
		SendResponse(c, errno.DBError, err)
		return
	}

	SendResponse(c, nil, fullIntents)
}