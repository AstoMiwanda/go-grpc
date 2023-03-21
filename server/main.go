package main

import (
	"context"
	"encoding/json"
	pb "go-grpc/student"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"sync"
)

type dataStudentServer struct {
	pb.UnimplementedDataStudentServer
	mu       sync.Mutex
	students []*pb.Student
}

func (d *dataStudentServer) FindStudentByEmail(ctx context.Context, student *pb.Student) (*pb.Student, error) {
	for _, v := range d.students {
		if v.Email == student.Email {
			return v, nil
		}
	}
	return nil, nil
}

func (d *dataStudentServer) loadData() {
	file, err := os.ReadFile("data/students.json")
	if err != nil {
		log.Fatalln("Error read file", err)
	}

	if err := json.Unmarshal(file, &d.students); err != nil {
		log.Fatalln("Error unmarshal data json", err)
	}
}

func newServer() *dataStudentServer {
	s := dataStudentServer{}
	s.loadData()
	return &s
}

func main() {
	listen, err := net.Listen("tcp", ":4000")
	if err != nil {
		log.Fatalln("error listen", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterDataStudentServer(grpcServer, newServer())

	log.Println("Running on :4000")
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalln("err server", err)
	}
}
