package freshman

import (
	"github.com/gin-gonic/gin"
	"rop/pkg/errno"
	. "rop/handler"
	"strconv"
	"rop/model"
)


func Submit(c *gin.Context) {
	instanceId, err := strconv.ParseUint(c.Query("instanceId"), 10, 64)
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
		Name: "罗文卿",
		//MainStage: "Public Sea",
		//SubStage: "None",
		OtherInfo: "{a json here}",
	}

	if err := freshman.Create(); err != nil {
		SendResponse(c, errno.DBError, err)
		return
	}

	intent1 := &model.IntentModel{
		FreshmanId: freshman.ID,
		Department: "技术研发中心",
		//GroupId: 0,
		MainStage: "Public Sea",
		SubStage: 1,
	}

	intent2 := &model.IntentModel{
		FreshmanId: freshman.ID,
		Department: "人力资源部门",
		//GroupId: 0,
		MainStage: "Public Sea",
		SubStage: 1,
	}

	intents := []*model.IntentModel{intent1, intent2}

	if err := model.CreateIntents(intents); err != nil {
		SendResponse(c, errno.DBError, err.Error())
		return
	}

	SendResponse(c, nil, nil)
}


//
//res1, err1 := service.SendRejectNotice("18867136212", "博亚", "求是潮2018秋纳", "求是潮")
//if err1 != nil {
//	SendResponse(c, errno.ErrSMS, err.Error())
//	return
//}
//SendResponse(c, nil, res1)