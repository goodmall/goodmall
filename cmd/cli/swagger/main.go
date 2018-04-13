// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html

package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"

	gingen "github.com/goodmall/goodmall/base/docgen/gin"

	//	"github.com/gin-gonic/gin"
	"github.com/goodmall/goodmall/cmd/api/gin/engine"
)

func main() {

	currentdir, _ := os.Getwd()

	currPath := filepath.Join(currentdir, "main.go")

	printImports(currPath)

	// 获取控制器 package
	r := engine.GetMainEngine()

	routes := r.Routes()
	for _, rf := range routes {
		fmt.Printf("%#v \n", rf)
	}

	gingen.AddControllerPackage("", "github.com/gin-gonic/gin")
	gingen.AddControllerPackage("", "github.com/goodmall/goodmall/pods/demo/adapters/api/gin")
	genDocs(currentdir)

}

func printImports(gofile string) {
	fset := token.NewFileSet()

	f, err := parser.ParseFile(fset, gofile, nil, parser.ParseComments)
	if err != nil {
		log.Fatalf("Error while parsing %s: %s", gofile, err)
	}

	imps := f.Imports
	for _, im := range imps {
		fmt.Println("\n \t name:", im.Name, " \t path: ", im.Path.Value)
	}
}

func genDocs(dir string) {
	gingen.GenerateDocs(dir)
}
