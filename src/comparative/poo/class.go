// Very Pythonic
package main

import "fmt"

type GoEmployee struct {
	Name string
	Id   int
}

func (e GoEmployee) String() string {
	return fmt.Sprintf("Id: %d, Name: %s", e.Id, e.Name)
}

func (e *GoEmployee) SetName(name string) {
	e.Name = name
}

func (e *GoEmployee) SetId(id int) {
	e.Id = id
}

func main() {
	e := GoEmployee{Name: "Zero", Id: 8965}

        e.SetName("Eliaz")
        e.SetId(44)

	fmt.Println(e)
}
