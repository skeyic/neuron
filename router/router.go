package router

import (
	"github.com/gin-gonic/gin"
	"github.com/skeyic/neuron/app/control"
	_ "github.com/skeyic/neuron/docs"
)

// InitRouter ...
func InitRouter() *gin.Engine {
	//if config.Config.DebugMode {
	//	gin.SetMode(gin.DebugMode)
	//} else {
	//	gin.SetMode(gin.ReleaseMode)
	//}

	r := gin.Default()

	// Basic
	r.GET("/", control.Index)

	return r
}
