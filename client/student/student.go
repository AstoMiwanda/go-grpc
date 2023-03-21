package student

import (
	"context"
	"fmt"
	"go-grpc/proto/student"
	"log"
	"time"
)

func GetDataStudentByEmail(client student.DataStudentClient, email string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	s := student.Student{Email: email}
	data, err := client.FindStudentByEmail(ctx, &s)
	if err != nil {
		log.Fatalln("err get student", err)
	}

	fmt.Println(data)
}
