package main

import "time"

type Person struct {
	Name string
	DNI  int
	Age  int
}

type Employee struct {
	ID       int
	Position string
}

type FullTimeEmployee struct {
	Employee
	Person
}

var GetPersonByDNI = func(dni int) (Person, error) {
	time.Sleep(5 * time.Second)
	// SELECT * FROM Persona where ...
	return Person{}, nil
}

var GetEmployeeByID = func(id int) (Employee, error) {
	time.Sleep(5 * time.Second)
	return Employee{}, nil
}

func GetFullTimeEmployeeByID(id int, dni int) (FullTimeEmployee, error) {
	var ftEmployee FullTimeEmployee

	employee, err := GetEmployeeByID(id)

	if err != nil {
		return ftEmployee, err
	}

	ftEmployee.Employee = employee

	person, err := GetPersonByDNI(dni)
	if err != nil {
		return FullTimeEmployee{}, err
	}

	ftEmployee.Person = person

	return ftEmployee, nil

}
