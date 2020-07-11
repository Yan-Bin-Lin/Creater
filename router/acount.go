package router

import (
	"app/middleware"
	"app/serve"
	"app/setting"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func AccountRouter() *gin.Engine {
	r := gin.New()
	gin.SetMode(setting.Servers["account"].RunMode)

	r.Use(middleware.Logging())
	r.Use(middleware.ErrorHandle())

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{fmt.Sprintf("http://%s:%d", setting.Servers["main"].Host, setting.Servers["main"].Port)}
	config.AllowCredentials = true
	r.Use(cors.New(config))

	r.LoadHTMLGlob("view/html/*/*")

	r.POST("/login", serve.Login)
	//r.POST("/user/:user/*owner", serve.Login)

	return r
}
