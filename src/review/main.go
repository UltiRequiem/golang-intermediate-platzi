package main

import (
	"fmt"
	"strconv"
)

func assigningVariables() {
	var x int = 8
	y := 8

	fmt.Println(fmt.Sprintf("I have %d apples and %d oranges.", x, y))
}

func managingErrors() {
	value, err := strconv.ParseInt("89e", 0, 64)

	if err != nil {
		fmt.Printf("%v", err)
                return
	}

        fmt.Println(value)
}
