package serve

import (
	"app/database"
	"app/logger"
	"app/setting"
	"app/util/hash"
	"app/util/random"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
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
		log.Warn(c, 1500001, err, "Sorry, something error", "database error of login")
		return
	} else if salt == "" {
		log.Warn(c, 2400001, err, "worng username or password", "")
		return
	}

	// get hash password
	pw, err := hash.GetPWHashString(c.PostForm("password"), salt, Params.iterations, Params.memory, Params.parallelism, Params.keyLength)
	if err != nil {
		log.Warn(c, 1500001, err, "Sorry, something error", "base64 decode error")
		return
	}

	// login
	userData, err := database.Login(userName, pw)
	if err != nil {
		log.Warn(c, 1500001, err, "Sorry, something error", "database error of login")
		return
	} else if userData == nil {
		log.Warn(c, 2400001, err, "worng username or password")
		return
	}

	// generate new token
	uid := strconv.Itoa(userData.Uid)
	code, err := newAccessToken(uid)
	if err != nil {
		log.Warn(c, 1500001, err, "Sorry, something error", "database error of create token")
	}

	// encode uid and token to cookie
	params := url.Values{}
	params.Add("AccessCode", code)
	params.Add("uid", uid)

	// add to header
	cookie := &http.Cookie{
		Name:     "AccessToken",
		Value:    params.Encode(),
		MaxAge:   2592000, // 30 day
		Path:     "/",
		Domain:   "." + setting.Servers["main"].Host,
		SameSite: http.SameSiteLaxMode,
		Secure:   false,
		HttpOnly: true,
	}
	c.Writer.Header().Add("Set-Cookie", cookie.String())
	c.JSON(http.StatusOK, gin.H{
		"user": userData,
	})
}

// insert an user if oid is 0 else update
func PutUser(c *gin.Context) {

	pw, salt, err := hash.NewPWHashString(c.PostForm("password"), Params.iterations, Params.memory, Params.parallelism, Params.keyLength)
	if err != nil {
		log.Error(c, 1500001, err, 0, "Sorry, something error", "rand function error")
		return
	}
	err = database.PutUser(c.PostForm("uid"), c.PostForm("username"), pw, c.PostForm("email"), salt)
	if err != nil {
		log.Warn(c, 2400001, err, "sorry, something error. try again", "insert new user fail")
	}

	c.Redirect(http.StatusSeeOther, setting.Servers["main"].Host+":8000")
}

// delect an user
func DelUser(c *gin.Context) {
	if err := database.DelUser(c.PostForm("uid"), c.PostForm("username"), c.PostForm("password")); err != nil {
		log.Warn(c, 2400001, err, "sorry, something error. try again", "delete user fail")
	}

	c.Redirect(http.StatusSeeOther, setting.Servers["main"].Host+":8000")
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

	var (
		accessCookie *http.Cookie
		err          error
	)

	// check cookie
	if accessCookie, err = c.Request.Cookie("AccessToken"); err != nil {
		log.Warn(c, 2401504, err, "access token not found in cookie")
		c.Abort()
		return
	}

	// check access
	param, err := url.ParseQuery(accessCookie.Value)
	if err != nil {
		log.Warn(c, 2401504, err, "access token parse fail")
		c.Abort()
		return
	}
	if has, err := database.CheckAccessToken(param.Get("uid"), param.Get("AccessCode"), c.PostForm("oid")); err != nil {
		log.Warn(c, 1500006, err, "Sorry, something error", "database error of check access token")
		c.Abort()
		return
	} else if !has {
		log.Warn(c, 2401504, err, "access token parse fail")
		c.Abort()
		return
	}
}
