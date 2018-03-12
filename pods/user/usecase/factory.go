package usecase

import (
	"github.com/goodmall/goodmall/pods/user/infra/repository/mysql"
)

type (
	// UsecaseFactory interface allows us to provide
	// other parts of the system with a way to make
	// instances of our use-case / interactors when
	// they need to
	UsecaseFactory interface {
		// NewUserInteractor creates a new UserInteractor interactor
		NewUserInteractor() UserInteractor // 此处考虑是否返回引用？

		// ...  其他本包的Interactor
	}

	// usecaseFactory stores the state of our usecase
	// which only involves a storage factory instance
	usecaseFactory struct {
		// StorageFactory  // TODO 后期添加这个依赖
	}
)

// NewEngine creates a new engine factory that will
// make use of the passed in StorageFactory for any
// data persistence needs.
// 该方法作为本层的 访问根  所有的访问开始于这里 然后你便需要“顺藤摸瓜”式的创建其他需要的对象了
func NewUsecase( /*s StorageFactory */ ) UsecaseFactory {
	return &usecaseFactory{ /*s*/ }
}

// NewUserInteractor  实现工厂接口 创造用户Interactor
// 该方法的位置 可以考虑下
//  - 提到interactors 文件去  代替常规的构造器实现方法 不过receiver是工厂而已 在常规的构造器模式中是没有接收者的
//  - 统一放在这里 该层每需要暴露一个公共类型 则添加一个构造器方法 用来创建该类型的实例
//    这种方法 感觉上使得interactor不知道存在工厂 所以这种方法感觉优于方法一
func (uf *usecaseFactory) NewUserInteractor() UserInteractor {
	// return NewUserInteractor()  // 代理到常规构造器实现
	return UserInteractor{
		userRepo: mysql.NewUserRepository(),
	}
}
