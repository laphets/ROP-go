package user

import (
	"github.com/gin-gonic/gin"
	. "rop/handler"
	"rop/pkg/errno"
	"net/http"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"github.com/lexkong/log"
	"github.com/buger/jsonparser"
	"rop/model"
	"rop/pkg/token"
)


func Login(c *gin.Context) {
	qscCookie, err := c.Cookie("qp2gl_sesstok")
	if err != nil {
		SendResponse(c, errno.NoEnoughAuth, nil)
		log.Debugf(err.Error())
		return
	}
	//fmt.Print("https://api.zjuqsc.com/passport/get_member_by_token?appid=%s&appsecret=%s&token=%s",viper.GetString("passport.appid"), viper.GetString("passport.appsecret"), qscCookie)
	resp, err := http.Get(fmt.Sprintf("https://api.zjuqsc.com/passport/get_member_by_token?appid=%s&appsecret=%s&token=%s",viper.GetString("passport.appid"), viper.GetString("passport.appsecret"), qscCookie))
	if err != nil {
		SendResponse(c, errno.RemoteError, nil)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		SendResponse(c, errno.RemoteError, nil)
		return
	}

	ZJUid := ""
	if value, err := jsonparser.GetString(body, "stuid"); err == nil {
		ZJUid = value
	}
	if ZJUid == "" {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}


	resp, err = http.Get(fmt.Sprintf("https://hr.zjuqsc.com/api/get_info_by_stuid?appid=%s&appsecret=%s&stuid=%s",viper.GetString("hr.appid"), viper.GetString("hr.appsecret"), ZJUid))
	if err != nil {
		SendResponse(c, errno.RemoteError, nil)
		return
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		SendResponse(c, errno.RemoteError, nil)
		return
	}

	innerId, err := jsonparser.GetString(body, "info", "bbs")
	name, err := jsonparser.GetString(body, "info", "name")
	department, err := jsonparser.GetString(body, "info", "department")
	position, err := jsonparser.GetString(body, "info", "position")

	existing, err := model.GetUserByZJUid(ZJUid)

	if err != nil {
		u := model.UserModel{
			ZJUid:ZJUid,
			InnerId:innerId,
			Name:name,
			Department:department,
			Position:position,
		}
		if err := u.Create(); err != nil {
			log.Debugf(err.Error())
			SendResponse(c, errno.DBError, nil)
			return
		}
	} else {
		existing.InnerId = innerId
		existing.Name = name
		existing.Department = department
		existing.Position = position
		if err := existing.Update(); err != nil {
			log.Debugf(err.Error())
			SendResponse(c, errno.DBError, nil)
			return
		}
	}
	JWT, err := token.Sign(token.Context{ZJUid:ZJUid}, "")
	if err != nil {
		SendResponse(c, errno.ErrToken, nil)
		return
	}
	SendResponse(c, nil, JWT)
}