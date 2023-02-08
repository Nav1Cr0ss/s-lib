package interceptor

import (
	"context"
	"time"

	"github.com/Nav1Cr0ss/s-lib/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func LoggerInterceptor(log *logger.Logger) grpc.UnaryServerInterceptor {

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

		h, err := handler(ctx, req)
		if err != nil {
			log.Infow("Invoked GRPC",
				"method", info.FullMethod,
				"time", time.Now(),
				"response_status", codes.OK,
				"error", err,
			)
			return h, err
		}
		log.Infow("Invoked GRPC",
			"method", info.FullMethod,
			"time", time.Now(),
			"response_status", codes.OK,
			"error", "",
		)
		return h, err
	}
}
