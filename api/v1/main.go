package api

import (
	"KVdb/api/v1/handlers"
	"KVdb/repo/store"
	"net/http"

	"github.com/labstack/echo/v4"
)

// set time out for 15 second
var TimeOut = 15

var Mux http.ServeMux
var e = echo.New()
var s store.Store

func init() {
	s = store.NewStore()
	e.GET("/:key", handlers.GetHandler(&s))
	e.POST("/", handlers.Sethandler(&s))
}

func GetMux() *echo.Echo {
	return e
}
