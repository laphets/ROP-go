package instance

import (
	"github.com/gin-gonic/gin"
	"rop/model"
	. "rop/handler"
	"rop/pkg/errno"
)

func List(c *gin.Context) {
	instances, err := model.ListInstance()
	if err != nil {
		SendResponse(c, errno.DBError, err.Error())
	}
	SendResponse(c, nil, instances)
}