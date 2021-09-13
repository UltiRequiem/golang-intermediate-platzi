package main

import (
	"fmt"
	"log"
	"time"
)

func doSomething(c chan string, index int) {
	time.Sleep(3 * time.Second)
	c <- fmt.Sprintf("I've already slept 3 seconds %d times.", index)
}

func main() {
	c := make(chan string)

	for index := range [10]int{} {
		go doSomething(c, index+1)
	}

	for range [10]int{} {
		log.Println(<-c)
	}

}
