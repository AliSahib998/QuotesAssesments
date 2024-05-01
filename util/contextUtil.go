package util

import (
	"context"
	"github.com/sirupsen/logrus"
	"net/http"
)

const (
	contextLogger = "contextLogger"
	contextHeader = "contextHeader"
)

func WithLogger(ctx context.Context, logger *logrus.Entry) context.Context {
	return context.WithValue(ctx, contextLogger, logger)
}
func WithHeader(ctx context.Context, header http.Header) context.Context {
	return context.WithValue(ctx, contextHeader, header)
}
