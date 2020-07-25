package error

import (
	"app/setting"
	"app/util/debug"
	"github.com/gin-gonic/gin"
)

// return error for logger
type ErrorReturn struct {
	ErrorData *ErrorDataStruct
	ErrorMeta *ErrorMetaStruct
}

func (e *ErrorReturn) Error() string {
	return e.ErrorData.Msg
}

// implement Unwrap for error in go 1.13
func (e *ErrorReturn) Unwrap() error {
	return e.ErrorMeta.Error
}

// struct of error message to let out
type ErrorDataStruct struct {
	Code int    `zap:"code"`
	Msg  string `zap:"msg"`
}

// error message for debug
type ErrorMetaStruct struct {
	Msg   string
	Error error // origin error
	Stack []*debug.FuncDataStruct
}

func (ems *ErrorMetaStruct) GetStack() []*debug.FuncDataStruct {
	return ems.Stack
}

// implement Unwrap for error in go 1.13
func (e *ErrorMetaStruct) Unwrap() error {
	return e.Error
}

// check if error code exist
func GetMsg(code int) string {
	if val, ok := setting.ErrorMap[code]; ok {
		return val
	} else {
		return setting.ErrorMap[0]
	}
}

//generate new error data
func NewErrorData(code int, customMsg string) (eData *ErrorDataStruct) {
	if customMsg != "" {
		eData = &ErrorDataStruct{
			Code: code,
			Msg:  customMsg,
		}
	} else {
		eData = &ErrorDataStruct{
			Code: code,
			Msg:  GetMsg(code),
		}
	}
	return
}

// split error code
// return errorType, httpStatus, customCode
func SplitCode(code int) (int, int, int) {
	if _, ok := setting.ErrorMap[code]; !ok {
		//error code not found
		return 0, 0, 0
	}

	return code / 1000000, (code % 1000000) / 1000, code % 1000
}

//generate new error meta
func NewErrorMeta(err error, stack []*debug.FuncDataStruct, customMsg string) *ErrorMetaStruct {
	return &ErrorMetaStruct{
		Msg:   customMsg,
		Error: err,
		Stack: stack,
	}
}

// new a error return struct
func NewErrorReturn(code int, err error, stack []*debug.FuncDataStruct, customMsg ...string) *ErrorReturn {
	var (
		errorData *ErrorDataStruct
		errorMeta *ErrorMetaStruct
	)

	// check error msg
	if len(customMsg) > 0 {
		errorData = NewErrorData(code, customMsg[0])
	} else {
		errorData = NewErrorData(code, "")
	}
	/*
		if errorData.Msg == setting.ErrorMap[0] {
			if err != nil {
				errorData.Msg = err.Error()
			}
		}*/
	if len(customMsg) > 1 {
		errorMeta = NewErrorMeta(err, stack, customMsg[1])
	} else {
		errorMeta = NewErrorMeta(err, stack, "")
	}

	return &ErrorReturn{errorData, errorMeta}
}

// return wrap error type
func New(code int, err error, skip int, customMsg ...string) (er error) {
	// NewErrorReturn
	if setting.Servers["main"].RunMode == gin.DebugMode {
		er = NewErrorReturn(code, err, debug.GetCallStack(skip+1), customMsg...) // skip this level
	} else {
		er = NewErrorReturn(code, err, nil, customMsg...)
	}

	return
}
