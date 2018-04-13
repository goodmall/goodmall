package usecase

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/goodmall/goodmall/pods/user/domain"
)

type AuthInteractor interface {
	GetToken(u domain.User) (string, error)
}

// NewAuthInteractor  生成验证服务
func NewAuthInteractor() AuthInteractor {
	return &authInteractor{}
}

var _ AuthInteractor = &authInteractor{}

type authInteractor struct {
	// 签名key 私钥 由外部配置文件传入
	SigningKey    []byte
	SigningMethod string

	// TODO 依赖UserInteractor 或者是 UserRepo ？ 两个都可以哦！

}

// TODO 方法签名变为 GetToken(u User) (token string, error)
func (itr *authInteractor) GetToken(u domain.User) (string, error) {
	/*
		   Here I just have to search in my database (SQL, I know how to do it).
		If the user is registered, I create a token and give it to him,

	*/

	// Create token
	// token := jwt.New(jwt.SigningMethodHS256)
	token := jwt.New(jwt.GetSigningMethod(itr.SigningMethod))

	/*
		// Try to log in the user
		user, err := s.UserService.Read(u.ID)
		if err != nil {
			return "", errors.New("Failed to retrieve user")
		}
		if user == nil {
			return "", errors.New("Failed to retrieve user")
		}
	*/

	//set claims

	claims := make(jwt.MapClaims)
	claims["iss"] = "admin"
	claims["exp"] = time.Now().Add(time.Minute * 20).Unix()
	claims["CustomUserInfo"] = struct {
		Name string
		// Role string
	}{u.Name /*, "Member" */}

	token.Claims = claims

	// Sign token with key
	tokenString, err := token.SignedString(itr.SigningKey)
	if err != nil {
		return "", errors.New("Failed to sign token")
	}

	return tokenString, nil

}

func (itr *authInteractor) RefreshToken() {

}

func (itr *authInteractor) Logout() {
	/*
		   I get a token and stop/delete it?

		client 端的很容易 登出后只需要销毁掉token就行了
		server 端  可能需要类似的session_id 机制 在token中存一个id  服务端只需要重生id 就可以失效掉token
		@see https://stackoverflow.com/questions/21978658/invalidating-json-web-tokens
	*/
}

// FIXME  不属于这里 属于用户功能！
func (itr *authInteractor) Register() {
	/*
	   I search if the user isn't register and then, if it isn't, I create a user in the database (I know how to do it). I connect him but again, how to make a new token?
	*/
}
