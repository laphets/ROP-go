package ssr

import (
	"github.com/gin-gonic/gin"
	. "rop/handler"
	"rop/service"
	"rop/pkg/errno"
	"rop/model"
	"strconv"
)

func Schedule(c *gin.Context)  {
	uid := c.DefaultQuery("uid", "")
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

	intents, err := model.ListIntentByFreshman(uint(freshmanId))
	if err != nil {
		SendResponse(c ,errno.DBError, err)
		return
	}
	freshman, err := model.GetFreshmanById(uint(freshmanId))
	if err != nil {
		SendResponse(c ,errno.DBError, err)
		return
	}

	intentGroups := make([]*IntentGroup, 0)

	for _, intent := range intents {

		if intent.SubStage != 2 {
			continue
		}

		if intent.TargetInterviewId == 0 {
			// can select any(auto_joinable)
			interviews, err := model.ListFulInterview(freshman.InstanceId, &model.InterviewModel{
				InterviewType: service.StateInNum(intent.MainStage)+1,
				Department: intent.Department,
				AutoJoinable: 1,
			})
			if err != nil {
				continue
			}
			intentGroup := &IntentGroup{
				IntentModel: intent,
				ChineseStage: service.StateInChinese(service.NextState(intent.MainStage)),
				Interviews:interviews,
			}
			intentGroups = append(intentGroups, intentGroup)
		} else {
			// can select only one
			interview, err := model.GetFulInterviewByID(intent.TargetInterviewId)
			if err != nil {
				continue
			}
			interviewArray := make([]*model.FullInterview, 0)
			interviewArray = append(interviewArray, interview)
			intentGroup := &IntentGroup{
				IntentModel: intent,
				ChineseStage: service.StateInChinese(service.NextState(intent.MainStage)),
				Interviews: interviewArray,
			}
			intentGroups = append(intentGroups, intentGroup)
		}
	}

	scheduleRes := &ScheduleResponse{
		FreshmanModel: freshman,
		IntentGroups: intentGroups,
	}

	SendResponse(c, nil, scheduleRes)
}