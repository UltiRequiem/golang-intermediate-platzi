package main

import "fmt"

func sum(numbers ...int) int {
	var total int
	for _, num := range numbers {
		total += num
	}

	return total
}

func main() {
	// Both Valid
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(1, 2, 5))
}
