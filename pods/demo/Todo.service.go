package demo

// "github.com/goodmall/goodmall/base"

// TodoInteractor represents a service for managing todos.
// alias : TodoService
// dependencies : TodoRepo
type TodoInteractor interface {
	//

	Create(td *Todo) (*Todo, error)

	Update(id int, td *Todo) (*Todo, error)

	Delete(id int) (*Todo, error)

	//
	Get(id int) (*Todo, error)

	// Query(q base.Query) ([]Todo, error) //  签名修改为使用搜索模型
	Query(q TodoSearch) ([]Todo, error)

	Count() (int, error)
}
