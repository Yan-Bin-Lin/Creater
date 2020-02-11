package router

import (
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
	/*
		logger, _ := zap.NewDevelopment(zap.AddCaller())

		r.Use(middleware.Ginzap(logger, time.RFC3339, true))
		r.Use(middleware.RecoveryWithZap(logger, true))

		r.Use(middleware.ErrorHandling())

		r.LoadHTMLGlob("view/html/*")
	*/

	gin.SetMode(setting.Servers["main"].RunMode)

	/*home page*/
	r.GET("/", serve.GetRoot)

	/*owner*/
	r.GET("/:owner", serve.GetOwner)

	/*works*/
	r.GET("/:owner/*project", serve.GetProject)

	/*auth*/
	/*
		owner := r.Group("/:owner")
		{
			owner.Use(middleware.Auth())
			// post a site, project or blog under owner
			owner.POST("/:site", PostUnderOwner)

			site := owner.Group("/:site")
			{
				// project or blog
				project := site.Group("/:project")
				{
					blog := project.Group("/:blog")
					{
						// blog
						blog.POST("/")
					}
				}
			}

		}
	*/
	return r
}