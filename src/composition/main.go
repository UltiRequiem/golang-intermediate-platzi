package main

import (
	"fmt"
	. "github.com/UltiRequiem/golang-intermediate-platzi/src/composition/employee"
)

type Human interface {
	GetName() string
}

func PrintName(h Human) {
	fmt.Println(fmt.Sprintf("Hi, my name is %s.", h.GetName()))
}

func main() {
	aFullTimeEmployee := NewFullTimeEmployee("Zero", 14, 99)
	aTemporaryEmployee := NewTemporaryEmployee("One", 16, 34, 3)

	// Both works:
	// fmt.Println(aFullTimeEmployee.Person.Age)
	// fmt.Println(aFullTimeEmployee.Age)

	PrintName(aFullTimeEmployee)
	PrintName(aTemporaryEmployee)
}
