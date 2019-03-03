package association

import (
	"fmt"
	. "git.zjuqsc.com/rop/ROP-go/handler"
	"git.zjuqsc.com/rop/ROP-go/model"
	"git.zjuqsc.com/rop/ROP-go/pkg/errno"
	"git.zjuqsc.com/rop/ROP-go/service"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func SendNotice(c *gin.Context) {
	req := SendNoticeRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		SendResponse(c, errno.ErrBind, err.Error())
		return
	}

	if len(req.UserId) == 0 {
		SendResponse(c, errno.ErrBind, "UserId should be unempty list.")
		return
	}

	success := true
	errorList := make([]string, 0)

	for _, userId := range req.UserId {
		user, err := model.GetUserById(userId)
		if err != nil {
			// Here we should handle error
			success = false
			errorList = append(errorList, strconv.Itoa(int(userId)))
			continue
		}
		if user.Mobile == "" {
			success = false
			errorList = append(errorList, user.Name)
			continue
		}
		if _, err := service.SendInterviewerNotice(user.Mobile, "求是潮", user.Name, user.Department, "纳新"); err != nil {
			success = false
			errorList = append(errorList, user.Name)
			continue
		}
	}

	if !success {
		SendResponse(c, errno.ErrSMS, fmt.Sprintf("Error occurs when sending SMS to : %s", strings.Join(errorList, ",")))
	}

	SendResponse(c, nil, nil)
}