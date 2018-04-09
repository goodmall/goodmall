package main

import (
	"fmt"
	"reflect"
	"runtime"
)

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
