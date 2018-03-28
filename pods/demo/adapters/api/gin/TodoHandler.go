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
)

// TODO  to be continue
// TodoHandler represent a Todo-Resource ,we can also named it TodoResource .(for restful)
type TodoHandler struct {
	ts demo.TodoInteractor
}

func (tdh *TodoHandler) Todos(c *gin.Context) {

	todos, err := tdh.ts.Todos()
	if err != nil {
		panic(err)
	}

	c.JSON(200, todos)

}

func (tdh *TodoHandler) CreateTodo(c *gin.Context) {
	//c.JSON(200, "create")

	var todo demo.Todo

	c.Bind(&todo)

	todo.Status = "creating" // demo.TodoStatus
	todo.Created = int32(time.Now().Unix())

	// tdh.db.Save(&todo)
	tdh.ts.CreateTodo(&todo)
	c.JSON(201, todo)
	// fmt.Println(c.Request)
	fmt.Println("yaya")

}

func (tdh *TodoHandler) GetTodo(c *gin.Context) {
	// 注意Params 和 Query的 区别
	// idStr := c.Params.ByName("id")
	idStr := c.Query("id")
	idInt, _ := strconv.Atoi(idStr)
	// id := int32(idInt)

	var todo *demo.Todo

	todo, err := tdh.ts.Todo(idInt)
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

func (tdh *TodoHandler) DeleteTodo(c *gin.Context) {
	//	panic("yayyayyayy")
	// idStr := c.Params.ByName("id")
	idStr := c.Query("id")
	idInt, _ := strconv.Atoi(idStr)
	// id := int32(idInt)
	fmt.Println("请求id是： " + idStr)
	// var todo demo.Todo

	err := tdh.ts.DeleteTodo(idInt)
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
