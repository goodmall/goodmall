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

	"github.com/goodmall/goodmall/base"
	"github.com/goodmall/goodmall/pods/demo"
)

type Todo struct {
	ID          uint   `gorm:"primary_key"`
	Created     int    `json:"created"`
	Status      string `json:"status"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// Ensure 接口被实现了.
var _ demo.TodoRepo = &todoRepo{}

//
type todoRepo struct {
}

func NewTodoRepo() demo.TodoRepo {

}

func (rp *todoRepo) Create(td *Todo) error {

}

//
func (rp *todoRepo) Update(id int, td *Todo) error {

}

// Remove(td *Todo)
func (rp *todoRepo) Remove(id int) error {

}

// ## Query methods:

// ##  finder methods (latestItems(since) )
func (rp *todoRepo) Load(id int) (*Todo, error) {

}

// Query(spec Specification)
// 实现方法 可以参考 https://www.sohamkamani.com/blog/2017/10/18/golang-adding-database-to-web-application/
// 有人用string来表示查询串  这个有点跟url中的query串类似 ：?page=0&per-page=10&name=someName&age=10&title=...
func (rp *todoRepo) Query(criteria base.Query) ([]Todo, error) {

}

// ## Extra Behavior
// Size()
func (rp *todoRepo) Count() (int, error) {

}
