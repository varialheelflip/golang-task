package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

type Employee struct {
	Id         uint   `db:"id"`
	Name       string `db:"name"`
	Department string `db:"department"`
	Salary     uint   `db:"salary"`
}

func getEmployees(db *sqlx.DB) ([]Employee, error) {
	var employees []Employee
	err := db.Select(&employees, "SELECT * FROM employees where department = ?", "技术部")
	return employees, err
}

func getMaxSalary(db *sqlx.DB) ([]Employee, error) {
	var employees []Employee
	var maxSalary uint
	err := db.Get(&maxSalary, "SELECT max(salary) FROM employees")
	if err != nil {
		return employees, err
	}
	err = db.Select(&employees, "SELECT * FROM employees where salary = ?", maxSalary)
	return employees, err
}

func main() {
	db, err := sqlx.Connect("mysql", "root:root@tcp(192.168.200.130:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	//employees, err := getEmployees(db)
	//fmt.Println(employees)
	employees, err := getMaxSalary(db)
	fmt.Println(employees)
	fmt.Println(err)
}
