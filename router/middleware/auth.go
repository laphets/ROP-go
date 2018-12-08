package middleware

import (
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
		c.Next()
	}
}