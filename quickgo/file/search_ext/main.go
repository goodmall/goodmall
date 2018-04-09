package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/beego/bee/utils"
)

// 文件操作  https://xojoc.pw/justcode/golang-file-tree-traversal.html 比如查找特定后缀的 找出重复文件

func main() {
	fmt.Println(checkExt(".txt"))

	//
	fmt.Println("<<all files in project --- ", strings.Repeat("--", 60), " project files :")
	defer fmt.Println("--- ", strings.Repeat("--", 60), "---   the end >>")
	run()
}

func checkExt(ext string) []string {
	pathS, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	var files []string
	filepath.Walk(pathS, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			/*
				if filepath.Ext(path) == ext {
					files = append(files, f.Name())
				}
			*/
			r, err := regexp.MatchString(ext, f.Name())
			if err == nil && r {
				files = append(files, f.Name())
			}
		}
		return nil
	})
	return files
}

// 扫描项目go文件
func run() ([]string, error) {
	//	searchDir := "c:/path/to/dir"
	gpsArr := utils.GetGOPATHs()
	projRelPath := "src/github.com/goodmall/goodmall"

	var searchDir string

	fmt.Printf("%v", gpsArr)

	for _, gps := range gpsArr {

		tgt := filepath.Join(gps, strings.Replace(projRelPath, "/", string(filepath.Separator), -1))

		fmt.Printf("\n\n current candidate dir is :  %s \n\n ", tgt)

		if utils.IsExist(tgt) {
			searchDir = tgt
		}
	}
	if searchDir == "" {
		panic("can't find project ")
	} else {
		fmt.Println("project path is :", searchDir)
	}

	fileList := make([]string, 0)
	e := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return err
	})

	if e != nil {
		panic(e)
	}

	for _, file := range fileList {
		fmt.Println(file)
	}

	return fileList, nil
}
