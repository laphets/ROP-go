package router

import (
	"github.com/gin-gonic/gin"
	"rop/router/middleware"
	"net/http"
	"rop/handler/sd"
	"github.com/spf13/pflag"
)

var (
	cfg = pflag.StringP("config", "c", "", "ROP config path.")
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	pflag.Parse()

	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)

	g.Use(mw...)

	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}