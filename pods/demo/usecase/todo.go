package usecase

import (
	"github.com/goodmall/goodmall/base"

	"github.com/goodmall/goodmall/pods/demo"

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

/*
func (interactor *todoInteractor) Help() string {
	return "hi this is help method of TodoInteractor!"
}
*/

func (itr *todoInteractor) Get(id int) (*demo.Todo, error) {
	return itr.TodoRepo.Load(id)
}
func (itr *todoInteractor) Query(q base.Query) ([]demo.Todo, error) {

	return itr.TodoRepo.Query(base.Query{})
	// return []*demo.Todo{}, nil
}

func (itr *todoInteractor) Create(model *demo.Todo) (*demo.Todo, error) {
	/*
		if err := model.Validate(); err != nil {
			return nil, err
		}
	*/
	if err := itr.TodoRepo.Create(model); err != nil {
		return nil, err
	}
	return itr.TodoRepo.Load(model.Id)
}

func (itr *todoInteractor) Update(id int, model *demo.Todo) (*demo.Todo, error) {
	/*
		if err := model.Validate(); err != nil {
			return nil, err
		}
	*/
	if err := itr.TodoRepo.Update(id, model); err != nil {
		return nil, err
	}
	return itr.TodoRepo.Load(id)
}

func (itr *todoInteractor) Delete(id int) (*demo.Todo, error) {
	obj, err := itr.TodoRepo.Load(id)
	if err != nil {
		return nil, err
	}
	err = itr.TodoRepo.Remove(id)
	return obj, err

}

func (itr *todoInteractor) Count() (int, error) {
	return itr.TodoRepo.Count()
}
