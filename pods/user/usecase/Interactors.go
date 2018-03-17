package usecase

import (
	"github.com/goodmall/goodmall/pods/user"
	"github.com/goodmall/goodmall/pods/user/domain"
	// "github.com/goodmall/goodmall/pods/user/domain"
)

// NewUserInteractor interactor creator
// 注意该方法 在[clean-go](https://github.com/CaptainCodeman/clean-go/blob/master/engine/greeter.go#L26)
// 项目中归属于工厂方法
func NewUserInteractor() UserInteractor {
	return UserInteractor{}
}

// UserInteractor allows to interact with user
// FIXME 此处用例也可以定义为接口 外层也可以使用接口类型 这样增加可替换性 和易测试性
type UserInteractor struct {
	userPod user.UserPod

	userRepo domain.UserRepository
}

// FindByUsername  通过用户名查找一个用户
func (interactor *UserInteractor) FindByUsername(username string) (*domain.User, error) {
	return interactor.userRepo.FindByUsername(username)
}

// Register register a new user in system with the given name
func (interactor *UserInteractor) Register(username string, password string) error {

	// 触发领域事件  UserCreated
	return nil
}

func (interactor *UserInteractor) Help() string {
	return "hi this is help method of UserInteractor!"
}

func (interactor *UserInteractor) ForgotPassword(user *domain.User) error {
	return nil
}
func (interactor *UserInteractor) ChangePassword(user *domain.User, password string) error {
	return nil
}
func (interactor *UserInteractor) Validate(user *domain.User) error {
	return nil
}
func (interactor *UserInteractor) Auth(user *domain.User, password string) error {
	return nil
}
func (interactor *UserInteractor) IsValid(user *domain.User) bool {
	return false
}
func (interactor *UserInteractor) GetRepo() *domain.UserRepository {
	return nil
}
