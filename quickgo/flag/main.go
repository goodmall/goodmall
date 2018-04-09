package main

import (
	"flag"
	"fmt"

	"log"
	"os"
	//	"strings"
	//	"golang.org/x/net/context"
	// "golang.org/x/sync/errgroup"
)

var (
	Version   = "0.0.1"
	BuildTime = "N/A"
)

func main() {
	dir := flag.String("dir", "", "扫描目录 默认是当前目录")
	flag.Usage = func() {
		fmt.Printf("%s by yiqing\n", os.Args[0])
		fmt.Printf("Version %s, Built: %s \n", Version, BuildTime)
		fmt.Println("Usage:")
		fmt.Printf("	flag [flags] arg1 arg2 \n")
		fmt.Println("Flags:")
		flag.PrintDefaults()
	}
	flag.Parse()
	if flag.NArg() != 2 {
		flag.Usage()
		os.Exit(-1)
	}
	a1 := flag.Arg(0)
	a2 := flag.Arg(1)
	fmt.Println("arg1: ", a1, " arg2:", a2)

	rtn, err := run(*dir)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("result is ", rtn)
}

func run(any string) (string, error) {
	var _ = any

	return "hi this is run func", nil
}
