package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect"

	"github.com/gin-gonic/gin"
	"github.com/goodmall/goodmall/cmd/api/gin/engine"
)

var eng *httpexpect.Expect

func GetEngine(t *testing.T) *httpexpect.Expect {
	gin.SetMode(gin.TestMode)
	if eng == nil {
		server := httptest.NewServer(engine.GetMainEngine())
		eng = httpexpect.New(t, server.URL)
	}
	return eng
}

/*
func TestArticles(t *testing.T) {
	e := GetEngine(t)
	e.GET("/api/v1/articles").
		Expect().
		Status(http.StatusOK).
		JSON().Object().ContainsKey("data").Keys().Length().Ge(0)
}
*/

func TestTodoHandler_Query(t *testing.T) {
	e := GetEngine(t)
	e.GET("/todos").
		Expect().
		Status(http.StatusOK)
	/*.
	JSON().Object().ContainsKey("data").Keys().Length().Ge(0) */
}
