package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/goodmall/goodmall/pods/user/usecase"

	"github.com/goodmall/goodmall/pods/user/adapters/api/gin/controller"
	//	"github.com/dgrijalva/jwt-go"
)

// InitPod 集成入口  系统应用（SysApp）可用通过此方法把该模块的功能集成到系统总体版图去

//  TODO 重构为类型的方法 NewUserPod( /* 依赖传入 */ ).Init()
//       pod := NewUserPod()
//       pod.Xxx = Xxx
//       pod.Init()

func InitPod(engine *gin.Engine) {

	r := engine

	r.GET("/userhelp", func(c *gin.Context) {

		// userInteractor := usecase.NewUserInteractor()
		userInteractor := usecase.NewUsecase( /* 依赖暂缺 */ ).NewUserInteractor() // usecase.NewUserInteractor()
		response := userInteractor.Help()

		c.JSON(200, gin.H{
			"message": response,
		})

	})

	r.GET("/username", func(c *gin.Context) {

		// userInteractor := usecase.NewUserInteractor()
		userInteractor := usecase.NewUsecase( /* 依赖暂缺 */ ).NewUserInteractor() // usecase.NewUserInteractor()

		un := c.Query("username")
		user, err := userInteractor.FindByUsername(un)
		if err != nil {

			c.JSON(200, gin.H{
				"message": "查不到呀!",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "" + un + " : " + fmt.Sprint(user),
		})

	})

	// TODO 通过 jwt-token-validation middleware 来保护需要的功能访问
	authCtrl := &controller.AuthController{}
	r.POST("/token-auth", authCtrl.Login)
	r.GET("/refresh-token-auth", authCtrl.RefreshToken)
	r.GET("/logout", authCtrl.Logout)

	r.GET("/status", NotImplemented)

	// TODO  我们可以在初始化方法中 触发一些事件 供内部钩子注册
}

/**
* 			 这是个不错的技巧 可以先占位：
*
*   r.GET("/some-route", NotImplemented)
*
*      当你先买了个鸟笼 你迟早会忍不住买个鸟关里面的
**/
var NotImplemented = func(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"note": " this route is not implemented yet ! ",
	})
}

/* Middleware handler to handle all requests for authentication */
func JWTAuthenticationMiddleware(next func(c *gin.Context)) func(c *gin.Context) {

	return func(c *gin.Context) {
		/*
			authorizationHeader := c.GetHeader("authorization")
			if authorizationHeader != "" {
				bearerToken := strings.Split(authorizationHeader, " ")
				if len(bearerToken) == 2 {
					token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
						if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
							// return nil, fmt.Errorf("There was an error")
						}
						return []byte("secret"), nil
					})

					if error != nil {
						// json.NewEncoder(w).Encode(Exception{Message: error.Error()})
						return
					}

					if token.Valid {
						log.Println("TOKEN WAS VALID")
						context.Set(req, "decoded", token.Claims)

						next(c)
					} else {
						// Exception{Message: "Invalid authorization token"}
					}
				}
			} else {
				//  "An authorization header is required"
			}
		*/
	}

}
