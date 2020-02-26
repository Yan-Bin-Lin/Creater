package router

import (
	"app/log"
	"app/middleware"
	"app/serve"
	"app/setting"
	"app/util/file"
	"app/util/hash"

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

	r.LoadHTMLGlob("view/html/*")

	/*home page*/
	r.GET("/", serve.GetRoot)

	/*owner*/
	r.GET("/:owner", serve.GetOwner)

	/*works*/
	r.GET("/:owner/*work", GetWork)

	/*auth*/
	owner := r.Group("/:owner")
	{
		// owner.Use(middleware.Auth())
		owner.PUT("", serve.PutOwner)
		owner.DELETE("", serve.DelOwner)
		work := owner.Group("/*work")
		{
			work.PUT("", PutWork)
			work.DELETE("", DelWork)
		}
		//owner.POST("/*project", serve.PostProject)
		//owner.PUT("/*work", PutWork)
		//owner.DELETE("/*work", DelWork)
	}

	return r
}

// route to get blog or project
func GetWork(c *gin.Context) {
	dir := hash.GetHashString(c.Param("owner"), c.Param("work"))
	if file.IsExist(setting.Servers["main"].FilePath, dir) {
		serve.GetBlog(c, dir)
	} else {
		serve.GetProject(c, dir)
	}
}

// route to put blog or project
func PutWork(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		log.Warn(c, 2400001, err, "binding error of put multipart form", "binding error of put multipart form")
	}

	// check prom kew exist exist
	if len(form.Value["WorkType"]) == 0 {
		log.Warn(c, 2400001, nil, "multy part form miss match key WorkType")
		return
	}

	if form.Value["WorkType"][0] == "blog" {
		serve.PutBlog(c, form, hash.GetHashString(c.Param("owner"), c.Param("work")))
	} else if form.Value["WorkType"][0] == "project" {
		log.Debug(c.Param("work"))
		serve.PutProject(c, form, hash.GetHashString(c.Param("owner"), c.Param("work")))
	} else {
		log.Warn(c, 2400001, err, "multy part form miss match key WorkType")
	}
}

// route to delete blog or project
func DelWork(c *gin.Context) {
	dir := hash.GetHashString(c.Param("owner"), c.Param("work"))
	if file.IsExist(setting.Servers["main"].FilePath, dir) {
		serve.DelBlog(c, dir)
	} else {
		serve.DelProject(c, dir)
	}
}
