package debug

import (
	"testing"
)

func wrap1() []*FuncDataStruct {
	return wrap2()
}

func wrap2() []*FuncDataStruct {
	return wrap3()
}

func wrap3() []*FuncDataStruct {
	return GetCallStack()
}

func Testdebug_GetCallStack(t *testing.T) {
	funcDatas := wrap1()

	var expect = []FuncDataStruct{
		{"", "app/util/debug.wrap3", 16},
		{"", "app/util/debug.wrap2", 12},
		{"", "app/util/debug.wrap1", 8},
		{"", "app/util/debug.TestGetCallStack", 20},
	}

	for i, l := 0, len(funcDatas); i < l; i++ {
		if funcDatas[i].Function != expect[i].Function || funcDatas[i].Line != expect[i].Line {
			t.Fatalf("wrong result. expect function: %s, line: %d. Get function: %s, line: %d",
				expect[i].Function, expect[i].Line, funcDatas[i].Function, funcDatas[i].Line)
		}
	}
}

func Testdebug_GetFuncData(t *testing.T) {
	funcData := GetFuncData()

	var expect = &FuncDataStruct{"", "app/util/debug.TestGetFuncData", 38}

	if funcData.Function != expect.Function || funcData.Line != expect.Line {
		t.Fatalf("wrong result. expect function: %s, line: %d. Get function: %s, line: %d",
			expect.Function, expect.Line, funcData.Function, funcData.Line)
	}
}
