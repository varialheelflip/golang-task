package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID int
}

func (e *Employee) PrintInfo() {
	fmt.Println("EmployeeID:", e.EmployeeID)
	fmt.Println("Name:", e.Name)
	fmt.Println("Age:", e.Age)
}
