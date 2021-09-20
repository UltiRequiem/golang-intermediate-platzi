package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := make(chan int, 10)

	var wg sync.WaitGroup
	defer wg.Wait()

	for i := 0; i < 10; i++ {
		c <- i
		wg.Add(1)
		go doSomething(i, &wg, c)
	}

}
func doSomething(i int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()

	fmt.Println(fmt.Sprintf("ID %d started.", i))

	time.Sleep(2 * time.Second)

	fmt.Println(fmt.Sprintf("ID %d ended.", i))

	<-c
}
