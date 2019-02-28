package interview

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"git.zjuqsc.com/rop/ROP-go/pkg/errno"
	. "git.zjuqsc.com/rop/ROP-go/handler"
	"git.zjuqsc.com/rop/ROP-go/model"
	"git.zjuqsc.com/rop/ROP-go/service"
)

func Join(c *gin.Context) {
	interviewId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		SendResponse(c, errno.ErrParam, err)
		return
	}

	req := &JoinRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		SendResponse(c, errno.ErrBind, err.Error())
		return
	}

	uid := req.Uid

	if uid == "" {
		SendResponse(c, errno.ErrParam, "param is null")
		return
	}
	freshmanIdString, err := service.Decrypt(uid)
	if err != nil {
		SendResponse(c, errno.ErrParam, err)
		return
	}

	freshmanId, err := strconv.ParseUint(freshmanIdString, 10, 64)
	if err != nil {
		SendResponse(c, errno.ErrParam, err)
		return
	}

	freshmanInfo ,err := model.GetFreshmanById(uint(freshmanId));
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, err.Error())
		return
	}

	// Check for interview
	// TODO: Del this extra query
	interview, err := model.GetInterviewByID(uint(interviewId))
	if err != nil {
		SendResponse(c, errno.DBError, err.Error())
		return
	}
	if interview.NotAvailable == 1 {
		SendResponse(c, errno.ErrInterviewNotAble, "This interview is not available")
		return
	}


	fulInterview, err := model.GetFulInterviewByID(uint(interviewId))
	if err != nil {
		SendResponse(c, errno.DBError, err.Error())
		return
	}

	if len(fulInterview.Participants) + len(req.Intents) > fulInterview.Capacity {
		SendResponse(c, errno.ErrInterviewFull, nil)
		return
	}

	instance, err := model.GetInstanceById(interview.InstanceId)
	if err != nil {
		SendResponse(c, errno.DBError, err.Error())
		return
	}

	intent := &model.IntentModel{}
	for _, item := range req.Intents {

		fulIntent, err := model.GetFullIntentByID(item)
		if err != nil {
			SendResponse(c, errno.DBError, err.Error())
			return
		}

		if fulIntent.InstanceId != interview.InstanceId || fulIntent.Department != interview.Department || service.StateInNum(service.NextState(fulIntent.MainStage)) != interview.InterviewType {
			//log.Debugf("%s||%s", fulIntent.Department, interview.Department)
			SendResponse(c, errno.ErrOperation, "Instance or Department or stage Not Match")
			return
		}
		//log.Debugf("%+v", fulIntent.IntentModel)
		//log.Debugf("%d", item)
		intent.ID = item
		intent.InterviewId = uint(interviewId)
		intent.MainStage = service.NextState(fulIntent.MainStage)
		intent.SubStage = 1
		intent.TargetInterviewId = 0
		if err := intent.Update(); err != nil {
			SendResponse(c, errno.DBError, err.Error())
			return
		}
		go service.SendInterviewConfirm(freshmanInfo.Mobile, freshmanInfo.Name, intent.Department, service.StateInChinese(intent.MainStage), interview.StartTime.Format("2006-01-02 15:04"), interview.Location, instance.Name)
	}
	SendResponse(c, nil, nil)
}