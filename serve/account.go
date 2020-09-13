package serve

import (
	"app/apperr"
	"app/common"
	"app/database"
	"app/log"
	"app/setting"
	"app/util/cookie"
	"app/util/hash"
	"app/util/random"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// set for hash parameter
var Params = struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}{
	memory:      65536,
	iterations:  10,
	parallelism: 2,
	saltLength:  16,
	keyLength:   32,
}

func Login(c *gin.Context) {
	// get salt
	userName := c.PostForm("username")
	salt, err := database.GetSalt(userName)
	if err != nil {
		log.Warn(c, apperr.ErrPermissionDenied, err, "Sorry, something error", "database error of login")
		return
	} else if salt == "" {
		log.Warn(c, apperr.ErrWrongArgument, err, "wrong username or password")
		return
	}

	// get hash password
	pw, err := hash.GetPWHashString(c.PostForm("password"), salt, Params.iterations, Params.memory, Params.parallelism, Params.keyLength)
	if err != nil {
		log.Warn(c, apperr.ErrPermissionDenied, err, "Sorry, something error", "base64 decode error")
		return
	}

	// login
	userData, err := database.Login(userName, pw)
	if err != nil {
		log.Warn(c, apperr.ErrPermissionDenied, err, "Sorry, something error", "database error of login")
		return
	} else if userData == nil {
		log.Warn(c, apperr.ErrWrongArgument, err, "wrong username or password")
		return
	}

	// generate new token
	uid := strconv.Itoa(userData.Uid)
	code, err := newAccessToken(uid)
	if err != nil {
		log.Warn(c, apperr.ErrPermissionDenied, err, "Sorry, something error", "database error of create token")
	}

	// delete uid
	userData.Uid = 0

	// create new cookie
	c.Writer.Header().Add("Set-Cookie", cookie.CreateCookie("AccessToken", []string{"AccessCode", "uid"},
	[]string{code, uid}, 2592000, "/", "."+setting.Servers["main"].Host, http.SameSiteLaxMode,
	true, true))

	c.JSON(http.StatusOK, gin.H{
		"user": userData,
	})
}

// get root page
func GetUser(c *gin.Context) {
	// return root page data
	c.HTML(http.StatusOK, "index/account", gin.H{
		"title":       "DCreater",
		"description": "create your blog",
		"author":      "林彥賓, https://github.com/Yan-Bin-Lin",
		"account":     true,
	})
}

// insert an user if oid is 0 else update
func PutUser(c *gin.Context) {
	pw, salt, err := hash.NewPWHashString(c.PostForm("password"), Params.iterations, Params.memory, Params.parallelism, Params.keyLength)
	if err != nil {
		log.Error(c, apperr.ErrPermissionDenied, err, 0, "Sorry, something error", "rand function error")
		return
	}
	err = database.PutUser(c.PostForm("uid"), c.PostForm("username"), pw, c.PostForm("email"), salt)
	if err != nil {
		log.Warn(c, apperr.ErrWrongArgument, err, "sorry, something error. try again", "insert new user fail")
		return
	}

	c.Redirect(http.StatusSeeOther, setting.Servers["main"].Host+strconv.Itoa(setting.Servers["main"].Port))
}

// delete an user
func DelUser(c *gin.Context) {
	if err := database.DelUser(c.PostForm("uid"), c.PostForm("username"), c.PostForm("password")); err != nil {
		log.Warn(c, apperr.ErrWrongArgument, err, "sorry, something error. try again", "delete user fail")
	}

	c.Redirect(http.StatusSeeOther, setting.Servers["in"].Host+strconv.Itoa(setting.Servers["main"].Port))
}

// generate a new token and save it
func newAccessToken(uid string) (string, error) {
	code, err := random.GetRandomString(32)
	if err != nil {
		return "", err
	}
	return code, database.NewAccessToken(uid, code)
}

// check access token vaild
func CheckAccessToken(c *gin.Context) {
	// get cookie param
	param, err := common.GetCookieParam(c, "AccessToken");
	if err != nil {
		return
	}

	if has, err := database.CheckAccessToken(param.Get("uid"), param.Get("AccessCode"), c.PostForm("oid")); err != nil {
		log.Warn(c, 1500006, err, "Sorry, something error", "database error of check access token")
		c.Abort()
		return
	} else if !has {
		log.Warn(c, apperr.ErrPermissionDenied, err, "access token parse fail")
		c.Abort()
		return
	}
}