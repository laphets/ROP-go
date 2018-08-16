package intent

import (
	"github.com/gin-gonic/gin"
	. "rop/handler"
	"rop/pkg/errno"
	"rop/model"
	"rop/service"
	"fmt"
)

func Assign(c *gin.Context) {
	req := &AssignRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		SendResponse(c, errno.ErrBind, err.Error())
		return
	}

	if req.AssignMode == "auto" || req.AssignMode == "manual" {
		for _, intentId := range req.Intents {
			fulIntent, err := model.GetFullIntentByID(intentId)
			if err != nil {
				SendResponse(c, errno.DBError, err.Error())
				return
			}

			instance, err := model.GetInstanceById(fulIntent.InstanceId)
			if err != nil {
				SendResponse(c, errno.DBError, err.Error())
				return
			}

			if req.AssignMode == "manual" {
				// Manual assign autojoinable should be -1
				targetInterview, err := model.GetInterviewByID(req.TargetInterviewId)
				if err != nil {
					SendResponse(c, errno.DBError, err)
					return
				}
				if targetInterview.AutoJoinable != -1 {
					SendResponse(c, errno.ErrOperation, "This interview can't be joined.")
					return
				}
				intent := &model.IntentModel{
					ID: intentId,
					SubStage:1,
					TargetInterviewId:targetInterview.ID,
				}
				if err := intent.Update(); err != nil {
					SendResponse(c, errno.DBError, err)
					return
				}
			} else {
				// Auto assign autojoinable should be 1
				intent := &model.IntentModel{
					ID:intentId,
					SubStage:1,
					//TargetInterviewId: 0,
				}
				if err := intent.Update(); err != nil {
					SendResponse(c, errno.DBError, err)
					return
				}
			}

			// Send SMS
			_, err = service.SendRecruitTime(fulIntent.Mobile, fulIntent.Name, fulIntent.Department+service.NextState(fulIntent.MainStage), instance.Name, fmt.Sprintf("https://rop.zjuqsc.com/intent?id=%d",intentId))
			if err != nil {
				SendResponse(c, errno.ErrSMS, err.Error())
				return
			}
		}
	} else {
		SendResponse(c, errno.ErrOperation, "Unsupported operation.")
		return
	}

	SendResponse(c, nil, nil)
}