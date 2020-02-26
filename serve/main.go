package serve

import (
	"app/database"
	"app/log"
	"app/setting"
	"app/util/file"
	"app/util/hash"
	"errors"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
	"strings"
)

var (
	putBlogParam = []string{"category", "bid", "oid", "pid", "num", "descript", "BlogType"}
	putProjParam = []string{"oid", "superid", "pid", "descript"}
	blogType     = map[string]string{"article": "1"}
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
	if v := checkParam(putProjParam, form); v != "" {
		log.Warn(c, 2400001, nil, "multy part form key miss match "+v)
		return
	}

	projs := strings.Split(c.Param("work"), "/")
	superUrl := "0"
	if len(projs) > 2 {
		superUrl = hash.GetHashString(c.Param("owner"), strings.Join(projs[:len(projs)-1], "/"))
	}
	log.Debug(dir)

	err := database.PutProject(c.Param("owner"), projs[1:], form.Value["oid"][0],
		form.Value["superid"][0], superUrl, form.Value["pid"][0], form.Value["descript"][0], dir)
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
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": blogData.Name,
		"body":  blogData,
	})
}

func PutBlog(c *gin.Context, form *multipart.Form, dir string) {
	// check form
	if v := checkParam(putBlogParam, form); v != "" {
		log.Warn(c, 2400001, nil, "multy part form miss match key "+v)
		return
	}
	if len(form.File["content"]) == 0 {
		log.Warn(c, 2400001, nil, "multy part form miss match key content")
		return
	}

	cat := &form.Value["category"][0]
	if *cat == "0" {
		cat = nil
	}
	//(oid, owner, projUrl, bid, blog, pid, num, descript, typeid, filepath string, catid *string)
	// put to database
	projs, blog := splitWork(c.Param("work"))
	log.Debug("", "/"+strings.Join(projs, "/"))
	err := database.PutBlog(form.Value["oid"][0], c.Param("owner"),
		hash.GetHashString(c.Param("owner"), "/"+strings.Join(projs, "/")),
		form.Value["bid"][0], blog, form.Value["pid"][0], form.Value["num"][0],
		form.Value["descript"][0], blogType[form.Value["BlogType"][0]], dir, cat)
	if err != nil {
		log.Warn(c, 1500001, err, "Sorry, something error. please try again", "database error of update blog")
		return
	}

	// write file
	if form.Value["BlogType"][0] == "article" {
		fileHeader := form.File["content"][0]
		filename := hash.GetHashString(fileHeader.Filename)
		go file.SaveMarkdown2Html(fileHeader, setting.Servers["main"].FilePath+"/"+dir, filename+".html")
		if err := file.SaveMulipart(fileHeader, setting.Servers["main"].FilePath+"/"+dir, filename+".md"); err != nil {
			log.Warn(c, 1500001, err, "something error in write file")
			return
		}
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

// check form key match or not
func checkParam(param []string, form *multipart.Form) string {
	for _, v := range param {
		if len(form.Value[v]) == 0 {
			return v
		}
	}

	return ""
}
