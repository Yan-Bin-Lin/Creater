package log

import (
	appErr "app/error"
	"app/util/debug"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	TaskOK      = "OK"    // success
	TaskFail    = "Fail"  // fail but not error (warn)
	TaskError   = "Error" // panic error occure
	TaskOnGoing = "OnGoing"
	TaskUnknow  = "Unknow"
)

var Logger *zap.Logger

// config logger here
func newLogger() *zap.Logger {
	logger, _ := zap.NewDevelopment()
	return logger
}

func init() {
	Logger = newLogger()
}

// fuction data record for log
type FuncDataStruct struct {
	debug.FuncDataStruct
}

// implment zap json-encode for FuncData
func (f *FuncDataStruct) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("locate", f.Loacate)
	enc.AddString("func", f.Function)
	return nil
}

type FuncDataStructs []*debug.FuncDataStruct

// implment zap json-encode for []FuncData
func (f FuncDataStructs) MarshalLogArray(enc zapcore.ArrayEncoder) error {
	for _, funcData := range f {
		if err := enc.AppendObject(&FuncDataStruct{*funcData}); err != nil {
			return err
		}
	}
	return nil
}

type LogErrorMeta struct {
	appErr.ErrorMetaStruct
}

// implement zap json-encode for ErrorMetaStruct
func (f *LogErrorMeta) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	if f.Error != nil {
		enc.AddString("Error", f.Error.Error())
	}
	if len(f.Stack) != 0 {
		// implement array marshal
		if err := enc.AddArray("CallStack", FuncDataStructs(f.Stack)); err != nil {
			return err
		}
	}
	return nil
}

type LogErrorData struct {
	appErr.ErrorDataStruct
}

// implment zap json-encode for ErrorDataStruct
func (f *LogErrorData) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddInt("code", f.Code)
	enc.AddString("msg", f.Msg)
	return nil
}

// recored msg in function for log
type LogDataStruct struct {
	TaskStatus string        `zap:"taskStatus"`
	ErrorData  *LogErrorData `zap:"ErrorData, omitempty"`
	ErrorMeta  *LogErrorMeta `zap:"ErrorMeta, omitempty"`
}

// implment zap json-encode for  InfoMsg
func (f *LogDataStruct) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("taskStatus", f.TaskStatus)
	if f.ErrorData != nil {
		if err := enc.AddObject("ErrorData", f.ErrorData); err != nil {
			return err
		}
	}
	if f.ErrorMeta != nil {
		if err := enc.AddObject("ErrorMeta", f.ErrorMeta); err != nil {
			return err
		}
	}
	return nil
}

// new a log data struct
func NewLogData(taskStatus string, errorData *LogErrorData, errorMeta *LogErrorMeta) *LogDataStruct {
	return &LogDataStruct{
		TaskStatus: taskStatus,
		ErrorData:  errorData,
		ErrorMeta:  errorMeta,
	}
}
