package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/skeyic/neuron/config"
	"github.com/skeyic/neuron/router"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title The Neuron Server
// @version 1.0
// @description The Neuron REST API

// @host
// @BasePath /

func main() {
	flag.Parse()

	r := router.InitRouter()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	glog.V(4).Info("Neuron Server starts...")
	glog.V(4).Infof("CONFIG: %+V", config.Config)
	glog.Fatal(r.Run(fmt.Sprintf(":%d", config.Config.Port)))
}
