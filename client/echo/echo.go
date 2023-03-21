package echo

import (
	"context"
	"fmt"
	"go-grpc/proto/echo"
	"google.golang.org/grpc/metadata"
)

func CallerUnaryEcho(client echo.EchoClient, message string) {
	fmt.Printf("\n--CallerUnaryEcho--\n")

	md := metadata.Pairs("id", "123")
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	unaryEcho, err := client.UnaryEcho(ctx, &echo.EchoRequest{Message: message})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response: %s", unaryEcho.Message)
}
