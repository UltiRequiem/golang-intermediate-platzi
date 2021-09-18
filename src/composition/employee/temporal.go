package employee

import "github.com/UltiRequiem/golang-intermediate-platzi/src/composition/person"

type TemporaryEmployee struct {
	person.Person
	Employee
	TimeToWork int
}

func NewTemporaryEmployee(name string, age int, id int, timeToWork int) *TemporaryEmployee {
	return &TemporaryEmployee{
		Person:     person.Person{Name: name, Age: age},
		Employee:   Employee{Id: id},
		TimeToWork: timeToWork,
	}
}
