package student

import (
	"context"
	"encoding/json"
	"go-grpc/proto/student"
	"log"
	"os"
	"sync"
)

type DataStudentServer struct {
	student.UnimplementedDataStudentServer
	mu       sync.Mutex
	students []*student.Student
}

func (d *DataStudentServer) FindStudentByEmail(ctx context.Context, student *student.Student) (*student.Student, error) {
	for _, v := range d.students {
		if v.Email == student.Email {
			return v, nil
		}
	}
	return nil, nil
}

func (d *DataStudentServer) loadData() {
	file, err := os.ReadFile("data/students.json")
	if err != nil {
		log.Fatalln("Error read file", err)
	}

	if err := json.Unmarshal(file, &d.students); err != nil {
		log.Fatalln("Error unmarshal data json", err)
	}
}

func NewStudentServer() *DataStudentServer {
	s := DataStudentServer{}
	s.loadData()
	return &s
}
