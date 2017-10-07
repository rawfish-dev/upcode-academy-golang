package main

import (
	"fmt"
)

type Student struct {
	Name string
}

// Receives a copy of the student instance and hence will not modify
func ChangeStudentName(student Student, newName string) {
	// Student is not the same instance that we passed in
	student.Name = newName
}

// Will modify the same student instance
// func ChangeStudentName(student *Student, newName string) {
// 	student.Name = newName
// }

func main() {
	student := Student{
		Name: "Kevin",
	}

	fmt.Println("Before:", student.Name)

	// Call the correct version
	ChangeStudentName(student, "Diana")
	// ChangeStudentName(&student, "Diana")

	fmt.Println("After:", student.Name)
}
