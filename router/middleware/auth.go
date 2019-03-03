package middleware

import (
	"git.zjuqsc.com/rop/ROP-go/handler"
	"git.zjuqsc.com/rop/ROP-go/model"
	"git.zjuqsc.com/rop/ROP-go/pkg/errno"
	"git.zjuqsc.com/rop/ROP-go/pkg/token"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		JWTpayload, err := token.ParseRequest(c)
		if err != nil {
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		user, err := model.GetUserById(uint(JWTpayload.UserId))
		if err != nil {
			handler.SendResponse(c, errno.ErrUserNotFound, err.Error())
			c.Abort()
			return
		}

		c.Set("UserId", JWTpayload.UserId)
		c.Set("UserName", user.Name)
		c.Set("ZJUid", user.ZJUid)
		c.Set("AssociationId", int(user.AssociationId))
		c.Set("AdminLevel", int(user.AdminLevel))

		// Update last seen time
		go model.UpdateLastSeen(uint(JWTpayload.UserId))

		// Log begin here
		log4Save := model.LogModel{
			UserId: uint(JWTpayload.UserId),
			IP: c.ClientIP(),
			URL: c.Request.RequestURI,
			UA: c.GetHeader("User-Agent"),
		}
		go log4Save.Create()

		c.Next()
	}
}