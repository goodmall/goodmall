package main

import (
	"fmt"
	"net/http"
	
	"github.com/labstack/echo"

	"github.com/goodmall/goodmall/pods/user/usecase"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
// Route => handler
e.GET("/userhelp", func(c echo.Context) error {

	 
		// userInteractor := usecase.NewUserInteractor()
		userInteractor :=  usecase.NewUsecase(/* 依赖暂缺 */).NewUserInteractor() // usecase.NewUserInteractor()
		response := userInteractor.Help() 
		fmt.Println(response)

	return c.String(http.StatusOK, response)
})

	e.Logger.Fatal(e.Start(":1323"))
}
