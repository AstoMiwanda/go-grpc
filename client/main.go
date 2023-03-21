package main

import (
	"context"
	"fmt"
	pb "go-grpc/student"
	"google.golang.org/grpc"
	"log"
	"time"
)

func getDataStudentByEmail(client pb.DataStudentClient, email string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	s := pb.Student{Email: email}
	student, err := client.FindStudentByEmail(ctx, &s)
	if err != nil {
		log.Fatalln("err get student", err)
	}

	fmt.Println(student)
}

func main() {
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	dial, err := grpc.Dial(":4000", opts...)
	if err != nil {
		log.Fatalln("err dial", err)
	}
	defer dial.Close()

	client := pb.NewDataStudentClient(dial)
	getDataStudentByEmail(client, "azza@mail.com")
	getDataStudentByEmail(client, "asto@mail.com")
}
