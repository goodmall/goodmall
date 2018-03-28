package demo

// TodoInteractor represents a service for managing todos.
// alias : TodoService
// dependencies : TodoRepo
type TodoInteractor interface {
	//

	CreateTodo(td *Todo) error

	UpdateTodo(td *Todo) error

	DeleteTodo(id int) error

	Todos() ([]*Todo, error)

	Todo(id int) (*Todo, error)
}
