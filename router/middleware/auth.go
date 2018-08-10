package middleware

import (
	"github.com/gin-gonic/gin"
	"rop/pkg/token"
	"rop/handler"
	"rop/pkg/errno"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := token.ParseRequest(c); err != nil {
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}