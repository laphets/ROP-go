package form

import (
	"github.com/gin-gonic/gin"
	. "git.zjuqsc.com/rop/ROP-go/handler"
	"git.zjuqsc.com/rop/ROP-go/pkg/errno"
	"git.zjuqsc.com/rop/ROP-go/model"
	json2 "encoding/json"
)

func GetRoot(formTemplate []*DataItem) int {
	vis := make(map[int]int)
	for _, item := range formTemplate {
		vis[item.Next] = 1
		if item.Choices != nil {
			for _, choice := range item.Choices {
				vis[choice.Next] = 1
			}
		}
	}
	for _, item := range formTemplate {
		if vis[item.Tag] == 0 {
			return item.Tag
		}
	}
	return -1
}

func Create(c *gin.Context) {
	req := CreateRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		SendResponse(c, errno.ErrBind, err.Error())
		return
	}

	json, err := json2.Marshal(req.Data)
	if err != nil {
		SendResponse(c, errno.ErrParam, nil)
		return
	}

	form := model.FormModel{
		Name:req.Name,
		RootTag: GetRoot(req.Data),
		Data:string(json),
	}

	if err := form.Create(); err != nil {
		SendResponse(c, errno.DBError, err.Error())
		return
	}
	SendResponse(c, nil, string(json))
}