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

	fmt.Println(fmt.Sprintf("The formated number is %d", value))
}

func maps() {
	m := make(map[string]int)

	m["myFavoriteNumber"] = 6

	fmt.Println(fmt.Sprintf("My favorite number is %d.", m["myFavoriteNumber"]))
}

func slices() {
	myNumbers := []int{1, 2, 3}

	// Add a value like:
	myNumbers = append(myNumbers, 4)

	for index, value := range myNumbers {
		fmt.Println(fmt.Sprintf("The index of %d is %d.", value, index))
	}

}
