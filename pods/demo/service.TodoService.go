package demo

// TodoInteractor represents a service for managing todos.
// alias : TodoService
// dependencies : tiedot-repo
type TodoInteractor interface {
	//
	Todos() ([]*Todo, error)

	Todo(id int) (*Todo, error)

	CreateTodo(td *Todo) error

	UpdateTodo(td *Todo) error

	DeleteTodo(id int) error
}
