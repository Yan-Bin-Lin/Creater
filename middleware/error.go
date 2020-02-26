package middleware

import (
	appErr "app/error"
	"app/log"
	"github.com/gin-gonic/gin"
	"net"
	"os"
	"strings"
)

// ErrorHandling returns a middleware that recovers from any panics and writes a 500 if there was one
// if no panic but there is error in context error, handle for warning to cleint side
func ErrorHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			var brokenPipe bool
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				// detect error type
				if brokenPipe {
					// connect error, can't send to front
					log.Error(c, 1500003, err.(error), 1)
				} else {
					// system error, need to send to front
					log.Error(c, 1500001, err.(error), 1, "Sorry, something error")
				}

				c.Abort()
			}

			// if no error
			err := c.Errors.Last()
			if err == nil {
				// log success
				log.Success(c)
				return
			}

			// If the connection is dead, we can't write a status to it.
			if !brokenPipe {
				// get error meta
				var (
					code int
					msg  string
				)
				switch err.Meta.(type) {
				case *appErr.ErrorDataStruct:
					meta := err.Meta.(*appErr.ErrorDataStruct)
					code = meta.Code
					msg = meta.Msg
				default:
					// worng type or something error
					code = 1500004
					msg = "Sorry, Something error"
					log.Warn(c, code, nil, msg)
				}

				// return to client
				_, httpStatus, _ := appErr.SplitCode(code)
				c.JSON(httpStatus, gin.H{
					"Code": code,
					"Msg":  msg,
				})
			}
		}()
		c.Next()
	}
}
