package usecase

import (
	"github.com/goodmall/goodmall/base"
	"github.com/goodmall/goodmall/pods/demo"
	// "github.com/goodmall/goodmall/pods/user/domain"
	"github.com/asaskevich/EventBus"
)

// NewTodoInteractor interactor creator
// 注意该方法 在[clean-go](https://github.com/CaptainCodeman/clean-go/blob/master/engine/greeter.go#L26)
// 项目中归属于工厂方法
func NewTodoInteractor(todoRepo demo.TodoRepo, eb EventBus.Bus) demo.TodoInteractor {
	return &todoInteractor{
		TodoRepo: todoRepo,
		EventBus: eb,
	}
}

// TodoInteractor allows to interact with user
// FIXME 此处用例也可以定义为接口 外层也可以使用接口类型 这样增加可替换性 和易测试性
type todoInteractor struct {
	TodoRepo demo.TodoRepo
	EventBus EventBus.Bus
}

func (interactor *todoInteractor) Help() string {
	return "hi this is help method of TodoInteractor!"
}

func (interactor *todoInteractor) Todo(id int) (*demo.Todo, error) {
	// return &demo.Todo{Description: "hi funny!"}, nil
	return interactor.TodoRepo.Load(id)
}
func (interactor *todoInteractor) Todos() ([]*demo.Todo, error) {

	return interactor.TodoRepo.Query(base.Query{})
	// return []*demo.Todo{}, nil
}

func (interactor *todoInteractor) CreateTodo(td *demo.Todo) error {
	return interactor.TodoRepo.Store(td)
}

func (interactor *todoInteractor) UpdateTodo(td *demo.Todo) error {
	return nil
}
func (interactor *todoInteractor) DeleteTodo(id int) error {
	return interactor.TodoRepo.Remove(id)
}
