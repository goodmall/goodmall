package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct{

}

func main()  {
	fmt.Println("hi")

	argTest := MyStruct{}
	typ := reflect.TypeOf(argTest)
	 
	fmt.Println(typ)

	fmt.Println(typ.PkgPath())
	fmt.Println(typ.Kind())
	fmt.Println(typ.PkgPath())
}