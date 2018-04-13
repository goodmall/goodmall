package controller

import (
	"fmt"
	"time"
	//	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	//	"github.com/gorilla/schema"

	"github.com/goodmall/goodmall/app"
	"github.com/goodmall/goodmall/base/errors"
	"github.com/goodmall/goodmall/pods/user/domain"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-ozzo/ozzo-validation"
)

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c Credential) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Username, validation.Required, validation.Length(3, 50)),

		validation.Field(&c.Password, validation.Required, validation.Length(6, 50)),
	)
}

/**
*     TODO to be continue : https://medium.com/@raul_11817/securing-golang-api-using-json-web-token-jwt-2dc363792a48
**/
type AuthController struct {
}

func (auth *AuthController) Login(c *gin.Context) {

	m := Credential{}

	c.Bind(&m)

	log.Printf("%#v \n", m)

	// 验证 用户输入
	// TODO 做用户名唯一性检测
	err := m.Validate()
	if err != nil {

		c.JSON(http.StatusBadRequest, err)
		return
	}
	// TODO 此处需要依赖auth-service|auth-interactor 来实现 暂时使用帮助方法模拟逻辑
	identity := authenticate(m)
	if identity == nil {
		e := errors.Unauthorized("invalid credential")
		c.JSON(http.StatusNonAuthoritativeInfo, e)
		return
	}

	signKey := app.Config.JWTSigningKey
	signMethod := app.Config.JWTSigningMethod

	// “HS256”说明这个令牌是通过HMAC-SHA256签名的。对称加密 加密和验证使用相同的钥匙
	// “RS256”非对称加密   加密跟验证key不同  私钥用来加密 公钥用来验证
	//create a rsa 256 signer
	signer := jwt.New(jwt.GetSigningMethod(signMethod))

	//set claims
	claims := make(jwt.MapClaims)
	claims["iss"] = "admin"
	claims["exp"] = time.Now().Add(time.Minute * 20).Unix()
	claims["CustomUserInfo"] = struct {
		Name string
		// Role string
	}{identity.GetName() /*, "Member" */}
	signer.Claims = claims

	tokenString, err := signer.SignedString([]byte(signKey))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, nil)
		//		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error while signing the token")
		log.Printf("Error signing token: %v\n", err)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

func (auth *AuthController) Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"anyKey": " Logout ",
	})
}

func (auth *AuthController) RefreshToken(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"anyKey": "refresh ",
	})
}

// --------------   私有 帮助方法 --------------  |

func authenticate(c Credential) domain.Identity {
	if c.Username == "demo" && c.Password == "pass" {
		return &domain.User{ID: "100", Name: "demo"}
	}
	return nil
}
