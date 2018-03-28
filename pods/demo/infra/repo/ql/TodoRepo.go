package ql

import (
	_ "encoding/json"
	"fmt"
	"log"
	_ "math/rand"
	// "path/filepath"
	_ "strconv"

	// "github.com/fatih/structs"
	//"github.com/mitchellh/mapstructure"

	_ "os"

	// "database/sql"

	qlBase "github.com/cznic/ql"
	_ "github.com/cznic/ql/driver"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/ql"

	"github.com/goodmall/goodmall/base"
	"github.com/goodmall/goodmall/pods/demo"
)

// Ensure 接口被实现了.
var _ demo.TodoRepo = &todoRepo{}

// use this https://upper.io/db.v3/ql as our orm lib
type todoRepo struct {
	// db *sql.DB
	sess *sqlbuilder.Database
}

/*
type MyTodo struct {
	demo.Todo
	ID int
}
*/

func NewTodoRepo() demo.TodoRepo {
	initSchema()
	// ------------------------------------------------------
	myDBDir := "./tmp/MyDatabase"
	//myDBDir := "/tmp/MyDatabase"
	dbPath := myDBDir + "/ql.db"
	var _ string = dbPath
	/*
		dbPath, err := filepath.Abs(dbPath)
		if err != nil {
			panic(err)
		}
		// log.Fatal(dbPath)
		dbPath = filepath.ToSlash(dbPath)
	*/
	// os.RemoveAll(myDBDir)
	// defer os.RemoveAll(myDBDir)
	/*
		db, err := sql.Open("ql", myDBDir+"/ql.db")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("open db successful !")
	*/
	var settings = ql.ConnectionURL{
		Database: `/tmp/ql.db`, // dbPath, //`/path/to/example.db`, // Path to a QL database file.
	}

	sess, err := ql.Open(settings)
	if err != nil {
		panic(err)
	}
	fmt.Println("open db successful ! path:", dbPath)
	log.Println("yes  init it ")

	// ------------------------------------------------------
	/*
		qlDb, err := ql.OpenFile(myDBDir+"/ql.db", nil)
		if err != nil {
			panic(err)
		}
		schema := ql.MustSchema((*MyTodo)(nil), "todos", nil)

		fmt.Print(schema)

		if _, _, err = qlDb.Execute(ql.NewRWCtx(), schema); err != nil {
			panic(err)
		}
	*/
	// ------------------------------------------------------

	return &todoRepo{sess: &sess}
}

// 只能调用一次 只是用来生成表结构
func initSchema() {
	dbPath := "/tmp/ql.db" // `/tmp/ql.db`

	qlDb, err := qlBase.OpenFile(dbPath, &qlBase.Options{})
	if err != nil {
		panic(err)
	}
	schema := qlBase.MustSchema((*demo.Todo)(nil), "todos", nil)

	fmt.Print(schema)

	if _, _, err = qlDb.Execute(qlBase.NewRWCtx(), schema); err != nil {
		panic(err)
	}
	log.Println("ok init it ")
}

func (todoRepo *todoRepo) Load(id int) (*demo.Todo, error) {
	// fmt.Print("come here todoRepo:create")
	td := demo.Todo{
		Description: "this is yes ",
	}
	/*
		col := todoRepo.db.Use("todos")
		// Read document
		// id = 3320033619915385300
		readBack, err := col.Read(id)
		if err != nil {
			fmt.Println(readBack)

			// panic(err)
			return &td, err
		}
		err2 := mapstructure.Decode(readBack, &td)
		if err2 != nil {
			panic(err2)
		}
	*/
	return &td, nil
	// return &demo.Todo{}, nil
}

//
func (todoRepo *todoRepo) Store(td *demo.Todo) error {
	/*
		docID, err := todoRepo.db.Use("todos").Insert(structs.Map(td))
		if err != nil {
			panic(err)
		}
		//fmt.Println(key)
		f*/
	//	mt.Println(docID)

	return nil
}
func (todoRepo *todoRepo) Remove(id int) error {
	/*
		col := todoRepo.db.Use("todos")
		// Delete document
		return col.Delete(id)
	*/
	return nil
}

// ## Extra Behavior
// Size()

// ## Query
// Query(spec Specification)
func (todoRepo *todoRepo) Query(criteria base.Query) ([]*demo.Todo, error) {
	var returns []*demo.Todo
	/*

		todos := todoRepo.db.Use("todos")
		//todos.ApproxDocCount()
		var query interface{}
		// json.Unmarshal([]byte(`[{"eq": "New Go release", "in": ["Title"]}, {"eq": "slackware.com", "in": ["Source"]}]`), &query)

		queryResult := make(map[int]struct{}) // query result (document IDs) goes into map keys

		query = "all"
		if err := db.EvalQuery(query, todos, &queryResult); err != nil {
			panic(err)
		}


		var td *demo.Todo

		// Query result are document IDs
		for id := range queryResult {
			// To get query result document, simply read it
			readBack, err := todos.Read(id)
			if err != nil {
				panic(err)
			}
			// fmt.Printf("Query returned document %v\n", readBack)

			// 注意用法 每次都要用新的
			td = &demo.Todo{}
			err2 := mapstructure.Decode(readBack, &td)
			if err2 != nil {
				panic(err2)
			}
			td.Id = id //int32(id)
			// fmt.Print(td)
			returns = append(returns, td)

		}
		// fmt.Println(json.Marshal(returns))
		// fmt.Println("query from docs!")
	*/
	return returns, nil
}
