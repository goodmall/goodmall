package main

import (
	"fmt"
	"reflect"
	"runtime"
)

//
/**
此代码的用途：
	现在各种web框架 都会注册对route的处理器的  这些处理器的文件路径可以通过此方式获取
	有这些处理器的路径后 我们可以扫描他们来 获取api文档注释！
*/
func main() {
	funcInfo(funcInfo)
}

func funcInfo(f interface{}) {
	fc := runtime.FuncForPC(reflect.ValueOf(f).Pointer())
	//.Name()
	// pc := make([]uintptr, 15)
	var pc uintptr
	file, l := fc.FileLine(pc)

	s := struct {
		File string
		Line int
	}{
		File: file,
		Line: l,
	}
	fmt.Printf("%#v \n\n", s)
}
