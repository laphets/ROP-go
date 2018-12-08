package intent

import (
	"github.com/gin-gonic/gin"
	"strconv"
	. "git.zjuqsc.com/rop/ROP-go/handler"
	"git.zjuqsc.com/rop/ROP-go/pkg/errno"
	"git.zjuqsc.com/rop/ROP-go/model"
	"git.zjuqsc.com/rop/ROP-go/service"
)

// Reject an intent by force
func Reject(c *gin.Context)  {
	intentId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		SendResponse(c, errno.ErrParam, err.Error())
		return
	}

	fulIntent, err := model.GetFullIntentByID(uint(intentId))
	if err != nil {
		SendResponse(c, errno.DBError, err.Error())
		return
	}
	if err := model.DeleteIntent(uint(intentId)); err != nil {
		SendResponse(c, errno.DBError, err.Error())
		return
	}
	service.SendRejectNotice(fulIntent.Mobile, fulIntent.Name, fulIntent.Department, "求是潮")
	SendResponse(c, nil, nil)
}