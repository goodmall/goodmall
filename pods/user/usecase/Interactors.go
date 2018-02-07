package usecase

// NewUserInteractor interactor creator
func NewUserInteractor() *UserInteractor{
	return &UserInteractor{}
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