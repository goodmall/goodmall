package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// @see https://stackoverflow.com/questions/37194739/how-check-a-file-contain-string-or-not-in-golang
//Usage: go run filename -text=dataYouAreLookingfor
//并行扫描可以使用这个技巧： https://gobyexample.com/worker-pools
func main() {
	/*
		var text string
		fmt.Print("Enter text: ")
		// get the sub string to search from the user
		fmt.Scanln(&text)
	*/
	var text string
	// use it as cmdline argument
	textArg := flag.String("text", "", "Text to search for")
	flag.Parse()
	if fmt.Sprintf("%s", *textArg) == "" {
		// input nothing or blank space
		for strings.Trim(text, " ") == "" {
			fmt.Print("Enter text: ")
			// get the sub string to search from the user
			fmt.Scanln(&text)
		}

	} else {
		text = fmt.Sprintf("%s", *textArg)
	}
	if contains("inputs.txt", text) {
		fmt.Println("found it !")
	} else {
		fmt.Println("not contain text:", text)
	}
	// -------------------------------------------------------------  ++ |

	if contains2("inputs.txt", text) {
		fmt.Println("v2 found it !")
	} else {
		fmt.Println("v2 not contain text:", text)
	}
}

func contains(filepath string, content string) bool {
	f, err := os.Open(filepath)
	if err != nil {
		// return 0, err
		panic(err)
	}
	defer f.Close()
	/**
		reader := bufio.NewReader(os.Stdin)
	    fmt.Print("Enter text: ")
	    text, _ := reader.ReadString('\n')
	*/

	// Splits on newlines by default.
	scanner := bufio.NewScanner(f)

	line := 1
	// https://golang.org/pkg/bufio/#Scanner.Scan
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), content) {
			// return line, nil
			fmt.Println("find it in line ", line)
			return true
		}

		line++
	}

	if err := scanner.Err(); err != nil {
		// Handle the error
	}
	//
	return false

}

func contains2(filepath string, content string) bool {
	// read the whole file at once

	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	s := string(b)
	length := len(s)
	/**
	Split the file in lines first (can be done using strings.Split or bytes.Split),
	*/
	// 另一个手法是 用搜索的内容来分割字符串源  最后统计数目 如果是0表示没找到！

	fmt.Println("\n file length is ", length)
	// fmt.Println("content is :", s, " file length is :", length)
	//check whether s contains substring text
	// fmt.Println(strings.Contains(s, text))
	return strings.Contains(s, content)

}

func readDir() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
