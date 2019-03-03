package user

import (
	"git.zjuqsc.com/rop/ROP-go/service"
	"github.com/gin-gonic/gin"
	. "git.zjuqsc.com/rop/ROP-go/handler"
	"git.zjuqsc.com/rop/ROP-go/pkg/errno"
	"net/http"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"github.com/lexkong/log"
	"github.com/buger/jsonparser"
	"git.zjuqsc.com/rop/ROP-go/model"
	"git.zjuqsc.com/rop/ROP-go/pkg/token"
)

func LoginByPassword(c *gin.Context) {
	req := LoginByPasswordRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		SendResponse(c, errno.ErrBind, err.Error())
		return
	}

	association, err := model.GetAssociationById(req.AssociationId)
	if err != nil {
		SendResponse(c, errno.ErrAssociationNotExist, err.Error())
		return
	}

	user, err := model.GetUserByAssociationAndZJUid(association.ID, req.ZJUid)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, err.Error())
		return
	}

	if user.Password == "" {
		SendResponse(c, errno.ErrPasswordNotSet, "Password is not set, try another way to login")
		return
	}

	if user.Password != service.SHAHash(req.Password) {
		SendResponse(c, errno.ErrPasswordWrong, "Password does not match")
		return
	}

	JWT, err := token.Sign(token.Context{UserId:int(user.ID)}, "")
	if err != nil {
		SendResponse(c, errno.ErrToken, nil)
		return
	}

	SendResponse(c, nil, JWT)
}

func LoginByQSC(c *gin.Context) {
	qscCookie, err := c.Cookie("qp2gl_sesstok")
	if err != nil {
		SendResponse(c, errno.NoCookie, err.Error())
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

	// Get Association
	association, err := model.GetAssociationByName("求是潮")
	if err != nil {
		SendResponse(c, errno.ErrAssociationNotExist, err.Error())
		return
	}

	innerId, err := jsonparser.GetString(body, "info", "bbs")
	name, err := jsonparser.GetString(body, "info", "name")
	department, err := jsonparser.GetString(body, "info", "department")
	position, err := jsonparser.GetString(body, "info", "position")
	mobile, err := jsonparser.GetString(body, "info", "mobile")

	existing, err := model.GetUserByAssociationAndZJUid(association.ID, ZJUid)

	var userId uint

	if err != nil {
		u := model.UserModel{
			ZJUid:ZJUid,
			AssociationId:association.ID,
			InnerId:innerId,
			Name:name,
			Department:department,
			Position:position,
			Mobile:mobile,
		}
		if err := u.Create(); err != nil {
			log.Debugf(err.Error())
			SendResponse(c, errno.DBError, nil)
			return
		}
		userId = u.ID
	} else {
		existing.InnerId = innerId
		existing.Name = name
		existing.Department = department
		existing.Position = position
		existing.Mobile = mobile
		if err := existing.Update(); err != nil {
			log.Debugf(err.Error())
			SendResponse(c, errno.DBError, nil)
			return
		}
		userId = existing.ID
	}
	JWT, err := token.Sign(token.Context{UserId:int(userId)}, "")
	if err != nil {
		SendResponse(c, errno.ErrToken, nil)
		return
	}
	SendResponse(c, nil, JWT)
}