package main

import (
	"encoding/json"
	"os"
	"reflect"
	"time"
)

// import "github.com/tonnerre/golang-pretty"
import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"

	"github.com/davecgh/go-spew/spew"
	myorm "github.com/goodmall/goodmall/quickgo/xorm"
)

var engine *xorm.Engine

type Todo struct {
	Id   int64
	Name string `xorm:"varchar(25) notnull unique 'todo_name'"`
}

func (td Todo) TableName() string {
	return "todos"
}

type MyTodo struct {
	Id    int64
	Title string `xorm:"varchar(150) notnull  "`
}

func main() {
	insertObj()
	// getBeanId()
}

func getBeanId() {
	log.Println("<< ==================enter getBeanID ======================")
	defer log.Println(" ==================end getBeanID ======================>>")
	engine, err := xorm.NewEngine("mysql", "root:@/test?charset=utf8")
	checkErr(err)
	engine.Ping()

	arg := Todo{
		Id: 10,
	}

	tbl := engine.TableInfo(arg)
	pks := tbl.PKColumns()
	// spew.Dump(pks)

	// 获取主键们
	for _, pkCol := range pks {
		nm := pkCol.Name

		log.Println("table pk name is :", nm)
		log.Println("field name is :", pkCol.FieldName)

		v := reflect.ValueOf(arg)
		log.Println(" id value of the struct is : ", v.FieldByName(pkCol.FieldName))
		log.Println(" id value of the struct is : ", v.FieldByName(pkCol.FieldName).Int())
		log.Println(" id value of the struct is : ", v.FieldByName(pkCol.FieldName).Interface())

		// engine.ID(v.FieldByName(pkCol.FieldName).Int()).Delete( reflect.Zero(reflect.TypeOf(arg)).Interface())
	}

}

func basicWork() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:@/test?charset=utf8")
	checkErr(err)
	engine.Ping()

	//	spew.Dump(engine.DBMetas())
	engine.SetTableMapper(core.SameMapper{})
	ti := engine.TableInfo(Todo{})
	spew.Dump(ti)

	if exist, _ := engine.IsTableExist(Todo{}); exist {
		log.Println("todo", "table exists! we will drop it")
		if err := engine.DropTables(Todo{}); err != nil {
			log.Println("drop table success ")
		}
	}
	log.Println("create todo table: ")
	// engine.CreateTables(Todo{})
	if err := engine.Table("todos").CreateTable(Todo{}); err != nil {
		log.Fatalln("create table failure ! ", err)
	}
	log.Println("create table scucess !")

	engine.DumpAll(os.Stdout)

	// ------------------------------------------------------------------------- +|
	engine2, err := xorm.NewEngine("sqlite3", "./test.db")
	checkErr(err)
	engine2.Ping()
	// PrettyPrint(engine2)
	// pretty.Println(struct{ Name string }{Name: "hi"})
	// myorm.Test()
}

func insertObj() {
	log.Println("<<< enter insert obj ...")
	defer log.Println(" ... exit insert obj >>>")

	var err error
	engine, err = xorm.NewEngine("mysql", "root:@/test?charset=utf8")
	checkErr(err)
	engine.Ping()

	td := Todo{
		Name: "hi this is test",
	}
	/*
		aff, err := engine.InsertOne(td)
		checkErr(err)
		PrettyPrint(aff)
	*/

	engine.ShowSQL(true)

	uow := myorm.NewUintOfWork(engine)
	uow.RegisterNew(&td)

	td.Name = "yes this is updated !"

	uow.RegisterDirty(&td)

	uow.Commit()

	time.Sleep(10 * time.Second)
	log.Println("10 sencond passed ! the record will be deleted now")
	uow.RegisterDeleted(&td)

	uow.Commit()
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func PrettyPrint(v interface{}) {
	//	 fmt.Printf("%+v\n", p) //With name and value
	//   fmt.Printf("%#v", p) //with name, value and type
	b, _ := json.MarshalIndent(v, "", "  ")
	println(string(b))
}
