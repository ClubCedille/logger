package logger

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	RequestPathField       = "requestPath"
	RequestQueryParamField = "queryPath"
	HttpMethodField        = "httpMethod"
	GrpcMethod             = "grpcMethod"
	Logger                 = "logger"
)

// HTTPLoggerMiddleware
// Middleware for HTTP requests adding a logger to each.
func HTTPLoggerMiddleware(logger logrus.FieldLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		entry := logger.
			WithField(RequestPathField, c.Request.URL.Path).
			WithField(RequestQueryParamField, c.Request.URL.RawQuery).
			WithField(HttpMethodField, c.Request.Method)

		c.Set(Logger, entry)

		c.Next()
	}
}

// GrpcLoggerMiddleware
// Middleware for gRPC calls adding a logger to each.
func GrpcLoggerMiddleware(logger logrus.FieldLogger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		grpcRequestLogger := logger.WithField(GrpcMethod, info.FullMethod)
		ctxWithLogger := context.WithValue(ctx, CtxKey, grpcRequestLogger)
		return handler(ctxWithLogger, req)
	}
}
