// Very Pythonic
package main

import "fmt"

type GoEmployee struct {
	Name   string
	Id     int
	canFly bool
}

func (e GoEmployee) String() string {
	return fmt.Sprintf("Id: %d, Name: %s, Can Fly: %t.", e.Id, e.Name, e.canFly)
}

func (e *GoEmployee) SetName(name string) {
	e.Name = name
}

func (e *GoEmployee) SetId(id int) {
	e.Id = id
}

func newGoEmployee(name string, id int, canFly bool) *GoEmployee {
	return &GoEmployee{
		Name:   name,
		Id:     id,
		canFly: canFly,
	}
}

func main() {
	e := newGoEmployee("Zero", 34, true)

	e.SetName("Eliaz")
	e.SetId(44)

	fmt.Println(e)
}
