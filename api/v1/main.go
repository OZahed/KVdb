package api

import (
	"KVdb/api/v1/handlers"
	"KVdb/repo/store"
	"KVdb/scadule"
	"KVdb/tasks"
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

// set time out for 15 second
var (
	TimeOut  = 15
	CleaFreq = 10 * time.Minute
	Mux      http.ServeMux
	e        = echo.New()
	s        store.Store
	l        = log.New(os.Stdout, "kv DB:: ", log.LstdFlags)
	path     = "./snapp.josn"
)

func init() {
	s = store.NewStore()
	err := scadule.ScaduleTask(tasks.CleanTask(context.Background(), CleaFreq, &s), "Clean", l)
	if err != nil {
		log.Fatal(err)
	}
	err = scadule.ScaduleTask(tasks.SnapShotTask(context.Background(), CleaFreq, &s, path), "Snap", l)
	if err != nil {
		log.Fatal(err)
	}
	e.GET("/:key", handlers.GetHandler(&s))
	e.POST("/", handlers.Sethandler(&s))
}

func GetMux() *echo.Echo {
	return e
}
