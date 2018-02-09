package main

import (
	"fmt"

	"github.com/hprose/hprose-golang/rpc"
)

type HelloService struct {
    Hello func(string) (string, error)
    Hello2 func(string) string `name:"hello"`
}

func main() {
	client := rpc.NewHTTPClient("http://127.0.0.1:8080/")
	var helloService *HelloService
	client.UseService(&helloService)
	fmt.Println(helloService.Hello("world"))
	fmt.Println(helloService.Hello2("world"))
}