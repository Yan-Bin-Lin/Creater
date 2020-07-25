package router

import (
	"app/middleware"
	"app/serve"
	"app/setting"
	"github.com/gin-gonic/gin"
)

func AccountRouter() *gin.Engine {
	r := gin.New()
	gin.SetMode(setting.Servers["account"].RunMode)

	r.Use(middleware.Logging())
	r.Use(middleware.ErrorHandle())

	r.LoadHTMLGlob("view/html/*/*")

	r.POST("/login", serve.Login)

	return r
}
