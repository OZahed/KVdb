package middlewares

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type MKey string

func ReqIdmiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		reqId := uuid.New()
		ctx := r.Context()
		ctx = context.WithValue(ctx, MKey("reqId"), reqId)
		r = r.WithContext(ctx)
		next(rw, r)
	}
}
