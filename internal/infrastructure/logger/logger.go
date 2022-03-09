package logger

import (
	"context"
	"os"

	"github.com/hizzuu/plate-backend/graph/model"
	"github.com/rs/zerolog"
)

type logger struct {
	log zerolog.Logger
}

func New() *logger {
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()
	return &logger{log: log}
}

type Logger interface {
	Infof(template string, args ...interface{})
	InfofWithCtx(ctx context.Context, template string, args ...interface{})
	Fatal(err error)
	Error(err error)
}

func (l *logger) Infof(template string, args ...interface{}) {
	l.log.Info().Msgf(template, args...)
}

func (l *logger) InfofWithCtx(ctx context.Context, template string, args ...interface{}) {
	l.log.Info().Interface("request_id", ctx.Value(model.RequestIDCtxKey)).Msgf(template, args...)
}

func (l *logger) Error(err error) {
	l.log.Error().Err(err).Send()
}

func (l *logger) Fatal(err error) {
	l.log.Fatal().Err(err).Send()
}
