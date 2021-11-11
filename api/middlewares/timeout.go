package middlewares

import (
	"context"
	"net/http"
	"time"
)

func ReqTimeOut(next http.HandlerFunc, tout int) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cln := context.WithCancel(r.Context())
		defer cln()
		// try-recive model
		r = r.WithContext(ctx)
		select {
		case <-time.After(time.Duration(tout) * time.Second):
			cln()
			return
		default:
			next(rw, r)
		}
	}
}
