package main

import (
	"github.com/go-ozzo/ozzo-config"

	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//
import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"

	"github.com/davecgh/go-spew/spew"
)

// -------------------------------------------------------------  ++|
// TODO    将来提取到单独的app 目录去  作为全局共享对象
var Config *config.Config

func LoadConfig(files ...string) {
	// create a Config object
	Config = config.New()
	Config.Load(files...)
}

// -------------------------------------------------------------  ++|

func main() {
	// LOAD Config
	LoadConfig("./config/app.toml")
	// PrettyPrint(Config)

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	buildRoutes(e)

	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func buildRoutes(e *echo.Echo) {
	g := e.Group("/gii")
	/*
		g.Use(middleware.BasicAuth(func(username, password string) bool {
			if username == "joe" && password == "secret" {
				return true
			}
			return false
		}))
	*/

	g.GET("/table/:table", func(c echo.Context) error {
		// NOTE 获取路径参数 注意前面没有冒号哦！
		name := c.Param("table")

		cols := getColumnsForTable(name)

		// return c.String(http.StatusOK, "tableName: "+name)
		return c.JSON(http.StatusOK, cols)

	})

	g.GET("/*", func(c echo.Context) error {

		return c.String(http.StatusOK, " from gii , pls give some valid route ")
	})
}

// Handler
func hello(c echo.Context) error {

	data, err := json.MarshalIndent(c.Echo().Routes(), "", "  ")
	if err != nil {
		return err
	}

	msg := fmt.Sprintf(" HELLO GII \n %s  %s ", "available routs : \n", string(data))

	return c.String(http.StatusOK, msg)
}

// =========================================================================  +|
// ##              core engin        -------------  +|
//             TODO 有空了提出到其他目录去

type MyColumn struct {
	core.Column

	GoType string
}

func getColumnsForTable(name string) map[string]*MyColumn /**core.Column*/ {
	var err error
	//	engine, err := xorm.NewEngine("mysql", "root:@/test?charset=utf8")
	engine, err := xorm.NewEngine(Config.GetString("db.driver", "mysql"),
		Config.GetString("db.dataSourceName", "root:@/test?charset=utf8"))

	// PrettyPrint(Config)

	checkErr(err)
	err2 := engine.Ping()
	if err2 != nil {
		log.Fatal(err2)
	}

	/*
		db := engine.DB()
		tables := db.
	*/
	dlc := engine.Dialect()
	log.Println(" db name : ", dlc.URI().DbName)

	tables, err := dlc.GetTables()
	checkErr(err)

	var tbl *core.Table
	for _, t := range tables {

		// for i, t := range tables {
		/*
			log.Printf("\n <--    table:%d    \t name: %s    --> \n", i, tbl.Name)

			colSeq, cols, err := dlc.GetColumns(tbl.Name)
			checkErr(err)

			PrettyPrint(colSeq)

			for nm, col := range cols {
				// PrettyPrint(col)
				fmt.Printf("\n\n name: %s  \t sql-type: %s  \t go-type: %s \n",
					nm,
					col.SQLType.Name,
					core.SQLType2Type(col.SQLType).Name())
			}
		*/
		if t.Name == name {
			tbl = t
			break
		}
	}
	if tbl == nil {
		log.Println("no such table :", name)
		// panic(name + " does not exists !")
		return nil // TODO 后期需要返回特定结构啦！
	}

	// 处理列
	colSeq, cols, err := dlc.GetColumns(tbl.Name)
	checkErr(err)

	var _ = colSeq
	//	PrettyPrint(colSeq)

	fmt.Printf("\n  name  \t  sql-type  \t  go-type  \n")
	fmt.Printf("================================================")

	var results = make(map[string]*MyColumn, len(cols))

	for nm, col := range cols {
		// PrettyPrint(col)
		fmt.Printf("\n  %s  \t  %s  \t  %s  ",
			nm,
			col.SQLType.Name,
			core.SQLType2Type(col.SQLType).Name())

		results[nm] = &MyColumn{
			Column: *col,
			GoType: core.SQLType2Type(col.SQLType).Name(),
		}

	}
	fmt.Println("\n")

	return results //cols

}
func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func PrettyPrint(v interface{}) {
	//   fmt.Printf("%#v", p) //with name, value and type
	// b, _ := json.MarshalIndent(v, "", "  ")
	// println(string(b))
	spew.Dump(v)
}
