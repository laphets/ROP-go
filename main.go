package main

import (
	"errors"
	"git.zjuqsc.com/rop/ROP-go/config"
	"git.zjuqsc.com/rop/ROP-go/model"
	"git.zjuqsc.com/rop/ROP-go/router"
	"git.zjuqsc.com/rop/ROP-go/router/middleware"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
	"net/http"
	"time"
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

		log.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}

func main() {
	//pflag.Parse()

	if err := config.Init(""); err != nil {
		panic(err)
	}

	model.DB.Init()
	defer model.DB.Close()

	gin.SetMode(viper.GetString("runmode"))
	g := gin.New()

	router.Load(
		g,
		middleware.RequestId(),
	)

	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Info("The router has been deployed successfully.")
	}()

	log.Infof("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	log.Info(http.ListenAndServe(viper.GetString("addr"), g).Error())
}
