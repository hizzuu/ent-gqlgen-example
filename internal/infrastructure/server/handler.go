package server

import (
	"context"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/hizzuu/plate-backend/graph/model"
	"github.com/hizzuu/plate-backend/internal/domain"
	"github.com/hizzuu/plate-backend/internal/infrastructure/firebase"
)

func authentication(next http.Handler, client firebase.AuthClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := client.VerifyIDToken(r.Context(), getIDTokenFromHeader(r))
		if err != nil {
			next.ServeHTTP(w, setErrContext(r, err))
			return
		}

		next.ServeHTTP(w, setUIDToContext(r, token.UID))
	})
}

func getIDTokenFromHeader(r *http.Request) string {
	authorization := r.Header.Get("Authorization")
	return strings.Replace(authorization, "Bearer ", "", 1)
}

func setErrContext(r *http.Request, err error) *http.Request {
	return r.WithContext(
		context.WithValue(
			r.Context(),
			model.AuthErrorCtxKey,
			domain.NewUnauthorizedError(err),
		),
	)
}

func setUIDToContext(r *http.Request, uid string) *http.Request {
	return r.WithContext(
		context.WithValue(
			r.Context(),
			model.UIDCtxKey,
			uid,
		),
	)
}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		next.ServeHTTP(w, r)
	})
}

func requestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get("X-Request-Id")
		if requestID == "" {
			requestID = uuid.NewString()
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), model.RequestIDCtxKey, requestID)))
	})
}
