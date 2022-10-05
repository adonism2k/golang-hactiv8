package main

import (
	"fmt"
	"os"
	"strconv"
)

type Student struct {
	ID  int
	Name string
	Address string
	Job string
	Reason string
}


func main() {
	// make 20 students data
	var students [20]Student
	for i := 0; i < 20; i++ {
		students[i] = Student{
			ID: i,
			Name: "Student " + strconv.Itoa(i),
			Address: "Address " + strconv.Itoa(i),
			Job: "Job " + strconv.Itoa(i),
			Reason: "Reason " + strconv.Itoa(i),
		}
	}

	// get student id from command line
	id, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// print Student data formatted
	fmt.Println(students[id].StudentFormatted())
}

// format Student data to string
func (s Student) StudentFormatted() string {
	return fmt.Sprintf("ID: %d\nName: %s\nAddress: %s\nJob: %s\nReason: %s", s.ID, s.Name, s.Address, s.Job, s.Reason)
}