package interceptor

import (
	"context"
	"encoding/json"

	"github.com/Nav1Cr0ss/s-lib/domains/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func SetUserInterceptor() grpc.UnaryServerInterceptor {

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		var (
			u   user.User
			md  metadata.MD
			err error
			ok  bool
		)

		md, ok = metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(codes.PermissionDenied, "user doesn't have permissions to do this")
		}

		uRaw := md.Get("user")
		if uRaw == nil {
			return nil, status.Errorf(codes.PermissionDenied, "user doesn't have permissions to do this")
		}

		if err = json.Unmarshal([]byte(uRaw[0]), &u); err != nil {
			return nil, status.Errorf(codes.PermissionDenied, "user doesn't have permissions to do this")
		}

		newCtx := context.WithValue(ctx, "user", u)

		return handler(newCtx, req)
	}
}
