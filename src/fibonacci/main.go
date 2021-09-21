package main

import "fmt"

func FibonnaciWorker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Println(fmt.Sprintf("Worker %d started with %d.", id, job))
		fib := Fibonacci(job)
		fmt.Println(fmt.Sprintf("Worker %d, job %d has gotten %d.", id, job, fib))
		results <- fib
	}
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {
	tasks := []int{2, 4, 3, 4, 5, 7, 8, 18, 12, 3, 4, 5, 5, 6, 66}
	tasklength := len(tasks)

	jobs := make(chan int, tasklength)
	results := make(chan int, tasklength)

	nWorkers := 3

	for i := 1; i < nWorkers+1; i++ {
		go FibonnaciWorker(i, jobs, results)
	}

	for _, value := range tasks {
		jobs <- value
	}

	for r := 0; r < tasklength; r++ {
		<-results
	}

	close(jobs)
}
