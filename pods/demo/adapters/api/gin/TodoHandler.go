package gin

import (
	"fmt"
	_ "strconv"
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

	/*
		idStr := c.Params.ByName("id")
		idInt, _ := strconv.Atoi(idStr)
		id := int32(idInt)

		var todo demo.Todo

		if tdh.db.First(&todo, id).RecordNotFound()false {
			c.JSON(404, gin.H{"error": "not found " + strconv.Itoa(int(id))})
		} else {
			todo = demo.Todo{Description: "hiiii"}
			c.JSON(200, todo)
		}
	*/
	c.JSON(200, "hello")

}

func (tdh *TodoHandler) DeleteTodo(c *gin.Context) {

	/*
		idStr := c.Params.ByName("id")
		idInt, _ := strconv.Atoi(idStr)
		id := int32(idInt)

		var todo demo.Todo

		if tdh.db.First(&todo, id).RecordNotFound() {
			c.JSON(404, gin.H{"error": "not found"})
		} else {
			tdh.db.Delete(&todo)
			c.Data(204, "application/json", make([]byte, 0))
		}
	*/
}
