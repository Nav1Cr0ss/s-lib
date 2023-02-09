package interceptor

import (
	"context"
	"encoding/json"

	"github.com/Nav1Cr0ss/s-lib/domains/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func UserInterceptor() grpc.UnaryServerInterceptor {

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		var (
			u   user.User
			err error
		)

		if uRaw := metadata.ValueFromIncomingContext(ctx, "user"); len(uRaw) > 0 {
			if err = json.Unmarshal([]byte(uRaw[0]), &u); err == nil {
				newCtx := context.WithValue(ctx, "user", u)

				return handler(newCtx, req)
			}
		}

		return handler(ctx, req)
	}
}
