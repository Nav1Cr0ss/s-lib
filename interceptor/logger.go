package interceptor

import (
	"context"
	"time"

	"github.com/Nav1Cr0ss/s-lib/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CodeToLevel(code codes.Code) zapcore.Level {
	switch code {
	case codes.OK:
		return zap.InfoLevel
	case codes.Canceled:
		return zap.InfoLevel
	case codes.Unknown:
		return zap.ErrorLevel
	case codes.InvalidArgument:
		return zap.InfoLevel
	case codes.DeadlineExceeded:
		return zap.WarnLevel
	case codes.NotFound:
		return zap.InfoLevel
	case codes.AlreadyExists:
		return zap.InfoLevel
	case codes.PermissionDenied:
		return zap.WarnLevel
	case codes.Unauthenticated:
		return zap.InfoLevel
	case codes.ResourceExhausted:
		return zap.WarnLevel
	case codes.FailedPrecondition:
		return zap.WarnLevel
	case codes.Aborted:
		return zap.WarnLevel
	case codes.OutOfRange:
		return zap.WarnLevel
	case codes.Unimplemented:
		return zap.ErrorLevel
	case codes.Internal:
		return zap.ErrorLevel
	case codes.Unavailable:
		return zap.WarnLevel
	case codes.DataLoss:
		return zap.ErrorLevel
	default:
		return zap.ErrorLevel
	}
}

func ErrorToCode(err error) codes.Code {
	return status.Code(err)
}

func LoggerInterceptor(log *logger.Logger) grpc.UnaryServerInterceptor {

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

		h, err := handler(ctx, req)
		code := ErrorToCode(err)
		level := CodeToLevel(code)

		log.CL.Log(level, "invoked GRPC",
			zap.String("method", info.FullMethod),
			zap.String("code", code.String()),
			zap.Time("time", time.Now()),
			zap.Error(err),
		)

		return h, err
	}
}
