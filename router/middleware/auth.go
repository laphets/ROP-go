package middleware

import (
	"github.com/gin-gonic/gin"
	"rop/pkg/token"
	"rop/handler"
	"rop/pkg/errno"
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