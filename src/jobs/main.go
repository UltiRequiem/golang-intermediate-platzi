package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

type Job struct {
	Name   string
	Delay  time.Duration
	Number int
}

type Worker struct {
	ID         int
	JobQueue   chan Job
	WorkerPool chan chan Job
	QuitChan   chan bool
}

type Dispatcher struct {
	WorkerPool chan chan Job
	MaxWorkers int
	JobQueue   chan Job
}

func NewWorker(id int, workerPool chan chan Job) *Worker {
	return &Worker{
		ID: id, JobQueue: make(chan Job), WorkerPool: workerPool, QuitChan: make(chan bool),
	}
}

func (w Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobQueue

			select {
			case job := <-w.JobQueue:
				fmt.Println(fmt.Sprintf("Worker with id %d started.", w.ID))
				fib := Fibonacci(job.Number)
				time.Sleep(job.Delay)
				fmt.Println(fmt.Sprintf("Worker with id %d finished with result %d.", w.ID, fib))

			case <-w.QuitChan:
				fmt.Println(fmt.Sprintf("Worker with id %d finished.", w.ID))
			}

		}
	}()

}

func (w Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}

func NewDispatcher(jobQueue chan Job, maxWorkers int) *Dispatcher {
	return &Dispatcher{
		JobQueue:   jobQueue,
		MaxWorkers: maxWorkers,
		WorkerPool: make(chan chan Job, maxWorkers),
	}
}

func (d *Dispatcher) Dispatch() {
	for {
		select {
		case job := <-d.JobQueue:
			go func() {
				workerJobQueue := <-d.WorkerPool
				workerJobQueue <- job
			}()
		}
	}
}

func (d *Dispatcher) Run() {
	for i := 0; i < d.MaxWorkers; i++ {
		worker := NewWorker(i, d.WorkerPool)
		worker.Start()
	}

	go d.Dispatch()
}

func RequestHandler(w http.ResponseWriter, r *http.Request, joQueue chan Job) {
	w.Header().Set("Allow", "POST")

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	delay, err := time.ParseDuration(r.FormValue("delay"))

	if err != nil {
		http.Error(w, "Invalid delay.", http.StatusBadRequest)
		return
	}

	value, err := strconv.Atoi(r.FormValue("value"))

	if err != nil {
		http.Error(w, "Incorrect value.", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")

	if name == "" {
		http.Error(w, "Invalid Name.", http.StatusBadRequest)
		return
	}

	job := Job{name, delay, value}

	joQueue <- job

	w.WriteHeader(http.StatusCreated)

}

func main() {
	const (
		maxWorkers   = 4
		maxQueueSize = 20
		port         = ":3000"
	)

	jobQueue := make(chan Job, maxQueueSize)

	dispatcher := NewDispatcher(jobQueue, maxWorkers)

	dispatcher.Run()

	http.HandleFunc("/fib", func(w http.ResponseWriter, r *http.Request) {
		RequestHandler(w, r, jobQueue)
	})

	log.Fatal(http.ListenAndServe(port, nil))
}
