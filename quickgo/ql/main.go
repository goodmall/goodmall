package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"time"

	"database/sql"

	qlBase "github.com/cznic/ql"
	_ "github.com/cznic/ql/driver"

	"github.com/go-ozzo/ozzo-dbx"

	"github.com/goodmall/goodmall/pods/demo"
)

// 属性类型不能是int  必须是确定的int32|int64！
type Todo struct {
	// demo.Todo
	// Id          int64  `json:"id" ql:"index xFoo, name Bar"` // int32
	ID int64 `json:"id" ql:"index pk_todos "` //
	// Created     int64  `json:"created"`
	// Status      string `json:"status"`
	Title string `json:"title"`
	// Description string `json:"description"`
}

func main() {
	log.Println("hello ql")
	// initSchema()
	// testQl()
	log.Println(Schema((*Todo)(nil), "todos", nil))

	testCrud()

}

func playSqlx() {
	db, err := dbx.Open("ql", "./tmp/db.ql")
	if err != nil {
		panic(err)
	}

	// CREATE TABLE `users` (`id` int primary key, `name` varchar(255))
	q := db.CreateTable("users", map[string]string{
		"id":   "int",
		"name": "string",
	})
	q.Execute()
	rslt, err := db.Insert("users", dbx.Params{
		// "id":   1,
		"name": "James",
	}).Execute()
	if err != nil {
		panic(err)
	}
	log.Println(rslt)
}

func Schema(v interface{}, name string, opt *qlBase.SchemaOptions) string {
	schema := qlBase.MustSchema(v, name, opt)

	return schema.String()
}

func testCrud() {
	dbPath := "./tmp/db.ql" // db.ql
	if err := os.Remove(dbPath); err != nil && !os.IsNotExist(err) {
		log.Fatal(err)
	}
	os.MkdirAll("./tmp", 0777)
	db, err := sql.Open("ql", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	if _, err := tx.Exec(Schema((*Todo)(nil), "todos", nil)); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		if _, err := tx.Exec(`INSERT INTO todos(Title) VALUES (?1 )`, "hello "+strconv.Itoa(i)); err != nil {
			log.Fatal(err)
		}
	}

	// -----------------------------------------------------  |

	rs, err := tx.Query(`select * from todos`)

	if err != nil {
		panic(err)
	}
	var td Todo
	for rs.Next() {
		td = Todo{}
		rs.Scan(&td.Title)
		log.Println(td)
		log.Printf("the struct is : %v ", td.ID)
	}
	// -----------------------------------------------------  |

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}

}

func testQl() {
	dbPath := "./tmp/db.ql" // db.ql
	if err := os.Remove(dbPath); err != nil && !os.IsNotExist(err) {
		log.Fatal(err)
	}
	os.MkdirAll("./tmp", 0777)
	db, err := sql.Open("ql", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()
	if _, err := tx.Exec(`CREATE TABLE test (
		_date time,
		str string)`); err != nil {
		log.Fatal(err)
	}
	want := "1234"
	if _, err := tx.Exec(`INSERT INTO test VALUES (?1, ?2)`, time.Now(), want); err != nil {
		log.Fatal(err)
	}
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
	var got string
	// Test that fields can appear in ORDER by only
	if err := db.QueryRow(`SELECT str FROM test ORDER BY _date`).Scan(&got); err != nil {
		log.Println(err)
	}
	if want != got {
		log.Printf("TestIndex: got %s exp %s", got, want)
	}
	// Test that ORDER BY works even without indexes
	var d time.Time
	if err := db.QueryRow(`SELECT str, _date FROM test ORDER BY _date`).Scan(&got, &d); err != nil {
		log.Println(err)
	}
	if want != got {
		log.Printf("TestIndex: got %s exp %s", got, want)
	}
	// Test that ORDER BY works at all
	tx, err = db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()
	if _, err := tx.Exec(`CREATE INDEX TestDate ON test (_date)`); err != nil {
		log.Fatal(err)
	}
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
	if err := db.QueryRow(`SELECT str, _date FROM test ORDER BY _date`).Scan(&got, &d); err != nil {
		log.Println(err)
	}
	if want != got {
		log.Printf("TestIndex: got %s exp %s", got, want)
	}
}

func initSchema() {
	dbPath := "./tmp/q2.db" // `/tmp/ql.db`

	// sql.Open("ql", "")

	qlDb, err := qlBase.OpenFile(dbPath, &qlBase.Options{CanCreate: true})
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
