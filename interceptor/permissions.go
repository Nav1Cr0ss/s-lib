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

func PermissionInterceptor() grpc.UnaryServerInterceptor {

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		var (
			u   user.User
			md  metadata.MD
			err error
			ok  bool
		)

		md, ok = metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(codes.Internal, "internal error")
		}

		if uRaw := md.Get("user"); uRaw != nil {
			if err = json.Unmarshal([]byte(uRaw[0]), &u); err == nil {
				newCtx := context.WithValue(ctx, "user", u)

				return handler(newCtx, req)
			}
		}

		return handler(ctx, req)
	}
}

func TestInterceptor(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	// Logic before invoking the invoker
	// Calls the invoker to execute RPC
	err := invoker(ctx, method, req, reply, cc, opts...)
	// Logic after invoking the invoker
	return err
}
