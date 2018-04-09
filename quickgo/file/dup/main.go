package main

// @SEE https://xojoc.pw/justcode/golang-file-tree-traversal.html
// @see https://flaviocopes.com/go-list-files/

import (
	"crypto/sha512"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

//  这里 检测可用性    常用的可能是 map[...]struct{}   这种惯用法！
var files = make(map[[sha512.Size]byte]string)

func checkDuplicate(path string, info os.FileInfo, err error) error {
	if err != nil {
		log.Print(err)
		return nil
	}
	if info.IsDir() {
		return nil
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Print(err)
		return nil
	}
	digest := sha512.Sum512(data)
	if v, ok := files[digest]; ok {
		fmt.Printf("%q is a duplicate of %q\n", path, v)
	} else {
		files[digest] = path
	}

	return nil
}

func main() {
	log.SetFlags(log.Lshortfile)
	dir := os.Args[1]
	err := filepath.Walk(dir, checkDuplicate)
	if err != nil {
		log.Fatal(err)
	}
}

//
/*
dir := os.Args[1]
	ignoreDirs := []string{".bzr", ".hg", ".git"}
	err := filepath.Walk(dir, printFile(ignoreDirs))
	if err != nil {
		log.Fatal(err)
	}
*/
func printFile(ignoreDirs []string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Print(err)
			return nil
		}
		if info.IsDir() {
			dir := filepath.Base(path)
			for _, d := range ignoreDirs {
				if d == dir {
					return filepath.SkipDir
				}
			}
		}
		fmt.Println(path)
		return nil
	}
}

/**
log.SetFlags(log.Lshortfile)
	dir := os.Args[1]
	info, err := os.Lstat(dir)
	if err != nil {
		log.Fatal(err)
	}
	du(dir, info)

*/
func filezize(currentPath string, info os.FileInfo) int64 {
	size := info.Size()
	if !info.IsDir() {
		return size
	}

	dir, err := os.Open(currentPath)
	if err != nil {
		log.Print(err)
		return size
	}
	defer dir.Close()

	fis, err := dir.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}
	for _, fi := range fis {
		if fi.Name() == "." || fi.Name() == ".." {
			continue
		}
		size += filezize(currentPath+"/"+fi.Name(), fi)
	}

	fmt.Printf("%d %s\n", size, currentPath)

	return size
}
