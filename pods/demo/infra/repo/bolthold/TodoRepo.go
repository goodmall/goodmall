package bolthold

import (
	"encoding/json"
	_ "encoding/json"
	"fmt"
	_ "math/rand"
	"os"
	_ "regexp"
	_ "strconv"
	_ "unsafe"

	_ "github.com/fatih/structs"
	_ "github.com/mitchellh/mapstructure"

	_ "os"

	"github.com/goodmall/goodmall/pods/demo"
	"github.com/timshannon/bolthold"

	"github.com/jinzhu/copier"
)

// Ensure 接口被实现了.
var _ demo.TodoRepo = &todoRepo{}

type MyTodo struct {
	demo.Todo
	Id int `boltholdKey:"Id"` // the tagName isn't required, but some linters will complain without it
}

type todoRepo struct {
	store *bolthold.Store
}

func NewTodoRepo() demo.TodoRepo {
	// TODO 后期要通过注入来完成！
	// ------------------------------------------------------
	myDBDir := "./tmp/bolthold" // 这里啥情况？  文件路径
	// os.RemoveAll(myDBDir)
	// defer os.RemoveAll(myDBDir)
	filename := myDBDir + "/todos.db"

	if _, err := os.Stat(myDBDir); os.IsNotExist(err) {
		// os.Create(filename)
		if err := os.MkdirAll(myDBDir, 0700); err != nil {
			panic(err)
		}
	}
	/*
		if _, err := os.Open(filename); err != nil {
			os.Create(filename)
		}
	*/

	// store, err := bolthold.Open(filename, 0666, nil)
	store, err := bolthold.Open(filename, 0666, &bolthold.Options{
		Encoder: json.Marshal,
		Decoder: json.Unmarshal,
	})
	if err != nil {
		panic(err)
	}

	// ------------------------------------------------------

	return &todoRepo{store: store}
}

//
func (todoRepo *todoRepo) Store(td *demo.Todo) error {

	var tmpTd demo.Todo
	/*tmpTd = demo.Todo(td)
	 */
	copier.Copy(&tmpTd, td)

	// p := (*MyTemplate)(unsafe.Pointer(t))
	myTd := MyTodo{Todo: tmpTd}

	key := bolthold.NextSequence()
	// item := structs.Map(td)
	err := todoRepo.store.Insert(key, myTd /*td*/)
	if err != nil {
		panic(err)
	}

	fmt.Println("id:", key, "document:", myTd)

	return nil
}
func (todoRepo *todoRepo) Remove(id int) error {
	/*
		col := todoRepo.store.Use("todos")
		// Delete document
		return col.Delete(id)
	*/
	return nil
}

func (todoRepo *todoRepo) Load(id int) (*demo.Todo, error) {
	// fmt.Print("come here todoRepo:create")
	td := demo.Todo{}

	return &td, nil
	// return &demo.Todo{}, nil
}

// ## Extra Behavior
// Size()

// ## Query
// Query(spec Specification)
func (todoRepo *todoRepo) Query(criteria demo.Query) ([]*demo.Todo, error) {
	/*
		todos := todoRepo.store.Use("todos")
		//todos.ApproxDocCount()
		var query interface{}
		// json.Unmarshal([]byte(`[{"eq": "New Go release", "in": ["Title"]}, {"eq": "slackware.com", "in": ["Source"]}]`), &query)

		queryResult := make(map[int]struct{}) // query result (document IDs) goes into map keys

		query = "all"
		if err := store.EvalQuery(query, todos, &queryResult); err != nil {
			panic(err)
		}
	*/

	// var returns []*demo.Todo
	var returns []MyTodo
	/*
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
	*/
	// fmt.Println(json.Marshal(returns))
	// fmt.Println("query from docs!")
	err := todoRepo.store.Find(&returns, &bolthold.Query{})
	// err := todoRepo.store.Find(&returns, bolthold.Where("Title").Eq("我是新来的当当当的 你好"))
	if err != nil {
		panic(err)
	}
	fmt.Println(returns)
	return []*demo.Todo{}, nil
}
