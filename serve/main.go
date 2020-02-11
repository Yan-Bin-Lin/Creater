package serve

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

/*home page*/

// get root
func GetRoot(c *gin.Context) {

}

/*owner*/

func GetOwner(c *gin.Context) {

}

/*work*/

func GetProject(c *gin.Context) {
	projs, blog := splitProject(c.Param("project"))
	c.JSON(http.StatusOK, gin.H{
		"owner" : c.Param("owner"),
		"project" : projs,
		"blog" : blog,
	})
}

// split "/*project" to slice project and last blog
func splitProject(url string) ([]string, string) {
	projects := strings.Split(url, "/")
	return projects[1:len(projects) - 1], projects[len(projects) - 1]
}