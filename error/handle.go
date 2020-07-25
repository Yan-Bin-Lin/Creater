package error

import (
	"errors"
	"github.com/gin-gonic/gin"
)

// set context.error to handle to front end
// skip will be set for stack level(start from 0)
func ErrorHandle(c *gin.Context, code int, err error, skip int, customMsg ...string) (er *ErrorReturn) {
	er = New(code, err, skip, customMsg...).(*ErrorReturn)
	_ = c.Error(errors.New("")).SetMeta(er.ErrorData)
	return
}

// set context.error to handle to front end
// skip will be set for stack level(start from 0)
func ErrorHandleErr(c *gin.Context, err error) *ErrorReturn {
	er, ok := err.(*ErrorReturn)
	if !ok {
		return ErrorHandle(c, 1500001, errors.New("WRONG ERROR TYPE!!"), 0)
	}
	_ = c.Error(errors.New("")).SetMeta(er.ErrorData)
	return er
}
