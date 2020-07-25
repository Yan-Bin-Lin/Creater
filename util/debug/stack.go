package debug

import (
	"runtime"
	"strconv"
)

const WrapLv = 0

type FuncDataStruct struct {
	Loacate  string
	Function string
}

// get the function data of caller, start from 0
func GetFuncData(skipNum ...int) *FuncDataStruct {
	// check skip number
	var skip int
	if len(skipNum) > 0 {
		skip = skipNum[0]
	} else {
		skip = WrapLv
	}

	pc, file, line, ok := runtime.Caller(skip + 1) // add one to skip this function
	if !ok {
		return nil
	}

	return &FuncDataStruct{
		Loacate:  file + ":" + strconv.Itoa(line),
		Function: runtime.FuncForPC(pc).Name(),
	}
}

// stack returns a nicely formatted stack frame, skipping skip frames.
func GetCallStack(skipNum ...int) []*FuncDataStruct {
	// check skip number
	var skip int
	if len(skipNum) > 0 {
		skip = skipNum[0]
	} else {
		skip = WrapLv
	}

	var stack []*FuncDataStruct

	for i := skip + 1; ; i++ { // Skip the expected number of frames, add 1 to skip this function
		funcData := GetFuncData(i)
		if funcData == nil {
			break
		}

		stack = append(stack, funcData)
	}

	return stack[:len(stack)-2] //skip runtime.main and exist
}
