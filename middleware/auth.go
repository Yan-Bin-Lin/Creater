package middleware

import (
	"app/log"
	"github.com/gin-gonic/gin"
	"net/http"

	"echotest.com/app/common"
	"echotest.com/app/service"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		var (
			refreshCookie *http.Cookie
			accessCookie  *http.Cookie
			err           error
		)

		if accessCookie, err = c.Request.Cookie("AccessToken"); err == nil {
			// check access token
			if _, err := service.GetuidByRefreshToken(accessCookie.Value); err == nil {
				c.Next()
				return
			}
		}

		// access token not found or invalid, check refresh token
		if refreshCookie, err = c.Request.Cookie("RefreshToken"); err != nil {
			// refresh token not found
			log.Warn(c, 2401504, err, "refresh token not found", "you should login first!")
			c.Abort()
			return
		} else if _, e := service.GetuidByRefreshToken(refreshCookie.Value); e != nil {
			// refresh token invalid
			log.Warn(c, 2401504, err, "refresh token invalid", "time out! please login again")
			c.Abort()
			return
		}

		// create an access token
		accessToken, err := service.NewAccessToken(refreshCookie.Value)
		if err != nil {
			log.Warn(c, 1401001, err, "generate access token fail", "Sorry, something error. Please login again")
			c.Abort()
			return
		}
		// add to header
		cookie := &http.Cookie{
			Name:     "AccessToken",
			Value:    accessToken,
			MaxAge:   86400, // one day
			Path:     "/",
			Domain:   "local.com",
			SameSite: http.SameSiteLaxMode,
			Secure:   false,
			HttpOnly: true,
		}
		c.Writer.Header().Add("Set-Cookie", cookie.String())

		c.Next()
	}
}
