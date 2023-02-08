package interceptor

import (
	"context"

	"google.golang.org/grpc"
)

type User struct {
	Id int `json:"id"`
}

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		newCtx := context.WithValue(ctx, "User", User{Id: 123})
		return handler(newCtx, req)
	}
}
