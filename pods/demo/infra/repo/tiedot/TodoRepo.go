package tiedot

import (
	_ "encoding/json"
	"fmt"
	_ "math/rand"
	_ "strconv"

	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"

	_ "os"

	"github.com/HouzuoGuo/tiedot/db"
	_ "github.com/HouzuoGuo/tiedot/dberr"

	"github.com/goodmall/goodmall/pods/demo"
)

// Ensure 接口被实现了.
var _ demo.TodoRepo = &todoRepo{}

type todoRepo struct {
	db *db.DB
}

func NewTodoRepo() demo.TodoRepo {
	// ------------------------------------------------------
	myDBDir := "./tmp/MyDatabase"
	// os.RemoveAll(myDBDir)
	// defer os.RemoveAll(myDBDir)

	// (Create if not exist) open a database
	myDB, err := db.OpenDB(myDBDir)
	if err != nil {
		panic(err)
	}
	todoCol := myDB.Use("todos")
	if todoCol == nil {
		if err := myDB.Create("todos"); err != nil {
			panic(err)
		}
	}

	// fmt.Print(todoCol)
	/*
		// Create two collections: Feeds and Votes
		if err := myDB.Create("Feeds"); err != nil {
			panic(err)
		}
		if err := myDB.Create("Votes"); err != nil {
			panic(err)
		}
	*/

	// ------------------------------------------------------

	return &todoRepo{db: myDB}
}

func (todoRepo *todoRepo) Load(id int) (*demo.Todo, error) {
	// fmt.Print("come here todoRepo:create")
	td := demo.Todo{}
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
	return &td, nil
	// return &demo.Todo{}, nil
}
func (todoRepo *todoRepo) FindById(id int) demo.Todo {

	return demo.Todo{}
}

//
func (todoRepo *todoRepo) Store(td *demo.Todo) error {
	/*docID, err := todoRepo.db.Use("todos").Insert(map[string]interface{}{
	"name": "Go 1.2 is released",
	"url":  "golang.org"})
	/*

	key := "id_" + strconv.Itoa(rand.Int())

		docID, err := todoRepo.db.Use("todos").Insert(map[string]interface{}{key: td})
	*/
	// tdDoc, _ := json.Unmarshal(td)
	docID, err := todoRepo.db.Use("todos").Insert(structs.Map(td))
	if err != nil {
		panic(err)
	}
	//fmt.Println(key)
	fmt.Println(docID)
	return nil
}
func (todoRepo *todoRepo) Remove(id int) error {

	col := todoRepo.db.Use("todos")
	// Delete document
	return col.Delete(id)
}

// ## Extra Behavior
// Size()

// ## Query
// Query(spec Specification)
func (todoRepo *todoRepo) Query(criteria demo.Query) ([]*demo.Todo, error) {

	todos := todoRepo.db.Use("todos")
	//todos.ApproxDocCount()
	var query interface{}
	// json.Unmarshal([]byte(`[{"eq": "New Go release", "in": ["Title"]}, {"eq": "slackware.com", "in": ["Source"]}]`), &query)

	queryResult := make(map[int]struct{}) // query result (document IDs) goes into map keys

	query = "all"
	if err := db.EvalQuery(query, todos, &queryResult); err != nil {
		panic(err)
	}

	var returns []*demo.Todo

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
	return returns, nil
}

// 工具类
// http://blog.csdn.net/dongfengkuayue/article/details/52473512
// github.com/fatih/structs
/*
func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[strings.ToLower(t.Field(i).Name)] = v.Field(i).Interface()
	}
	return data
}
*/
