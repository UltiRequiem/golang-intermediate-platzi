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

func newFullTimeEmployee(name string, age int, id int) *FullTimeEmployee {
	return &FullTimeEmployee{
		Person{Name: name, Age: age},
		Employee{Id: id},
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
	PrintName(aFullTimeEmployee)
}
