package main

import (
	"fmt"
	"time"
)

func DoSomething(i time.Duration, c chan<- int, param int) {
	time.Sleep(i)
	c <- param * 2
}

func main() {
	chaOne, chaTwo := make(chan int), make(chan int)
	duraOne, duraTwo := 4*time.Second, 2*time.Second
	valuOne, valuTwo := 2, 1

	go DoSomething(duraOne, chaOne, valuTwo)
	go DoSomething(duraTwo, chaTwo, valuOne)

	for i := 0; i < 2; i++ {
		select {
		case channelMsgOne := <-chaOne:
			fmt.Println(channelMsgOne)
		case channelMsgTwo := <-chaTwo:
			fmt.Println(channelMsgTwo)
		}
	}
}
