package repository

import "context"

type Logger interface {
	Info(template string, args ...interface{})
	InfoWithCtx(ctx context.Context, template string, args ...interface{})
	Fatal(err error)
	Error(err error)
}
