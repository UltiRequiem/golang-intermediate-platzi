package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	os.Stdout.Write([]byte("Main Running..." + "\n"))

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go doSomething(i, &wg)
	}

	os.Stdout.Write([]byte("Ending the proces..." + "\n"))

	wg.Wait()
}

func doSomething(i int, wg *sync.WaitGroup) {

	os.Stdout.Write([]byte(fmt.Sprintf("Started %d.\n", i)))

	time.Sleep(time.Second * 2)

	os.Stdout.Write([]byte("End\n"))

	wg.Done()
}
