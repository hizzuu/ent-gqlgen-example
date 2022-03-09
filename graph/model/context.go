package model

type contextKey string

const (
	ClaimsCtxKey      contextKey = "claims"
	RequestIDCtxKey   contextKey = "request_id"
	UIDCtxKey         contextKey = "uid"
	CurrentUserCtxKey contextKey = "currentUser"
	AuthErrorCtxKey   contextKey = "authError"
)
