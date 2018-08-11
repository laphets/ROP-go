package freshman

import (
	"github.com/gin-gonic/gin"
	"rop/pkg/errno"
	. "rop/handler"
	"strconv"
	"github.com/lexkong/log"
)

func Submit(c *gin.Context) {
	instanceId, err := strconv.ParseUint(c.Param("instanceId"), 10, 64)
	if err != nil {
		SendResponse(c, errno.ErrParam, err)
		return
	}
	log.Debugf("%d", instanceId)
	req := &SubmitRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		SendResponse(c, errno.ErrBind, err.Error())
		return
	}
	SendResponse(c, nil, req)
}