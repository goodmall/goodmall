package usecase

// NewUserInteractor interactor creator
// 注意该方法 在[clean-go](https://github.com/CaptainCodeman/clean-go/blob/master/engine/greeter.go#L26)
// 项目中归属于工厂方法
func NewUserInteractor() UserInteractor{
	return UserInteractor{}
}

// UserInteractor allows to interact with user 
type UserInteractor struct{

}

// Register register a new user in system with the given name
func (interactor *UserInteractor) Register(	username string)(error){
	return nil 
}

func (interactor *UserInteractor) Help()(string){
	return "hi this is help method of UserInteractor!"
}