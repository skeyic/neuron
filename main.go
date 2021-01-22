package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/skeyic/neuron/config"
	"github.com/skeyic/neuron/router"
	"github.com/skeyic/neuron/startup"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title The Neuron Server
// @version 1.0
// @description The Neuron REST API

// @host
// @BasePath /

func main() {
	var (
		err error
	)

	flag.Parse()

	err = startup.StartUp()
	if err != nil {
		glog.Error("failed to startup, err: %v", err)
		panic(err)
	}

	r := router.InitRouter()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	glog.V(4).Info("Neuron Server starts...")
	glog.Fatal(r.Run(fmt.Sprintf(":%d", config.Config.Port)))
}
