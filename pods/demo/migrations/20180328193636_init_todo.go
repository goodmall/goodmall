package migration

import (
	"database/sql"
	"reflect"
	"time"

	//	"github.com/pressly/goose"
	"github.com/goyes/goose"

	"github.com/goodmall/goodmall/app"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var mysqldb *sql.DB

var db *gorm.DB
var err error

var _isInited = false

func init() {

	goose.AddMigration(Up20180328193636, Down20180328193636)
}

// 相当于另一种方式写了个表定义哦！
type MyTodo struct {
	gorm.Model
	Id          int    `json:"id"` // int32
	Created     int32  `json:"created"`
	Status      string `json:"status"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
type User2 struct {
	gorm.Model
	Name         string
	Age          sql.NullInt64
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`        // set field size to 255
	MemberNumber *string `gorm:"unique;not null"` // set member number to unique and not null
	Num          int     `gorm:"AUTO_INCREMENT"`  // set num to auto incrementable
	Address      string  `gorm:"index:addr"`      // create index with name `addr` for address
	IgnoreMe     int     `gorm:"-"`               // ignore this field
}

// 确保里面的代码只执行一次哦
func InitContext() {
	// SEARCH how-to-know-if-a-variable-of-arbitrary-type-is-zero-in-golang
	if _isInited == true {
		return
	}
	// 判断是否赋值了 ： https://github.com/golang/go/issues/7501
	if !reflect.ValueOf(db).IsNil() {

		return
	}

	mysqldb = app.Config.MysqlDB()

	err = mysqldb.Ping()

	if err != nil {
		panic("failed to connect database" + app.Config.DSNMysql)
	} else {
		//	println("ok db get successfully ")
	}

	db, err = gorm.Open("mysql", app.Config.DSNMysql)
	if err != nil {
		panic("failed to connect database" + app.Config.DSNMysql)
	} else {
		// println("ok db get successfully ")
	}
	// Enable Logger, show detailed log
	db.LogMode(true)

	_isInited = true
}

func Up20180328193636(tx *sql.Tx) error {
	InitContext()

	println(" init todo ")

	// Migrate the schema
	db.AutoMigrate(&User2{})

	// This code is executed when the migration is applied.
	return nil
}

func Down20180328193636(tx *sql.Tx) error {
	InitContext()

	println("hiii down!")

	db.DropTableIfExists(&User2{})

	// This code is executed when the migration is rolled back.
	return nil
}
