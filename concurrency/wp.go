package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	tasks := rand.Perm(50)
	fmt.Println(tasks)
	nWorkers := 10
	jobs := make(chan int, len(tasks))
	results := make(chan int, len(tasks))

	for i := 0; i < nWorkers; i++ {
		go Worker(i, jobs, results)
	}

	for _, value := range tasks {
		jobs <- value
	}

	close(jobs)

	for r := 0; r < len(tasks); r++ {
		<-results
	}
}

func Worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Started worker: '%d' with fib '%d'\n", id, job)
		fib := Fibonacci(job)
		fmt.Printf("Completed worker '%d' with job '%d' and fib '%d'\n", id, job, fib)
		results <- fib
	}
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}
