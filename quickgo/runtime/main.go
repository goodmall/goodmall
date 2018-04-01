package main

import (
	"fmt"
	"runtime"
)

func main() {
	realMain()
}

func realMain() {
	_, filename, _, _ := runtime.Caller(0)
	fmt.Println("Current test filename: " + filename)
}
