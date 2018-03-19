package mock

import (
	"github.com/goodmall/goodmall/pods/demo"
	// "github.com/goodmall/goodmall/pods/demo/usecase"
)

// Ensure TodoInteractor implements demo.TodoInteractor.
var _ demo.TodoInteractor = &TodoInteractor{}

type TodoInteractor struct {
	TodosFn func()[]*demo.Todo, error)
	TodosInvoked bool

	TodoFn func(id int) (*demo.Todo, error)
	TodoInvoked bool

	CreateTodoFn func(td *demo.Todo) error
	CreateTodoInvoked bool

	UpdateTodoFn func (td *demo.Todo) error
	UpdateTodoInvoked bool

	DeleteTodoFn func(id int) error
	DeleteTodoInvoked bool
}

func (ti *TodoInteractor) Todos() ([]*demo.Todo, error) {
	ti.TodosInvoked = true
	return ti.TodosFn()
}

func (ti *TodoInteractor) demo.Todo(id int) (*demo.Todo, error) {
	ti.TodoInvoked = true 
	return ti.TodoFn(id)
}

func (ti *TodoInteractor) CreateTodo(td *demo.Todo) error {
	ti.CreateTodoInvoked = true
	return ti.CreateTodoFn(td)
}

func (ti *TodoInteractor) UpdateTodo(td *demo.Todo) error {
	ti.UpdateTodoInvoked = true
	return ti.UpdateTodoFn(id)
}

func (ti *TodoInteractor) DeleteTodo(id int) error {
	ti.DeleteTodoInvoked = true
	return ti.DeleteTodoFn(id)
}
