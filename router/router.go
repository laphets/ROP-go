package router

import (
	"github.com/gin-gonic/gin"
	"git.zjuqsc.com/rop/ROP-go/router/middleware"
	"net/http"
	"git.zjuqsc.com/rop/ROP-go/handler/sd"
	"git.zjuqsc.com/rop/ROP-go/handler/user"
	"git.zjuqsc.com/rop/ROP-go/handler/instance"
	"git.zjuqsc.com/rop/ROP-go/handler/form"
	"git.zjuqsc.com/rop/ROP-go/handler/freshman"
	"git.zjuqsc.com/rop/ROP-go/handler/interview"
	"git.zjuqsc.com/rop/ROP-go/handler/intent"
	"git.zjuqsc.com/rop/ROP-go/handler/association"
	"git.zjuqsc.com/rop/ROP-go/handler/ssr"
	"git.zjuqsc.com/rop/ROP-go/handler/file"
)


func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {

	//g.MaxMultipartMemory = 8 << 20

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
		userGroup.POST("/avatar", user.Avatar)
		userGroup.GET("/sms", user.GetSMSAccount)
	}

	insGroup := g.Group("/v1/instance")
	insGroup.Use(middleware.AuthMiddleware())
	{
		insGroup.POST("", instance.Create)
		insGroup.GET("", instance.List)
		insGroup.PUT("/:id", instance.Update)
	}

	formGroup := g.Group("/v1/form")
	formGroup.Use(middleware.AuthMiddleware())
	{
		formGroup.POST("", form.Create)
		formGroup.GET("", form.List)
		formGroup.PUT("/:id", form.Update)
	}

	freGroup := g.Group("/v1/freshman")
	//freGroup.Use(middleware.AuthMiddleware())
	{
		freGroup.POST("/submit", freshman.Submit)
	}

	intentGroup := g.Group("/v1/intent")
	intentGroup.Use(middleware.AuthMiddleware())
	{
		intentGroup.POST("/assign", intent.Assign)
		intentGroup.POST("/reject/:id", intent.Reject)
		intentGroup.GET("", intent.List)
	}

	interviewGroup := g.Group("/v1/interview")
	interviewGroup.Use(middleware.AuthMiddleware())
	{
		interviewGroup.POST("", interview.Create)
		interviewGroup.PUT("/:id", interview.Update)
		interviewGroup.GET("", interview.List)
	}

	associationGroup := g.Group("/v1/association")
	associationGroup.Use(middleware.AuthMiddleware())
	{
		associationGroup.POST("", association.Create)
		associationGroup.GET("", association.Get)
		associationGroup.GET("/user", association.GetUser)
		associationGroup.POST("/notice", association.SendNotice)
	}

	ssrGroup := g.Group("/v1/ssr")
	{
		ssrGroup.GET("/schedule", ssr.Schedule)
		ssrGroup.POST("/join/:id", interview.Join)
		ssrGroup.GET("/form", ssr.GetFormByIns)
		ssrGroup.POST("/reject/:id", intent.Cancel)
		ssrGroup.GET("/instance", instance.List)
	}

	fileGroup := g.Group("/v1/file")
	{
		fileGroup.POST("/upload/img", file.UploadImage)
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