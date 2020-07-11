package router

import (
	"app/middleware"
	"app/serve"
	"app/setting"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AssetRouter() *gin.Engine {
	r := gin.New()
	gin.SetMode(setting.Servers["account"].RunMode)

	r.Use(middleware.Logging())
	r.Use(middleware.ErrorHandle())

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{fmt.Sprintf("http://%s:%d", setting.Servers["main"].Host, setting.Servers["main"].Port)}
	config.AllowCredentials = true
	r.Use(cors.New(config))

	// serve js and css file
	r.StaticFS("/static", http.Dir("./view"))

	r.GET("/file/:type/:oid/:name", serve.GetFile)

	r.Use(middleware.Auth())
	r.POST("/file/:type/:oid/:name", serve.UploadFile)

	return r
}
