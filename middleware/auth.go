package middleware

import (
	"app/serve"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		serve.CheckAccessToken(c)
		if c.IsAborted() {
			return
		}

		c.Next()
	}
}
