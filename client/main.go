package main

import (
	"go-grpc/client/echo"
	"go-grpc/client/student"
	pbEcho "go-grpc/proto/echo"
	pbStudent "go-grpc/proto/student"
	"google.golang.org/grpc"
	"log"
)

func main() {
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	dial, err := grpc.Dial(":4000", opts...)
	if err != nil {
		log.Fatalln("err dial", err)
	}
	defer func() {
		if err := dial.Close(); err != nil {
			panic(err)
		}
	}()

	clientStudent := pbStudent.NewDataStudentClient(dial)
	student.GetDataStudentByEmail(clientStudent, "azza@mail.com")
	student.GetDataStudentByEmail(clientStudent, "asto@mail.com")

	clientEcho := pbEcho.NewEchoClient(dial)
	echo.CallerUnaryEcho(clientEcho, "message-1")
	echo.CallerUnaryEcho(clientEcho, "message-2")
}
