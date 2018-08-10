package router

import (
	"github.com/gin-gonic/gin"
	"rop/router/middleware"
	"net/http"
	"rop/handler/sd"
	"rop/handler/user"
	"rop/handler/instance"
)


func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {

	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)

	g.Use(mw...)

	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	auth := g.Group("/v1/auth")
	{
		auth.POST("/login", user.Login)
	}

	u := g.Group("/v1/user")
	u.Use(middleware.AuthMiddleware())
	{
		u.GET("/info", user.Info)
	}

	ins := g.Group("/v1/instance")
	ins.Use(middleware.AuthMiddleware())
	{
		ins.POST("", instance.Create)
		ins.GET("", instance.List)
	}

	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}