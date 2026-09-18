package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

const requestHeaderKey = "X-Request-ID"

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get(requestHeaderKey)
		if requestID == "" {
			requestID = uuid.New().String()
			w.Header().Set(requestHeaderKey, requestID)
		}
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "requestID", requestID)))
	})
}
