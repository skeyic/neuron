package router

import (
	"github.com/gin-gonic/gin"
	"github.com/skeyic/neuron/app/control"
	"github.com/skeyic/neuron/config"
	_ "github.com/skeyic/neuron/docs"
)

// InitRouter ...
func InitRouter() *gin.Engine {
	if config.Config.DebugMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	// Basic
	r.GET("/", control.Index)

	// User
	user := r.Group("/users")
	{
		user.POST("", control.NewUser)
		user.GET("", control.GetUsers)
		user.GET(":id", control.GetUser)
		user.DELETE(":id", control.NotFinished)
		user.POST(":id/send", control.NotFinished)

		bark := user.Group(":id/bark")
		{
			bark.POST("", control.NotFinished)
			bark.GET("", control.NotFinished)
			bark.GET(":id", control.NotFinished)
			bark.DELETE(":id", control.NotFinished)
			bark.POST(":id/send", control.NotFinished)
		}
	}

	return r
}
