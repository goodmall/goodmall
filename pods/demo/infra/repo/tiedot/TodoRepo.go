package tiedot

import (
	"github.com/goodmall/goodmall/pods/demo"
)

type todoRepo struct {
}

func NewTodoRepo() demo.TodoRepo {
	return &todoRepo{}
}

func (todoRepo *todoRepo) Load(id int) (demo.Todo, error) {
	return demo.Todo{}, nil
}
func (todoRepo *todoRepo) FindById(id int) demo.Todo {
	return demo.Todo{}
}

//
func (todoRepo *todoRepo) Store(td *demo.Todo) error {
	return nil
}
func (todoRepo *todoRepo) Remove(td *demo.Todo) {}

// ## Extra Behavior
// Size()

// ## Query
// Query(spec Specification)
func (todoRepo *todoRepo) Query(criteria demo.Query) {

}
