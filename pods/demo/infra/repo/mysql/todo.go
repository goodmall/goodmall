package mysql

import (
	_ "encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	// "fmt"
	// "log"
	_ "math/rand"
	// "path/filepath"
	_ "strconv"

	// "github.com/fatih/structs"
	//"github.com/mitchellh/mapstructure"

	_ "os"

	// "github.com/goodmall/goodmall/base"
	"github.com/goodmall/goodmall/pods/demo"

	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	. "github.com/go-xorm/builder"
	. "github.com/goodmall/goodmall/base/xorm/builder"
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
	db *gorm.DB
}

func NewTodoRepo(db *gorm.DB) demo.TodoRepo {
	return &todoRepo{
		db: db,
	}
}

func (rp *todoRepo) Create(td *demo.Todo) error {

	td2 := Todo{}

	copier.Copy(&td2, td)

	fmt.Printf("%#v", td2)

	rp.db.Create(&td2)

	td.Id = int(td2.ID)

	return nil
}

//
func (rp *todoRepo) Update(id int, model *demo.Todo) error {

	modelTmp := demo.Todo{}
	rp.db.First(&modelTmp, id)
	// fmt.Printf("%#v \n\n", model)
	// 一定要确保没有查到时的检测 不然全表被修改 太危险啦！
	if modelTmp.Id == 0 {
		return errors.New("Not Found the model for id :" + strconv.Itoa(id))
	}

	model.Id = id
	// return rs.Tx().Model(artist).Exclude("Id").Update()
	rp.db.Save(&model)

	return nil

}

// Remove(td *Todo)
func (rp *todoRepo) Remove(id int) error {

	model := demo.Todo{}
	rp.db.First(&model, id)

	fmt.Printf("%#v \n\n", model)
	if model.Id == 0 {
		return errors.New("Not Found the model for id :" + strconv.Itoa(id))
	}
	rp.db.Delete(&model)

	return nil
}

// ## Query methods:

// ##  finder methods (latestItems(since) )
func (rp *todoRepo) Load(id int) (*demo.Todo, error) {

	model := demo.Todo{}
	rp.db.First(&model, id)

	fmt.Printf("%#v \n\n", model)
	if model.Id == 0 {
		return &model, errors.New("Not Found the model for id :" + strconv.Itoa(id))
	}
	return &model, nil
}

// Query(spec Specification)
// 实现方法 可以参考 https://www.sohamkamani.com/blog/2017/10/18/golang-adding-database-to-web-application/
// 有人用string来表示查询串  这个有点跟url中的query串类似 ：?page=0&per-page=10&name=someName&age=10&title=...&sort=key1,key2 desc,key3
func (rp *todoRepo) Query(sm demo.TodoSearch, offset, limit int, sort string) ([]demo.Todo, error) {
	fmt.Println("sort is :", sort)
	rslt := []demo.Todo{}

	// rp.db.Find(&rslt)
	// 带条件查询
	fmt.Printf("search model : %#v \n\n", sm)

	// 构造条件子句
	sql, args := rp.buildSearchCond(sm)

	if len(strings.Trim(sort, " ")) == 0 {
		sort = "id desc"
	}

	// fmt.Println(sql, args)
	if len(sql) != 0 {
		rp.db.Where(sql, args...).
			Offset(offset).Limit(limit).
			Order(sort).
			Find(&rslt)
	} else {
		rp.db.Where(&sm).
			Offset(offset).Limit(limit).
			Order(sort).
			Find(&rslt) // NOTE 表名隐藏在Find 参数的类型中哦 因为是复数 所以用元素类型 即demo.Todo 来推断表名
	}

	return rslt, nil
}

// ## Extra Behavior
// Size()
func (rp *todoRepo) Count(sm demo.TodoSearch) (int, error) {

	cnt := 0

	// 构造条件子句
	sql, args := rp.buildSearchCond(sm)

	if len(sql) != 0 {
		rp.db.Model(&Todo{}).Where(sql, args...).Count(&cnt)
	} else {
		rp.db.Model(&Todo{}).Count(&cnt)
	}

	return cnt, nil

}

func (rp *todoRepo) buildSearchCond(sm demo.TodoSearch) (sql string, args []interface{}) {
	// 构造条件子句
	sql, args, _ = ToSQL(
		And(
			FilterCond(Like{"title", sm.Title}),
			FilterCond(Like{"description", sm.Description})))

	return sql, args

}
