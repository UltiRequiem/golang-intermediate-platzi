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

func main() {
	e := GoEmployee{name: "Zero", id: 0}

	e.id = 44

	e.name = "Eliaz"

	fmt.Println(fmt.Sprintf("%v", e))

	e.doWork()
}
