package router

import (
	"app/middleware"
	"app/serve"
	"app/setting"
	//ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	//"go.uber.org/zap"
	//"time"
	//"net/http"
	//"echotest.com/app/middleware"
)

func MainRouter() *gin.Engine {
	r := gin.New()
	gin.SetMode(setting.Servers["main"].RunMode)

	r.Use(middleware.Logging())
	r.Use(middleware.ErrorHandle())

	r.LoadHTMLGlob("view/html/**/*")

	/*home page*/
	r.GET("/", serve.GetRoot)

	/*owner*/
	r.GET("/:owner", serve.GetOwner)

	/*works*/
	r.GET("/:owner/*work", serve.GetBlog)

	/*auth*/
	owner := r.Group("/:owner")
	{
		owner.Use(middleware.Auth())
		owner.POST("", serve.CreateOwner)
		owner.PUT("", serve.UpdateOwner)
		owner.DELETE("", serve.DelOwner)
		work := owner.Group("/*work")
		{
			work.POST("", serve.CreateBlog)
			work.PUT("", serve.UpdateBlog)
			work.DELETE("", serve.DelBlog)
		}
	}

	return r
}