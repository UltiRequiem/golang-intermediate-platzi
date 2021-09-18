package main

import "fmt"

var (
	area = func(l int, b int) int {
		return l * b
	}
)

func main() {
	fmt.Println(area(20, 30))

	fmt.Println(fmt.Sprintf("100 (°F) = %.2f (°C).",
		func(f float64) float64 {
			return (f - 32.0) * (5.0 / 9.0)
		}(100),
	))
}
