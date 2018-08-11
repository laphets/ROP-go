package freshman

import (
	"github.com/gin-gonic/gin"
	"rop/pkg/errno"
	. "rop/handler"
	"strconv"
	"github.com/lexkong/log"
	"rop/service"
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

	_, err = service.SendRecruitTime("18888922004", "罗文卿", "一面", "求是潮2018秋纳测试", "https://rop.zjuqsc.com/7643ghrydst63teayd7")
	if err != nil {
		SendResponse(c, errno.ErrSMS, err.Error())
		return
	}

	res1, err1 := service.SendRejectNotice("18888922004", "罗文卿", "求是潮2018秋纳", "求是潮")
	if err1 != nil {
		SendResponse(c, errno.ErrSMS, err.Error())
		return
	}
	SendResponse(c, nil, res1)
}