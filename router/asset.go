package router

import (
	"app/middleware"
	"app/serve"
	"app/setting"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AssetRouter() *gin.Engine {
	r := gin.New()
	gin.SetMode(setting.Servers["account"].RunMode)

	r.Use(middleware.Logging())
	r.Use(middleware.ErrorHandle())

	// serve js and css file
	r.StaticFS("/static", http.Dir("./view"))

	r.GET("/file/:type/:oid/:name", serve.GetFile)

	r.Use(middleware.Auth())
	r.POST("/file/:type/:oid/:name", serve.UploadFile)

	return r
}
