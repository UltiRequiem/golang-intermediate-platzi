package main

import "fmt"

func sum(numbers ...int) int {
	var total int
	for _, num := range numbers {
		total += num
	}

	return total
}

func areFriends(names ...string) {
	var message string
	var totalFriends int = len(names)

	switch totalFriends {

	case 1:
		message = fmt.Sprintf("%s has no friends.", names[0])
	case 2:
		message = fmt.Sprintf("%s and %s are friends.", names[0], names[1])
	default:
		for index, name := range names {
			switch index {
			case totalFriends - 1:
				message += fmt.Sprintf("and %s", name)
			case totalFriends - 2:
				message += fmt.Sprintf("%s ", name)
			default:
				message += fmt.Sprintf("%s, ", name)

			}

		}
		message = fmt.Sprintf("%s are friends.", message)
	}

	fmt.Println(message)
}

func main() {
	// All Valid
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(1, 2, 5))
	fmt.Println(sum(1))

	areFriends("Zero", "Pedro", "Juanita")
	areFriends("Suzaku", "Golem")
	areFriends("Rockman")
}
