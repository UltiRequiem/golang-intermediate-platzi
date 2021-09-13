package main

import "fmt"

type Base struct {
	b   int
	tag string
}

func (base Base) DescribeTag() string {
	return fmt.Sprintf("Base tag is %s.", base.tag)
}

type Container struct {
	base Base
	c    string
	tag  string
}

func main() {

	b := Base{b: 10, tag: "b's tag"}
	fmt.Println(b.DescribeTag())

	co := Container{base: b, c: "foo", tag: "co's tag"}
	fmt.Println(co.base.DescribeTag())
}
