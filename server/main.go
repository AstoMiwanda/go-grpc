package main

import (
	"go-grpc/proto/echo"
	"go-grpc/proto/student"
	pbEcho "go-grpc/server/echo"
	"go-grpc/server/middleware"
	pbStudent "go-grpc/server/student"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":4000")
	if err != nil {
		log.Fatalln("error listen", err)
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(middleware.Auth))
	student.RegisterDataStudentServer(grpcServer, pbStudent.NewStudentServer())
	echo.RegisterEchoServer(grpcServer, pbEcho.NewEchoServer())

	log.Println("Running on :4000")
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalln("err server", err)
	}
}
