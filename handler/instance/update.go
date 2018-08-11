package instance

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"rop/model"
	. "rop/handler"
	"rop/pkg/errno"
)

func Update(c *gin.Context) {
	instanceId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		SendResponse(c, errno.ErrParam, err)
		return
	}

	ins := model.InstanceModel{}
	if err := c.ShouldBindJSON(&ins); err != nil {
		SendResponse(c, errno.ErrBind, err)
		return
	}

	ins.ID = uint(instanceId)

	if err := ins.Update(); err!= nil {
		SendResponse(c, errno.DBError, err.Error())
		return
	}
	SendResponse(c, nil, nil)
}