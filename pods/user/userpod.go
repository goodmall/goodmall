package user

type UserPod struct {
}

func NewUserPod(options ...func(*UserPod)) (*UserPod, error) {
	pod := UserPod{}

	for _, opt := range options {
		opt(&pod)
	}

	return &pod, nil
}

/*
func (up *UserPod) Configure() {

}
func (up *UserPod) Init() {

}
*/
