package main

// import "github.com/tonnerre/golang-pretty"
import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	tableMeta()
}

func tableMeta() {
	var err error
	engine, err := xorm.NewEngine("mysql", "root:@/test?charset=utf8")
	checkErr(err)
	engine.Ping()

	/*
		db := engine.DB()
		tables := db.
	*/
	dlc := engine.Dialect()

	tables, err := dlc.GetTables()
	checkErr(err)

	for i, tbl := range tables {
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

	}

}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func PrettyPrint(v interface{}) {
	//	 fmt.Printf("%+v\n", p) //With name and value
	//   fmt.Printf("%#v", p) //with name, value and type
	// b, _ := json.MarshalIndent(v, "", "  ")
	// println(string(b))
	spew.Dump(v)
}
