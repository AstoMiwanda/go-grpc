package student

import (
	"context"
	"fmt"
	"go-grpc/proto/echo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type DataEchoServer struct {
	echo.UnimplementedEchoServer
}

func (d *DataEchoServer) UnaryEcho(ctx context.Context, in *echo.EchoRequest) (*echo.EchoResponse, error) {
	fmt.Printf("\n--UnaryEcho--\n")
	fmt.Printf("incoming request:\n")
	fmt.Printf("message: %s", in.Message)

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.DataLoss, "error unary echo")
	}
	ids := md.Get("id")
	if len(ids) > 0 {
		fmt.Printf("\nfrom metadata context id\n")
		for i, id := range ids {
			fmt.Printf("index: %d, value: %s\n", i, id)
		}
	}

	return &echo.EchoResponse{Message: in.Message}, nil
}

func NewEchoServer() *DataEchoServer {
	return &DataEchoServer{}
}
