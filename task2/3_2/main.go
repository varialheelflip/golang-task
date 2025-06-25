package main

func main() {
	e := &Employee{EmployeeID: 123, Person: Person{Name: "小明", Age: 10}}
	e.PrintInfo()
}
