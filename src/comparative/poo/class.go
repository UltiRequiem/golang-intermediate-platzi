// Very Pythonic
package main

import "fmt"

type GoEmployee struct {
	name string
	id   int
}

func (e GoEmployee) doWork() {
	fmt.Printf(fmt.Sprintf("%s (%d) is working with Golang code.", e.name, e.id))
}

func (e GoEmployee) String() string {
	return fmt.Sprintf("ID: %d, Name: %s", e.id, e.name)
}

func main() {
	e := GoEmployee{name: "Zero", id: 0}

	e.id = 44

	e.name = "Eliaz"

	fmt.Println(e)

	e.doWork()
}
