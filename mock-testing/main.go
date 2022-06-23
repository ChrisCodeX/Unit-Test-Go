package main

import "time"

type Person struct {
	DNI  string
	Name string
	Age  int
}
type Employee struct {
	Id       int
	Position string
}
type FullTimeEmployee struct {
	Employee
	Person
}

func GetPersonByDNI(dni string) (Person, error) {
	time.Sleep(5 * time.Second)
	return Person{}, nil
}

func GetEmployeeById(id string) (Employee, error) {
	time.Sleep(5 * time.Second)
	return Employee{}, nil
}

func GetFullTimeEmployeeById(id int, dni string) (FullTimeEmployee, error) {
	var ftEmployee FullTimeEmployee
}
