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

	// TODO 增加fields字段 允许查询指定属性列表
	Query(sm TodoSearch, offset, limit int, sort string) ([]Todo, error)

	Count(sm TodoSearch) (int, error)
}
