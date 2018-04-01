package gin

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/goodmall/goodmall/pods/demo"
	// "github.com/goodmall/goodmall/pods/demo/usecase"
	// "github.com/goodmall/goodmall/pods/demo/usecase"
	"github.com/goodmall/goodmall/base"
)

// TODO  to be continue
// TodoHandler represent a Todo-Resource ,we can also named it TodoResource .(for restful)
type TodoHandler struct {
	ts demo.TodoInteractor
}

func (tdh *TodoHandler) Query(c *gin.Context) {
	// c.Request.
	//  借鉴yii的 其实query 可以是一个模型 跟常规model类似 或者就更泛化点  map[string]string
	todos, err := tdh.ts.Query(base.Query{})
	if err != nil {
		panic(err)
	}

	c.JSON(200, todos)

}

func (tdh *TodoHandler) Create(c *gin.Context) {
	//c.JSON(200, "create")

	var todo demo.Todo

	c.Bind(&todo)

	todo.Status = "creating" // demo.TodoStatus
	//	todo.Created = int32(time.Now().Unix())
	todo.Created = int(time.Now().Unix())

	// tdh.db.Save(&todo)
	tdh.ts.Create(&todo)
	c.JSON(201, todo)
	// fmt.Println(c.Request)
	fmt.Println("yaya")

}

func (tdh *TodoHandler) Get(c *gin.Context) {
	// 注意Params 和 Query的 区别
	// idStr := c.Params.ByName("id")
	idStr := c.Query("id")
	idInt, _ := strconv.Atoi(idStr)
	// id := int32(idInt)

	var todo *demo.Todo

	todo, err := tdh.ts.Get(idInt)
	if err != nil {
		c.JSON(404, gin.H{"error": "not found " + strconv.Itoa(idInt)})
		return
	}
	/*
		if tdh.ts.First(&todo, id).RecordNotFound()false {
			c.JSON(404, gin.H{"error": "not found " + strconv.Itoa(int(id))})
		} else {
			todo = demo.Todo{Description: "hiiii"}
			c.JSON(200, todo)
		}
	*/
	c.JSON(http.StatusFound, todo)

}

// FIXME  注意在Repo层 又实现了查询逻辑 这里有重复查询的问题！
func (tdh *TodoHandler) Update(c *gin.Context) {
	// 注意Params 和 Query的 区别
	// idStr := c.Params.ByName("id")
	idStr := c.Query("id")
	idInt, _ := strconv.Atoi(idStr)
	// id := int32(idInt)

	var todo *demo.Todo

	todo, err := tdh.ts.Get(idInt)
	if err != nil {
		c.JSON(404, gin.H{"error": "not found " + strconv.Itoa(idInt)})
		return
	}

	c.Bind(&todo)

	_, err2 := tdh.ts.Update(idInt, todo)

	if err2 != nil {
		c.JSON(http.StatusFound, err2)
	}

	c.JSON(http.StatusAccepted, gin.H{"error": "nil", "msg": "update success"})

}

func (tdh *TodoHandler) Delete(c *gin.Context) {
	//	panic("yayyayyayy")
	// idStr := c.Params.ByName("id")
	idStr := c.Query("id")
	idInt, _ := strconv.Atoi(idStr)
	// id := int32(idInt)
	fmt.Println("请求id是： " + idStr)
	// var todo demo.Todo

	_, err := tdh.ts.Delete(idInt)
	if err != nil {
		c.JSON(404, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusFound, "ok")
	/*
		if tdh.db.First(&todo, id).RecordNotFound() {
			c.JSON(404, gin.H{"error": "not found"})
		} else {
			tdh.db.Delete(&todo)
			c.Data(204, "application/json", make([]byte, 0))
		}
	*/
}
