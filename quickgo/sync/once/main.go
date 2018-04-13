package main

import (
	"fmt"
	"sync"
)

var m *Single

type Single struct {
	name string
	age  int
	sex  bool
}

var once sync.Once

func GetInstance() *Single {
	once.Do(func() {
		m = &Single{}
	})
	return m
}
func main() {
	a := GetInstance()
	a.name = "a"
	a.age = 1
	a.sex = true
	fmt.Println(a)
	b := GetInstance()
	b.name = "b"
	b.age = 2
	b.sex = false

	fmt.Println("should be same!")
	fmt.Println(a, b)
	fmt.Println(*a, *b)
}
