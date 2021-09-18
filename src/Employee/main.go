package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) getName() string {
	return p.Name
}

type Employee struct {
	Id int
}

type FullTimeEmployee struct {
	Person
	Employee
}

type TemporaryEmployee struct {
	Person
	Employee
	TimeToWork int
}

func newFullTimeEmployee(name string, age int, id int) *FullTimeEmployee {
	return &FullTimeEmployee{
		Person{Name: name, Age: age},
		Employee{Id: id},
	}
}

func newTemporaryEmployee(name string, age int, id int, timeToWork int) *TemporaryEmployee {
	return &TemporaryEmployee{
		Person:     Person{Name: name, Age: age},
		Employee:   Employee{Id: id},
		TimeToWork: timeToWork,
	}
}

type Human interface {
	getName() string
}

func PrintName(h Human) {
	fmt.Println(fmt.Sprintf("Hi, my name is %s.", h.getName()))
}

func main() {
	aFullTimeEmployee := newFullTimeEmployee("Zero", 14, 99)
	aTemporaryEmployee := newTemporaryEmployee("One", 16, 34, 3)

	PrintName(aFullTimeEmployee)
	PrintName(aTemporaryEmployee)
}
