package freshman

import (
	"github.com/gin-gonic/gin"
	"rop/pkg/errno"
	. "rop/handler"
	"strconv"
	"rop/model"
)

func Submit(c *gin.Context) {
	instanceId, err := strconv.ParseUint(c.Param("instanceId"), 10, 64)
	if err != nil {
		SendResponse(c, errno.ErrParam, err)
		return
	}

	if _, err := model.GetInstanceById(uint(instanceId)); err != nil {
		SendResponse(c, errno.ErrInstanceNotFound, nil)
		return
	}

	req := &SubmitRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		SendResponse(c, errno.ErrBind, err.Error())
		return
	}

	freshman := &model.FreshmanModel{
		InstanceId: uint(instanceId),
		ZJUid: "3170111705",
		Mobile: "18888922004",
		MainStage: "Public Sea",
		SubStage: "None",
		OtherInfo: "{a json here}",
	}

	if err := freshman.Create(); err != nil {
		SendResponse(c, errno.DBError, err)
		return
	}
	SendResponse(c, nil, nil)
}

//_, err = service.SendRecruitTime("18867136212", "博亚", "一面", "求是潮2018秋纳", "https://rop.zjuqsc.com/7643ghrydst63teayd7")
//if err != nil {
//	SendResponse(c, errno.ErrSMS, err.Error())
//	return
//}
//
//res1, err1 := service.SendRejectNotice("18867136212", "博亚", "求是潮2018秋纳", "求是潮")
//if err1 != nil {
//	SendResponse(c, errno.ErrSMS, err.Error())
//	return
//}
//SendResponse(c, nil, res1)