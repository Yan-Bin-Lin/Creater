package serve

import (
	"app/database"
	"app/log"
	"errors"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
	"strings"
)

/*home page*/

// get root
func GetRoot(c *gin.Context) {
	log.Error(c, 1500001, errors.New("test error"), 0, "public msg", "private msg")
}

/*owner*/

func GetOwner(c *gin.Context) {
	// check project exist
	ownerData, err := database.GetOwner(c.Param("owner"))
	if err != nil {
		log.Warn(c, 1500001, err, "Sorry, something error", "database error of get owner")
		return
	}

	// blog not found, get project
	c.JSON(http.StatusOK, gin.H{
		"ownerData": ownerData,
	})
}

// put a new owner
func PutOwner(c *gin.Context) {
	if err := database.PutOwner(c.PostForm("oid"), c.PostForm("uid"), c.PostForm("uniquename"), c.PostForm("nickname")); err != nil {
		if err == database.ERR_NAME_CONFLICT {
			log.Warn(c, 2400001, err, "Name conflict of put owner")
		} else if err == database.ERR_TASK_FAIL {
			log.Warn(c, 2400001, err, "put owner fail, please check oid and uid correct")
		} else {
			log.Warn(c, 1500001, err, "Sorry, something error. please try again", "database error of put owner")
		}
	}

	c.Header("Location", "/.")
	c.Status(http.StatusCreated)
}

func DelOwner(c *gin.Context) {
	err := database.DelOwner(c.PostForm("oid"), c.PostForm("uid"), c.PostForm("uniquename"))
	if err != nil {
		if err == database.ERR_TASK_FAIL {
			log.Warn(c, 2400001, err, "delete owner fail, please check oid and owner name correct")
		} else {
			log.Warn(c, 1500001, err, "Sorry, something error. please try again", "database error of delete owner")
		}
		return
	}

	c.Status(http.StatusResetContent)
}

/*project*/

func GetProject(c *gin.Context, dir string) {
	// check project exist
	projDatas, err := database.GetProject(dir)
	if err != nil {
		log.Warn(c, 1500001, err, "Sorry, something error", "database error")
		return
	}

	// blog not found, get project
	c.JSON(http.StatusOK, gin.H{
		"project": projDatas,
	})
}

func PutProject(c *gin.Context, form *multipart.Form, dir string) {
	//func PutProject(owner string, proj []string, oid, superid, super_url, pid, descript, url string) error {
	projs := strings.Split(c.Param("work"), "/")[1:]
	err := database.PutProject(c.Param("owner"), projs, form.Value["owner"][0],
		form.Value["super"][0], form.Value["super"][1], form.Value["project"][0], form.Value["project"][1], dir)
	if err != nil {
		log.Warn(c, 1500001, err, "Sorry, something error", "database error")
		return
	}

	c.Header("Location", "/.")
	c.Status(http.StatusCreated)
}

func DelProject(c *gin.Context, dir string) {
	if err := database.DelProject(c.PostForm("oid"), c.PostForm("pid"), dir); err != nil {
		if err == database.ERR_TASK_FAIL {
			log.Warn(c, 2400001, err, "delete owner fail, please check oid and owner name correct")
		} else {
			log.Warn(c, 1500001, err, "Sorry, something error. please try again", "database error of delete owner")
		}
		return
	}

	c.Status(http.StatusResetContent)
}

/*blog*/

func GetBlog(c *gin.Context, dir string) {
	blogData, err := database.GetBlog(dir)
	if err != nil {
		log.Warn(c, 1500001, err, "Sorry, something error", "database error")
		return
	}

	// blog not found, get project
	c.JSON(http.StatusOK, gin.H{
		"blog": blogData,
	})
}

func PutBlog(c *gin.Context, form *multipart.Form, dir string) {
	// check category
	cat := &form.Value["category"][0]
	if *cat == "0" {
		cat = nil
	}

	// put to database
	_, blog := splitWork(c.Param("work"))
	err := database.PutBlog(form.Value["owner"][0], c.Param("owner"), form.Value["project"][0],
		form.Value["blog"][0], blog, form.Value["project"][1], form.Value["blog"][1], form.Value["blog"][2],
		form.Value["BlogType"][0], dir, cat)
	if err != nil {
		log.Warn(c, 1500001, err, "Sorry, something error. please try again", "database error of update blog")
		return
	}

	c.Header("Location", "/.")
	c.Status(http.StatusCreated)
}

func DelBlog(c *gin.Context, dir string) {
	if err := database.DelBlog(c.PostForm("oid"), c.PostForm("pid"), c.PostForm("bid"), dir); err != nil {
		if err == database.ERR_TASK_FAIL {
			log.Warn(c, 2400001, err, "delete owner fail, please check oid and owner name correct")
		} else {
			log.Warn(c, 1500001, err, "Sorry, something error. please try again", "database error of delete owner")
		}
		return
	}

	c.Status(http.StatusResetContent)
}

// split url to slice of projects and last project or blog
func splitWork(url string) ([]string, string) {
	works := strings.Split(url, "/")[1:]
	return works[:len(works)-1], works[len(works)-1]
}
