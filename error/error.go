package error

import (
	"app/setting"
	"app/util/debug"
)

// return error for log
type ErrorReturn struct {
	ErrorData *ErrorDataStruct
	ErrorMeta *ErrorMetaStruct
}

// struct of error message
type ErrorDataStruct struct {
	Code int    `zap:"code"`
	Msg  string `zap:"msg"`
}

// error message for debug
type ErrorMetaStruct struct {
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
func NewErrorData(code int) *ErrorDataStruct {
	return &ErrorDataStruct{
		Code: code,
		Msg:  GetMsg(code),
	}
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
func NewErrorMeta(err error, stack []*debug.FuncDataStruct) *ErrorMetaStruct {
	return &ErrorMetaStruct{
		Error: err,
		Stack: stack,
	}
}

// new a error return struct
func NewErrorReturn(code int, err error, stack []*debug.FuncDataStruct) *ErrorReturn {
	// check error msg
	errorData := NewErrorData(code)
	if errorData.Msg == setting.ErrorMap[0] {
		if err != nil {
			errorData.Msg = err.Error()
		}
	}

	return &ErrorReturn{errorData, NewErrorMeta(err, stack)}
}
