package demo

import (
	"github.com/goodmall/goodmall/base"
)

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

	Query(q base.Query) ([]Todo, error)

	Count() (int, error)
}
