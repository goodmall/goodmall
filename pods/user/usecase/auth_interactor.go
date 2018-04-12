package usecase

type AuthInteractor struct {
}

func (itr *AuthInteractor) Login() {
	/*
		   Here I just have to search in my database (SQL, I know how to do it).
		If the user is registered, I create a token and give it to him,
		 but how can I do it?
	*/
}

func (itr *AuthInteractor) RefreshToken() {

}

func (itr *AuthInteractor) Logout() {
	/*
	   I get a token and stop/delete it?
	*/
}

// FIXME  不属于这里 属于用户功能！
func (itr *AuthInteractor) Register() {
	/*
	   I search if the user isn't register and then, if it isn't, I create a user in the database (I know how to do it). I connect him but again, how to make a new token?
	*/
}
