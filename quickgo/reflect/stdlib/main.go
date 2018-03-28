package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

// https://stackoverflow.com/questions/7850140/how-do-you-create-a-new-instance-of-a-struct-from-its-type-at-run-time-in-go
// https://blog.golang.org/laws-of-reflection
type MyStruct struct {
	Name string "this is the tag string anything here"
}

func main() {
	fmt.Println("hi")

	argTest := MyStruct{
		Name: "yes this is me",
	}
	typ := reflect.TypeOf(argTest)

	fmt.Println(typ)

	fmt.Println(typ.PkgPath())
	fmt.Println(typ.Kind())
	fmt.Println(typ.PkgPath())

	obj2 := reflect.Zero(typ).Interface()
	fmt.Printf("the obj is : %#v \n", obj2)

	// 获取对象名称
	// https://stackoverflow.com/questions/24318389/golang-elem-vs-indirect-in-the-reflect-package
	fmt.Println(reflect.Indirect(reflect.ValueOf(argTest)).Type().Name())

	fmt.Println(strings.Repeat("=", 80))
	printTagOf(argTest)

	obj3 := struct {
		Name string `username:"yiqing" sex:"1"  `
	}{
		Name: "hi",
	}

	printStructVal(obj3)

	printTagOf(obj3)

	printTagKey(obj3, "Name", "username")
	printTagKey(obj3, "Name", "sex")
	obj4 := &obj3
	printTagOf(*obj4)
	printTagOf(obj4)

}

func printTagOf(obj interface{}) {
	fmt.Println("\n <<< enter  ", WhereAmI(), strings.Repeat("=", 80), "\n\n")
	defer fmt.Println("\n\n", strings.Repeat("=", 80), " exit ", WhereAmI(), " >>>\n")

	val := reflect.Indirect(reflect.ValueOf(obj))

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		tag := typeField.Tag

		// fmt.Printf("Field Name: %s,\t Field Value: %v,\t Tag Value: %s\n", typeField.Name, valueField.Interface(), tag.Get("tag_name"))
		fmt.Printf("Field Name: %s,\t Field Value: %v,\t Tag Value: %#v \n", typeField.Name, valueField.Interface(), tag)
	}
}

func printTagKey(obj interface{}, fieldName string, tagKey string) {
	// 类型是所有结构体共享的东西
	typ := reflect.TypeOf(obj)

	fieldVal := reflect.Indirect(reflect.ValueOf(obj)).FieldByName(fieldName)
	// structVal := reflect.Indirect(reflect.ValueOf(obj))

	fieldTyp, _ := typ.FieldByName(fieldName)
	fmt.Printf("Field Name: %s,\t Field Value: %v,\t Tag Value for key %s is: %#v \n",
		fieldTyp.Name,
		fieldVal.Interface(),
		tagKey,
		fieldTyp.Tag.Get(tagKey))
	// fmt.Printf("%#v \n", fieldTyp)
	// fmt.Printf("Field Name: %s,\t Field Value: %v,\t Tag Value: %#v \n", typeField.Name, valueField.Interface(), tag)
	// fmt.Println(fieldName, " tag for key: ", tagKey, "  is :", tag.Get(tagKey))
}

func printStructVal(obj interface{}) {
	fmt.Println("\n <<< enter printStructVal ", strings.Repeat("=", 80), "\n\n")
	defer fmt.Println("\n\n", strings.Repeat("=", 80), " exit printStructVal >>>\n")

	val := reflect.Indirect(reflect.ValueOf(obj))
	fmt.Printf("%#v \n", val)

	// caller := runtime.Caller(0)
	// fmt.Printf("%#v \n", runtime.GOOS)

	// trace()
	// trace2()
	fmt.Println("from function :", getFuncName(), WhereAmI())
}

// ------------------------------------------------------------------------- ++|
func trace() {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	fmt.Printf("%s:%d %s\n", file, line, f.Name())
}
func trace2() {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	fmt.Printf("%s,:%d %s\n", frame.File, frame.Line, frame.Function)
}

func getFuncName() string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	// fmt.Printf("%s,:%d %s\n", frame.File, frame.Line, frame.Function)
	return frame.Function
}

// copy from https://github.com/jimlawless/whereami/blob/master/whereami.go
// return a string containing the file name, function name
// and the line number of a specified entry on the call stack
func WhereAmI(depthList ...int) string {
	var depth int
	if depthList == nil {
		depth = 1
	} else {
		depth = depthList[0]
	}
	function, file, line, _ := runtime.Caller(depth)
	return fmt.Sprintf("File: %s  Function: %s Line: %d", chopPath(file), runtime.FuncForPC(function).Name(), line)
}

// return the source filename after the last slash
func chopPath(original string) string {
	i := strings.LastIndex(original, "/")
	if i == -1 {
		return original
	} else {
		return original[i+1:]
	}
}

// ------------------------------------------------------------------------- ++|
