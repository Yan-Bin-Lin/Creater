package log

import (
	appErr "app/error"
	"app/util/debug"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const LogCtxKey = "LogData"

// set logger data
func LogHandle(c *gin.Context, logData *LogDataStruct) {
	c.Set(LogCtxKey, logData)
}

type V struct {
	Key   string
	Value interface{}
}

// wrap logger debug
func Debug(msg string, params ...interface{}) {
	Logger.Debug(msg,
		zap.Any("FuncData", debug.GetFuncData(1)),
		zap.Any("param", params),
	)
}

// wrap logger info
func Info(c *gin.Context, taskStatus string, msg string) {
	LogHandle(c, NewLogData(taskStatus, &LogErrorData{appErr.ErrorDataStruct{-1, msg}}, nil))
}

// a short cup of request success info
func Success(c *gin.Context) {
	LogHandle(c, NewLogData(TaskOK, nil, nil))
}

// handle for custom private message for logger
func LogErrorHandle(c *gin.Context, taskStatus string, errorData *LogErrorData, errorMeta *LogErrorMeta, CustomMsg ...string) {
	nld := NewLogData(taskStatus, errorData, errorMeta)
	if len(CustomMsg) > 1 {
		nld.ErrorData.Msg = CustomMsg[1]
	}
	LogHandle(c, nld)
}

// wrap logger warn, warn error.Error() can be ""
// first custom msg should be message return to front end
// second msg should be message for logging
func Warn(c *gin.Context, code int, err error, CustomMsg ...string) {
	errorReturn := appErr.ErrorHandle(c, code, err, 1, CustomMsg...) // skip this wrap
	LogErrorHandle(c, TaskFail, &LogErrorData{*errorReturn.ErrorData}, &LogErrorMeta{*errorReturn.ErrorMeta}, CustomMsg...)
}

// wrap logger warn, warn error.Error() can be ""
// first custom msg should be message return to front end
// second msg should be message for logging
func WarnErr(c *gin.Context, err error) {
	errorReturn := appErr.ErrorHandleErr(c, err) // skip this wrap
	LogErrorHandle(c, TaskFail, &LogErrorData{*errorReturn.ErrorData}, &LogErrorMeta{*errorReturn.ErrorMeta})
}

// wrap logger error
// first custom msg should be message return to front end
// second msg should be message for logging
// if stack is true, the call stack will start form skip(start from 0)
func Error(c *gin.Context, code int, err error, skip int, CustomMsg ...string) {
	errorReturn := appErr.ErrorHandle(c, code, err, skip+1, CustomMsg...) // skip this wrap
	LogErrorHandle(c, TaskError, &LogErrorData{*errorReturn.ErrorData}, &LogErrorMeta{*errorReturn.ErrorMeta}, CustomMsg...)
}
