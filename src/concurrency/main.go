package main

import (
	"fmt"
	"time"
)

func chanela() {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func start() {
	c := make(chan int, 3)

	go func() {
		c <- 1
		time.Sleep(time.Second * 1)
		c <- 2
		c <- 3

		fmt.Println("I'll never be seen :(")
		c <- 80085
	}()

	for i := 0; i < 3; i++ {
		fmt.Println(<-c)
	}

	fmt.Println("End")
}
