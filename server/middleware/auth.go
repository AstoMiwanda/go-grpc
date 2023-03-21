package middleware

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func isValidKey(keys []string) bool {
	return keys[0] == "astoazza"
}

func Auth(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	fmt.Printf("LOG: Middleware\n")
	fmt.Printf("LOG: incoming request: %s\n", info.FullMethod)
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "ApiKey not found")
	}

	keys := md.Get("apikey")
	if len(keys) < 1 {
		return nil, status.Error(codes.InvalidArgument, "ApiKey not found")
	}

	if !isValidKey(keys) {
		return nil, status.Error(codes.Unauthenticated, "Apikey not valid")
	}

	next, err := handler(ctx, req)
	if err != nil {
		return nil, err
	}

	return next, nil
}
