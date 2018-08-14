package router

import (
	"github.com/gin-gonic/gin"
	"rop/router/middleware"
	"net/http"
	"rop/handler/sd"
	"rop/handler/user"
	"rop/handler/instance"
	"rop/handler/form"
	"rop/handler/freshman"
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

	userGroup := g.Group("/v1/user")
	userGroup.Use(middleware.AuthMiddleware())
	{
		userGroup.GET("/info", user.Info)
	}

	insGroup := g.Group("/v1/instance")
	insGroup.Use(middleware.AuthMiddleware())
	{
		insGroup.POST("", instance.Create)
		insGroup.GET("", instance.List)
		insGroup.PUT("/:id", instance.Update)
	}

	formGroup := g.Group("/v1/form")
	insGroup.Use(middleware.AuthMiddleware())
	{
		formGroup.POST("", form.Create)
		formGroup.GET("", form.List)
		formGroup.PUT("/:id", form.Update)
	}

	freGroup := g.Group("/v1/freshman")
	freGroup.Use(middleware.AuthMiddleware())
	{
		freGroup.POST("/submit/:instanceId", freshman.Submit)
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