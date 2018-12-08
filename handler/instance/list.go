package instance

import (
	"github.com/gin-gonic/gin"
	"git.zjuqsc.com/rop/ROP-go/model"
	"git.zjuqsc.com/rop/ROP-go/pkg/errno"
	"time"
	"git.zjuqsc.com/rop/ROP-go/pkg/timerange"
	. "git.zjuqsc.com/rop/ROP-go/handler"
)

func List(c *gin.Context) {
	instances, err := model.ListInstance()
	if err != nil {
		SendResponse(c, errno.DBError, err.Error())
		return
	}
	listRes := make([]*ListResponse, 0)
	now := time.Now()

	for _, value := range instances {
		tmp := &ListResponse{InstanceModel: value}
		tmp.Status = timerange.GetStatusSync(now, tmp.StartTime, tmp.EndTime)
		count, err := model.GetFreshmanCountByID(tmp.ID)
		if err != nil {
			SendResponse(c, errno.DBError, err.Error())
			return
		}
		tmp.FreshmanCount = count
		form, err := model.GetFormByID(tmp.FormId)
		if err != nil {
			SendResponse(c, errno.ErrFormNotFound, err)
			return
		}
		tmp.FormName = form.Name
		listRes = append(listRes, tmp)
	}
	SendResponse(c, nil, listRes)
}