package mysql

import (
	_ "encoding/json"
	"errors"
	"fmt"
	"strconv"
	// "fmt"
	// "log"
	_ "math/rand"
	// "path/filepath"
	_ "strconv"

	// "github.com/fatih/structs"
	//"github.com/mitchellh/mapstructure"

	_ "os"

	"github.com/goodmall/goodmall/base"
	"github.com/goodmall/goodmall/pods/demo"

	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
// 有人用string来表示查询串  这个有点跟url中的query串类似 ：?page=0&per-page=10&name=someName&age=10&title=...
func (rp *todoRepo) Query(criteria base.Query) ([]demo.Todo, error) {

	rslt := []demo.Todo{}

	rp.db.Find(&rslt)

	return rslt, nil
}

// ## Extra Behavior
// Size()
func (rp *todoRepo) Count() (int, error) {
	return 0, nil

}
