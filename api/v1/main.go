package api

import (
	"KVdb/api/middlewares"
	"net/http"
)

// set time out for 15 second
var TimeOut = 15

var Mux http.ServeMux

func init() {
	Mux.HandleFunc("/", middlewares.ReqTimeOut(func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hello world"))
	}, TimeOut))
}

func GetMux() http.ServeMux {
	return Mux
}
