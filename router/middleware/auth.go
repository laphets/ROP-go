package middleware

import (
	"git.zjuqsc.com/rop/ROP-go/model"
	"github.com/gin-gonic/gin"
	"git.zjuqsc.com/rop/ROP-go/pkg/token"
	"git.zjuqsc.com/rop/ROP-go/handler"
	"git.zjuqsc.com/rop/ROP-go/pkg/errno"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		JWTpayload, err := token.ParseRequest(c)
		if err != nil {
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}
		c.Set("ZJUid", JWTpayload.ZJUid)

		// Update last seen time
		go model.UpdateLastSeen(JWTpayload.ZJUid)

		// Log begin here
		log4Save := model.LogModel{
			ZJUid: JWTpayload.ZJUid,
			IP: c.ClientIP(),
			URL: c.Request.RequestURI,
			UA: c.GetHeader("User-Agent"),
		}
		go log4Save.Create()

		c.Next()
	}
}