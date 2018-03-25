package main

import (
	"fmt"
	"reflect"
)

// https://stackoverflow.com/questions/7850140/how-do-you-create-a-new-instance-of-a-struct-from-its-type-at-run-time-in-go
// https://blog.golang.org/laws-of-reflection
type MyStruct struct {
	Name string
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
	fmt.Printf("the obj is : %#v", obj2)

}
