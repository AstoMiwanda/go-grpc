package student

import (
	"context"
	"encoding/json"
	"fmt"
	"go-grpc/proto/student"
	"log"
	"os"
	"path/filepath"
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
	// Get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	// Build the path to the data.json file using filepath.Join()
	dataPath := filepath.Join(wd, "server/data", "students.json")
	file, err := os.ReadFile(dataPath)
	if err != nil {
		log.Fatalln("Error read file", err)
	}

	if err := json.Unmarshal(file, &d.students); err != nil {
		log.Fatalln("Error unmarshal data json", err)
	}
	fmt.Printf("INFO: %v\n", d.students)
}

func NewStudentServer() *DataStudentServer {
	s := DataStudentServer{}
	s.loadData()
	return &s
}
