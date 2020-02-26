package error

import (
	"app/setting"
	"app/util/debug"
	"errors"
	"github.com/gin-gonic/gin"
)

// set context.error to handle to front end
// skip will be set for stack level(start from 0)
func ErrorHandle(c *gin.Context, code int, err error, skip int, CustomMsg ...string) (er *ErrorReturn) {
	// NewErrorReturn
	if setting.Servers["main"].RunMode == gin.DebugMode {
		er = NewErrorReturn(code, err, debug.GetCallStack(skip+1)) // skip this level
	} else {
		er = NewErrorReturn(code, err, nil)
	}

	// set gin error for front end
	if len(CustomMsg) > 0 {
		er.ErrorData.Msg = CustomMsg[0]
	}
	_ = c.Error(errors.New("")).SetMeta(er.ErrorData)

	return
}
