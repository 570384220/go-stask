package main

import "fmt"

func stu20() {

	employee := Employee{EmployeeId: 1, Person1: Person1{Name: "Tom", Age: 19}}
	fmt.Println(employee)
}

type Person1 struct {
	Name string
	Age  int
}

type Employee struct {
	EmployeeId int
	Person1
}
