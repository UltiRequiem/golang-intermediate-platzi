package employee

import "github.com/UltiRequiem/golang-intermediate-platzi/src/composition/person"

type FullTimeEmployee struct {
	person.Person
	Employee
}

func NewFullTimeEmployee(name string, age int, id int) *FullTimeEmployee {
	return &FullTimeEmployee{
		person.Person{Name: name, Age: age},
		Employee{Id: id},
	}
}
