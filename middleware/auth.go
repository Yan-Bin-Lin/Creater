package middleware

import (
	"app/log"
	"app/serve"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

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
		if has, err := serve.CheckAccessToken(param.Get("uid"), param.Get("AccessCode"), c.Param("owner")); err != nil {
			log.Warn(c, 1500006, err, "Sorry, something error", "database error of check access token")
			c.Abort()
			return
		} else if !has {
			log.Warn(c, 2401504, err, "access token parse fail")
			c.Abort()
			return
		}

		c.Next()
	}
}
