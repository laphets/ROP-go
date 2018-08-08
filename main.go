package main

import (
	"github.com/gin-gonic/gin"
	"rop/router"
	"net/http"
	"log"
	"time"
	"errors"
	"rop/config"
	"github.com/spf13/viper"
)

//var (
//	cfg = pflag.StringP("config", "c", "", "ROP config file path.")
//)

func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		log.Print("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}

func main() {
	//pflag.Parse()
	if err := config.Init(""); err != nil {
		panic(err)
	}

	gin.SetMode(viper.GetString("runmode"))
	g := gin.New()

	middlewares := []gin.HandlerFunc{}

	router.Load(
		g,
		middlewares...,
	)

	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Print("The router has been deployed successfully.")
	}()

	log.Printf("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	log.Printf(http.ListenAndServe(viper.GetString("addr"), g).Error())
}