package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/schema"
)

type Person struct {
	Name  string `schema:"name"` // custom name
	Admin bool   `schema:"-"`    // this field is never set
}

// see https://github.com/google/go-querystring/issues/7
// see https://stackoverflow.com/questions/40564842/convert-url-query-map-of-slices-to-struct-golang
func MyHandler(w http.ResponseWriter, r *http.Request) {

	var person Person
	decoder := schema.NewDecoder()
	if err := decoder.Decode(&person, r.URL.Query()); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(person)

	b, err := json.Marshal(person)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Fprintf(w, string(b)) //这个写入到w的是输出到客户端的

	/**
		   w.Header().Set("Content-Type", "application/json")
	    json.NewEncoder(w).Encode(employeeStruct)
	*/
}

func main() {
	http.HandleFunc("/", MyHandler)          //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
