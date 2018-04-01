package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	cwd_arg = flag.String("cwd", "", "set cwd")
	h       = flag.Bool("h", false, "help for this app , ")
)

func init() {
	flag.Parse()
	if *cwd_arg != "" {
		/*
			if err := os.Chdir(*cwd_arg); err != nil {
				fmt.Println("Chdir error:", err)
			}
		*/
		fmt.Println("chdir to some specified directory ：", *cwd_arg)
	} else {
		// fmt.Println("给我个配置文件位置呢 !")
		// 	os.Exit(2)
	}
	// 改变默认的 Usage
	flag.Usage = usage

}

func main() {
	args := os.Args //获取用户输入的所有参数
	if args == nil || len(args) < 2 {
		flag.Usage() //如果用户没有输入,或参数个数不够,则调用该函数提示用户
		return
	}
	if *h {
		flag.Usage()
		return
	}
	// printConfigFile()
}

func usage() {
	fmt.Fprintf(os.Stderr, `nginx version: nginx/1.10.0
Usage: appname|go run main.go [-h]   [-cwd <<current work dirictory>>] 

Options:
`)
	flag.PrintDefaults()
}

/**
* https://stackoverflow.com/questions/23847003/golang-tests-and-working-directory
**/
func printConfigFile() {
	wd, _ := os.Getwd()
	for !strings.HasSuffix(wd, "goodmall") {
		wd = filepath.Dir(wd)
	}

	// raw, err := ioutil.ReadFile(fmt.Sprintf("%s/src/conf/conf.dev.json", wd))
	raw, err := ioutil.ReadFile(fmt.Sprintf("%s/config/app.yaml", wd))
	if err != nil {
		panic(err)
	}
	fmt.Println(string(raw))
	fmt.Println(wd)
}
